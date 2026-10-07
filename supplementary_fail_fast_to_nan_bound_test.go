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

func supplementaryFailFastCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:     "fail_fast/coordinates_both",
			message:  &supplementaryv1.ParityCoordinates{},
			text:     `lat:999.999 lng:-999.999`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gte_lte", message: "must be greater than or equal to -90 and less than or equal to 90", field: "lat", rule: "double.gte"},
			}},
		},
		{
			name:     "fail_fast/enum_rule_order",
			message:  &supplementaryv1.ParityEnumRuleOrder{},
			text:     `val:99`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
			}},
		},
		{
			name:     "fail_fast/int32_rule_order",
			message:  &supplementaryv1.ParityInt32RuleOrder{},
			text:     `val:3`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val", rule: "int32.gt"},
			}},
		},
		{
			name:     "fail_fast/person_only_id_900",
			message:  &supplementaryv1.ParityPerson{},
			text:     `id:900`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "uint64.gt", message: "must be greater than 999", field: "id", rule: "uint64.gt"},
			}},
		},
		{
			name:     "fail_fast/wrapper_list_items",
			message:  &supplementaryv1.WrapperListAndItems{},
			text:     `val:{value:5} val:{value:5}`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
			}},
		},
		{
			name:     "fail_fast/map_values",
			message:  &supplementaryv1.WrapperMapInt32Values{},
			text:     `val:{key:"a" value:{}} val:{key:"b" value:{value:15}}`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt_lt", message: "must be greater than 10 and less than 20", field: "val[\"a\"]", rule: "map.values.int32.gt"},
			}},
		},
		{
			name:     "fail_fast/message_cel_then_fields",
			message:  &supplementaryv1.OrderMsgField{},
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "m", message: "m failed", field: "", rule: ""},
			}},
		},
		{
			name:     "fail_fast/violation_before_error",
			message:  &supplementaryv1.ParityViolationBeforeError{},
			text:     `a:"x"`,
			failFast: true,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.min_len", message: "must be at least 5 characters", field: "a", rule: "string.min_len"},
			}},
		},
	}
}

func supplementaryNanBoundCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "nan_bound/gt_lt_nan/double/nan",
			message: &supplementaryv1.NanBoundGtLtNanDouble{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gt_lt_nan/double/0",
			message: &supplementaryv1.NanBoundGtLtNanDouble{},
		},
		{
			name:    "nan_bound/gt_lt_nan/double/3",
			message: &supplementaryv1.NanBoundGtLtNanDouble{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gt_lt_nan/float/nan",
			message: &supplementaryv1.NanBoundGtLtNanFloat{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gt_lt_nan/float/0",
			message: &supplementaryv1.NanBoundGtLtNanFloat{},
		},
		{
			name:    "nan_bound/gt_lt_nan/float/3",
			message: &supplementaryv1.NanBoundGtLtNanFloat{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gt_lt_nan/double_value/nan",
			message: &supplementaryv1.NanBoundGtLtNanDoubleValue{},
			text:    `val:{value:nan}`,
		},
		{
			name:    "nan_bound/gt_lt_nan/double_value/0",
			message: &supplementaryv1.NanBoundGtLtNanDoubleValue{},
			text:    `val:{}`,
		},
		{
			name:    "nan_bound/gt_lt_nan/double_value/3",
			message: &supplementaryv1.NanBoundGtLtNanDoubleValue{},
			text:    `val:{value:3}`,
		},
		{
			name:    "nan_bound/gte_lte_nan/double/nan",
			message: &supplementaryv1.NanBoundGteLteNanDouble{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gte_lte_nan/double/0",
			message: &supplementaryv1.NanBoundGteLteNanDouble{},
		},
		{
			name:    "nan_bound/gte_lte_nan/double/3",
			message: &supplementaryv1.NanBoundGteLteNanDouble{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gte_lte_nan/float/nan",
			message: &supplementaryv1.NanBoundGteLteNanFloat{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gte_lte_nan/float/0",
			message: &supplementaryv1.NanBoundGteLteNanFloat{},
		},
		{
			name:    "nan_bound/gte_lte_nan/float/3",
			message: &supplementaryv1.NanBoundGteLteNanFloat{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gte_lte_nan/double_value/nan",
			message: &supplementaryv1.NanBoundGteLteNanDoubleValue{},
			text:    `val:{value:nan}`,
		},
		{
			name:    "nan_bound/gte_lte_nan/double_value/0",
			message: &supplementaryv1.NanBoundGteLteNanDoubleValue{},
			text:    `val:{}`,
		},
		{
			name:    "nan_bound/gte_lte_nan/double_value/3",
			message: &supplementaryv1.NanBoundGteLteNanDoubleValue{},
			text:    `val:{value:3}`,
		},
		{
			name:    "nan_bound/gt_nan_lt/double/nan",
			message: &supplementaryv1.NanBoundGtNanLtDouble{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gt_nan_lt/double/0",
			message: &supplementaryv1.NanBoundGtNanLtDouble{},
		},
		{
			name:    "nan_bound/gt_nan_lt/double/3",
			message: &supplementaryv1.NanBoundGtNanLtDouble{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gt_nan_lt/float/nan",
			message: &supplementaryv1.NanBoundGtNanLtFloat{},
			text:    `val:nan`,
		},
		{
			name:    "nan_bound/gt_nan_lt/float/0",
			message: &supplementaryv1.NanBoundGtNanLtFloat{},
		},
		{
			name:    "nan_bound/gt_nan_lt/float/3",
			message: &supplementaryv1.NanBoundGtNanLtFloat{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/gt_nan_lt/double_value/nan",
			message: &supplementaryv1.NanBoundGtNanLtDoubleValue{},
			text:    `val:{value:nan}`,
		},
		{
			name:    "nan_bound/gt_nan_lt/double_value/0",
			message: &supplementaryv1.NanBoundGtNanLtDoubleValue{},
			text:    `val:{}`,
		},
		{
			name:    "nan_bound/gt_nan_lt/double_value/3",
			message: &supplementaryv1.NanBoundGtNanLtDoubleValue{},
			text:    `val:{value:3}`,
		},
		{
			name:    "nan_bound/lt_nan_only/double/nan",
			message: &supplementaryv1.NanBoundLtNanOnlyDouble{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.lt", message: "must be less than NaN", field: "val", rule: "double.lt"},
			}},
		},
		{
			name:    "nan_bound/lt_nan_only/double/0",
			message: &supplementaryv1.NanBoundLtNanOnlyDouble{},
		},
		{
			name:    "nan_bound/lt_nan_only/double/3",
			message: &supplementaryv1.NanBoundLtNanOnlyDouble{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/lt_nan_only/float/nan",
			message: &supplementaryv1.NanBoundLtNanOnlyFloat{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "float.lt", message: "must be less than NaN", field: "val", rule: "float.lt"},
			}},
		},
		{
			name:    "nan_bound/lt_nan_only/float/0",
			message: &supplementaryv1.NanBoundLtNanOnlyFloat{},
		},
		{
			name:    "nan_bound/lt_nan_only/float/3",
			message: &supplementaryv1.NanBoundLtNanOnlyFloat{},
			text:    `val:3`,
		},
		{
			name:    "nan_bound/lt_nan_only/double_value/nan",
			message: &supplementaryv1.NanBoundLtNanOnlyDoubleValue{},
			text:    `val:{value:nan}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.lt", message: "must be less than NaN", field: "val", rule: "double.lt"},
			}},
		},
		{
			name:    "nan_bound/lt_nan_only/double_value/0",
			message: &supplementaryv1.NanBoundLtNanOnlyDoubleValue{},
			text:    `val:{}`,
		},
		{
			name:    "nan_bound/lt_nan_only/double_value/3",
			message: &supplementaryv1.NanBoundLtNanOnlyDoubleValue{},
			text:    `val:{value:3}`,
		},
	}
}
