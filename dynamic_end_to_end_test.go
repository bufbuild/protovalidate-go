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
	"fmt"
	"math"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestDynamicRulesEndToEnd(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		typ  *descriptorpb.FieldDescriptorProto_Type
		rule *validate.FieldRules
		info dynamicMessageTesterInfo
	}{
		{
			name: "int32_gt",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Int32: validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfInt32(1),
				badValue:          protoreflect.ValueOfInt32(0),
				failedRuleID:      "int32.gt",
				failedRuleMessage: "must be greater than 0",
			},
		},
		{
			name: "uint64_gte",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum(),
			rule: validate.FieldRules_builder{
				Uint64: validate.UInt64Rules_builder{Gte: proto.Uint64(10)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfUint64(10),
				badValue:          protoreflect.ValueOfUint64(9),
				failedRuleID:      "uint64.gte",
				failedRuleMessage: "must be greater than or equal to 10",
			},
		},
		{
			name: "double_lt",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
			rule: validate.FieldRules_builder{
				Double: validate.DoubleRules_builder{Lt: proto.Float64(100)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfFloat64(50),
				badValue:          protoreflect.ValueOfFloat64(100),
				failedRuleID:      "double.lt",
				failedRuleMessage: "must be less than 100",
			},
		},
		{
			name: "double_finite",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
			rule: validate.FieldRules_builder{
				Double: validate.DoubleRules_builder{Finite: proto.Bool(true)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfFloat64(50),
				badValue:          protoreflect.ValueOfFloat64(math.Inf(1)),
				failedRuleID:      "double.finite",
				failedRuleMessage: "must be finite",
			},
		},
		{
			name: "min_len",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			rule: validate.FieldRules_builder{
				String: validate.StringRules_builder{MinLen: proto.Uint64(3)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfString("abc"),
				badValue:          protoreflect.ValueOfString("ab"),
				failedRuleID:      "string.min_len",
				failedRuleMessage: "must be at least 3 characters",
			},
		},
		{
			name: "prefix",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			rule: validate.FieldRules_builder{
				String: validate.StringRules_builder{Prefix: proto.String("hello")}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfString("hello world"),
				badValue:          protoreflect.ValueOfString("world"),
				failedRuleID:      "string.prefix",
				failedRuleMessage: "does not have prefix `hello`",
			},
		},
		{
			name: "bool_const",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum(),
			rule: validate.FieldRules_builder{
				Bool: validate.BoolRules_builder{Const: proto.Bool(true)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfBool(true),
				badValue:          protoreflect.ValueOfBool(false),
				failedRuleID:      "bool.const",
				failedRuleMessage: "must equal true",
			},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMessageType(t, "test.native", "TestMessage", &descriptorpb.FieldDescriptorProto{
				Name:    proto.String("value"),
				Number:  proto.Int32(1),
				Type:    d.typ,
				Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				Options: fieldOpts(d.rule),
			})
			d.info.msgType = msgType
			// first with CEL rules
			dynamicMessageTester(t, d.info, "value", false)
			// now with native rules to validate they produce identical results
			dynamicMessageTester(t, d.info, "value", true)
		})
	}
}

func TestDynamicRepeatedRulesEndToEnd(t *testing.T) {
	t.Parallel()
	data := []struct {
		name      string
		typ       *descriptorpb.FieldDescriptorProto_Type
		rule      *validate.FieldRules
		info      dynamicMessageTesterInfo
		goodValue []int32
		badValue  []int32
	}{
		{
			name: "repeated_min_items",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{MinItems: proto.Uint64(2)}.Build(),
			}.Build(),
			goodValue: []int32{1, 2, 3},
			badValue:  []int32{1},
			info: dynamicMessageTesterInfo{
				failedRuleID:      "repeated.min_items",
				failedRuleMessage: "must contain at least 2 item(s)",
			},
		},
		{
			name: "repeated_max_items",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{MaxItems: proto.Uint64(2)}.Build(),
			}.Build(),
			goodValue: []int32{1},
			badValue:  []int32{1, 2, 3},
			info: dynamicMessageTesterInfo{
				failedRuleID:      "repeated.max_items",
				failedRuleMessage: "must contain no more than 2 item(s)",
			},
		},
		{
			name: "repeated_unique",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{Unique: ptr(true)}.Build(),
			}.Build(),
			goodValue: []int32{1, 2, 3},
			badValue:  []int32{1, 2, 1},
			info: dynamicMessageTesterInfo{
				failedRuleID:      "repeated.unique",
				failedRuleMessage: "repeated value must contain unique items",
			},
		},
		{
			name: "repeated_unique_items",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{Unique: proto.Bool(true), Items: validate.FieldRules_builder{Int32: validate.Int32Rules_builder{Gte: ptr(int32(2))}.Build()}.Build()}.Build(),
			}.Build(),
			goodValue: []int32{2, 4, 6},
			badValue:  []int32{2, 1, 3},
			info: dynamicMessageTesterInfo{
				failedRuleID:      "int32.gte",
				failedRuleMessage: "must be greater than or equal to 2",
			},
		},
		{
			name: "repeated_unique_max_items",
			typ:  descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
			rule: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{Unique: proto.Bool(true), MaxItems: proto.Uint64(3), Items: validate.FieldRules_builder{Int32: validate.Int32Rules_builder{Gte: ptr(int32(2))}.Build()}.Build()}.Build(),
			}.Build(),
			goodValue: []int32{2, 4, 6},
			badValue:  []int32{2, 6, 3, 5},
			info: dynamicMessageTesterInfo{
				failedRuleID:      "repeated.max_items",
				failedRuleMessage: "must contain no more than 3 item(s)",
			},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMessageType(t, "test.native", "TestMessage", &descriptorpb.FieldDescriptorProto{
				Name:    proto.String("value"),
				Number:  proto.Int32(1),
				Type:    d.typ,
				Label:   descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
				Options: fieldOpts(d.rule),
			})
			d.info.msgType = msgType
			// make lists
			template := dynamicpb.NewMessage(msgType.Descriptor())
			goodList := template.NewField(msgType.Descriptor().Fields().ByName("value"))
			for _, v := range d.goodValue {
				goodList.List().Append(protoreflect.ValueOfInt32(v))
			}
			d.info.goodValue = goodList
			badList := template.NewField(msgType.Descriptor().Fields().ByName("value"))
			for _, v := range d.badValue {
				badList.List().Append(protoreflect.ValueOfInt32(v))
			}
			d.info.badValue = badList

			// first with CEL rules
			dynamicMessageTester(t, d.info, "value", false)
			// now with native rules to validate they produce identical results
			dynamicMessageTester(t, d.info, "value", true)
		})
	}
}

