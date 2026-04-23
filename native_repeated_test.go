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
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestNativeRepeatedMinItems(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{MinItems: proto.Uint64(3)}.Build())
	require.NotNil(t, eval)

	// 3 items passes
	list := newInt32List(t, 1, 2, 3)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{}))

	// 2 items fails
	list = newInt32List(t, 1, 2)
	err := eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "repeated.min_items", valErr.Violations[0].Proto.GetRuleId())
}

func TestNativeRepeatedMaxItems(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{MaxItems: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	// 2 items passes
	list := newInt32List(t, 1, 2)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{}))

	// 3 items fails
	list = newInt32List(t, 1, 2, 3)
	err := eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "repeated.max_items", valErr.Violations[0].Proto.GetRuleId())
}

func TestNativeRepeatedUnique(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{Unique: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	// unique items passes
	list := newInt32List(t, 1, 2, 3)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{}))

	// duplicate fails
	list = newInt32List(t, 1, 2, 1)
	err := eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "repeated.unique", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "repeated value must contain unique items", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeRepeatedUnique_Empty(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{Unique: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	list := newInt32List(t)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{}))
}

func TestNativeRepeatedUnique_Single(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{Unique: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	list := newInt32List(t, 42)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfList(list), &validationConfig{}))
}

func TestTryNativeRepeatedRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.RepeatedRules
	}{
		{"nil_rules", nil},
		{"empty_rules", validate.RepeatedRules_builder{}.Build()},
		{"items_only", validate.RepeatedRules_builder{
			Items: validate.FieldRules_builder{
				Int32: validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
			}.Build(),
		}.Build()},
		{"unique_false", validate.RepeatedRules_builder{Unique: proto.Bool(false)}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryNativeRepeatedRules(base{}, tt.rules))
		})
	}
}

func TestNativeUniqueEnums(t *testing.T) {
	t.Parallel()
	// should fail
	{
		msg := examplev1.TestUnique_builder{
			Enums: []examplev1.TestEnum{
				examplev1.TestEnum_TEST_ENUM_VAL1,
				examplev1.TestEnum_TEST_ENUM_VAL1,
				examplev1.TestEnum_TEST_ENUM_VAL2,
				examplev1.TestEnum_TEST_ENUM_VAL3,
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.Error(t, validator.Validate(msg))
	}

	// should pass
	{
		msg := examplev1.TestUnique_builder{
			Enums: []examplev1.TestEnum{
				examplev1.TestEnum_TEST_ENUM_VAL1,
				examplev1.TestEnum_TEST_ENUM_VAL2,
				examplev1.TestEnum_TEST_ENUM_VAL3,
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.NoError(t, validator.Validate(msg))
	}
}

func TestNativeUniqueStrings(t *testing.T) {
	t.Parallel()
	// should fail
	{
		msg := examplev1.TestUnique_builder{
			Strings: []string{
				"a",
				"b",
				"a",
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.Error(t, validator.Validate(msg))
	}

	// should pass
	{
		msg := examplev1.TestUnique_builder{
			Strings: []string{
				"a",
				"b",
				"c",
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.NoError(t, validator.Validate(msg))
	}
}

func TestNativeUniqueBytes(t *testing.T) {
	t.Parallel()
	// should fail
	{
		msg := examplev1.TestUnique_builder{
			Bytes: [][]byte{
				[]byte("a"),
				[]byte("b"),
				[]byte("a"),
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.Error(t, validator.Validate(msg))
	}

	// should pass
	{
		msg := examplev1.TestUnique_builder{
			Bytes: [][]byte{
				[]byte("a"),
				[]byte("b"),
				[]byte("c"),
			},
		}.Build()
		validator, err := New(WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor()))
		require.NoError(t, err)
		require.NoError(t, validator.Validate(msg))
	}
}

func TestNativeRepeatedTautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeRepeated(t, validate.RepeatedRules_builder{MinItems: proto.Uint64(1)}.Build())
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}

// --- Helpers ---

func buildNativeRepeated(t testing.TB, rules *validate.RepeatedRules) evaluator {
	t.Helper()
	fdesc := newFieldDescriptor(t, descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum())
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryNativeRepeatedRules(b, rules)
}

// newInt32List creates a protoreflect.List containing the given int32 values.
func newInt32List(t testing.TB, vals ...int32) protoreflect.List {
	t.Helper()

	msgType := newDynamicMessageType(t, "test.native", "TestMessage", &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("value"),
		Number: proto.Int32(1),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
	})

	template := dynamicpb.NewMessage(msgType.Descriptor())
	list := template.NewField(msgType.Descriptor().Fields().ByName("value"))
	for _, v := range vals {
		list.List().Append(protoreflect.ValueOfInt32(v))
	}
	return list.List()
}

type testList[T any] struct {
	protoreflect.List
	vals []T
}

// These are the only methods we care about.
func (l testList[T]) Len() int {
	return len(l.vals)
}

func (l testList[T]) Get(i int) protoreflect.Value {
	return protoreflect.ValueOf(l.vals[i])
}

func Test_isUniqueList(t *testing.T) {
	t.Parallel()
	// try lengths of 0, 1, 10, 16, and 20
	// don't care about type, so just use int
	// need both positive and negative tests
	data := []struct {
		name   string
		vals   []int64
		result bool
	}{
		{"empty", []int64{}, true},
		{"single", []int64{1}, true},
		{"10", []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{"16", []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, true},
		{"20", []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, true},
		{"10_dup", []int64{1, 2, 3, 4, 10, 6, 7, 8, 9, 10}, false},
		{"16", []int64{1, 2, 3, 4, 5, 6, 16, 8, 9, 10, 11, 12, 13, 14, 15, 16}, false},
		{"20", []int64{1, 2, 3, 4, 5, 6, 7, 20, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, false},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			l := testList[int64]{
				vals: d.vals,
			}
			result := isUniqueList[int64](l)
			assert.Equal(t, d.result, result)
		})
	}
}

func Test_isUniqueBytes(t *testing.T) {
	t.Parallel()
	// try lengths of 0, 1, 3, 16, and 20
	// need both positive and negative tests
	data := []struct {
		name   string
		vals   [][]byte
		result bool
	}{
		{"empty", [][]byte{}, true},
		{"single", [][]byte{[]byte("1")}, true},
		{"3", [][]byte{[]byte("1"), []byte("2"), []byte("3")}, true},
		{"16", [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("7"), []byte("8"), []byte("9"), []byte("10"), []byte("11"), []byte("12"), []byte("13"), []byte("14"), []byte("15"), []byte("16")}, true},
		{"20_dup", [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("7"), []byte("8"), []byte("9"), []byte("10"), []byte("11"), []byte("12"), []byte("13"), []byte("14"), []byte("15"), []byte("16"), []byte("17"), []byte("18"), []byte("19"), []byte("20")}, true},
		{"3_dup", [][]byte{[]byte("1"), []byte("2"), []byte("1")}, false},
		{"16_dup", [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("16"), []byte("8"), []byte("9"), []byte("10"), []byte("11"), []byte("12"), []byte("13"), []byte("14"), []byte("15"), []byte("16")}, false},
		{"20_dup", [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6"), []byte("20"), []byte("8"), []byte("9"), []byte("10"), []byte("11"), []byte("12"), []byte("13"), []byte("14"), []byte("15"), []byte("16"), []byte("17"), []byte("18"), []byte("19"), []byte("20")}, false},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			l := testList[[]byte]{
				vals: d.vals,
			}
			result := isUniqueBytes(l)
			assert.Equal(t, d.result, result)
		})
	}
}
