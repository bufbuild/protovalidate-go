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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func buildNativeString(t testing.TB, rules *validate.StringRules) evaluator {
	t.Helper()
	fdesc := newFieldDescriptor(t, descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum())
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryBuildNativeStringRules(b, rules)
}

func TestNativeStringConst(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Const: proto.String("hello")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("hello"), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfString("world"), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "string.const", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must equal `hello`", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeStringLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Len: proto.Uint64(3)}.Build())
	require.NotNil(t, eval)

	// ASCII
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	// Unicode: 3 code points
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("αβγ"), &validationConfig{}))

	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("ab"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abcd"), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfString("ab"), &validationConfig{})
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "string.len", valErr.Violations[0].Proto.GetRuleId())
}

func TestNativeStringMinLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{MinLen: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("ab"), &validationConfig{}))
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("a"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString(""), &validationConfig{}))
}

func TestNativeStringMaxLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{MaxLen: proto.Uint64(3)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString(""), &validationConfig{}))
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abcd"), &validationConfig{}))
}

func TestNativeStringLenBytes(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{LenBytes: proto.Uint64(4)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abcd"), &validationConfig{}))
	// "αβ" is 4 bytes (2 bytes per Greek letter)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("αβ"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abcde"), &validationConfig{}))
}

func TestNativeStringMinBytes(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{MinBytes: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("ab"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("a"), &validationConfig{}))
}

func TestNativeStringMaxBytes(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{MaxBytes: proto.Uint64(3)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abcd"), &validationConfig{}))
}

func TestNativeStringPattern(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Pattern: proto.String("^[a-z]+$")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("ABC"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("abc1"), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfString("ABC"), &validationConfig{})
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "string.pattern", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "does not match regex pattern `^[a-z]+$`", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeStringPattern_InvalidRegex(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Pattern: proto.String("[invalid")}.Build())
	assert.Nil(t, eval, "invalid regex should bail to CEL")
}

func TestNativeStringPrefix(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Prefix: proto.String("foo")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("foobar"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("barfoo"), &validationConfig{}))
}

func TestNativeStringSuffix(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Suffix: proto.String("bar")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("foobar"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("barbaz"), &validationConfig{}))
}

func TestNativeStringContains(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{Contains: proto.String("mid")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("amidst"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("absent"), &validationConfig{}))
}

func TestNativeStringNotContains(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{NotContains: proto.String("bad")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("good"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("badger"), &validationConfig{}))
}

func TestNativeStringIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{In: []string{"a", "b", "c"}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("a"), &validationConfig{}))
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("b"), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfString("d"), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "string.in", valErr.Violations[0].Proto.GetRuleId())
}

func TestNativeStringNotIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{NotIn: []string{"x", "y"}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("a"), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfString("x"), &validationConfig{}))
}

func TestNativeStringUnicode(t *testing.T) {
	t.Parallel()

	// min_len counts runes, not bytes
	eval := buildNativeString(t, validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build())
	require.NotNil(t, eval)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString("α"), &validationConfig{})) // 1 rune, 2 bytes

	// max_bytes counts bytes
	eval2 := buildNativeString(t, validate.StringRules_builder{MaxBytes: proto.Uint64(1)}.Build())
	require.NotNil(t, eval2)
	require.Error(t, eval2.Evaluate(nil, protoreflect.ValueOfString("α"), &validationConfig{})) // 2 bytes > 1
}

// --- Bail-out tests ---

func TestTryBuildNativeStringRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.StringRules
	}{
		{"nil_rules", nil},
		{"empty_rules", validate.StringRules_builder{}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryBuildNativeStringRules(base{}, tt.rules))
		})
	}
}

