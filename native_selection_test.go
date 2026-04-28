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
	"reflect"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestNativeEvaluatorSelection verifies native rules.
func TestNativeEvaluatorSelection(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		fieldType  descriptorpb.FieldDescriptorProto_Type
		label      descriptorpb.FieldDescriptorProto_Label
		rules      *validate.FieldRules
		nativeType string // expected evaluator type name
	}{
		{
			name:      "int32_gt",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_INT32,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Int32: validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[int32]",
		},
		{
			name:      "int64_gte",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_INT64,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Int64: validate.Int64Rules_builder{Gte: proto.Int64(0)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[int64]",
		},
		{
			name:      "uint32_lt",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_UINT32,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Uint32: validate.UInt32Rules_builder{Lt: proto.Uint32(100)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[uint32]",
		},
		{
			name:      "uint64_const",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_UINT64,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Uint64: validate.UInt64Rules_builder{Const: proto.Uint64(42)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[uint64]",
		},
		{
			name:      "float_lt",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_FLOAT,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Float: validate.FloatRules_builder{Lt: proto.Float32(1.0)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[float32]",
		},
		{
			name:      "double_gte",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Double: validate.DoubleRules_builder{Gte: proto.Float64(0)}.Build(),
			}.Build(),
			nativeType: "nativeNumericCompare[float64]",
		},
		{
			name:      "string_min_len",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				String: validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build(),
			}.Build(),
			nativeType: "nativeStringEval",
		},
		{
			name:      "bool_const",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_BOOL,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Bool: validate.BoolRules_builder{Const: proto.Bool(true)}.Build(),
			}.Build(),
			nativeType: "nativeBoolEval",
		},
		{
			name:      "bytes_min_len",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_BYTES,
			label:     descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL,
			rules: validate.FieldRules_builder{
				Bytes: validate.BytesRules_builder{MinLen: proto.Uint64(1)}.Build(),
			}.Build(),
			nativeType: "nativeBytesEval",
		},
		{
			name:      "repeated_min_items",
			fieldType: descriptorpb.FieldDescriptorProto_TYPE_INT32,
			label:     descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
			rules: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{MinItems: proto.Uint64(1)}.Build(),
			}.Build(),
			nativeType: "nativeRepeatedEval",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMessageType(t, "test.sel", "Msg", &descriptorpb.FieldDescriptorProto{
				Name:    proto.String("value"),
				Number:  proto.Int32(1),
				Type:    tt.fieldType.Enum(),
				Label:   tt.label.Enum(),
				Options: fieldOpts(tt.rules),
			})
			v, err := New(WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor()))
			require.NoError(t, err)

			evals := findFieldEvaluators(t, v, msgType.Descriptor(), "value")
			require.NotEmpty(t, evals, "expected at least one evaluator for field")
			assertHasEvaluatorType(t, evals, tt.nativeType)
			assertNoCELPrograms(t, evals)
		})
	}
}

// TestNativeEvaluatorSelection_Enum tests enum separately since it needs
// an enum descriptor in the file.
func TestNativeEvaluatorSelection_Enum(t *testing.T) {
	t.Parallel()
	enumDesc := &descriptorpb.EnumDescriptorProto{
		Name: proto.String("TestEnum"),
		Value: []*descriptorpb.EnumValueDescriptorProto{
			{Name: proto.String("UNSPECIFIED"), Number: proto.Int32(0)},
			{Name: proto.String("ONE"), Number: proto.Int32(1)},
		},
	}
	msgType := newDynamicMessageTypeWithEnum(t, "test.sel", "EnumMsg", enumDesc, &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("value"),
		Number:   proto.Int32(1),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
		TypeName: proto.String(".test.sel.TestEnum"),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Options: fieldOpts(validate.FieldRules_builder{
			Enum: validate.EnumRules_builder{In: []int32{0, 1}}.Build(),
		}.Build()),
	})

	v, err := New(WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor()))
	require.NoError(t, err)

	evals := findFieldEvaluators(t, v, msgType.Descriptor(), "value")
	require.NotEmpty(t, evals)
	assertHasEvaluatorType(t, evals, "nativeEnumEval")
	assertNoCELPrograms(t, evals)
}

// TestNativeEvaluatorSelection_Map tests map separately since it needs
// a map field descriptor.
func TestNativeEvaluatorSelection_Map(t *testing.T) {
	t.Parallel()
	msgType := newDynamicMapMessageType(t, "test.sel", "MapMsg",
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		validate.FieldRules_builder{
			Map: validate.MapRules_builder{MinPairs: proto.Uint64(1)}.Build(),
		}.Build(),
	)

	v, err := New(WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor()))
	require.NoError(t, err)

	evals := findFieldEvaluators(t, v, msgType.Descriptor(), "entries")
	require.NotEmpty(t, evals)
	assertHasEvaluatorType(t, evals, "nativeMapEval")
	assertNoCELPrograms(t, evals)
}

// TestCELFallbackWithoutNativeFlag verifies that without native rules,
// the builder uses CEL programs.
func TestCELFallbackWithoutNativeFlag(t *testing.T) {
	t.Parallel()
	msgType := newDynamicMessageType(t, "test.sel", "CELMsg", &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("value"),
		Number: proto.Int32(1),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Options: fieldOpts(validate.FieldRules_builder{
			Int32: validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
		}.Build()),
	})

	v, err := New(WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor()), WithDisableNativeRules())
	require.NoError(t, err)

	evals := findFieldEvaluators(t, v, msgType.Descriptor(), "value")
	require.NotEmpty(t, evals)
	assertHasEvaluatorType(t, evals, "celPrograms")
}

// --- Helpers ---

// findFieldEvaluators extracts the Rules evaluators for a named field
// from a built validator's evaluator tree.
func findFieldEvaluators(
	t testing.TB,
	v Validator,
	desc protoreflect.MessageDescriptor,
	fieldName string,
) evaluators {
	t.Helper()
	val, ok := v.(*validator)
	require.True(t, ok, "expected *validator")

	msgEval := val.builder.Load(desc)
	msg, ok := msgEval.(*message)
	require.True(t, ok, "expected *message evaluator")

	for _, ne := range msg.nestedEvaluators {
		f, ok := ne.(field)
		if !ok {
			continue
		}
		if f.Value.Descriptor.Name() == protoreflect.Name(fieldName) {
			return f.Value.Rules
		}
	}
	t.Fatalf("field %q not found in evaluator tree", fieldName)
	return nil
}

// assertHasEvaluatorType checks that at least one evaluator in the slice
// has the expected type name.
func assertHasEvaluatorType(t testing.TB, evals evaluators, typeName string) {
	t.Helper()
	for _, e := range evals {
		if evaluatorTypeName(e) == typeName {
			return
		}
	}
	var actual []string
	for _, e := range evals {
		actual = append(actual, evaluatorTypeName(e))
	}
	t.Errorf("expected evaluator type %q, got %v", typeName, actual)
}

// assertNoCELPrograms checks that no evaluator in the slice is a celPrograms.
func assertNoCELPrograms(t testing.TB, evals evaluators) {
	t.Helper()
	for _, e := range evals {
		assert.NotEqual(t, "celPrograms", evaluatorTypeName(e),
			"expected no CEL programs when native rules are enabled")
	}
}

// evaluatorTypeName returns a string identifying the evaluator's concrete type
// using reflection, so it automatically handles new evaluator types.
func evaluatorTypeName(e evaluator) string {
	return reflect.TypeOf(e).Name()
}