func TestNativeEnum_EndToEnd(t *testing.T) {
	t.Parallel()
	// Build a proto with an enum field and const rule.
	enumDesc := &descriptorpb.EnumDescriptorProto{
		Name: proto.String("TestEnum"),
		Value: []*descriptorpb.EnumValueDescriptorProto{
			{Name: proto.String("UNSPECIFIED"), Number: proto.Int32(0)},
			{Name: proto.String("VALUE_ONE"), Number: proto.Int32(1)},
			{Name: proto.String("VALUE_TWO"), Number: proto.Int32(2)},
		},
	}
	data := []struct {
		name string
		rule *validate.FieldRules
		info dynamicMessageTesterInfo
	}{
		{
			name: "enum_const",
			rule: validate.FieldRules_builder{
				Enum: validate.EnumRules_builder{Const: proto.Int32(1)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfEnum(1),
				badValue:          protoreflect.ValueOfEnum(2),
				failedRuleID:      "enum.const",
				failedRuleMessage: "must equal 1",
			},
		},
		{
			name: "enum_in",
			rule: validate.FieldRules_builder{
				Enum: validate.EnumRules_builder{In: []int32{1, 2}}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				goodValue:         protoreflect.ValueOfEnum(1),
				badValue:          protoreflect.ValueOfEnum(3),
				failedRuleID:      "enum.in",
				failedRuleMessage: "must be in list [1, 2]",
			},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMessageTypeWithEnum(t, "test.native", "EnumMsg", enumDesc, &descriptorpb.FieldDescriptorProto{
				Name:     proto.String("value"),
				Number:   proto.Int32(1),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
				TypeName: proto.String(".test.native.TestEnum"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				Options:  fieldOpts(d.rule),
			})
			d.info.msgType = msgType
			// first with CEL rules
			dynamicMessageTester(t, d.info, "value", false)
			// now with native rules to validate they produce identical results
			dynamicMessageTester(t, d.info, "value", true)
		})
	}
}

func TestNativeMap_EndToEnd(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		rule    *validate.FieldRules
		info    dynamicMessageTesterInfo
		goodMap map[string]string
		badMap  map[string]string
	}{
		{
			name: "map_min_pairs",
			rule: validate.FieldRules_builder{
				Map: validate.MapRules_builder{MinPairs: proto.Uint64(2)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				failedRuleID:      "map.min_pairs",
				failedRuleMessage: "map must be at least 2 entries",
			},
			goodMap: map[string]string{"a": "1", "b": "2"},
			badMap:  map[string]string{"a": "1"},
		},
		{
			name: "map_max_pairs",
			rule: validate.FieldRules_builder{
				Map: validate.MapRules_builder{MaxPairs: proto.Uint64(2)}.Build(),
			}.Build(),
			info: dynamicMessageTesterInfo{
				failedRuleID:      "map.max_pairs",
				failedRuleMessage: "map must be at most 2 entries",
			},
			goodMap: map[string]string{"a": "1", "b": "2"},
			badMap:  map[string]string{"a": "1", "b": "2", "c": "3"},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMapMessageType(t, "test.native", "MapMsg",
				descriptorpb.FieldDescriptorProto_TYPE_STRING,
				descriptorpb.FieldDescriptorProto_TYPE_STRING,
				d.rule,
			)

			d.info.msgType = msgType

			template := dynamicpb.NewMessage(msgType.Descriptor())
			fd := msgType.Descriptor().Fields().ByName("entries")
			maker := func(m map[string]string) protoreflect.Value {
				mapField := template.NewField(fd)
				for k, v := range m {
					mapField.Map().Set(
						protoreflect.ValueOfString(k).MapKey(),
						protoreflect.ValueOfString(v),
					)
				}
				return mapField
			}
			// make a good map
			d.info.goodValue = maker(d.goodMap)
			// make a bad map
			d.info.badValue = maker(d.badMap)

			// first with CEL rules
			dynamicMessageTester(t, d.info, "entries", true)
			// now with native rules to validate they produce identical results
			dynamicMessageTester(t, d.info, "entries", false)
		})
	}
}

