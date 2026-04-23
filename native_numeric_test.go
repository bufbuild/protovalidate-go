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
	"math"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// numericTestCase defines a single constraint test case.
type numericTestCase[T numericValue] struct {
	name   string
	pass   []T
	fail   []T
	ruleID string
}

// runNumericCases runs a set of test cases against the given evaluator.
func runNumericCases[T numericValue](
	t *testing.T,
	eval evaluator,
	toValue func(T) protoreflect.Value,
	cases []numericTestCase[T],
) {
	t.Helper()
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for _, v := range tt.pass {
				err := eval.Evaluate(nil, toValue(v), &validationConfig{})
				require.NoError(t, err, "expected %v to pass", v)
			}
			for _, v := range tt.fail {
				err := eval.Evaluate(nil, toValue(v), &validationConfig{})
				require.Error(t, err, "expected %v to fail", v)
				var valErr *ValidationError
				require.ErrorAs(t, err, &valErr)
				require.Len(t, valErr.Violations, 1)
				assert.Equal(t, tt.ruleID, valErr.Violations[0].Proto.GetRuleId())
			}
		})
	}
}

// --- Int32 ---

func TestNativeInt32Compare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.Int32Rules
		cases []numericTestCase[int32]
	}{
		{
			name:  "gt_only",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{6}, []int32{4, 5}, "int32.gt"},
			},
		},
		{
			name:  "gte_only",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{5, 6}, []int32{4}, "int32.gte"},
			},
		},
		{
			name:  "lt_only",
			rules: validate.Int32Rules_builder{Lt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4}, []int32{5, 6}, "int32.lt"},
			},
		},
		{
			name:  "lte_only",
			rules: validate.Int32Rules_builder{Lte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4, 5}, []int32{6}, "int32.lte"},
			},
		},
		{
			name:  "gt_lt",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(0), Lt: proto.Int32(10)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{1, 9}, []int32{-1, 0, 10, 11}, "int32.gt_lt"},
			},
		},
		{
			name:  "gt_lt_exclusive",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(10), Lt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4, 11}, []int32{5, 6, 9, 10}, "int32.gt_lt_exclusive"},
			},
		},
		{
			name:  "gt_lte",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(0), Lte: proto.Int32(10)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{1, 9, 10}, []int32{-1, 0, 11}, "int32.gt_lte"},
			},
		},
		{
			name:  "gt_lte_exclusive",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(10), Lte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4, 5, 11}, []int32{6, 9, 10}, "int32.gt_lte_exclusive"},
			},
		},
		{
			name:  "gte_lt",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(0), Lt: proto.Int32(10)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{0, 1, 9}, []int32{-1, 10, 11}, "int32.gte_lt"},
			},
		},
		{
			name:  "gte_lt_exclusive",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(10), Lt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4, 10, 11}, []int32{5, 6, 9}, "int32.gte_lt_exclusive"},
			},
		},
		{
			name:  "gte_lte",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(0), Lte: proto.Int32(10)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{0, 1, 9, 10}, []int32{-1, 11}, "int32.gte_lte"},
			},
		},
		{
			name:  "gte_lte_exclusive",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(10), Lte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{4, 5, 10, 11}, []int32{6, 9}, "int32.gte_lte_exclusive"},
			},
		},
		{
			name:  "gte_eq_lt",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(5), Lt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"all_fail", nil, []int32{4, 5, 6}, "int32.gte_lt"},
			},
		},
		{
			name:  "gt_eq_lt",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(5), Lt: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"all_fail", nil, []int32{4, 5, 6}, "int32.gt_lt"},
			},
		},
		{
			name:  "gt_eq_lte",
			rules: validate.Int32Rules_builder{Gt: proto.Int32(5), Lte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"all_fail", nil, []int32{4, 5, 6}, "int32.gt_lte"},
			},
		},
		{
			name:  "gte_eq_lte",
			rules: validate.Int32Rules_builder{Gte: proto.Int32(5), Lte: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{5}, []int32{4, 6}, "int32.gte_lte"},
			},
		},
		{
			name:  "const",
			rules: validate.Int32Rules_builder{Const: proto.Int32(5)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{5}, []int32{4, 6}, "int32.const"},
			},
		},
		{
			name:  "const_with_gte",
			rules: validate.Int32Rules_builder{Const: proto.Int32(5), Gte: proto.Int32(4)}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{5}, []int32{7, 4, 6}, "int32.const"},
			},
		},
		{
			name:  "in",
			rules: validate.Int32Rules_builder{In: []int32{1, 3, 5}}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{1, 3, 5}, []int32{0, 2, 4, 6}, "int32.in"},
			},
		},
		{
			name:  "not_in",
			rules: validate.Int32Rules_builder{NotIn: []int32{1, 3, 5}}.Build(),
			cases: []numericTestCase[int32]{
				{"pass_fail", []int32{0, 2, 4, 6}, []int32{1, 3, 5}, "int32.not_in"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			eval := buildNativeNumeric(t, tt.rules, int32Config, descriptorpb.FieldDescriptorProto_TYPE_INT32)
			require.NotNil(t, eval)
			runNumericCases(t, eval, protoreflect.ValueOfInt32, tt.cases)
		})
	}
}

