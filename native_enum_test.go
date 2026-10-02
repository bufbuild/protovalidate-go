// Copyright 2023-2026 Buf Technologies, Inc.
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
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	examplev1 "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNativeEnumConst(t *testing.T) {
	t.Parallel()
	eval := buildNativeEnum(t, validate.EnumRules_builder{Const: proto.Int32(1)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfEnum(1), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfEnum(2), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "enum.const", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must equal 1", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeEnumIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeEnum(t, validate.EnumRules_builder{In: []int32{1, 2}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfEnum(1), &validationConfig{}))
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfEnum(2), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfEnum(3), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "enum.in", valErr.Violations[0].Proto.GetRuleId())
}

func TestNativeEnumNotIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeEnum(t, validate.EnumRules_builder{NotIn: []int32{0}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfEnum(1), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfEnum(0), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "enum.not_in", valErr.Violations[0].Proto.GetRuleId())
}

func TestTryBuildNativeEnumRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.EnumRules
	}{
		{"nil_rules", nil},
		{"empty_rules", validate.EnumRules_builder{}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryBuildNativeEnumRules(base{}, tt.rules, nil))
		})
	}
}

func TestNativeEnumTautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeEnum(t, validate.EnumRules_builder{Const: proto.Int32(1)}.Build())
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}

func buildNativeEnum(t testing.TB, rules *validate.EnumRules) evaluator {
	t.Helper()
	fdesc := newEnumFieldDescriptor(t)
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryBuildNativeEnumRules(b, rules, fdesc.Enum().Values())
}

// newEnumFieldDescriptor creates a minimal enum field descriptor for testing.
func newEnumFieldDescriptor(t testing.TB) protoreflect.FieldDescriptor {
	t.Helper()
	fileProto := &descriptorpb.FileDescriptorProto{
		Name:    new("test_enum.proto"),
		Package: new("test"),
		EnumType: []*descriptorpb.EnumDescriptorProto{{
			Name: new("TestEnum"),
			Value: []*descriptorpb.EnumValueDescriptorProto{
				{Name: new("UNSPECIFIED"), Number: proto.Int32(0)},
				{Name: new("ONE"), Number: proto.Int32(1)},
				{Name: new("TWO"), Number: proto.Int32(2)},
			},
		}},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: new("Msg"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:     new("val"),
				Number:   proto.Int32(1),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
				TypeName: new(".test.TestEnum"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			}},
		}},
		Syntax: new("proto3"),
	}
	file, err := protodesc.NewFile(fileProto, nil)
	require.NoError(t, err)
	return file.Messages().Get(0).Fields().Get(0)
}

// newRegistryWithValidateProto creates a file registry containing the
// validate proto needed for field options.
func newRegistryWithValidateProto(t testing.TB) *protoregistry.Files {
	t.Helper()
	registry := new(protoregistry.Files)
	require.NoError(t, registry.RegisterFile(validate.File_buf_validate_validate_proto))
	require.NoError(t, registry.RegisterFile(wrapperspb.File_google_protobuf_wrappers_proto))
	return registry
}

func TestNativeEnumViolationOrder(t *testing.T) {
	t.Parallel()
	msg := examplev1.EnumRuleOrder_builder{Val: examplev1.TestEnum(99)}.Build()
	for _, disableNative := range []bool{false, true} {
		var opts []ValidatorOption
		if disableNative {
			opts = append(opts, WithDisableNativeRules())
		}
		val, err := New(opts...)
		require.NoError(t, err)
		assert.Equal(t, []string{"enum.defined_only", "enum.in", "enum.not_in"},
			violationRuleIDs(t, val.Validate(msg)), "disableNative=%v", disableNative)
		val, err = New(append(opts, WithFailFast())...)
		require.NoError(t, err)
		assert.Equal(t, []string{"enum.defined_only"},
			violationRuleIDs(t, val.Validate(msg)), "disableNative=%v", disableNative)
	}
}

func TestNativeEnumDefinedOnly(t *testing.T) {
	t.Parallel()
	fdesc := newEnumFieldDescriptor(t)
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	eval := tryBuildNativeEnumRules(b, validate.EnumRules_builder{DefinedOnly: new(true)}.Build(), fdesc.Enum().Values())
	require.NotNil(t, eval)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfEnum(2), &validationConfig{}))

	undefined := protoreflect.ValueOfEnum(99)
	nativeErr := eval.Evaluate(nil, undefined, &validationConfig{})
	definedErr := definedEnum{base: b, ValueDescriptors: fdesc.Enum().Values()}.Evaluate(nil, undefined, &validationConfig{})
	var nativeValErr, definedValErr *ValidationError
	require.ErrorAs(t, nativeErr, &nativeValErr)
	require.ErrorAs(t, definedErr, &definedValErr)
	assert.True(t, proto.Equal(definedValErr.ToProto(), nativeValErr.ToProto()))
}

func TestNativeEnumViolationOrderAllRules(t *testing.T) {
	t.Parallel()
	eval := buildNativeEnum(t, validate.EnumRules_builder{
		Const:       proto.Int32(1),
		DefinedOnly: new(true),
		In:          []int32{1},
		NotIn:       []int32{99},
	}.Build())
	require.NotNil(t, eval)
	undefined := protoreflect.ValueOfEnum(99)
	assert.Equal(t, []string{"enum.const", "enum.defined_only", "enum.in", "enum.not_in"},
		violationRuleIDs(t, eval.Evaluate(nil, undefined, &validationConfig{})))
	assert.Equal(t, []string{"enum.const"},
		violationRuleIDs(t, eval.Evaluate(nil, undefined, &validationConfig{failFast: true})))
}