func TestNativeStringWellKnowns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		rules       *validate.StringRules
		valid       string
		invalid     string
		ruleID      string
		message     string
		emptyRuleID string
		emptyMsg    string
	}{
		{
			name:        "email",
			rules:       validate.StringRules_builder{Email: proto.Bool(true)}.Build(),
			valid:       "user@example.com",
			invalid:     "not-an-email",
			ruleID:      "string.email",
			message:     "must be a valid email address",
			emptyRuleID: "string.email_empty",
			emptyMsg:    "value is empty, which is not a valid email address",
		},
		{
			name:        "hostname",
			rules:       validate.StringRules_builder{Hostname: proto.Bool(true)}.Build(),
			valid:       "example.com",
			invalid:     "-invalid",
			ruleID:      "string.hostname",
			message:     "must be a valid hostname",
			emptyRuleID: "string.hostname_empty",
			emptyMsg:    "value is empty, which is not a valid hostname",
		},
		{
			name:        "ip",
			rules:       validate.StringRules_builder{Ip: proto.Bool(true)}.Build(),
			valid:       "192.168.1.1",
			invalid:     "not-valid",
			ruleID:      "string.ip",
			message:     "must be a valid IP address",
			emptyRuleID: "string.ip_empty",
			emptyMsg:    "value is empty, which is not a valid IP address",
		},
		{
			name:        "ipv4",
			rules:       validate.StringRules_builder{Ipv4: proto.Bool(true)}.Build(),
			valid:       "192.168.1.1",
			invalid:     "::1",
			ruleID:      "string.ipv4",
			message:     "must be a valid IPv4 address",
			emptyRuleID: "string.ipv4_empty",
			emptyMsg:    "value is empty, which is not a valid IPv4 address",
		},
		{
			name:        "ipv6",
			rules:       validate.StringRules_builder{Ipv6: proto.Bool(true)}.Build(),
			valid:       "::1",
			invalid:     "192.168.1.1",
			ruleID:      "string.ipv6",
			message:     "must be a valid IPv6 address",
			emptyRuleID: "string.ipv6_empty",
			emptyMsg:    "value is empty, which is not a valid IPv6 address",
		},
		{
			name:        "uri",
			rules:       validate.StringRules_builder{Uri: proto.Bool(true)}.Build(),
			valid:       "https://example.com",
			invalid:     "not a uri",
			ruleID:      "string.uri",
			message:     "must be a valid URI",
			emptyRuleID: "string.uri_empty",
			emptyMsg:    "value is empty, which is not a valid URI",
		},
		{
			name:    "uri_ref",
			rules:   validate.StringRules_builder{UriRef: proto.Bool(true)}.Build(),
			valid:   "/path/to/resource",
			invalid: "not valid ref",
			ruleID:  "string.uri_ref",
			message: "must be a valid URI Reference",
		},
		{
			name:        "address",
			rules:       validate.StringRules_builder{Address: proto.Bool(true)}.Build(),
			valid:       "example.com",
			invalid:     "!@#$%",
			ruleID:      "string.address",
			message:     "must be a valid hostname, or ip address",
			emptyRuleID: "string.address_empty",
			emptyMsg:    "value is empty, which is not a valid hostname, or ip address",
		},
		{
			name:        "uuid",
			rules:       validate.StringRules_builder{Uuid: proto.Bool(true)}.Build(),
			valid:       "550e8400-e29b-41d4-a716-446655440000",
			invalid:     "not-a-uuid",
			ruleID:      "string.uuid",
			message:     "must be a valid UUID",
			emptyRuleID: "string.uuid_empty",
			emptyMsg:    "value is empty, which is not a valid UUID",
		},
		{
			name:        "tuuid",
			rules:       validate.StringRules_builder{Tuuid: proto.Bool(true)}.Build(),
			valid:       "550e8400e29b41d4a716446655440000",
			invalid:     "not-a-tuuid",
			ruleID:      "string.tuuid",
			message:     "must be a valid trimmed UUID",
			emptyRuleID: "string.tuuid_empty",
			emptyMsg:    "value is empty, which is not a valid trimmed UUID",
		},
		{
			name:        "ip_with_prefixlen",
			rules:       validate.StringRules_builder{IpWithPrefixlen: proto.Bool(true)}.Build(),
			valid:       "192.168.0.1/24",
			invalid:     "not-valid",
			ruleID:      "string.ip_with_prefixlen",
			message:     "must be a valid IP prefix",
			emptyRuleID: "string.ip_with_prefixlen_empty",
			emptyMsg:    "value is empty, which is not a valid IP prefix",
		},
		{
			name:        "ipv4_with_prefixlen",
			rules:       validate.StringRules_builder{Ipv4WithPrefixlen: proto.Bool(true)}.Build(),
			valid:       "192.168.0.1/24",
			invalid:     "not-valid",
			ruleID:      "string.ipv4_with_prefixlen",
			message:     "must be a valid IPv4 address with prefix length",
			emptyRuleID: "string.ipv4_with_prefixlen_empty",
			emptyMsg:    "value is empty, which is not a valid IPv4 address with prefix length",
		},
		{
			name:        "ipv6_with_prefixlen",
			rules:       validate.StringRules_builder{Ipv6WithPrefixlen: proto.Bool(true)}.Build(),
			valid:       "::1/128",
			invalid:     "not-valid",
			ruleID:      "string.ipv6_with_prefixlen",
			message:     "must be a valid IPv6 address with prefix length",
			emptyRuleID: "string.ipv6_with_prefixlen_empty",
			emptyMsg:    "value is empty, which is not a valid IPv6 address with prefix length",
		},
		{
			name:        "ip_prefix",
			rules:       validate.StringRules_builder{IpPrefix: proto.Bool(true)}.Build(),
			valid:       "192.168.0.0/24",
			invalid:     "not-valid",
			ruleID:      "string.ip_prefix",
			message:     "must be a valid IP prefix",
			emptyRuleID: "string.ip_prefix_empty",
			emptyMsg:    "value is empty, which is not a valid IP prefix",
		},
		{
			name:        "ipv4_prefix",
			rules:       validate.StringRules_builder{Ipv4Prefix: proto.Bool(true)}.Build(),
			valid:       "192.168.0.0/24",
			invalid:     "not-valid",
			ruleID:      "string.ipv4_prefix",
			message:     "must be a valid IPv4 prefix",
			emptyRuleID: "string.ipv4_prefix_empty",
			emptyMsg:    "value is empty, which is not a valid IPv4 prefix",
		},
		{
			name:        "ipv6_prefix",
			rules:       validate.StringRules_builder{Ipv6Prefix: proto.Bool(true)}.Build(),
			valid:       "2001:db8::/32",
			invalid:     "not-valid",
			ruleID:      "string.ipv6_prefix",
			message:     "must be a valid IPv6 prefix",
			emptyRuleID: "string.ipv6_prefix_empty",
			emptyMsg:    "value is empty, which is not a valid IPv6 prefix",
		},
		{
			name:        "host_and_port",
			rules:       validate.StringRules_builder{HostAndPort: proto.Bool(true)}.Build(),
			valid:       "example.com:80",
			invalid:     "example.com",
			ruleID:      "string.host_and_port",
			message:     "must be a valid host (hostname or IP address) and port pair",
			emptyRuleID: "string.host_and_port_empty",
			emptyMsg:    "value is empty, which is not a valid host and port pair",
		},
		{
			name:        "ulid",
			rules:       validate.StringRules_builder{Ulid: proto.Bool(true)}.Build(),
			valid:       "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			invalid:     "not-a-ulid",
			ruleID:      "string.ulid",
			message:     "must be a valid ULID",
			emptyRuleID: "string.ulid_empty",
			emptyMsg:    "value is empty, which is not a valid ULID",
		},
		{
			name:        "well_known_regex_header_name",
			rules:       validate.StringRules_builder{WellKnownRegex: validate.KnownRegex_KNOWN_REGEX_HTTP_HEADER_NAME.Enum()}.Build(),
			valid:       "Content-Type",
			invalid:     "invalid header",
			ruleID:      "string.well_known_regex.header_name",
			message:     "must be a valid HTTP header name",
			emptyRuleID: "string.well_known_regex.header_name_empty",
			emptyMsg:    "value is empty, which is not a valid HTTP header name",
		},
		{
			name:    "well_known_regex_header_value",
			rules:   validate.StringRules_builder{WellKnownRegex: validate.KnownRegex_KNOWN_REGEX_HTTP_HEADER_VALUE.Enum()}.Build(),
			valid:   "application/json",
			invalid: "\x00",
			ruleID:  "string.well_known_regex.header_value",
			message: "must be a valid HTTP header value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			eval := buildNativeString(t, tt.rules)
			require.NotNil(t, eval)

			require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfString(tt.valid), &validationConfig{}))

			err := eval.Evaluate(nil, protoreflect.ValueOfString(tt.invalid), &validationConfig{})
			require.Error(t, err)
			var valErr *ValidationError
			require.ErrorAs(t, err, &valErr)
			require.Len(t, valErr.Violations, 1)
			assert.Equal(t, tt.ruleID, valErr.Violations[0].Proto.GetRuleId())
			assert.Equal(t, tt.message, valErr.Violations[0].Proto.GetMessage())

			if tt.emptyRuleID != "" {
				err = eval.Evaluate(nil, protoreflect.ValueOfString(""), &validationConfig{})
				require.Error(t, err)
				var emptyValErr *ValidationError
				require.ErrorAs(t, err, &emptyValErr)
				require.Len(t, emptyValErr.Violations, 1)
				assert.Equal(t, tt.emptyRuleID, emptyValErr.Violations[0].Proto.GetRuleId())
				assert.Equal(t, tt.emptyMsg, emptyValErr.Violations[0].Proto.GetMessage())
			}
		})
	}
}

func TestNativeStringTautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeString(t, validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build())
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}
