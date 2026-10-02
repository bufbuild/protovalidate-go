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

func supplementaryShapeCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "shape/string_rules_on_repeated",
			message: &supplementaryv1.ShapeRepStr{},
			text:    `val:"a"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/enum_rules_on_repeated",
			message: &supplementaryv1.ShapeRepEnum{},
			text:    `val:7`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/int32_rules_on_repeated_wrapper",
			message: &supplementaryv1.ShapeRepWrapper{},
			text:    `val:{value:100}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/duration_rules_on_repeated",
			message: &supplementaryv1.ShapeRepDuration{},
			text:    `val:{seconds:5}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/repeated_rules_on_singular",
			message: &supplementaryv1.ShapeSingularRepeated{},
			text:    `val:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/repeated_rules_on_map",
			message: &supplementaryv1.ShapeMapRepeated{},
			text:    `val:{key:"a" value:"b"}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/map_rules_on_repeated",
			message: &supplementaryv1.ShapeRepMap{},
			text:    `val:"a"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/string_rules_on_map",
			message: &supplementaryv1.ShapeMapStr{},
			text:    `val:{key:"abcd" value:"abcd"}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/map_key_rules_wrong_type",
			message: &supplementaryv1.ShapeMapKeys{},
			text:    `val:{key:1 value:"a"}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/map_value_rules_wrong_type",
			message: &supplementaryv1.ShapeMapValues{},
			text:    `val:{key:"a" value:"b"}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "shape/items_int32_on_repeated_wrapper/valid",
			message: &supplementaryv1.ItemsWrapper{},
			text:    `val:{value:100}`,
		},
		{
			name:    "shape/items_int32_on_repeated_wrapper/invalid",
			message: &supplementaryv1.ItemsWrapper{},
			text:    `val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val[0]", rule: "repeated.items.int32.gt"},
			}},
		},
	}
}

func supplementaryTypeCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "type/string_field_int32_rules",
			message: &supplementaryv1.Mismatch0{},
			text:    `val:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/bytes_field_string_rules",
			message: &supplementaryv1.Mismatch1{},
			text:    `val:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/bool_field_string_rules",
			message: &supplementaryv1.Mismatch2{},
			text:    `val:true`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/enum_field_int32_rules",
			message: &supplementaryv1.Mismatch3{},
			text:    `val:COLOR_RED`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/int32_field_int64_rules",
			message: &supplementaryv1.Mismatch4{},
			text:    `val:5`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/uint32_field_int32_rules",
			message: &supplementaryv1.Mismatch5{},
			text:    `val:5`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/sint32_field_int32_rules",
			message: &supplementaryv1.Mismatch6{},
			text:    `val:5`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/fixed32_field_uint32_rules",
			message: &supplementaryv1.Mismatch7{},
			text:    `val:5`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/string_field_bytes_rules",
			message: &supplementaryv1.Mismatch8{},
			text:    `val:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/message_field_string_rules",
			message: &supplementaryv1.MsgMismatch0{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/int32value_string_rules",
			message: &supplementaryv1.MsgMismatch1{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/int64value_int32_rules",
			message: &supplementaryv1.MsgMismatch2{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/stringvalue_bytes_rules",
			message: &supplementaryv1.MsgMismatch3{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/duration_field_timestamp_rules",
			message: &supplementaryv1.MsgMismatch4{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/field_mask_rules_on_duration",
			message: &supplementaryv1.MsgMismatch5{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "type/any_rules_on_item",
			message: &supplementaryv1.MsgMismatch6{},
			text:    `val:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
	}
}

func supplementaryUniqueCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "unique/bytes",
			message: &supplementaryv1.UniqueBytes{},
			text:    `val:"a" val:"a"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/bool",
			message: &supplementaryv1.UniqueBool{},
			text:    `val:true val:true`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/enum",
			message: &supplementaryv1.UniqueEnum{},
			text:    `val:COLOR_RED val:COLOR_RED`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/double_zero_negzero",
			message: &supplementaryv1.UniqueDoubleZeroNegzero{},
			text:    `val:0 val:-0`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/double_nan_nan",
			message: &supplementaryv1.UniqueDoubleNanNan{},
			text:    `val:nan val:nan`,
		},
		{
			name:    "unique/float_dup",
			message: &supplementaryv1.UniqueFloatDup{},
			text:    `val:1.5 val:1.5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/uint64_dup",
			message: &supplementaryv1.UniqueUint64Dup{},
			text:    `val:18446744073709551615 val:18446744073709551615`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/sint64_dup",
			message: &supplementaryv1.UniqueSint64Dup{},
			text:    `val:-1 val:-1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/fixed64_distinct",
			message: &supplementaryv1.UniqueFixed64Distinct{},
			text:    `val:1 val:2`,
		},
		{
			name:    "unique/message",
			message: &supplementaryv1.UniqueMsgMessage{},
			text:    `val:{x:1} val:{x:1}`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "unique/int32value",
			message: &supplementaryv1.UniqueMsgInt32Value{},
			text:    `val:{value:1} val:{value:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
			}},
		},
		{
			name:    "unique/duration",
			message: &supplementaryv1.UniqueMsgDuration{},
			text:    `val:{seconds:1} val:{seconds:1}`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
	}
}
