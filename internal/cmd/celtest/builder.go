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

package main

import (
	"fmt"
	"sync"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pvcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/bufbuild/protovalidate-go/resolve"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

// cache is a build-through cache to computed standard constraints.
// type cache struct {
// 	cache map[protoreflect.FieldDescriptor]astSet
// }

type builder struct {
	mtx sync.Mutex // serializes cache writes.
	env *cel.Env
	// constraints           cache
	extensionTypeResolver protoregistry.ExtensionTypeResolver
	// allowUnknownFields    bool
	// Load                  func(desc protoreflect.MessageDescriptor) messageEvaluator
}

// newBuilder initializes a new Builder.
func newBuilder(
	env *cel.Env,
	// disableLazy bool,
	// extensionTypeResolver protoregistry.ExtensionTypeResolver,
	// allowUnknownFields bool,
	// seedDesc ...protoreflect.MessageDescriptor,
) *builder {
	bldr := &builder{
		env: env,
		// constraints:           newCache(),
		// extensionTypeResolver: extensionTypeResolver,
		// allowUnknownFields:    allowUnknownFields,
	}
	return bldr
}

func (bldr *builder) build(desc protoreflect.MessageDescriptor) *message {
	msgEval := &message{}
	bldr.buildMessage(desc, msgEval)
	return msgEval
}

func (bldr *builder) load(desc protoreflect.MessageDescriptor) *message {
	return bldr.build(desc)
}

func (bldr *builder) buildMessage(
	desc protoreflect.MessageDescriptor, msgEval *message,
) {
	msgConstraints := resolve.MessageConstraints(desc)
	if msgConstraints.GetDisabled() {
		return
	}

	steps := []func(
		desc protoreflect.MessageDescriptor,
		msgConstraints *validate.MessageConstraints,
		msg *message,
	){
		bldr.processMessageExpressions,
		bldr.processFields,
	}

	for _, step := range steps {
		step(desc, msgConstraints, msgEval)
	}

	// bldr.processFields(desc, msgEval)
}

func (bldr *builder) processMessageExpressions(
	desc protoreflect.MessageDescriptor,
	msgConstraints *validate.MessageConstraints,
	msgEval *message,
) {
	exprs := expressions{
		Constraints: msgConstraints.GetCel(),
	}

	compiledProgram, err := compile(
		exprs,
		bldr.env,
		cel.Types(dynamicpb.NewMessage(desc)),
		cel.Variable("this", cel.ObjectType(string(desc.FullName()))),
	)
	if err != nil {
		msgEval.Err = err
		return
	}

	msgEval.program = compiledProgram
}

func (bldr *builder) processFields(
	desc protoreflect.MessageDescriptor,
	msgConstraints *validate.MessageConstraints,
	msgEval *message,
) {
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)
		fieldConstraints := resolve.FieldConstraints(fdesc)
		fld, err := bldr.buildField(fdesc, fieldConstraints)
		if err != nil {
			fld.Err = err
			return
		}
		msgEval.fieldEval = fld
	}
}

func (bldr *builder) buildField(
	fieldDescriptor protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
) (field, error) {
	fld := field{
		Value: value{
			Descriptor: fieldDescriptor,
		},
	}
	err := bldr.buildValue(fieldDescriptor, fieldConstraints, &fld.Value)
	return fld, err
}

func (bldr *builder) buildValue(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	valEval *value,
) (err error) {
	steps := []func(
		fdesc protoreflect.FieldDescriptor,
		fieldConstraints *validate.FieldConstraints,
		valEval *value,
	) error{
		bldr.processFieldExpressions,
		// bldr.processEmbeddedMessage,
		// bldr.processWrapperConstraints,
		// bldr.processStandardConstraints,
		// bldr.processMapConstraints,
	}

	for _, step := range steps {
		if err = step(fdesc, constraints, valEval); err != nil {
			return err
		}
	}
	return nil
}

type expressions struct {
	Constraints []*validate.Constraint
	RulePath    []*validate.FieldPathElement
}

func (bldr *builder) processFieldExpressions(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	eval *value,
) error {
	exprs := expressions{
		Constraints: fieldConstraints.GetCel(),
	}
	celTyp := pvcel.ProtoFieldToType(fieldDesc, false, false)
	fmt.Printf("%+v", celTyp)
	opts := append(
		pvcel.RequiredEnvOptions(fieldDesc),
		cel.Variable("this", celTyp),
	)
	compiledProgram, err := compile(exprs, bldr.env, opts...)
	if err != nil {
		return err
	}
	eval.program = compiledProgram
	return nil
}
