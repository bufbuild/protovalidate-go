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

package protovalidate

import (
	"fmt"
	"sync"
	"sync/atomic"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pvcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/bufbuild/protovalidate-go/resolve"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

//nolint:gochecknoglobals
var (
	celRuleDescriptor = (&validate.FieldConstraints{}).ProtoReflect().Descriptor().Fields().ByName("cel")
	celRuleField      = fieldPathElement(celRuleDescriptor)
)

// builder is a build-through cache of message evaluators keyed off the provided
// descriptor.
type builder struct {
	mtx                   sync.Mutex                   // serializes cache writes.
	cache                 atomic.Pointer[messageCache] // copy-on-write cache.
	env                   *cel.Env
	constraints           cache
	extensionTypeResolver protoregistry.ExtensionTypeResolver
	allowUnknownFields    bool
	Load                  func(desc protoreflect.MessageDescriptor) messageEvaluator
}

// newBuilder initializes a new Builder.
func newBuilder(
	env *cel.Env,
	disableLazy bool,
	extensionTypeResolver protoregistry.ExtensionTypeResolver,
	allowUnknownFields bool,
	seedDesc ...protoreflect.MessageDescriptor,
) *builder {
	bldr := &builder{
		env:                   env,
		constraints:           newCache(),
		extensionTypeResolver: extensionTypeResolver,
		allowUnknownFields:    allowUnknownFields,
	}

	if disableLazy {
		bldr.Load = bldr.load
	} else {
		bldr.Load = bldr.loadOrBuild
	}

	cache := make(messageCache, len(seedDesc))
	for _, desc := range seedDesc {
		bldr.build(desc, cache)
	}
	bldr.cache.Store(&cache)
	return bldr
}

// load returns a pre-cached MessageEvaluator for the given descriptor or, if
// the descriptor is unknown, returns an evaluator that always resolves to a
// errors.CompilationError.
func (bldr *builder) load(desc protoreflect.MessageDescriptor) messageEvaluator {
	if eval, ok := (*bldr.cache.Load())[desc]; ok {
		return eval
	}
	return unknownMessage{desc: desc}
}

// loadOrBuild either returns a memoized MessageEvaluator for the given
// descriptor or lazily constructs a new one. This method is thread-safe via
// locking.
func (bldr *builder) loadOrBuild(desc protoreflect.MessageDescriptor) messageEvaluator {
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

func (bldr *builder) build(
	desc protoreflect.MessageDescriptor,
	cache messageCache,
) *message {
	if eval, ok := cache[desc]; ok {
		return eval
	}
	msgEval := &message{}
	cache[desc] = msgEval
	bldr.buildMessage(desc, msgEval, cache)
	return msgEval
}

func (bldr *builder) buildMessage(
	desc protoreflect.MessageDescriptor, msgEval *message,
	cache messageCache,
) {
	msgConstraints := resolve.MessageConstraints(desc)
	if msgConstraints.GetDisabled() {
		return
	}

	steps := []func(
		desc protoreflect.MessageDescriptor,
		msgConstraints *validate.MessageConstraints,
		msg *message,
		cache messageCache,
	){
		bldr.processMessageExpressions,
		bldr.processOneofConstraints,
		bldr.processFields,
	}

	for _, step := range steps {
		step(desc, msgConstraints, msgEval, cache)
	}
}

func (bldr *builder) processMessageExpressions(
	desc protoreflect.MessageDescriptor,
	msgConstraints *validate.MessageConstraints,
	msgEval *message,
	_ messageCache,
) {
	exprs := expressions{
		Constraints: msgConstraints.GetCel(),
	}
	compiledExprs, err := compile(
		exprs,
		bldr.env,
		cel.Types(dynamicpb.NewMessage(desc)),
		cel.Variable("this", cel.ObjectType(string(desc.FullName()))),
	)
	if err != nil {
		msgEval.Err = err
		return
	}

	msgEval.Append(celPrograms{
		programSet: compiledExprs,
	})
}

func (bldr *builder) processOneofConstraints(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
	_ messageCache,
) {
	oneofs := desc.Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneofDesc := oneofs.Get(i)
		oneofConstraints := resolve.OneofConstraints(oneofDesc)
		oneofEval := oneof{
			Descriptor: oneofDesc,
			Required:   oneofConstraints.GetRequired(),
		}
		msgEval.AppendNested(oneofEval)
	}
}

func (bldr *builder) processFields(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
	cache messageCache,
) {
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)
		fieldConstraints := resolve.FieldConstraints(fdesc)
		fldEval, err := bldr.buildField(fdesc, fieldConstraints, cache)
		if err != nil {
			fldEval.Err = err
		}
		msgEval.AppendNested(fldEval)
	}
}

