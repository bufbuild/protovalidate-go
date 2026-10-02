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

func supplementaryRegexCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "regex/lookahead",
			message: &supplementaryv1.Re0{},
			text:    `val:"ab"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/lookbehind",
			message: &supplementaryv1.Re1{},
			text:    `val:"ab"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/backreference",
			message: &supplementaryv1.Re2{},
			text:    `val:"aa"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/any_byte_C",
			message: &supplementaryv1.Re3{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/quoted_literal",
			message: &supplementaryv1.Re4{},
			text:    `val:"abc"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^\\Qa.c\\E$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/repeat_1000",
			message: &supplementaryv1.Re5{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"`,
		},
		{
			name:    "regex/repeat_1001",
			message: &supplementaryv1.Re6{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/unicode_class_pL",
			message: &supplementaryv1.Re7{},
			text:    `val:"é"`,
		},
		{
			name:    "regex/unicode_class_greek",
			message: &supplementaryv1.Re8{},
			text:    `val:"λ"`,
		},
		{
			name:    "regex/casefold_unicode",
			message: &supplementaryv1.Re9{},
			text:    `val:"É"`,
		},
		{
			name:    "regex/casefold_sigma",
			message: &supplementaryv1.Re10{},
			text:    `val:"Σ"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/digit_nonascii",
			message: &supplementaryv1.Re11{},
			text:    `val:"٣"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^\\d$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/word_nonascii",
			message: &supplementaryv1.Re12{},
			text:    `val:"é"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^\\w$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/space_nbsp",
			message: &supplementaryv1.Re13{},
			text:    `val:" "`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^\\s$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/space_vtab",
			message: &supplementaryv1.Re14{},
			text:    `val:"\x0b"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^\\s$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/word_boundary_nonascii",
			message: &supplementaryv1.Re15{},
			text:    `val:"éb"`,
		},
		{
			name:    "regex/dollar_before_newline",
			message: &supplementaryv1.Re16{},
			text:    `val:"abc\n"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^abc$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/z_anchor",
			message: &supplementaryv1.Re17{},
			text:    `val:"abc"`,
		},
		{
			name:    "regex/dot_newline",
			message: &supplementaryv1.Re18{},
			text:    `val:"\n"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.pattern", message: "does not match regex pattern `^.$`", field: "val", rule: "string.pattern"},
			}},
		},
		{
			name:    "regex/dot_newline_s_flag",
			message: &supplementaryv1.Re19{},
			text:    `val:"\n"`,
		},
		{
			name:    "regex/multiline_flag",
			message: &supplementaryv1.Re20{},
			text:    `val:"a\nb"`,
		},
		{
			name:    "regex/empty_pattern",
			message: &supplementaryv1.Re21{},
			text:    `val:"abc"`,
		},
		{
			name:    "regex/posix_class",
			message: &supplementaryv1.Re22{},
			text:    `val:"123"`,
		},
		{
			name:    "regex/named_group_P",
			message: &supplementaryv1.Re23{},
			text:    `val:"a"`,
		},
		{
			name:    "regex/named_group_angle",
			message: &supplementaryv1.Re24{},
			text:    `val:"a"`,
		},
		{
			name:    "regex/ungreedy_flag",
			message: &supplementaryv1.Re25{},
			text:    `val:"aaa"`,
		},
		{
			name:    "regex/hex_escape_braces",
			message: &supplementaryv1.Re26{},
			text:    `val:"é"`,
		},
		{
			name:    "regex/octal_escape",
			message: &supplementaryv1.Re27{},
			text:    `val:"a"`,
		},
		{
			name:    "regex/invalid_unclosed_paren",
			message: &supplementaryv1.Re28{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/invalid_unclosed_bracket",
			message: &supplementaryv1.Re29{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/invalid_leading_star",
			message: &supplementaryv1.Re30{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/invalid_range",
			message: &supplementaryv1.Re31{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/hyphen_after_class_escape",
			message: &supplementaryv1.Re32{},
			text:    `val:"-"`,
		},
		{
			name:    "regex/possessive",
			message: &supplementaryv1.Re33{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/atomic_group",
			message: &supplementaryv1.Re34{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
		{
			name:    "regex/bytes_pattern_unicode",
			message: &supplementaryv1.ReBytesUnicode{},
			text:    `val:"é"`,
		},
		{
			name:    "regex/bytes_pattern_invalid_utf8",
			message: &supplementaryv1.ReBytesBadUtf8{},
			text:    `val:"\xc3"`,
			want:    supplementaryResult{err: "EvaluationError"},
		},
		{
			name:    "regex/cel_matches_any_byte_C",
			message: &supplementaryv1.ReCelMatches{},
			text:    `val:"a"`,
			want:    supplementaryResult{err: "CompilationError"},
		},
	}
}

func supplementaryFormatCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "format/hostname/253_plus_trailing_dot",
			message: &supplementaryv1.Fmt0{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa."`,
		},
		{
			name:    "format/hostname/254_plus_trailing_dot",
			message: &supplementaryv1.Fmt1{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa."`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.hostname", message: "must be a valid hostname", field: "val", rule: "string.hostname"},
			}},
		},
		{
			name:    "format/hostname/label_63_trailing_dot",
			message: &supplementaryv1.Fmt2{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa."`,
		},
		{
			name:    "format/hostname/punycode",
			message: &supplementaryv1.Fmt3{},
			text:    `val:"xn--bcher-kva.example"`,
		},
		{
			name:    "format/email/total_254",
			message: &supplementaryv1.Fmt4{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb.ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc.ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd.eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"`,
		},
		{
			name:    "format/email/total_255",
			message: &supplementaryv1.Fmt5{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb.ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc.ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd.eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"`,
		},
		{
			name:    "format/email/quoted_local_65",
			message: &supplementaryv1.Fmt6{},
			text:    `val:"\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"@b.com"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.email", message: "must be a valid email address", field: "val", rule: "string.email"},
			}},
		},
		{
			name:    "format/email/local_64",
			message: &supplementaryv1.Fmt7{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@b.com"`,
		},
		{
			name:    "format/email/trailing_dot_domain",
			message: &supplementaryv1.Fmt8{},
			text:    `val:"a@b.c."`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.email", message: "must be a valid email address", field: "val", rule: "string.email"},
			}},
		},
		{
			name:    "format/ip_prefix/v4_mapped_host_bits",
			message: &supplementaryv1.Fmt9{},
			text:    `val:"::ffff:1.2.3.4/120"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.ip_prefix", message: "must be a valid IP prefix", field: "val", rule: "string.ip_prefix"},
			}},
		},
		{
			name:    "format/ip_prefix/v4_mapped_network",
			message: &supplementaryv1.Fmt10{},
			text:    `val:"::ffff:1.2.3.0/120"`,
		},
		{
			name:    "format/ipv6_prefix/v6_prefix_v4_mapped",
			message: &supplementaryv1.Fmt11{},
			text:    `val:"::ffff:1.2.3.0/120"`,
		},
		{
			name:    "format/ip/zone_upper",
			message: &supplementaryv1.Fmt12{},
			text:    `val:"fe80::1%ETH0"`,
		},
		{
			name:    "format/ip/zone_percent_25",
			message: &supplementaryv1.Fmt13{},
			text:    `val:"fe80::1%25eth0"`,
		},
		{
			name:    "format/host_and_port/zone_percent_25_port",
			message: &supplementaryv1.Fmt14{},
			text:    `val:"[::1%25eth0]:80"`,
		},
		{
			name:    "format/host_and_port/zone_port",
			message: &supplementaryv1.Fmt15{},
			text:    `val:"[fe80::1%eth0]:80"`,
		},
		{
			name:    "format/host_and_port/v4_mapped_port",
			message: &supplementaryv1.Fmt16{},
			text:    `val:"[::ffff:1.2.3.4]:443"`,
		},
		{
			name:    "format/host_and_port/trailing_dot_port",
			message: &supplementaryv1.Fmt17{},
			text:    `val:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.:8080"`,
		},
		{
			name:    "format/uri/zone_upper",
			message: &supplementaryv1.Fmt18{},
			text:    `val:"http://[fe80::1%25ETH0]/"`,
		},
		{
			name:    "format/uri/empty_port",
			message: &supplementaryv1.Fmt19{},
			text:    `val:"http://example.com:/"`,
		},
		{
			name:    "format/uri/big_port",
			message: &supplementaryv1.Fmt20{},
			text:    `val:"http://example.com:65536/"`,
		},
		{
			name:    "format/uri/non_ascii_host",
			message: &supplementaryv1.Fmt21{},
			text:    `val:"http://ä.com/"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.uri", message: "must be a valid URI", field: "val", rule: "string.uri"},
			}},
		},
		{
			name:    "format/uri_ref/authority_zone",
			message: &supplementaryv1.Fmt22{},
			text:    `val:"//[fe80::1%25eth0]"`,
		},
		{
			name:    "format/ipv4/trailing_newline",
			message: &supplementaryv1.Fmt23{},
			text:    `val:"1.2.3.4\n"`,
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.ipv4", message: "must be a valid IPv4 address", field: "val", rule: "string.ipv4"},
			}},
		},
	}
}

func supplementaryHeaderCases() []supplementaryCase {
	return []supplementaryCase{
		{
			name:    "header/http_header_name/loose/empty",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			want: supplementaryResult{violations: []supplementaryViolation{
				{ruleID: "string.well_known_regex.header_name_empty", message: "value is empty, which is not a valid HTTP header name", field: "val", rule: "string.well_known_regex"},
			}},
		},
		{
			name:    "header/http_header_name/loose/del",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"\x7f"`,
		},
		{
			name:    "header/http_header_name/loose/c1_control",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"\u0085"`,
		},
		{
			name:    "header/http_header_name/loose/unicode",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"é"`,
		},
		{
			name:    "header/http_header_name/loose/colon_prefix",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:":a"`,
		},
		{
			name:    "header/http_header_name/loose/only_colon",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:":"`,
		},
		{
			name:    "header/http_header_name/loose/tab",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"a\tb"`,
		},
		{
			name:    "header/http_header_name/loose/vtab",
			message: &supplementaryv1.HeaderHttpHeaderName{},
			text:    `val:"a\x0bb"`,
		},
		{
			name:    "header/http_header_value/loose/empty",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
		},
		{
			name:    "header/http_header_value/loose/del",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"\x7f"`,
		},
		{
			name:    "header/http_header_value/loose/c1_control",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"\u0085"`,
		},
		{
			name:    "header/http_header_value/loose/unicode",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"é"`,
		},
		{
			name:    "header/http_header_value/loose/colon_prefix",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:":a"`,
		},
		{
			name:    "header/http_header_value/loose/only_colon",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:":"`,
		},
		{
			name:    "header/http_header_value/loose/tab",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"a\tb"`,
		},
		{
			name:    "header/http_header_value/loose/vtab",
			message: &supplementaryv1.HeaderHttpHeaderValue{},
			text:    `val:"a\x0bb"`,
		},
	}
}
