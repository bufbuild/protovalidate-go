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

func supplementaryFieldMaskCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "field_mask/const_order",
			message: &supplementaryv1.FieldMaskConst{},
			text:    `val:{paths:"b" paths:"a"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "field_mask.const", message: "must equal paths [a, b]", field: "val", rule: "field_mask.const"},
			}},
		},
		{
			name:    "field_mask/in_subpath",
			message: &supplementaryv1.FieldMaskIn{},
			text:    `val:{paths:"a.b"}`,
		},
		{
			name:    "field_mask/in_empty",
			message: &supplementaryv1.FieldMaskIn{},
			text:    `val:{}`,
		},
	}
}

func supplementaryContradictoryCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "contradictory/in_and_not_in_overlap",
			message: &supplementaryv1.Contra0{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.not_in", message: "must not be in list [1]", field: "val", rule: "int32.not_in"},
			}},
		},
		{
			name:    "contradictory/const_not_in_list",
			message: &supplementaryv1.Contra1{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.in", message: "must be in list [2]", field: "val", rule: "int32.in"},
			}},
		},
		{
			name:    "contradictory/gt_and_gte",
			message: &supplementaryv1.Contra2{},
			text:    `val:3`,
		},
		{
			name:    "contradictory/empty_in_list",
			message: &supplementaryv1.Contra3{},
			text:    `val:3`,
		},
	}
}

func supplementaryRecursionCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "recursion/depth_50_violation_at_bottom",
			message: &supplementaryv1.Recursive{},
			text:    `root:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:1 next:{v:-1}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt", message: "must be greater than 0", field: "root.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.v", rule: "int32.gt"},
			}},
		},
	}
}

func supplementaryParityCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "parity/go/person_valid",
			message: &supplementaryv1.ParityPerson{},
			text:    `id:1234 email:"protovalidate@buf.build" name:"Buf Build" home:{lat:27.38 lng:33.63}`,
		},
		{
			name:    "parity/go/person_bad_email",
			message: &supplementaryv1.ParityPerson{},
			text:    `id:1234 email:"not an email" name:"Buf Build"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.email", message: "must be a valid email address", field: "email", rule: "string.email"},
			}},
		},
		{
			name:    "parity/go/person_only_id_900",
			message: &supplementaryv1.ParityPerson{},
			text:    `id:900`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "uint64.gt", message: "must be greater than 999", field: "id", rule: "uint64.gt"},
				{ruleID: "string.email_empty", message: "value is empty, which is not a valid email address", field: "email", rule: "string.email"},
				{ruleID: "string.pattern", message: "does not match regex pattern `^[[:alpha:]]+( [[:alpha:]]+)*$`", field: "name", rule: "string.pattern"},
			}},
		},
		{
			name:    "parity/go/person_bad_home",
			message: &supplementaryv1.ParityPerson{},
			text:    `id:1234 email:"protovalidate@buf.build" name:"Buf Build" home:{lat:999.999}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gte_lte", message: "must be greater than or equal to -90 and less than or equal to 90", field: "home.lat", rule: "double.gte"},
			}},
		},
		{
			name:    "parity/go/coordinates_lat",
			message: &supplementaryv1.ParityCoordinates{},
			text:    `lat:999.999`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gte_lte", message: "must be greater than or equal to -90 and less than or equal to 90", field: "lat", rule: "double.gte"},
			}},
		},
		{
			name:    "parity/go/coordinates_both",
			message: &supplementaryv1.ParityCoordinates{},
			text:    `lat:999.999 lng:-999.999`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gte_lte", message: "must be greater than or equal to -90 and less than or equal to 90", field: "lat", rule: "double.gte"},
				{ruleID: "double.gte_lte", message: "must be greater than or equal to -180 and less than or equal to 180", field: "lng", rule: "double.gte"},
			}},
		},
		{
			name:    "parity/go/mismatch_rules",
			message: &supplementaryv1.ParityMismatchRules{},
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/go/mixed_valid_invalid_rules",
			message: &supplementaryv1.ParityMixedValidInvalidRules{},
			text:    `string_field_bool_rule:"foo" valid_string_rule:"bar"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/go/enum_rule_order",
			message: &supplementaryv1.ParityEnumRuleOrder{},
			text:    `val:99`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "enum.defined_only", message: "value must be one of the defined enum values", field: "val", rule: "enum.defined_only"},
				{ruleID: "enum.in", message: "must be in list [1]", field: "val", rule: "enum.in"},
				{ruleID: "enum.not_in", message: "must not be in list [99]", field: "val", rule: "enum.not_in"},
			}},
		},
		{
			name:    "parity/go/int32_rule_order",
			message: &supplementaryv1.ParityInt32RuleOrder{},
			text:    `val:3`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val", rule: "int32.gt"},
				{ruleID: "int32.in", message: "must be in list [1]", field: "val", rule: "int32.in"},
				{ruleID: "int32.not_in", message: "must not be in list [3]", field: "val", rule: "int32.not_in"},
			}},
		},
		{
			name:    "parity/go/wrapper_cel_expression/5",
			message: &supplementaryv1.ParityWrapperCelExpression{},
			text:    `val:{value:5}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "this > 100 ? '' : 'must be greater than 100'", message: "must be greater than 100", field: "val", rule: "cel_expression[0]"},
				{ruleID: "int32.gt", message: "must be greater than 10", field: "val", rule: "int32.gt"},
			}},
		},
		{
			name:    "parity/go/wrapper_cel_expression/50",
			message: &supplementaryv1.ParityWrapperCelExpression{},
			text:    `val:{value:50}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "this > 100 ? '' : 'must be greater than 100'", message: "must be greater than 100", field: "val", rule: "cel_expression[0]"},
			}},
		},
		{
			name:    "parity/go/wrapper_cel_expression/500",
			message: &supplementaryv1.ParityWrapperCelExpression{},
			text:    `val:{value:500}`,
		},
		{
			name:    "parity/go/double_finite_pos_inf",
			message: &supplementaryv1.ParityDoubleFinite{},
			text:    `val:inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.finite", message: "must be finite", field: "val", rule: "double.finite"},
			}},
		},
		{
			name:    "parity/py/double_finite_neg_inf",
			message: &supplementaryv1.ParityDoubleFinite{},
			text:    `val:-inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.finite", message: "must be finite", field: "val", rule: "double.finite"},
			}},
		},
		{
			name:    "parity/go/map_min_pairs",
			message: &supplementaryv1.ParityMapMinMax{},
			text:    `val:{key:"a" value:true}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "map.min_pairs", message: "map must be at least 2 entries", field: "val", rule: "map.min_pairs"},
			}},
		},
		{
			name:    "parity/py/map_max_pairs",
			message: &supplementaryv1.ParityMapMinMax{},
			text:    `val:{key:"a" value:true} val:{key:"b" value:true} val:{key:"c" value:true} val:{key:"d" value:true} val:{key:"e" value:true}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "map.max_pairs", message: "map must be at most 4 entries", field: "val", rule: "map.max_pairs"},
			}},
		},
		{
			name:    "parity/go/min_items_wrappers",
			message: &supplementaryv1.ParityMinItemsWrappers{},
			text:    `val:{value:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "repeated.min_items", message: "must contain at least 2 item(s)", field: "val", rule: "repeated.min_items"},
			}},
		},
		{
			name:    "parity/go/header_loose/name_space",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"a b"`,
		},
		{
			name:    "parity/go/header_loose/name_newline",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"a\nb"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.well_known_regex.header_name", message: "must be a valid HTTP header name", field: "val", rule: "string.well_known_regex"},
			}},
		},
		{
			name:    "parity/go/header_loose/value_control_char",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"a\x01b"`,
		},
		{
			name:    "parity/go/header_loose/value_carriage_return",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"a\rb"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.well_known_regex.header_value", message: "must be a valid HTTP header value", field: "val", rule: "string.well_known_regex"},
			}},
		},
		{
			name:    "parity/py/map_key_for_key",
			message: &supplementaryv1.ParityMapKeys{},
			text:    `val:{key:1 value:"a"}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "sint64.lt", message: "must be less than 0", field: "val[1]", rule: "map.keys.sint64.lt", forKey: true},
			}},
		},
		{
			name:    "parity/py/oneof/none",
			message: &supplementaryv1.ParityOneof{},
		},
		{
			name:    "parity/py/oneof/a",
			message: &supplementaryv1.ParityOneof{},
			text:    `a:"x"`,
		},
		{
			name:    "parity/py/oneof/both",
			message: &supplementaryv1.ParityOneof{},
			text:    `a:"x" b:"y"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "message.oneof", message: "only one of a, b can be set", field: "", rule: ""},
			}},
		},
		{
			name:    "parity/py/oneof_required/none",
			message: &supplementaryv1.ParityOneofRequired{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "message.oneof", message: "one of a, b must be set", field: "", rule: ""},
			}},
		},
		{
			name:    "parity/py/oneof_required/a",
			message: &supplementaryv1.ParityOneofRequired{},
			text:    `a:"x"`,
		},
		{
			name:    "parity/py/oneof_required/both",
			message: &supplementaryv1.ParityOneofRequired{},
			text:    `a:"x" b:"y"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "message.oneof", message: "only one of a, b can be set", field: "", rule: ""},
			}},
		},
		{
			name:    "parity/py/oneof_unknown_field/none",
			message: &supplementaryv1.ParityOneofUnknownField{},
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/oneof_unknown_field/a",
			message: &supplementaryv1.ParityOneofUnknownField{},
			text:    `a:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/oneof_unknown_field/both",
			message: &supplementaryv1.ParityOneofUnknownField{},
			text:    `a:"x" b:"y"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/mistyped_rule",
			message: &supplementaryv1.ParityMistypedRule{},
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/nested_mistyped_rule/unset",
			message: &supplementaryv1.ParityNestedMistypedRule{},
		},
		{
			name:    "parity/py/nested_mistyped_rule/set",
			message: &supplementaryv1.ParityNestedMistypedRule{},
			text:    `child:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/nested_message_rule_error/unset",
			message: &supplementaryv1.ParityNestedMessageRuleError{},
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/nested_message_rule_error/set",
			message: &supplementaryv1.ParityNestedMessageRuleError{},
			text:    `child:{}`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/violation_before_error",
			message: &supplementaryv1.ParityViolationBeforeError{},
			text:    `a:"x"`,
			expect:  "CompilationError",
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "parity/py/hostname/253",
			message: &supplementaryv1.ParityHostname{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"`,
		},
		{
			name:    "parity/py/hostname/253_trailing_dot",
			message: &supplementaryv1.ParityHostname{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa."`,
		},
		{
			name:    "parity/py/hostname/254",
			message: &supplementaryv1.ParityHostname{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.hostname", message: "must be a valid hostname", field: "val", rule: "string.hostname"},
			}},
		},
		{
			name:    "parity/py/hostname/254_trailing_dot",
			message: &supplementaryv1.ParityHostname{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa."`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.hostname", message: "must be a valid hostname", field: "val", rule: "string.hostname"},
			}},
		},
		{
			name:    "parity/py/double_const_inf",
			message: &supplementaryv1.ParityDoubleConstInf{},
			text:    `val:1`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.const", message: "must equal Infinity", field: "val", rule: "double.const"},
			}},
		},
		{
			name:    "parity/py/double_infinite_range",
			message: &supplementaryv1.ParityDoubleInfiniteRange{},
			text:    `val:inf`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "double.gt_lt", message: "must be greater than -Infinity and less than Infinity", field: "val", rule: "double.gt"},
			}},
		},
		{
			name:    "parity/py/cel_unique_wrappers/distinct",
			message: &supplementaryv1.ParityCelUniqueWrappers{},
			text:    `val:{value:1} val:{value:2}`,
		},
		{
			name:    "parity/py/cel_unique_wrappers/dup",
			message: &supplementaryv1.ParityCelUniqueWrappers{},
			text:    `val:{value:1} val:{value:1}`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "unique_wrappers", message: "must be unique", field: "val", rule: "cel[0]"},
			}},
		},
	}
}
