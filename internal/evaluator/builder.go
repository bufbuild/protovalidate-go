// Copyright 2023-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package evaluator

import (
	"sync"
	"sync/atomic"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go/celext"
	"github.com/bufbuild/protovalidate-go/internal/constraints"
	"github.com/bufbuild/protovalidate-go/internal/errors"
	"github.com/bufbuild/protovalidate-go/internal/expression"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// Builder is a build-through cache of message evaluators keyed off the provided
// descriptor.
type Builder struct {
	mtx         sync.Mutex                   // serializes cache writes.
	cache       atomic.Pointer[MessageCache] // copy-on-write cache.
	env         *cel.Env
	constraints constraints.Cache
	resolver    StandardConstraintResolver
	Load        func(desc protoreflect.MessageDescriptor) MessageEvaluator
}

type StandardConstraintResolver interface {
	ResolveMessageConstraints(desc protoreflect.MessageDescriptor) *validate.MessageConstraints
	ResolveOneofConstraints(desc protoreflect.OneofDescriptor) *validate.OneofConstraints
	ResolveFieldConstraints(desc protoreflect.FieldDescriptor) *validate.FieldConstraints
}

// NewBuilder initializes a new Builder.
func NewBuilder(
	env *cel.Env,
	disableLazy bool,
	res StandardConstraintResolver,
	seedDesc ...protoreflect.MessageDescriptor,
) *Builder {
	bldr := &Builder{
		env:         env,
		constraints: constraints.NewCache(),
		resolver:    res,
	}

	if disableLazy {
		bldr.Load = bldr.load
	} else {
		bldr.Load = bldr.loadOrBuild
	}

	cache := make(MessageCache, len(seedDesc))
	for _, desc := range seedDesc {
		bldr.build(desc, cache)
	}
	bldr.cache.Store(&cache)
	return bldr
}

// load returns a pre-cached MessageEvaluator for the given descriptor or, if
// the descriptor is unknown, returns an evaluator that always resolves to a
// errors.CompilationError.
func (bldr *Builder) load(desc protoreflect.MessageDescriptor) MessageEvaluator {
	if eval, ok := (*bldr.cache.Load())[desc]; ok {
		return eval
	}
	return unknownMessage{desc: desc}
}

// loadOrBuild either returns a memoized MessageEvaluator for the given
// descriptor, or lazily constructs a new one. This method is thread-safe via
// locking.
func (bldr *Builder) loadOrBuild(desc protoreflect.MessageDescriptor) MessageEvaluator {
	if eval, ok := (*bldr.cache.Load())[desc]; ok {
		return eval
	}
	bldr.mtx.Lock()
	defer bldr.mtx.Unlock()
	cache := *bldr.cache.Load()
	if eval, ok := cache[desc]; ok {
		return eval
	}
	newCache := cache.Clone()
	msgEval := bldr.build(desc, newCache)
	bldr.cache.Store(&newCache)
	return msgEval
}

func (bldr *Builder) build(
	desc protoreflect.MessageDescriptor,
	cache MessageCache,
) *message {
	if eval, ok := cache[desc]; ok {
		return eval
	}
	msgEval := &message{}
	cache[desc] = msgEval
	bldr.buildMessage(desc, msgEval, cache)
	return msgEval
}

func (bldr *Builder) buildMessage(
	desc protoreflect.MessageDescriptor, msgEval *message,
	cache MessageCache,
) {
	msgConstraints := bldr.resolver.ResolveMessageConstraints(desc)
	if msgConstraints.GetDisabled() {
		return
	}

	steps := []func(
		desc protoreflect.MessageDescriptor,
		msgConstraints *validate.MessageConstraints,
		msg *message,
		cache MessageCache,
	){
		bldr.processMessageExpressions,
		bldr.processOneofConstraints,
		bldr.processFields,
	}

	for _, step := range steps {
		if step(desc, msgConstraints, msgEval, cache); msgEval.Err != nil {
			break
		}
	}
}

func (bldr *Builder) processMessageExpressions(
	desc protoreflect.MessageDescriptor,
	msgConstraints *validate.MessageConstraints,
	msgEval *message,
	_ MessageCache,
) {
	compiledExprs, err := expression.Compile(
		msgConstraints.GetCel(),
		bldr.env,
		cel.Types(dynamicpb.NewMessage(desc)),
		cel.Variable("this", cel.ObjectType(string(desc.FullName()))),
	)
	if err != nil {
		msgEval.Err = err
		return
	}

	msgEval.Append(celPrograms(compiledExprs))
}

func (bldr *Builder) processOneofConstraints(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
	_ MessageCache,
) {
	oneofs := desc.Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneofDesc := oneofs.Get(i)
		oneofConstraints := bldr.resolver.ResolveOneofConstraints(oneofDesc)
		oneofEval := oneof{
			Descriptor: oneofDesc,
			Required:   oneofConstraints.GetRequired(),
		}
		msgEval.Append(oneofEval)
	}
}