func (bldr *builder) buildField(
	fieldDescriptor protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	cache messageCache,
) (field, error) {
	fld := field{
		Value: value{
			Descriptor: fieldDescriptor,
		},
		Required:     fieldConstraints.GetRequired(),
		IgnoreAlways: bldr.shouldIgnoreAlways(fieldConstraints),
		IgnoreEmpty: fieldDescriptor.HasPresence() ||
			bldr.shouldIgnoreEmpty(fieldConstraints),
		IgnoreDefault: fieldDescriptor.HasPresence() &&
			bldr.shouldIgnoreDefault(fieldConstraints),
	}
	if fld.IgnoreDefault {
		fld.Zero = bldr.zeroValue(fieldDescriptor, false)
	}
	err := bldr.buildValue(fieldDescriptor, fieldConstraints, &fld.Value, cache)
	return fld, err
}

func (bldr *builder) buildValue(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	valEval *value,
	cache messageCache,
) (err error) {
	if bldr.shouldIgnoreAlways(constraints) {
		return nil
	}

	steps := []func(
		fdesc protoreflect.FieldDescriptor,
		fieldConstraints *validate.FieldConstraints,
		valEval *value,
		cache messageCache,
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
		if err = step(fdesc, constraints, valEval, cache); err != nil {
			return err
		}
	}
	return nil
}

func (bldr *builder) processIgnoreEmpty(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	val *value,
	_ messageCache,
) error {
	// the only time we need to ignore empty on a value is if it's evaluating a
	// field item (repeated element or map key/value).
	val.IgnoreEmpty = val.NestedRule != nil && bldr.shouldIgnoreEmpty(constraints)
	if val.IgnoreEmpty {
		val.Zero = bldr.zeroValue(fdesc, val.NestedRule != nil)
	}
	return nil
}

func (bldr *builder) processFieldExpressions(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	eval *value,
	_ messageCache,
) error {
	exprs := expressions{
		Constraints: fieldConstraints.GetCel(),
	}

	celTyp := pvcel.ProtoFieldToType(fieldDesc, false, eval.NestedRule != nil)
	opts := append(
		pvcel.RequiredEnvOptions(fieldDesc),
		cel.Variable("this", celTyp),
	)
	compiledExpressions, err := compile(exprs, bldr.env, opts...)
	if err != nil {
		return err
	}
	for i := range compiledExpressions {
		compiledExpressions[i].Path = []*validate.FieldPathElement{
			{
				FieldNumber: proto.Int32(celRuleField.GetFieldNumber()),
				FieldType:   celRuleField.GetFieldType().Enum(),
				FieldName:   proto.String(celRuleField.GetFieldName()),
				Subscript: &validate.FieldPathElement_Index{
					Index: uint64(i), //nolint:gosec // indices are guaranteed to be non-negative
				},
			},
		}
		compiledExpressions[i].Descriptor = celRuleDescriptor
	}
	if len(compiledExpressions) > 0 {
		eval.Constraints = append(eval.Constraints,
			celPrograms{
				base:       newBase(eval),
				programSet: compiledExpressions,
			},
		)
	}
	return nil
}

func (bldr *builder) processEmbeddedMessage(
	fdesc protoreflect.FieldDescriptor,
	_ *validate.FieldConstraints,
	valEval *value,
	cache messageCache,
) error {
	if !isMessageField(fdesc) ||
		fdesc.IsMap() ||
		(fdesc.IsList() && valEval.NestedRule == nil) {
		return nil
	}

	embedEval := bldr.build(fdesc.Message(), cache)
	if err := embedEval.Err; err != nil {
		return &CompilationError{cause: fmt.Errorf(
			"failed to compile embedded type %s for %s: %w",
			fdesc.Message().FullName(), fdesc.FullName(), err)}
	}
	valEval.AppendNested(&embeddedMessage{
		base:    newBase(valEval),
		message: embedEval,
	})

	return nil
}

func (bldr *builder) processWrapperConstraints(
	fdesc protoreflect.FieldDescriptor,
	rules *validate.FieldConstraints,
	valEval *value,
	cache messageCache,
) error {
	if !isMessageField(fdesc) ||
		fdesc.IsMap() ||
		(fdesc.IsList() && valEval.NestedRule == nil) {
		return nil
	}

	expectedWrapperDescriptor, ok := expectedWrapperConstraints(fdesc.Message().FullName())
	if !ok || !rules.ProtoReflect().Has(expectedWrapperDescriptor) {
		return nil
	}
	unwrapped := value{
		Descriptor: valEval.Descriptor,
		NestedRule: valEval.NestedRule,
	}
	err := bldr.buildValue(fdesc.Message().Fields().ByName("value"), rules, &unwrapped, cache)
	if err != nil {
		return err
	}
	valEval.Append(unwrapped.Constraints)
	return nil
}

func (bldr *builder) processStandardConstraints(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	valEval *value,
	_ messageCache,
) error {
	stdConstraints, err := bldr.constraints.Build(
		bldr.env,
		fdesc,
		constraints,
		bldr.extensionTypeResolver,
		bldr.allowUnknownFields,
		valEval.NestedRule != nil,
	)
	if err != nil {
		return err
	}
	valEval.Append(celPrograms{
		base:       newBase(valEval),
		programSet: stdConstraints,
	})
	return nil
}