// --- Int64 ---

func TestNativeInt64Compare(t *testing.T) {
	t.Parallel()

	rules := validate.Int64Rules_builder{Gt: proto.Int64(0), Lt: proto.Int64(100)}.Build()
	eval := buildNativeNumeric(t, rules, int64Config, descriptorpb.FieldDescriptorProto_TYPE_INT64)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfInt64, []numericTestCase[int64]{
		{"pass_fail", []int64{1, 50, 99}, []int64{-1, 0, 100, 101}, "int64.gt_lt"},
	})
}

// --- Uint32 ---

func TestNativeUint32Compare(t *testing.T) {
	t.Parallel()

	rules := validate.UInt32Rules_builder{Gte: proto.Uint32(5), Lte: proto.Uint32(10)}.Build()
	eval := buildNativeNumeric(t, rules, uint32Config, descriptorpb.FieldDescriptorProto_TYPE_UINT32)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfUint32, []numericTestCase[uint32]{
		{"pass_fail", []uint32{5, 7, 10}, []uint32{4, 11}, "uint32.gte_lte"},
	})
}

// --- Uint64 ---

func TestNativeUint64Compare(t *testing.T) {
	t.Parallel()

	rules := validate.UInt64Rules_builder{Gt: proto.Uint64(0)}.Build()
	eval := buildNativeNumeric(t, rules, uint64Config, descriptorpb.FieldDescriptorProto_TYPE_UINT64)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfUint64, []numericTestCase[uint64]{
		{"pass_fail", []uint64{1, 100}, []uint64{0}, "uint64.gt"},
	})
}

// --- Sint32 ---

func TestNativeSint32Compare(t *testing.T) {
	t.Parallel()

	rules := validate.SInt32Rules_builder{Gt: proto.Int32(-10), Lt: proto.Int32(10)}.Build()
	eval := buildNativeNumeric(t, rules, sint32Config, descriptorpb.FieldDescriptorProto_TYPE_SINT32)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfInt32, []numericTestCase[int32]{
		{"pass_fail", []int32{-9, 0, 9}, []int32{-10, 10}, "sint32.gt_lt"},
	})
}

// --- Sint64 ---

func TestNativeSint64Compare(t *testing.T) {
	t.Parallel()

	rules := validate.SInt64Rules_builder{Gte: proto.Int64(0)}.Build()
	eval := buildNativeNumeric(t, rules, sint64Config, descriptorpb.FieldDescriptorProto_TYPE_SINT64)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfInt64, []numericTestCase[int64]{
		{"pass_fail", []int64{0, 1, 100}, []int64{-1}, "sint64.gte"},
	})
}

// --- Fixed32 ---

func TestNativeFixed32Compare(t *testing.T) {
	t.Parallel()

	rules := validate.Fixed32Rules_builder{Lt: proto.Uint32(100)}.Build()
	eval := buildNativeNumeric(t, rules, fixed32Config, descriptorpb.FieldDescriptorProto_TYPE_FIXED32)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfUint32, []numericTestCase[uint32]{
		{"pass_fail", []uint32{0, 99}, []uint32{100, 101}, "fixed32.lt"},
	})
}

