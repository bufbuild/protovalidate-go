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

func supplementaryNumericCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "numeric/int64_const_min",
			message: &supplementaryv1.Num0{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int64.const", message: "must equal -9223372036854775808", field: "val", rule: "int64.const"},
			}},
		},
		{
			name:    "numeric/int64_const_max",
			message: &supplementaryv1.Num1{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int64.const", message: "must equal 9223372036854775807", field: "val", rule: "int64.const"},
			}},
		},
		{
			name:    "numeric/int64_in_min_max",
			message: &supplementaryv1.Num2{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int64.in", message: "must be in list [-9223372036854775808, 9223372036854775807]", field: "val", rule: "int64.in"},
			}},
		},
		{
			name:    "numeric/int64_gt_max",
			message: &supplementaryv1.Num3{},
			text:    `val:9223372036854775807`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int64.gt", message: "must be greater than 9223372036854775807", field: "val", rule: "int64.gt"},
			}},
		},
		{
			name:    "numeric/int64_lt_min",
			message: &supplementaryv1.Num4{},
			text:    `val:-9223372036854775808`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int64.lt", message: "must be less than -9223372036854775808", field: "val", rule: "int64.lt"},
			}},
		},
		{
			name:    "numeric/uint64_const_max",
			message: &supplementaryv1.Num5{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "uint64.const", message: "must equal 18446744073709551615", field: "val", rule: "uint64.const"},
			}},
		},
		{
			name:    "numeric/uint64_gte_max",
			message: &supplementaryv1.Num6{},
			text:    `val:18446744073709551615`,
		},
		{
			name:    "numeric/int32_gt_lt_equal",
			message: &supplementaryv1.Num7{},
			text:    `val:5`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt_lt", message: "must be greater than 5 and less than 5", field: "val", rule: "int32.gt"},
			}},
		},
		{
			name:    "numeric/int32_gte_lte_equal/valid",
			message: &supplementaryv1.Num8{},
			text:    `val:5`,
		},
		{
			name:    "numeric/int32_gte_lte_equal/invalid",
			message: &supplementaryv1.Num9{},
			text:    `val:4`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gte_lte", message: "must be greater than or equal to 5 and less than or equal to 5", field: "val", rule: "int32.gte"},
			}},
		},
		{
			name:    "numeric/int32_exclusive_gt10_lt5",
			message: &supplementaryv1.Num10{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt_lt_exclusive", message: "must be greater than 10 or less than 5", field: "val", rule: "int32.gt"},
			}},
		},
		{
			name:    "numeric/double_const_nan/val_nan",
			message: &supplementaryv1.Num11{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal NaN", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "numeric/double_const_nan/val_1",
			message: &supplementaryv1.Num12{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal NaN", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "numeric/double_in_nan/val_nan",
			message: &supplementaryv1.Num13{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.in", message: "must be in list [NaN]", field: "val", rule: "double.in"},
			}},
		},
		{
			name:    "numeric/double_not_in_nan/val_nan",
			message: &supplementaryv1.Num14{},
			text:    `val:nan`,
		},
		{
			name:    "numeric/double_const_neg_inf",
			message: &supplementaryv1.Num15{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal -Infinity", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "numeric/double_lte_inf/val_inf",
			message: &supplementaryv1.Num16{},
			text:    `val:inf`,
		},
		{
			name:    "numeric/double_gt_negzero/val_zero",
			message: &supplementaryv1.Num17{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gt", message: "must be greater than -0", field: "val", rule: "double.gt"},
			}},
		},
		{
			name:    "numeric/double_const_tiny",
			message: &supplementaryv1.Num18{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal 0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "numeric/double_const_max",
			message: &supplementaryv1.Num19{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal 179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "numeric/float_const_0.1",
			message: &supplementaryv1.Num20{},
			text:    `val:0.2`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "float.const", message: "must equal 0.10000000149011612", field: "val", rule: "float.const"},
			}},
		},
		{
			name:    "numeric/float_in_0.1/val_0.1",
			message: &supplementaryv1.Num21{},
			text:    `val:0.1`,
		},
		{
			name:    "numeric/float_lt_max/val_inf",
			message: &supplementaryv1.Num22{},
			text:    `val:inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "float.lt", message: "must be less than 340282346638528860000000000000000000000", field: "val", rule: "float.lt"},
			}},
		},
		{
			name:    "numeric/float_const_inf",
			message: &supplementaryv1.Num23{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "float.const", message: "must equal Infinity", field: "val", rule: "float.const"},
			}},
		},
	}
}

func supplementaryTimeCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "time/timestamp_const_min",
			message: &supplementaryv1.Time0{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 0001-01-01T00:00:00Z", field: "val", rule: "timestamp.const"},
			}},
		},
		{
			name:    "time/timestamp_const_max",
			message: &supplementaryv1.Time1{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 9999-12-31T23:59:59.999999999Z", field: "val", rule: "timestamp.const"},
			}},
		},
		{
			name:    "time/timestamp_const_before_min",
			message: &supplementaryv1.Time2{},
			text:    `val:{}`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "time/timestamp_const_nanos_1e9",
			message: &supplementaryv1.Time3{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 1970-01-01T00:00:01Z", field: "val", rule: "timestamp.const"},
			}},
		},
		{
			name:    "time/timestamp_const_nanos_negative",
			message: &supplementaryv1.Time4{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 1969-12-31T23:59:59.999999999Z", field: "val", rule: "timestamp.const"},
			}},
		},
		{
			name:    "time/timestamp_gt_value_year_10000",
			message: &supplementaryv1.Time5{},
			text:    `val:{seconds:253402300800}`,
		},
		{
			name:    "time/timestamp_gt_value_nanos_invalid",
			message: &supplementaryv1.Time6{},
			text:    `val:{seconds:5 nanos:-5}`,
		},
		{
			name:    "time/timestamp_gt_now_and_const",
			message: &supplementaryv1.Time7{},
			text:    `val:{seconds:2}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 1970-01-01T00:00:01Z", field: "val", rule: "timestamp.const"},
				{ruleID: "timestamp.gt_now", message: "must be greater than now", field: "val", rule: "timestamp.gt_now"},
			}},
		},
		{
			name:    "time/timestamp_within_zero",
			message: &supplementaryv1.Time8{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.within", message: "must be within 0s of now", field: "val", rule: "timestamp.within"},
			}},
		},
		{
			name:    "time/timestamp_within_negative",
			message: &supplementaryv1.Time9{},
			text:    `val:{}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.within", message: "must be within -10s of now", field: "val", rule: "timestamp.within"},
			}},
		},
		{
			name:    "time/timestamp_lt_now_and_gt",
			message: &supplementaryv1.Time10{},
			text:    `val:{seconds:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.gt", message: "must be greater than 2100-01-01T00:00:00Z", field: "val", rule: "timestamp.gt"},
			}},
		},
		{
			name:    "time/duration_const_max",
			message: &supplementaryv1.Time11{},
			text:    `val:{seconds:1}`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "time/duration_const_past_max",
			message: &supplementaryv1.Time12{},
			text:    `val:{seconds:1}`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "time/duration_const_mixed_sign",
			message: &supplementaryv1.Time13{},
			text:    `val:{seconds:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "duration.const", message: "must equal 0.999999999s", field: "val", rule: "duration.const"},
			}},
		},
		{
			name:    "time/duration_in_fractional",
			message: &supplementaryv1.Time14{},
			text:    `val:{seconds:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "duration.in", message: "must be in list [0.000000001s, -1.5s]", field: "val", rule: "duration.in"},
			}},
		},
		{
			name:    "time/duration_gt_value_past_max",
			message: &supplementaryv1.Time15{},
			text:    `val:{seconds:315576000001}`,
		},
	}
}

func supplementaryOrderCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "order/string_len_pattern_prefix",
			message: &supplementaryv1.Order0{},
			text:    `val:"ab"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.min_len", message: "must be at least 5 characters", field: "val", rule: "string.min_len"},
				{ruleID: "string.pattern", message: "does not match regex pattern `^z`", field: "val", rule: "string.pattern"},
				{ruleID: "string.prefix", message: "does not have prefix `q`", field: "val", rule: "string.prefix"},
				{ruleID: "string.suffix", message: "does not have suffix `w`", field: "val", rule: "string.suffix"},
				{ruleID: "string.contains", message: "does not contain substring `k`", field: "val", rule: "string.contains"},
			}},
		},
		{
			name:    "order/string_min_len_min_bytes",
			message: &supplementaryv1.Order1{},
			text:    `val:"ab"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.min_len", message: "must be at least 5 characters", field: "val", rule: "string.min_len"},
				{ruleID: "string.min_bytes", message: "must be at least 5 bytes", field: "val", rule: "string.min_bytes"},
			}},
		},
		{
			name:    "order/string_const_len_in",
			message: &supplementaryv1.Order2{},
			text:    `val:"ab"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.const", message: "must equal `abcd`", field: "val", rule: "string.const"},
				{ruleID: "string.len", message: "must be 4 characters", field: "val", rule: "string.len"},
				{ruleID: "string.in", message: "must be in list [abcd]", field: "val", rule: "string.in"},
			}},
		},
		{
			name:    "order/int32_const_in_not_in",
			message: &supplementaryv1.Order3{},
			text:    `val:3`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.const", message: "must equal 1", field: "val", rule: "int32.const"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val", rule: "int32.gt"},
				{ruleID: "int32.in", message: "must be in list [1]", field: "val", rule: "int32.in"},
				{ruleID: "int32.not_in", message: "must not be in list [3]", field: "val", rule: "int32.not_in"},
			}},
		},
		{
			name:    "order/bytes_len_prefix_pattern",
			message: &supplementaryv1.Order4{},
			text:    `val:"ab"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "bytes.min_len", message: "must be at least 5 bytes", field: "val", rule: "bytes.min_len"},
				{ruleID: "bytes.pattern", message: "must match regex pattern `^z`", field: "val", rule: "bytes.pattern"},
				{ruleID: "bytes.prefix", message: "does not have prefix 71", field: "val", rule: "bytes.prefix"},
				{ruleID: "bytes.contains", message: "does not contain 6b", field: "val", rule: "bytes.contains"},
			}},
		},
		{
			name:    "order/enum_const_defined_in_not_in",
			message: &supplementaryv1.Order5{},
			text:    `val:7`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.const", message: "must equal 1", field: "val", rule: "enum.const"},
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
				{ruleID: "enum.in", message: "must be in list [1]", field: "val", rule: "enum.in"},
				{ruleID: "enum.not_in", message: "must not be in list [7]", field: "val", rule: "enum.not_in"},
			}},
		},
		{
			name:    "order/double_finite_and_bounds",
			message: &supplementaryv1.Order6{},
			text:    `val:inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal 2", field: "val", rule: "double.const"},
				{ruleID: "double.not_in", message: "must not be in list [Infinity]", field: "val", rule: "double.not_in"},
				{ruleID: "double.finite", message: "must be finite", field: "val", rule: "double.finite"},
			}},
		},
		{
			name:    "order/string_email_and_len",
			message: &supplementaryv1.Order7{},
			text:    `val:"x"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.min_len", message: "must be at least 10 characters", field: "val", rule: "string.min_len"},
				{ruleID: "string.email", message: "must be a valid email address", field: "val", rule: "string.email"},
			}},
		},
		{
			name:    "order/repeated_min_unique_items",
			message: &supplementaryv1.Order8{},
			text:    `val:"a" val:"a"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 3 item(s)", field: "val", rule: "repeated.min_items"},
				{ruleID: "repeated.unique", message: "repeated value must contain unique items", field: "val", rule: "repeated.unique"},
				{ruleID: "string.min_len", message: "must be at least 3 characters", field: "val[0]", rule: "repeated.items.string.min_len"},
				{ruleID: "string.min_len", message: "must be at least 3 characters", field: "val[1]", rule: "repeated.items.string.min_len"},
			}},
		},
		{
			name:    "order/timestamp_const_lt_now",
			message: &supplementaryv1.OrderTs{},
			text:    `val:{seconds:4102444800}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "timestamp.const", message: "must equal 1970-01-01T00:00:01Z", field: "val", rule: "timestamp.const"},
				{ruleID: "timestamp.lt_now", message: "must be less than now", field: "val", rule: "timestamp.lt_now"},
			}},
		},
		{
			name:    "order/message_cel_then_fields",
			message: &supplementaryv1.OrderMsgField{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "m", message: "m failed", field: "", rule: ""},
				{ruleID: "string.min_len", message: "must be at least 3 characters", field: "a", rule: "string.min_len"},
				{ruleID: "int32.gt", message: "must be greater than 1", field: "b", rule: "int32.gt"},
			}},
		},
	}
}