func (bldr *builder) processAnyConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	valEval *value,
	_ messageCache,
) error {
	if (fdesc.IsList() && valEval.NestedRule == nil) ||
		!isMessageField(fdesc) ||
		fdesc.Message().FullName() != "google.protobuf.Any" {
		return nil
	}

	typeURLDesc := fdesc.Message().Fields().ByName("type_url")
	anyPbDesc := (&validate.AnyRules{}).ProtoReflect().Descriptor()
	inField := anyPbDesc.Fields().ByName("in")
	notInField := anyPbDesc.Fields().ByName("not_in")
	anyEval := anyPB{
		base:              newBase(valEval),
		TypeURLDescriptor: typeURLDesc,
		In:                stringsToSet(fieldConstraints.GetAny().GetIn()),
		NotIn:             stringsToSet(fieldConstraints.GetAny().GetNotIn()),
		InValue:           fieldConstraints.GetAny().ProtoReflect().Get(inField),
		NotInValue:        fieldConstraints.GetAny().ProtoReflect().Get(notInField),
	}
	valEval.Append(anyEval)
	return nil
}

func (bldr *builder) processEnumConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	valEval *value,
	_ messageCache,
) error {
	if fdesc.Kind() != protoreflect.EnumKind {
		return nil
	}
	if fieldConstraints.GetEnum().GetDefinedOnly() {
		valEval.Append(definedEnum{
			base:             newBase(valEval),
			ValueDescriptors: fdesc.Enum().Values(),
		})
	}
	return nil
}

func (bldr *builder) processMapConstraints(
	fieldDesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	valEval *value,
	cache messageCache,
) error {
	if !fieldDesc.IsMap() {
		return nil
	}

	mapEval := newKVPairs(valEval)

	err := bldr.buildValue(
		fieldDesc.MapKey(),
		constraints.GetMap().GetKeys(),
		&mapEval.KeyConstraints,
		cache)
	if err != nil {
		return &CompilationError{cause: fmt.Errorf(
			"failed to compile key constraints for map %s: %w",
			fieldDesc.FullName(), err)}
	}

	err = bldr.buildValue(
		fieldDesc.MapValue(),
		constraints.GetMap().GetValues(),
		&mapEval.ValueConstraints,
		cache)
	if err != nil {
		return &CompilationError{cause: fmt.Errorf(
			"failed to compile value constraints for map %s: %w",
			fieldDesc.FullName(), err)}
	}

	valEval.Append(mapEval)
	return nil
}

func (bldr *builder) processRepeatedConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	valEval *value,
	cache messageCache,
) error {
	if !fdesc.IsList() || valEval.NestedRule != nil {
		return nil
	}

	listEval := newListItems(valEval)

	err := bldr.buildValue(fdesc, fieldConstraints.GetRepeated().GetItems(), &listEval.ItemConstraints, cache)
	if err != nil {
		return &CompilationError{cause: fmt.Errorf(
			"failed to compile items constraints for repeated %v: %w", fdesc.FullName(), err)}
	}

	valEval.Append(listEval)
	return nil
}

func (bldr *builder) shouldIgnoreAlways(constraints *validate.FieldConstraints) bool {
	return constraints.GetIgnore() == validate.Ignore_IGNORE_ALWAYS
}

func (bldr *builder) shouldIgnoreEmpty(constraints *validate.FieldConstraints) bool {
	return constraints.GetIgnore() == validate.Ignore_IGNORE_IF_UNPOPULATED ||
		constraints.GetIgnore() == validate.Ignore_IGNORE_IF_DEFAULT_VALUE
}

func (bldr *builder) shouldIgnoreDefault(constraints *validate.FieldConstraints) bool {
	return constraints.GetIgnore() == validate.Ignore_IGNORE_IF_DEFAULT_VALUE
}

func (bldr *builder) zeroValue(fdesc protoreflect.FieldDescriptor, forItems bool) protoreflect.Value {
	switch {
	case forItems && fdesc.IsList():
		msg := dynamicpb.NewMessage(fdesc.ContainingMessage())
		return msg.Get(fdesc).List().NewElement()
	case isMessageField(fdesc) &&
		fdesc.Cardinality() != protoreflect.Repeated:
		msg := dynamicpb.NewMessage(fdesc.Message())
		return protoreflect.ValueOfMessage(msg)
	default:
		return fdesc.Default()
	}
}

type messageCache map[protoreflect.MessageDescriptor]*message

func (c messageCache) Clone() messageCache {
	newCache := make(messageCache, len(c)+1)
	c.SyncTo(newCache)
	return newCache
}
func (c messageCache) SyncTo(other messageCache) {
	for k, v := range c {
		other[k] = v
	}
}

// isMessageField returns true if the field descriptor fdesc describes a field
// containing a submessage. Although they are represented differently on the
// wire, group fields are treated like message fields in protoreflect and have
// similar properties. In the 2023 edition of protobuf, message fields with the
// delimited encoding feature will be detected as groups, but should otherwise
// be treated the same.
func isMessageField(fdesc protoreflect.FieldDescriptor) bool {
	return fdesc.Kind() == protoreflect.MessageKind ||
		fdesc.Kind() == protoreflect.GroupKind
}
