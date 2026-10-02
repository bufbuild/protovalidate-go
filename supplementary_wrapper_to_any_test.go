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

import supplementaryv1 "buf.build/go/protovalidate/internal/gen/tests/supplementary/v1"

func supplementaryWrapperCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "wrapper/standard_and_cel/both_fail",
			message: &supplementaryv1.WrapperBoth{},
			text:    `val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "custom", message: "custom failed", field: "val", rule: "cel[0]"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val", rule: "int32.gt"},
			}},
		},
		{
			name:    "wrapper/standard_and_cel/cel_fails",
			message: &supplementaryv1.WrapperBoth{},
			text:    `val:{value:50}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "custom", message: "custom failed", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "wrapper/standard_and_cel/unset",
			message: &supplementaryv1.WrapperBoth{},
		},
		{
			name:    "wrapper/items_standard_and_cel",
			message: &supplementaryv1.WrapperItemsBoth{},
			text:    `val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "custom", message: "custom failed", field: "val[0]", rule: "repeated.items.cel[0]"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper/map_values_standard_and_cel",
			message: &supplementaryv1.WrapperMapValues{},
			text:    `val:{key:"k" value:{value:"a"}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "custom", message: "custom failed", field: "val[\"k\"]", rule: "map.values.cel[0]"},
				{ruleID: "string.min_len", message: "must be at least 3 characters", field: "val[\"k\"]", rule: "map.values.string.min_len"},
			}},
		},
		{
			name:    "wrapper/cel_only/unset",
			message: &supplementaryv1.WrapperCelOnlyUnset{},
		},
		{
			name:    "wrapper/cel_only/set_zero",
			message: &supplementaryv1.WrapperCelOnlyUnset{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "\"this > 1\" returned false", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "wrapper/oneof_member",
			message: &supplementaryv1.WrapperOneof{},
			text:    `a:{value:3}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "uint64.in", message: "must be in list [1, 2]", field: "a", rule: "uint64.in"},
			}},
		},
		{
			name:    "wrapper/double_rule_order",
			message: &supplementaryv1.WrapperDoubleOrder{},
			text:    `val:{value:inf}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal 2", field: "val", rule: "double.const"},
				{ruleID: "double.not_in", message: "must not be in list [Infinity]", field: "val", rule: "double.not_in"},
				{ruleID: "double.finite", message: "must be finite", field: "val", rule: "double.finite"},
			}},
		},
	}
}

func supplementaryWrapperListCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "wrapper_list/unique/int32_default_vs_zero",
			message: &supplementaryv1.UniqueWrapperInt32DefaultVsZero{},
			text:    `val:{} val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/int32_distinct",
			message: &supplementaryv1.UniqueWrapperInt32Distinct{},
			text:    `val:{value:1} val:{value:2}`,
		},
		{
			name:    "wrapper_list/unique/int64_dup",
			message: &supplementaryv1.UniqueWrapperInt64Dup{},
			text:    `val:{value:-9223372036854775808} val:{value:-9223372036854775808}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/uint32_dup",
			message: &supplementaryv1.UniqueWrapperUint32Dup{},
			text:    `val:{value:7} val:{value:7}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/uint64_dup",
			message: &supplementaryv1.UniqueWrapperUint64Dup{},
			text:    `val:{value:18446744073709551615} val:{value:18446744073709551615}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/float_dup",
			message: &supplementaryv1.UniqueWrapperFloatDup{},
			text:    `val:{value:1.5} val:{value:1.5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/double_zero_negzero",
			message: &supplementaryv1.UniqueWrapperDoubleZeroNegzero{},
			text:    `val:{} val:{value:-0}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/double_nan_nan",
			message: &supplementaryv1.UniqueWrapperDoubleNanNan{},
			text:    `val:{value:nan} val:{value:nan}`,
		},
		{
			name:    "wrapper_list/unique/bool_dup",
			message: &supplementaryv1.UniqueWrapperBoolDup{},
			text:    `val:{value:true} val:{value:true}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/string_dup",
			message: &supplementaryv1.UniqueWrapperStringDup{},
			text:    `val:{value:"a"} val:{value:"a"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/string_empty_vs_unset",
			message: &supplementaryv1.UniqueWrapperStringEmptyVsUnset{},
			text:    `val:{} val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/bytes_dup",
			message: &supplementaryv1.UniqueWrapperBytesDup{},
			text:    `val:{value:"a"} val:{value:"a"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/unique/bytes_distinct",
			message: &supplementaryv1.UniqueWrapperBytesDistinct{},
			text:    `val:{value:"a"} val:{value:"b"}`,
		},
		{
			name:    "wrapper_list/min_items_unique_items",
			message: &supplementaryv1.WrapperListAndItems{},
			text:    `val:{value:5} val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/min_items_unique_items/valid",
			message: &supplementaryv1.WrapperListAndItems{},
			text:    `val:{value:11} val:{value:12} val:{value:13}`,
		},
		{
			name:    "wrapper_list/min_items_items",
			message: &supplementaryv1.WrapperListMinItemsItems{},
			text:    `val:{value:5} val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/min_items_items/plain_int32",
			message: &supplementaryv1.WrapperListMinItemsItemsPlain{},
			text:    `val:5 val:5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/unique_items",
			message: &supplementaryv1.WrapperListUniqueItems{},
			text:    `val:{value:5} val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/unique_items/plain_int32",
			message: &supplementaryv1.WrapperListUniqueItemsPlain{},
			text:    `val:5 val:5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/min_items_unique",
			message: &supplementaryv1.WrapperListMinItemsUnique{},
			text:    `val:{value:5} val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/min_items_unique/plain_int32",
			message: &supplementaryv1.WrapperListMinItemsUniquePlain{},
			text:    `val:5 val:5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/max_items_items",
			message: &supplementaryv1.WrapperListMaxItemsItems{},
			text:    `val:{value:5} val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.max_items", message: "must contain no more than 1 item(s)", field: "val", rule: "repeated.max_items"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/max_items_items/plain_int32",
			message: &supplementaryv1.WrapperListMaxItemsItemsPlain{},
			text:    `val:5 val:5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.max_items", message: "must contain no more than 1 item(s)", field: "val", rule: "repeated.max_items"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[1]", rule: "repeated.items.int32.gt"},
			}},
		},
		{
			name:    "wrapper_list/min_items_unique/valid",
			message: &supplementaryv1.WrapperListMinItemsUnique{},
			text:    `val:{value:5} val:{value:6} val:{value:7}`,
		},
		{
			name:    "wrapper_list/max_items_unique",
			message: &supplementaryv1.WrapperListMaxItemsUnique{},
			text:    `val:{value:"a"} val:{value:"a"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.max_items", message: "must contain no more than 1 item(s)", field: "val", rule: "repeated.max_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/max_items_unique/valid",
			message: &supplementaryv1.WrapperListMaxItemsUnique{},
			text:    `val:{value:"a"}`,
		},
		{
			name:    "wrapper_list/max_items_unique/plain_string",
			message: &supplementaryv1.WrapperListMaxItemsUniquePlain{},
			text:    `val:"a" val:"a"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.max_items", message: "must contain no more than 1 item(s)", field: "val", rule: "repeated.max_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "wrapper_list/items_cel_only",
			message: &supplementaryv1.WrapperItemsCelOnly{},
			text:    `val:{value:5} val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "\"this > 1\" returned false", field: "val[1]", rule: "repeated.items.cel[0]"},
			}},
		},
		{
			name:    "wrapper_list/max_items_only",
			message: &supplementaryv1.WrapperListMaxItems{},
			text:    `val:{value:"a"} val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.max_items", message: "must contain no more than 1 item(s)", field: "val", rule: "repeated.max_items"},
			}},
		},
	}
}

func supplementaryWrapperMapCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "wrapper_map/values_and_min_pairs",
			message: &supplementaryv1.WrapperMapInt32Values{},
			text:    `val:{key:"k" value:{value:25}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "map.min_pairs", message: "map must be at least 2 entries", field: "val", rule: "map.min_pairs"},
				{ruleID: "int32.gt_lt", message: "must be greater than 10 and less than 20", field: "val[\"k\"]", rule: "map.values.int32.gt"},
			}},
		},
		{
			name:    "wrapper_map/values_unset_value",
			message: &supplementaryv1.WrapperMapInt32Values{},
			text:    `val:{key:"a" value:{}} val:{key:"b" value:{value:15}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt_lt", message: "must be greater than 10 and less than 20", field: "val[\"a\"]", rule: "map.values.int32.gt"},
			}},
		},
		{
			name:    "wrapper_map/bytes_const",
			message: &supplementaryv1.WrapperMapBytesConst{},
			text:    `val:{key:1 value:{value:"y"}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "bytes.const", message: "must be 78", field: "val[1]", rule: "map.values.bytes.const"},
			}},
		},
	}
}

func supplementaryEnumCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "enum/order/const_defined",
			message: &supplementaryv1.EnumOrderConstDefined{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.const", message: "must equal 1", field: "val", rule: "enum.const"},
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
			}},
		},
		{
			name:    "enum/order/defined_in",
			message: &supplementaryv1.EnumOrderDefinedIn{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
				{ruleID: "enum.in", message: "must be in list [1]", field: "val", rule: "enum.in"},
			}},
		},
		{
			name:    "enum/order/defined_not_in",
			message: &supplementaryv1.EnumOrderDefinedNotIn{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
				{ruleID: "enum.not_in", message: "must not be in list [7]", field: "val", rule: "enum.not_in"},
			}},
		},
		{
			name:    "enum/order/const_in",
			message: &supplementaryv1.EnumOrderConstIn{},
			text:    `val:COLOR_GREEN`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.const", message: "must equal 1", field: "val", rule: "enum.const"},
				{ruleID: "enum.in", message: "must be in list [1]", field: "val", rule: "enum.in"},
			}},
		},
		{
			name:    "enum/order/items",
			message: &supplementaryv1.EnumOrderItems{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.const", message: "must equal 1", field: "val[0]", rule: "repeated.items.enum.const"},
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val[0]", rule: "repeated.items.enum.defined_only"},
				{ruleID: "enum.not_in", message: "must not be in list [7]", field: "val[0]", rule: "repeated.items.enum.not_in"},
			}},
		},
		{
			name:    "enum/defined_only_negative",
			message: &supplementaryv1.EnumDefined{},
			text:    `val:-1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
			}},
		},
		{
			name:    "enum/defined_only_int32_min",
			message: &supplementaryv1.EnumDefined{},
			text:    `val:-2147483648`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
			}},
		},
		{
			name:    "enum/in_duplicates",
			message: &supplementaryv1.EnumInDup{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.in", message: "must be in list [1, 1, 2]", field: "val", rule: "enum.in"},
			}},
		},
	}
}

func supplementaryAnyCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "any/in_empty_type_url",
			message: &supplementaryv1.AnyIn{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "any.in", message: "type URL must be in the allow list", field: "val", rule: "any.in"},
			}},
		},
		{
			name:    "any/in_malformed_type_url",
			message: &supplementaryv1.AnyIn{},
			text:    `val:{type_url:"not a url"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "any.in", message: "type URL must be in the allow list", field: "val", rule: "any.in"},
			}},
		},
		{
			name:    "any/in_unresolvable_type_url",
			message: &supplementaryv1.AnyIn{},
			text:    `val:{type_url:"type.googleapis.com/x.Y"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "any.in", message: "type URL must be in the allow list", field: "val", rule: "any.in"},
			}},
		},
	}
}