// --- Fixed64 ---

func TestNativeFixed64Compare(t *testing.T) {
	t.Parallel()

	rules := validate.Fixed64Rules_builder{Lte: proto.Uint64(50)}.Build()
	eval := buildNativeNumeric(t, rules, fixed64Config, descriptorpb.FieldDescriptorProto_TYPE_FIXED64)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfUint64, []numericTestCase[uint64]{
		{"pass_fail", []uint64{0, 50}, []uint64{51}, "fixed64.lte"},
	})
}

// --- Sfixed32 ---

func TestNativeSfixed32Compare(t *testing.T) {
	t.Parallel()

	rules := validate.SFixed32Rules_builder{Const: proto.Int32(42)}.Build()
	eval := buildNativeNumeric(t, rules, sfixed32Config, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfInt32, []numericTestCase[int32]{
		{"pass_fail", []int32{42}, []int32{0, 41, 43}, "sfixed32.const"},
	})
}

// --- Sfixed64 ---

func TestNativeSfixed64Compare(t *testing.T) {
	t.Parallel()

	rules := validate.SFixed64Rules_builder{In: []int64{1, 2, 3}}.Build()
	eval := buildNativeNumeric(t, rules, sfixed64Config, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfInt64, []numericTestCase[int64]{
		{"pass_fail", []int64{1, 2, 3}, []int64{0, 4}, "sfixed64.in"},
	})
}

// --- Float ---

func TestNativeFloatCompare(t *testing.T) {
	t.Parallel()

	rules := validate.FloatRules_builder{Gt: proto.Float32(0), Lt: proto.Float32(10)}.Build()
	eval := buildNativeNumeric(t, rules, floatConfig, descriptorpb.FieldDescriptorProto_TYPE_FLOAT)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfFloat32, []numericTestCase[float32]{
		{"pass_fail", []float32{0.1, 5, 9.9}, []float32{-1, 0, 10, 11}, "float.gt_lt"},
	})
}

func TestNativeFloatFinite(t *testing.T) {
	t.Parallel()

	rules := validate.FloatRules_builder{Finite: proto.Bool(true)}.Build()
	eval := buildNativeNumeric(t, rules, floatConfig, descriptorpb.FieldDescriptorProto_TYPE_FLOAT)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfFloat32, []numericTestCase[float32]{
		{"pass_fail", []float32{0.1, 5, 9.9}, []float32{float32(math.NaN()), float32(math.Inf(1)), float32(math.Inf(-1))}, "float.finite"},
	})
}

func TestNativeDoubleFinite(t *testing.T) {
	t.Parallel()

	rules := validate.DoubleRules_builder{Finite: proto.Bool(true)}.Build()
	eval := buildNativeNumeric(t, rules, doubleConfig, descriptorpb.FieldDescriptorProto_TYPE_DOUBLE)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfFloat64, []numericTestCase[float64]{
		{"pass_fail", []float64{0.1, 5, 9.9}, []float64{math.NaN(), math.Inf(1), math.Inf(-1)}, "double.finite"},
	})
}

func TestNativeFloat_NaN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.FloatRules
	}{
		{"gt", validate.FloatRules_builder{Gt: proto.Float32(0)}.Build()},
		{"lt", validate.FloatRules_builder{Lt: proto.Float32(10)}.Build()},
		{"gte_lte", validate.FloatRules_builder{Gte: proto.Float32(0), Lte: proto.Float32(10)}.Build()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			eval := buildNativeNumeric(t, tt.rules, floatConfig, descriptorpb.FieldDescriptorProto_TYPE_FLOAT)
			require.NotNil(t, eval)
			nan := float32(math.NaN())
			err := eval.Evaluate(nil, protoreflect.ValueOfFloat32(nan), &validationConfig{})
			require.Error(t, err, "NaN should fail range checks")
		})
	}
}