func (bldr *Builder) processFields(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
	cache MessageCache,
) {
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)
		fieldConstraints := bldr.resolver.ResolveFieldConstraints(fdesc)
		fldEval, err := bldr.buildField(fdesc, fieldConstraints, cache)
		if err != nil {
			msgEval.Err = err
			return
		}
		msgEval.Append(fldEval)
	}
}

func (bldr *Builder) buildField(
	fieldDescriptor protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	cache MessageCache,
) (field, error) {
	fld := field{
		Descriptor: fieldDescriptor,
		Required:   fieldConstraints.GetRequired(),
		IgnoreEmpty: fieldDescriptor.HasPresence() ||
			bldr.shouldIgnoreEmpty(fieldConstraints),
		IgnoreDefault: fieldDescriptor.HasPresence() &&
			bldr.shouldIgnoreDefault(fieldConstraints),
	}
	if fld.IgnoreDefault {
		fld.Zero = bldr.zeroValue(fieldDescriptor, false)
	}
	err := bldr.buildValue(fieldDescriptor, fieldConstraints, false, &fld.Value, cache)
	return fld, err
}

func (bldr *Builder) buildValue(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	cache MessageCache,
) (err error) {
	steps := []func(
		fdesc protoreflect.FieldDescriptor,
		fieldConstraints *validate.FieldConstraints,
		forItems bool,
		valEval *value,
		cache MessageCache,
	) error{
		bldr.processIgnoreEmpty,
		bldr.processFieldExpressions,
		bldr.processEmbeddedMessage,
		bldr.processWrapperConstraints,
		bldr.processStandardConstraints,
		bldr.processAnyConstraints,
		bldr.processEnumConstraints,
		bldr.processMapConstraints,
		bldr.processRepeatedConstraints,
	}

	for _, step := range steps {
		if err = step(fdesc, constraints, forItems, valEval, cache); err != nil {
			return err
		}
	}
	return nil
}

func (bldr *Builder) processIgnoreEmpty(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	forItems bool,
	val *value,
	_ MessageCache,
) error {
	// the only time we need to ignore empty on a value is if it's evaluating a
	// field item (repeated element or map key/value).
	val.IgnoreEmpty = forItems && bldr.shouldIgnoreEmpty(constraints)
	if val.IgnoreEmpty {
		val.Zero = bldr.zeroValue(fdesc, forItems)
	}
	return nil
}

func (bldr *Builder) processFieldExpressions(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
	eval *value,
	_ MessageCache,
) error {
	exprs := fieldConstraints.GetCel()
	if len(exprs) == 0 {
		return nil
	}

	celTyp := celext.ProtoFieldToCELType(fieldDesc, false, forItems)
	opts := append(
		celext.RequiredCELEnvOptions(fieldDesc),
		cel.Variable("this", celTyp),
	)
	compiledExpressions, err := expression.Compile(exprs, bldr.env, opts...)
	if err != nil {
		return err
	}
	if len(compiledExpressions) > 0 {
		eval.Constraints = append(eval.Constraints, celPrograms(compiledExpressions))
	}
	return nil
}

func (bldr *Builder) processEmbeddedMessage(
	fdesc protoreflect.FieldDescriptor,
	rules *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	cache MessageCache,
) error {
	if fdesc.Kind() != protoreflect.MessageKind ||
		bldr.shouldSkip(rules) ||
		fdesc.IsMap() ||
		(fdesc.IsList() && !forItems) {
		return nil
	}

	embedEval := bldr.build(fdesc.Message(), cache)
	if err := embedEval.Err; err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile embedded type %s for %s: %w",
			fdesc.Message().FullName(), fdesc.FullName(), err)
	}
	valEval.Append(embedEval)

	return nil
}

func (bldr *Builder) processWrapperConstraints(
	fdesc protoreflect.FieldDescriptor,
	rules *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	cache MessageCache,
) error {
	if fdesc.Kind() != protoreflect.MessageKind ||
		bldr.shouldSkip(rules) ||
		fdesc.IsMap() ||
		(fdesc.IsList() && !forItems) {
		return nil
	}

	expectedWrapperDescriptor, ok := constraints.ExpectedWrapperConstraints(fdesc.Message().FullName())
	if !ok || !rules.ProtoReflect().Has(expectedWrapperDescriptor) {
		return nil
	}
	var unwrapped value
	err := bldr.buildValue(fdesc.Message().Fields().ByName("value"), rules, true, &unwrapped, cache)
	if err != nil {
		return err
	}
	valEval.Append(unwrapped.Constraints)
	return nil
}