type dynamicMessageTesterInfo struct {
	msgType           protoreflect.MessageType
	goodValue         protoreflect.Value
	badValue          protoreflect.Value
	failedRuleID      string
	failedRuleMessage string
}

func dynamicMessageTester(t *testing.T, info dynamicMessageTesterInfo, fieldName protoreflect.Name, disableNativeRules bool) {
	t.Helper()
	options := []ValidatorOption{WithDisableLazy(), WithMessageDescriptors(info.msgType.Descriptor())}
	if disableNativeRules {
		options = append(options, WithDisableNativeRules())
	}
	validator, err := New(options...)
	require.NoError(t, err)

	passing := dynamicpb.NewMessage(info.msgType.Descriptor())
	passing.Set(info.msgType.Descriptor().Fields().ByName(fieldName), info.goodValue)
	require.NoError(t, validator.Validate(passing))

	failing := dynamicpb.NewMessage(info.msgType.Descriptor())
	failing.Set(info.msgType.Descriptor().Fields().ByName(fieldName), info.badValue)
	err = validator.Validate(failing)
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, info.failedRuleID, valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, info.failedRuleMessage, valErr.Violations[0].Proto.GetMessage())
}

// TestEnumCombinedRules validates that enum fields work correctly when
// multiple rules are active: native const/in/not_in (via processStandardRules)
// combined with native defined_only (via processEnumRules).
func TestEnumCombinedRules(t *testing.T) {
	t.Parallel()
	enumDesc := &descriptorpb.EnumDescriptorProto{
		Name: proto.String("Status"),
		Value: []*descriptorpb.EnumValueDescriptorProto{
			{Name: proto.String("STATUS_UNSPECIFIED"), Number: proto.Int32(0)},
			{Name: proto.String("STATUS_ACTIVE"), Number: proto.Int32(1)},
			{Name: proto.String("STATUS_INACTIVE"), Number: proto.Int32(2)},
		},
	}

	tests := []struct {
		name           string
		rules          *validate.EnumRules
		value          protoreflect.EnumNumber
		wantErr        bool
		violationCount int
		ruleIDs        []string // expected rule IDs in order
	}{
		// --- defined_only + in ---
		{
			name:    "defined_only+in/pass: defined and in list",
			rules:   validate.EnumRules_builder{DefinedOnly: proto.Bool(true), In: []int32{1, 2}}.Build(),
			value:   1, // STATUS_ACTIVE: defined, in [1,2]
			wantErr: false,
		},
		{
			name:           "defined_only+in/fail_in_only: defined but not in list",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), In: []int32{1, 2}}.Build(),
			value:          0, // STATUS_UNSPECIFIED: defined, not in [1,2]
			wantErr:        true,
			violationCount: 1,
			ruleIDs:        []string{"enum.in"},
		},
		{
			name:           "defined_only+in/fail_both: undefined and not in list",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), In: []int32{1, 2}}.Build(),
			value:          99, // undefined, not in [1,2]
			wantErr:        true,
			violationCount: 2,
			ruleIDs:        []string{"enum.in", "enum.defined_only"},
		},

		// --- defined_only + const ---
		{
			name:    "defined_only+const/pass: defined and equals const",
			rules:   validate.EnumRules_builder{DefinedOnly: proto.Bool(true), Const: proto.Int32(1)}.Build(),
			value:   1, // STATUS_ACTIVE: defined, equals 1
			wantErr: false,
		},
		{
			name:           "defined_only+const/fail_const_only: defined but wrong value",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), Const: proto.Int32(1)}.Build(),
			value:          2, // STATUS_INACTIVE: defined, but not 1
			wantErr:        true,
			violationCount: 1,
			ruleIDs:        []string{"enum.const"},
		},
		{
			name:           "defined_only+const/fail_both: undefined and wrong value",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), Const: proto.Int32(1)}.Build(),
			value:          99, // undefined, not 1
			wantErr:        true,
			violationCount: 2,
			ruleIDs:        []string{"enum.const", "enum.defined_only"},
		},

		// --- defined_only + not_in ---
		{
			name:    "defined_only+not_in/pass: defined and not in exclusion list",
			rules:   validate.EnumRules_builder{DefinedOnly: proto.Bool(true), NotIn: []int32{0}}.Build(),
			value:   1, // STATUS_ACTIVE: defined, not in [0]
			wantErr: false,
		},
		{
			name:           "defined_only+not_in/fail_not_in_only: defined but in exclusion list",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), NotIn: []int32{0}}.Build(),
			value:          0, // STATUS_UNSPECIFIED: defined, but in [0]
			wantErr:        true,
			violationCount: 1,
			ruleIDs:        []string{"enum.not_in"},
		},
		{
			name:           "defined_only+not_in/fail_defined_only: undefined but not in exclusion list",
			rules:          validate.EnumRules_builder{DefinedOnly: proto.Bool(true), NotIn: []int32{0}}.Build(),
			value:          99, // undefined, but 99 is not in [0]
			wantErr:        true,
			violationCount: 1,
			ruleIDs:        []string{"enum.defined_only"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMessageTypeWithEnum(t, "test.combined", "EnumCombined", enumDesc, &descriptorpb.FieldDescriptorProto{
				Name:     proto.String("status"),
				Number:   proto.Int32(1),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
				TypeName: proto.String(".test.combined.Status"),
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				Options:  fieldOpts(validate.FieldRules_builder{Enum: tt.rules}.Build()),
			})

			// Run with both CEL and native to verify identical results.
			for _, mode := range []bool{false, true} {
				t.Run(fmt.Sprintf("native=%v", !mode), func(t *testing.T) {
					options := []ValidatorOption{WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor())}
					if mode {
						options = append(options, WithDisableNativeRules())
					}

					validator, err := New(options...)
					require.NoError(t, err)

					msg := dynamicpb.NewMessage(msgType.Descriptor())
					msg.Set(msgType.Descriptor().Fields().ByName("status"), protoreflect.ValueOfEnum(tt.value))

					err = validator.Validate(msg)
					if !tt.wantErr {
						require.NoError(t, err)
						return
					}
					require.Error(t, err)
					var valErr *ValidationError
					require.ErrorAs(t, err, &valErr)
					require.Len(t, valErr.Violations, tt.violationCount,
						"expected %d violations, got %d", tt.violationCount, len(valErr.Violations))
					for i, expectedID := range tt.ruleIDs {
						assert.Equal(t, expectedID, valErr.Violations[i].Proto.GetRuleId(),
							"violation[%d] rule ID mismatch", i)
					}
				})
			}
		})
	}
}

// newDynamicMessageTypeWithEnum creates a dynamic message type that includes
// an enum type definition.
func newDynamicMessageTypeWithEnum(
	t testing.TB,
	pkg, name string,
	enumDesc *descriptorpb.EnumDescriptorProto,
	field *descriptorpb.FieldDescriptorProto,
) protoreflect.MessageType {
	t.Helper()

	file := &descriptorpb.FileDescriptorProto{
		Name:    proto.String(pkg + "." + name + ".proto"),
		Package: proto.String(pkg),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"buf/validate/validate.proto",
		},
		EnumType: []*descriptorpb.EnumDescriptorProto{enumDesc},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:  proto.String(name),
			Field: []*descriptorpb.FieldDescriptorProto{field},
		}},
	}

	registry := newRegistryWithValidateProto(t)
	fd, err := protodesc.FileOptions{}.New(file, registry)
	require.NoError(t, err)

	desc := fd.Messages().ByName(protoreflect.Name(name))
	require.NotNil(t, desc)

	return dynamicpb.NewMessageType(desc)
}