func TestNativeFloat_FiniteDoesNotBailToCEL(t *testing.T) {
	t.Parallel()
	rules := validate.FloatRules_builder{Finite: proto.Bool(true), Gt: proto.Float32(0)}.Build()
	eval := tryBuildNativeFloatRules(base{}, rules)
	if eval == nil {
		t.Error("finite rules should have native implementation")
	}
}

// --- Double ---

func TestNativeDoubleCompare(t *testing.T) {
	t.Parallel()

	rules := validate.DoubleRules_builder{Gte: proto.Float64(-1.5), Lte: proto.Float64(1.5)}.Build()
	eval := buildNativeNumeric(t, rules, doubleConfig, descriptorpb.FieldDescriptorProto_TYPE_DOUBLE)
	require.NotNil(t, eval)

	runNumericCases(t, eval, protoreflect.ValueOfFloat64, []numericTestCase[float64]{
		{"pass_fail", []float64{-1.5, 0, 1.5}, []float64{-2, 2}, "double.gte_lte"},
	})
}

func TestNativeDouble_NaN(t *testing.T) {
	t.Parallel()

	rules := validate.DoubleRules_builder{Gt: proto.Float64(0)}.Build()
	eval := buildNativeNumeric(t, rules, doubleConfig, descriptorpb.FieldDescriptorProto_TYPE_DOUBLE)
	require.NotNil(t, eval)

	err := eval.Evaluate(nil, protoreflect.ValueOfFloat64(math.NaN()), &validationConfig{})
	require.Error(t, err, "NaN should fail range checks")
}

// --- Common tests ---

func TestNativeNumericCompare_FieldValue(t *testing.T) {
	t.Parallel()
	eval := buildNativeNumeric(t,
		validate.Int32Rules_builder{Gt: proto.Int32(5)}.Build(),
		int32Config,
		descriptorpb.FieldDescriptorProto_TYPE_INT32,
	)
	require.NotNil(t, eval)
	val := protoreflect.ValueOfInt32(3)
	err := eval.Evaluate(nil, val, &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, int64(3), valErr.Violations[0].FieldValue.Int())
}

func TestNativeNumericCompare_Tautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeNumeric(t,
		validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
		int32Config,
		descriptorpb.FieldDescriptorProto_TYPE_INT32,
	)
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}

func TestTryBuildNativeNumericRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fn   func() evaluator
	}{
		{"nil_int32", func() evaluator { return tryBuildNativeInt32Rules(base{}, nil) }},
		{"nil_int64", func() evaluator { return tryBuildNativeInt64Rules(base{}, nil) }},
		{"nil_uint32", func() evaluator { return tryBuildNativeUint32Rules(base{}, nil) }},
		{"nil_uint64", func() evaluator { return tryBuildNativeUint64Rules(base{}, nil) }},
		{"nil_float", func() evaluator { return tryBuildNativeFloatRules(base{}, nil) }},
		{"nil_double", func() evaluator { return tryBuildNativeDoubleRules(base{}, nil) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tt.fn())
		})
	}
}

// --- Helpers ---

// buildNativeNumeric constructs a nativeNumericCompare evaluator for testing.
func buildNativeNumeric[T numericValue, R numericRules[T]](
	t testing.TB,
	rules R,
	config numericTypeConfig[T],
	fieldType descriptorpb.FieldDescriptorProto_Type,
) evaluator {
	t.Helper()
	fdesc := newFieldDescriptor(t, fieldType,
		descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum())
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryBuildNativeNumericRules(b, rules, config)
}

// newFieldDescriptor creates a minimal field descriptor for a given proto type.
func newFieldDescriptor(t testing.TB, fieldType descriptorpb.FieldDescriptorProto_Type, label *descriptorpb.FieldDescriptorProto_Label) protoreflect.FieldDescriptor {
	t.Helper()
	fileProto := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("Msg"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("val"),
						Number: proto.Int32(1),
						Type:   fieldType.Enum(),
						Label:  label,
					},
				},
			},
		},
		Syntax: proto.String("proto3"),
	}
	file, err := protodesc.NewFile(fileProto, nil)
	require.NoError(t, err)
	return file.Messages().Get(0).Fields().Get(0)
}