func (bldr *Builder) processStandardConstraints(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	_ MessageCache,
) error {
	stdConstraints, err := bldr.constraints.Build(
		bldr.env,
		fdesc,
		constraints,
		forItems,
	)
	if err != nil {
		return err
	}
	valEval.Append(celPrograms(stdConstraints))
	return nil
}

func (bldr *Builder) processAnyConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	_ MessageCache,
) error {
	if (fdesc.IsList() && !forItems) ||
		fdesc.Kind() != protoreflect.MessageKind ||
		fdesc.Message().FullName() != "google.protobuf.Any" {
		return nil
	}

	typeURLDesc := fdesc.Message().Fields().ByName("type_url")
	anyEval := anyPB{
		TypeURLDescriptor: typeURLDesc,
		In:                stringsToSet(fieldConstraints.GetAny().GetIn()),
		NotIn:             stringsToSet(fieldConstraints.GetAny().GetNotIn()),
	}
	valEval.Append(anyEval)
	return nil
}

func (bldr *Builder) processEnumConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	_ bool,
	valEval *value,
	_ MessageCache,
) error {
	if fdesc.Kind() != protoreflect.EnumKind {
		return nil
	}
	if fieldConstraints.GetEnum().GetDefinedOnly() {
		valEval.Append(definedEnum{ValueDescriptors: fdesc.Enum().Values()})
	}
	return nil
}

func (bldr *Builder) processMapConstraints(
	fieldDesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	_ bool,
	valEval *value,
	cache MessageCache,
) error {
	if !fieldDesc.IsMap() {
		return nil
	}

	var mapEval kvPairs

	err := bldr.buildValue(
		fieldDesc.MapKey(),
		constraints.GetMap().GetKeys(),
		true,
		&mapEval.KeyConstraints,
		cache)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile key constraints for map %s: %w",
			fieldDesc.FullName(), err)
	}

	err = bldr.buildValue(
		fieldDesc.MapValue(),
		constraints.GetMap().GetValues(),
		true,
		&mapEval.ValueConstraints,
		cache)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile value constraints for map %s: %w",
			fieldDesc.FullName(), err)
	}

	valEval.Append(mapEval)
	return nil
}

func (bldr *Builder) processRepeatedConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
	cache MessageCache,
) error {
	if !fdesc.IsList() || forItems {
		return nil
	}

	var listEval listItems
	err := bldr.buildValue(fdesc, fieldConstraints.GetRepeated().GetItems(), true, &listEval.ItemConstraints, cache)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile items constraints for repeated %v: %w", fdesc.FullName(), err)
	}

	valEval.Append(listEval)
	return nil
}

func (bldr *Builder) shouldSkip(constraints *validate.FieldConstraints) bool {
	return constraints.GetSkipped() ||
		constraints.GetIgnore() == validate.Ignore_IGNORE_ALWAYS
}

func (bldr *Builder) shouldIgnoreEmpty(constraints *validate.FieldConstraints) bool {
	return constraints.GetIgnoreEmpty() ||
		constraints.GetIgnore() == validate.Ignore_IGNORE_IF_UNPOPULATED ||
		constraints.GetIgnore() == validate.Ignore_IGNORE_IF_DEFAULT_VALUE
}

func (bldr *Builder) shouldIgnoreDefault(constraints *validate.FieldConstraints) bool {
	return constraints.GetIgnore() == validate.Ignore_IGNORE_IF_DEFAULT_VALUE
}

func (bldr *Builder) zeroValue(fdesc protoreflect.FieldDescriptor, forItems bool) protoreflect.Value {
	switch {
	case forItems && fdesc.IsList():
		msg := dynamicpb.NewMessage(fdesc.ContainingMessage())
		return msg.Get(fdesc).List().NewElement()
	case fdesc.Kind() == protoreflect.MessageKind &&
		fdesc.Cardinality() != protoreflect.Repeated:
		msg := dynamicpb.NewMessage(fdesc.Message())
		return protoreflect.ValueOfMessage(msg)
	default:
		return fdesc.Default()
	}
}

type MessageCache map[protoreflect.MessageDescriptor]*message

func (c MessageCache) Clone() MessageCache {
	newCache := make(MessageCache, len(c)+1)
	c.SyncTo(newCache)
	return newCache
}
func (c MessageCache) SyncTo(other MessageCache) {
	for k, v := range c {
		other[k] = v
	}
}
