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

func supplementaryCelCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "cel/result_int",
			message: &supplementaryv1.Cel0{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_list",
			message: &supplementaryv1.Cel1{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_map",
			message: &supplementaryv1.Cel2{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_double",
			message: &supplementaryv1.Cel3{},
			text:    `val:1.5`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_bytes",
			message: &supplementaryv1.Cel4{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_null",
			message: &supplementaryv1.Cel5{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_duration",
			message: &supplementaryv1.Cel6{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/result_dyn_string",
			message: &supplementaryv1.Cel7{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_nan",
			message: &supplementaryv1.Cel8{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xNaN", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_neg_inf",
			message: &supplementaryv1.Cel9{},
			text:    `val:-inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x-Infinity", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_1e21",
			message: &supplementaryv1.Cel10{},
			text:    `val:1e+21`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x1000000000000000000000", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_1e-7",
			message: &supplementaryv1.Cel11{},
			text:    `val:1e-07`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x0.0000001", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_f_default",
			message: &supplementaryv1.Cel12{},
			text:    `val:0.3333333333333333`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x0.333333", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_duration",
			message: &supplementaryv1.Cel13{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x1.5s", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_timestamp",
			message: &supplementaryv1.Cel14{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x2020-01-01T00:00:00.5Z", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_d_with_string",
			message: &supplementaryv1.Cel15{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/format_too_few_args",
			message: &supplementaryv1.Cel16{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/format_uint",
			message: &supplementaryv1.Cel17{},
			text:    `val:18446744073709551615`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x18446744073709551615", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_bytes_s",
			message: &supplementaryv1.Cel18{},
			text:    `val:"é"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xé", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/format_map_s",
			message: &supplementaryv1.Cel19{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x{a: 2, b: 1}", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_lowerAscii",
			message: &supplementaryv1.Cel20{},
			text:    `val:"ÉA"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xÉa", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_upperAscii",
			message: &supplementaryv1.Cel21{},
			text:    `val:"éa"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xéA", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_split",
			message: &supplementaryv1.Cel22{},
			text:    `val:"a,b,c"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x3", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_join",
			message: &supplementaryv1.Cel23{},
			text:    `val:"a,b"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xa-b", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_replace",
			message: &supplementaryv1.Cel24{},
			text:    `val:"aaa"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xbbb", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_indexOf_unicode",
			message: &supplementaryv1.Cel25{},
			text:    `val:"éb"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x1", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_substring_unicode",
			message: &supplementaryv1.Cel26{},
			text:    `val:"éb"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xb", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_charAt",
			message: &supplementaryv1.Cel27{},
			text:    `val:"éb"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xé", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_trim",
			message: &supplementaryv1.Cel28{},
			text:    `val:" a "`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xax", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_reverse",
			message: &supplementaryv1.Cel29{},
			text:    `val:"abé"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xéba", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_quote",
			message: &supplementaryv1.Cel30{},
			text:    `val:"a\"b"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x\"a\\\"b\"", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/ext_lastIndexOf",
			message: &supplementaryv1.Cel31{},
			text:    `val:"aba"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x2", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/runtime_int_overflow",
			message: &supplementaryv1.Cel32{},
			text:    `val:9223372036854775807`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_div_zero",
			message: &supplementaryv1.Cel33{},
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_mod_zero",
			message: &supplementaryv1.Cel34{},
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_int_min_div_neg1",
			message: &supplementaryv1.Cel35{},
			text:    `val:-9223372036854775808`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_uint_conversion",
			message: &supplementaryv1.Cel36{},
			text:    `val:-1`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_int_conversion_double",
			message: &supplementaryv1.Cel37{},
			text:    `val:1e+19`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_string_to_int",
			message: &supplementaryv1.Cel38{},
			text:    `val:"x"`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/now_compare",
			message: &supplementaryv1.Cel39{},
			text:    `val:1`,
		},
		{
			name:    "cel/getField_missing",
			message: &supplementaryv1.Cel40{},
			text:    `val:1`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/size_bytes",
			message: &supplementaryv1.Cel41{},
			text:    `val:"é"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x2", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/size_string_astral",
			message: &supplementaryv1.Cel42{},
			text:    `val:"😀"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x1", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/double_equality_int",
			message: &supplementaryv1.Cel43{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/heterogeneous_eq",
			message: &supplementaryv1.Cel44{},
			text:    `val:1`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "cel/isNan",
			message: &supplementaryv1.Cel45{},
			text:    `val:nan`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/isInf_sign",
			message: &supplementaryv1.Cel46{},
			text:    `val:-inf`,
		},
		{
			name:    "cel/isHostAndPort_port_required",
			message: &supplementaryv1.Cel47{},
			text:    `val:"example.com"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/isIp_bad_version",
			message: &supplementaryv1.Cel48{},
			text:    `val:"1.2.3.4"`,
		},
		{
			name:    "cel/isIpPrefix_strict",
			message: &supplementaryv1.Cel49{},
			text:    `val:"1.2.3.4/24"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/unique_mixed_types",
			message: &supplementaryv1.Cel50{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "x", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/string_concat_unicode",
			message: &supplementaryv1.Cel51{},
			text:    `val:"é"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "xéé", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/runtime_missing_map_key",
			message: &supplementaryv1.CelMapKey{},
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/runtime_index_out_of_range",
			message: &supplementaryv1.CelIndex{},
			text:    `val:1`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "cel/timestamp_lt_now",
			message: &supplementaryv1.CelTsNow{},
			text:    `val:{seconds:4102444800}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "c", message: "\"this < now\" returned false", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "cel/cel_expression_string_result",
			message: &supplementaryv1.CelExpression{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "this > 5 ? '' : 'too small'", message: "too small", field: "val", rule: "cel_expression[0]"},
			}},
		},
		{
			name:    "cel/message_rule_string_result",
			message: &supplementaryv1.CelMsgString{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "m", message: "a must be set", field: "", rule: ""},
			}},
		},
		{
			name:    "cel/message_rule_int_result",
			message: &supplementaryv1.CelMsgInt{},
			want:    supplementaryResult{err: "CompilationError"},
		},
	}
}

func supplementaryIgnoreCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "ignore/message_items_default_skipped",
			message: &supplementaryv1.IgnoreMsgItems{},
			text:    `val:{} val:{x:-1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "x", message: "x must be positive", field: "val[1]", rule: "repeated.items.cel[0]"},
			}},
		},
		{
			name:    "ignore/message_values_default_skipped",
			message: &supplementaryv1.IgnoreMsgValues{},
			text:    `val:{key:"a" value:{}} val:{key:"b" value:{x:-1}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "x", message: "x must be positive", field: "val[\"b\"]", rule: "map.values.cel[0]"},
			}},
		},
		{
			name:    "ignore/message_with_unknown_field_only",
			message: &supplementaryv1.IgnoreMsgField{},
			// val is an Item whose only content is unknown field 99 (varint 1).
			raw: []byte{0x0a, 0x00, 0x0a, 0x03, 0x98, 0x06, 0x01},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "x", message: "x must be positive", field: "val", rule: "cel[0]"},
			}},
		},
		{
			name:    "ignore/always_with_required",
			message: &supplementaryv1.IgnoreAlwaysRequired{},
		},
		{
			name:    "ignore/oneof_member_zero_value",
			message: &supplementaryv1.IgnoreOneof{},
			text:    `a:""`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.min_len", message: "must be at least 3 characters", field: "a", rule: "string.min_len"},
			}},
		},
	}
}
