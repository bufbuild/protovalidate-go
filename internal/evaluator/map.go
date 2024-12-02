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
	"fmt"
	"strconv"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

//nolint:gochecknoglobals
var (
	mapRuleDescriptor     = (&validate.FieldConstraints{}).ProtoReflect().Descriptor().Fields().ByName("map")
	mapKeysRuleDescriptor = (&validate.MapRules{}).ProtoReflect().Descriptor().Fields().ByName("keys")
	mapKeysRulePath       = []*validate.FieldPathElement{
		errors.FieldPathElement(mapRuleDescriptor),
		errors.FieldPathElement(mapKeysRuleDescriptor),
	}
	mapValuesDescriptor = (&validate.MapRules{}).ProtoReflect().Descriptor().Fields().ByName("values")
	mapValuesRulePath   = []*validate.FieldPathElement{
		errors.FieldPathElement(mapRuleDescriptor),
		errors.FieldPathElement(mapValuesDescriptor),
	}
)

// kvPairs performs validation on a map field's KV Pairs.
type kvPairs struct {
	// Descriptor is the FieldDescriptor targeted by this evaluator
	Descriptor protoreflect.FieldDescriptor
	// KeyConstraints are checked on the map keys
	KeyConstraints value
	// ValueConstraints are checked on the map values
	ValueConstraints value
}

func (m kvPairs) Evaluate(val protoreflect.Value, failFast bool) (err error) {
	var ok bool
	val.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		evalErr := m.evalPairs(key, value, failFast)
		if evalErr != nil {
			element := errors.FieldPathElement(m.Descriptor)
			element.KeyType = descriptorpb.FieldDescriptorProto_Type(m.Descriptor.MapKey().Kind()).Enum()
			element.ValueType = descriptorpb.FieldDescriptorProto_Type(m.Descriptor.MapValue().Kind()).Enum()
			switch m.Descriptor.MapKey().Kind() {
			case protoreflect.BoolKind:
				element.Subscript = &validate.FieldPathElement_BoolKey{BoolKey: key.Bool()}
			case protoreflect.Int32Kind, protoreflect.Int64Kind,
				protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
				protoreflect.Sint32Kind, protoreflect.Sint64Kind:
				element.Subscript = &validate.FieldPathElement_IntKey{IntKey: key.Int()}
			case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
				protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
				element.Subscript = &validate.FieldPathElement_UintKey{UintKey: key.Uint()}
			case protoreflect.StringKind:
				element.Subscript = &validate.FieldPathElement_StringKey{StringKey: key.String()}
			case protoreflect.EnumKind, protoreflect.FloatKind, protoreflect.DoubleKind,
				protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
				err = errors.NewCompilationErrorf(
					"unexpected map key type %s",
					m.Descriptor.MapKey().Kind(),
				)
				return false
			}
			errors.AppendFieldPath(evalErr, element, false)
		}
		ok, err = errors.Merge(err, evalErr, failFast)
		return ok
	})
	return err
}

func (m kvPairs) evalPairs(key protoreflect.MapKey, value protoreflect.Value, failFast bool) (err error) {
	evalErr := m.KeyConstraints.Evaluate(key.Value(), failFast)
	errors.MarkForKey(evalErr)
	ok, err := errors.Merge(err, evalErr, failFast)
	if !ok {
		return err
	}

	evalErr = m.ValueConstraints.Evaluate(value, failFast)
	_, err = errors.Merge(err, evalErr, failFast)
	return err
}

func (m kvPairs) Tautology() bool {
	return m.KeyConstraints.Tautology() &&
		m.ValueConstraints.Tautology()
}

func (m kvPairs) formatKey(key any) string {
	switch k := key.(type) {
	case string:
		return strconv.Quote(k)
	default:
		return fmt.Sprintf("%v", key)
	}
}

// keysWrapper wraps the evaluation of nested map key rules.
type keysWrapper struct {
	evaluator
}

func newKeysWrapper(evaluator evaluator) evaluator { return keysWrapper{evaluator} }

func (e keysWrapper) Evaluate(val protoreflect.Value, failFast bool) error {
	err := e.evaluator.Evaluate(val, failFast)
	errors.PrependRulePath(err, mapKeysRulePath)
	return err
}

// valuesWrapper wraps the evaluation of nested map value rules.
type valuesWrapper struct {
	evaluator
}

func newValuesWrapper(evaluator evaluator) evaluator { return valuesWrapper{evaluator} }

func (e valuesWrapper) Evaluate(val protoreflect.Value, failFast bool) error {
	err := e.evaluator.Evaluate(val, failFast)
	errors.PrependRulePath(err, mapValuesRulePath)
	return err
}

var (
	_ evaluator = kvPairs{}
	_ evaluator = keysWrapper{}
	_ wrapper   = newKeysWrapper
	_ evaluator = valuesWrapper{}
	_ wrapper   = newValuesWrapper
)
