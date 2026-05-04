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
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate/internal/rules"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// tryBuildNativeStringRules attempts to build a native Go evaluator for
// string rules. Returns nil if the rules can't be handled natively.
func tryBuildNativeStringRules(base base, rules *validate.StringRules) evaluator {
	if rules == nil {
		return nil
	}

	// Bail out for custom predefined extensions.
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	var wellKnownRule *stringWellKnownRule
	var knownRegex validate.KnownRegex
	var strict bool

	if rules.HasWellKnown() {
		var err error
		wellKnownRule, knownRegex, strict, err = parseStringWellKnown(rules)
		if err != nil {
			return nil
		}
		if wellKnownRule != nil {
			rules.ProtoReflect().Clear(wellKnownRule.site.desc)
		}
		if knownRegex != validate.KnownRegex_KNOWN_REGEX_UNSPECIFIED {
			rules.ProtoReflect().Clear(strDescs.wellKnownRegexSite.desc)
		}
		hasRule = true
	}

	var constVal *string
	if rules.HasConst() {
		constVal = ptr(rules.GetConst())
		rules.ProtoReflect().Clear(strDescs.constSite.desc)
		hasRule = true
	}

	var exactLen *uint64
	if rules.HasLen() {
		exactLen = ptr(rules.GetLen())
		rules.ProtoReflect().Clear(strDescs.lenSite.desc)
		hasRule = true
	}

	var minLen *uint64
	if rules.HasMinLen() {
		minLen = ptr(rules.GetMinLen())
		rules.ProtoReflect().Clear(strDescs.minLenSite.desc)
		hasRule = true
	}

	var maxLen *uint64
	if rules.HasMaxLen() {
		maxLen = ptr(rules.GetMaxLen())
		rules.ProtoReflect().Clear(strDescs.maxLenSite.desc)
		hasRule = true
	}

	var exactBytes *uint64
	if rules.HasLenBytes() {
		exactBytes = ptr(rules.GetLenBytes())
		rules.ProtoReflect().Clear(strDescs.lenBytesSite.desc)
		hasRule = true
	}

	var minBytes *uint64
	if rules.HasMinBytes() {
		minBytes = ptr(rules.GetMinBytes())
		rules.ProtoReflect().Clear(strDescs.minBytesSite.desc)
		hasRule = true
	}

	var maxBytes *uint64
	if rules.HasMaxBytes() {
		maxBytes = ptr(rules.GetMaxBytes())
		rules.ProtoReflect().Clear(strDescs.maxBytesSite.desc)
		hasRule = true
	}

	var compiledPattern *regexp.Regexp
	var patternStr string
	if rules.HasPattern() {
		patternStr = rules.GetPattern()
		var err error
		compiledPattern, err = regexp.Compile(patternStr)
		if err != nil {
			// Invalid regex — bail to CEL which will also report a CompilationError.
			return nil
		}
		rules.ProtoReflect().Clear(strDescs.patternSite.desc)
		hasRule = true
	}

	var prefix *string
	if rules.HasPrefix() {
		prefix = ptr(rules.GetPrefix())
		rules.ProtoReflect().Clear(strDescs.prefixSite.desc)
		hasRule = true
	}

	var suffix *string
	if rules.HasSuffix() {
		suffix = ptr(rules.GetSuffix())
		rules.ProtoReflect().Clear(strDescs.suffixSite.desc)
		hasRule = true
	}

	var containsVal *string
	if rules.HasContains() {
		containsVal = ptr(rules.GetContains())
		rules.ProtoReflect().Clear(strDescs.containsSite.desc)
		hasRule = true
	}

	var notContains *string
	if rules.HasNotContains() {
		notContains = ptr(rules.GetNotContains())
		rules.ProtoReflect().Clear(strDescs.notContainsSite.desc)
		hasRule = true
	}

	var inVals []string
	if inVals = rules.GetIn(); len(inVals) > 0 {
		rules.ProtoReflect().Clear(strDescs.inSite.desc)
		hasRule = true
	}

	var notInVals []string
	if notInVals = rules.GetNotIn(); len(notInVals) > 0 {
		rules.ProtoReflect().Clear(strDescs.notInSite.desc)
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeStringEval{
		base:          base,
		constVal:      constVal,
		inVals:        inVals,
		notInVals:     notInVals,
		exactLen:      exactLen,
		minLen:        minLen,
		maxLen:        maxLen,
		exactBytes:    exactBytes,
		minBytes:      minBytes,
		maxBytes:      maxBytes,
		pattern:       compiledPattern,
		patternStr:    patternStr,
		prefix:        prefix,
		suffix:        suffix,
		contains:      containsVal,
		notContains:   notContains,
		wellKnownRule: wellKnownRule,
		knownRegex:    knownRegex,
		strict:        strict,
	}
}

// stringDescriptors bundles the field descriptors for StringRules.
type stringDescriptors struct {
	ruleDesc protoreflect.FieldDescriptor // FieldRules.string

	// Pre-built rule sites for the error path. Each pairs ruleDesc with the
	// corresponding leaf descriptor, built once at init.
	constSite          ruleSite
	lenSite            ruleSite
	minLenSite         ruleSite
	maxLenSite         ruleSite
	lenBytesSite       ruleSite
	minBytesSite       ruleSite
	maxBytesSite       ruleSite
	patternSite        ruleSite
	prefixSite         ruleSite
	suffixSite         ruleSite
	containsSite       ruleSite
	notContainsSite    ruleSite
	inSite             ruleSite
	notInSite          ruleSite
	wellKnownRegexSite ruleSite
}

func makeStringDescriptors() stringDescriptors {
	descriptors := stringDescriptors{
		ruleDesc: fieldRulesDesc.Fields().ByName("string"),
	}
	descriptors.constSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("const"), "string.const", "")
	descriptors.lenSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("len"), "string.len", "")
	descriptors.minLenSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("min_len"), "string.min_len", "")
	descriptors.maxLenSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("max_len"), "string.max_len", "")
	descriptors.lenBytesSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("len_bytes"), "string.len_bytes", "")
	descriptors.minBytesSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("min_bytes"), "string.min_bytes", "")
	descriptors.maxBytesSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("max_bytes"), "string.max_bytes", "")
	descriptors.patternSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("pattern"), "string.pattern", "")
	descriptors.prefixSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("prefix"), "string.prefix", "")
	descriptors.suffixSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("suffix"), "string.suffix", "")
	descriptors.containsSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("contains"), "string.contains", "")
	descriptors.notContainsSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("not_contains"), "string.not_contains", "")
	descriptors.inSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("in"), "string.in", "")
	descriptors.notInSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("not_in"), "string.not_in", "")
	descriptors.wellKnownRegexSite = makeRuleSite(descriptors.ruleDesc, rulesDesc.Fields().ByName("well_known_regex"), "", "")
	return descriptors
}

//nolint:gochecknoglobals
var strDescs = makeStringDescriptors()

var (
	uuidRegexp        = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	tuuidRegexp       = regexp.MustCompile(`^[0-9a-fA-F]{32}$`)
	ulidRegexp        = regexp.MustCompile(`^[0-7][0-9A-HJKMNP-TV-Za-hjkmnp-tv-z]{25}$`)
	looseRegexp       = regexp.MustCompile(`^[^\x00\x0A\x0D]+$`)
	headerNameRegexp  = regexp.MustCompile(`^:?[0-9a-zA-Z!#$%&\\'*+-.^_|~\x60]+$`)
	headerValueRegexp = regexp.MustCompile(`^[^\x00-\x08\x0A-\x1F\x7F]*$`)
)

// stringWellKnownRule describes a well-known string format constraint.
// It bundles the field descriptor, rule IDs, messages, and validation
// function so that all well-known checks share a single generic method.
type stringWellKnownRule struct {
	site      ruleSite // pre-built rule path site for the error path
	emptySite ruleSite // pre-built rule path site for the empty value check
	validate  func(string) bool
}

//nolint:gochecknoglobals
var (
	rulesDesc = (*validate.StringRules)(nil).ProtoReflect().Descriptor()

	stringRuleEmail = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("email"), "string.email", "must be a valid email address"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("email"), "string.email_empty", "value is empty, which is not a valid email address"),
		validate:  rules.IsEmail,
	}
	stringRuleHostname = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("hostname"), "string.hostname", "must be a valid hostname"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("hostname"), "string.hostname_empty", "value is empty, which is not a valid hostname"),
		validate:  rules.IsHostname,
	}
	stringRuleIP = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip"), "string.ip", "must be a valid IP address"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip"), "string.ip_empty", "value is empty, which is not a valid IP address"),
		validate:  func(s string) bool { return rules.IsIP(s, 0) },
	}
	stringRuleIPv4 = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4"), "string.ipv4", "must be a valid IPv4 address"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4"), "string.ipv4_empty", "value is empty, which is not a valid IPv4 address"),
		validate:  func(s string) bool { return rules.IsIP(s, 4) },
	}
	stringRuleIPv6 = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6"), "string.ipv6", "must be a valid IPv6 address"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6"), "string.ipv6_empty", "value is empty, which is not a valid IPv6 address"),
		validate:  func(s string) bool { return rules.IsIP(s, 6) },
	}
	stringRuleURI = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("uri"), "string.uri", "must be a valid URI"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("uri"), "string.uri_empty", "value is empty, which is not a valid URI"),
		validate:  rules.IsURI,
	}
	stringRuleURIRef = stringWellKnownRule{
		site: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("uri_ref"), "string.uri_ref", "must be a valid URI Reference"),
		// emptySite is unused
		validate: rules.IsURIRef,
	}
	stringRuleAddress = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("address"), "string.address", "must be a valid hostname, or ip address"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("address"), "string.address_empty", "value is empty, which is not a valid hostname, or ip address"),
		validate:  func(s string) bool { return rules.IsHostname(s) || rules.IsIP(s, 0) },
	}
	stringRuleUUID = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("uuid"), "string.uuid", "must be a valid UUID"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("uuid"), "string.uuid_empty", "value is empty, which is not a valid UUID"),
		validate:  uuidRegexp.MatchString,
	}
	stringRuleTUUID = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("tuuid"), "string.tuuid", "must be a valid trimmed UUID"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("tuuid"), "string.tuuid_empty", "value is empty, which is not a valid trimmed UUID"),
		validate:  tuuidRegexp.MatchString,
	}
	stringRuleIPPrefixLen = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip_with_prefixlen"), "string.ip_with_prefixlen", "must be a valid IP prefix"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip_with_prefixlen"), "string.ip_with_prefixlen_empty", "value is empty, which is not a valid IP prefix"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 0, false) },
	}
	stringRuleIPv4PrefixLen = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4_with_prefixlen"), "string.ipv4_with_prefixlen", "must be a valid IPv4 address with prefix length"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4_with_prefixlen"), "string.ipv4_with_prefixlen_empty", "value is empty, which is not a valid IPv4 address with prefix length"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 4, false) },
	}
	stringRuleIPv6PrefixLen = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6_with_prefixlen"), "string.ipv6_with_prefixlen", "must be a valid IPv6 address with prefix length"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6_with_prefixlen"), "string.ipv6_with_prefixlen_empty", "value is empty, which is not a valid IPv6 address with prefix length"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 6, false) },
	}
	stringRuleIPPrefix = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip_prefix"), "string.ip_prefix", "must be a valid IP prefix"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ip_prefix"), "string.ip_prefix_empty", "value is empty, which is not a valid IP prefix"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 0, true) },
	}
	stringRuleIPv4Prefix = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4_prefix"), "string.ipv4_prefix", "must be a valid IPv4 prefix"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv4_prefix"), "string.ipv4_prefix_empty", "value is empty, which is not a valid IPv4 prefix"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 4, true) },
	}
	stringRuleIPv6Prefix = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6_prefix"), "string.ipv6_prefix", "must be a valid IPv6 prefix"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ipv6_prefix"), "string.ipv6_prefix_empty", "value is empty, which is not a valid IPv6 prefix"),
		validate:  func(s string) bool { return rules.IsIPPrefix(s, 6, true) },
	}
	stringRuleHostAndPort = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("host_and_port"), "string.host_and_port", "must be a valid host (hostname or IP address) and port pair"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("host_and_port"), "string.host_and_port_empty", "value is empty, which is not a valid host and port pair"),
		validate:  func(s string) bool { return rules.IsHostAndPort(s, true) },
	}
	stringRuleULID = stringWellKnownRule{
		site:      makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ulid"), "string.ulid", "must be a valid ULID"),
		emptySite: makeRuleSite(strDescs.ruleDesc, rulesDesc.Fields().ByName("ulid"), "string.ulid_empty", "value is empty, which is not a valid ULID"),
		validate:  ulidRegexp.MatchString,
	}
)

// nativeStringEval is a native Go evaluator for string rules.
// It replaces CEL evaluation with direct Go operations for
// const, in, not_in, len, min_len, max_len, len_bytes, min_bytes,
// max_bytes, pattern, prefix, suffix, contains, and not_contains.
type nativeStringEval struct {
	base
	constVal      *string
	inVals        []string
	notInVals     []string
	exactLen      *uint64
	minLen        *uint64
	maxLen        *uint64
	exactBytes    *uint64
	minBytes      *uint64
	maxBytes      *uint64
	pattern       *regexp.Regexp
	patternStr    string
	prefix        *string
	suffix        *string
	contains      *string
	notContains   *string
	wellKnownRule *stringWellKnownRule
	knownRegex    validate.KnownRegex
	strict        bool
}

//nolint:gocyclo // this code has nested ifs but it's not hard to follow.
func (n nativeStringEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	strVal := val.String()
	var violations []*Violation

	if n.exactLen != nil || n.minLen != nil || n.maxLen != nil {
		runeCount := uint64(utf8.RuneCountInString(strVal)) //nolint:gosec // cannot be negative
		if vs := n.evaluateLength(runeCount, val); len(vs) > 0 {
			violations = append(violations, vs...)
			if cfg.failFast {
				return &ValidationError{Violations: violations[:1]}
			}
		}
	}

	if n.exactBytes != nil || n.minBytes != nil || n.maxBytes != nil {
		byteCount := uint64(len(strVal))
		if vs := n.evaluateByteLength(byteCount, val); len(vs) > 0 {
			violations = append(violations, vs...)
			if cfg.failFast {
				return &ValidationError{Violations: violations[:1]}
			}
		}
	}

	if n.constVal != nil && strVal != *n.constVal {
		violations = append(violations, n.newViolation(strDescs.constSite,
			"string.const", fmt.Sprintf("must equal `%s`", *n.constVal),
			val, protoreflect.ValueOfString(*n.constVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if n.pattern != nil && !n.pattern.MatchString(strVal) {
		violations = append(violations, n.newViolation(strDescs.patternSite,
			"string.pattern", fmt.Sprintf("does not match regex pattern `%s`", n.patternStr),
			val, protoreflect.ValueOfString(n.patternStr)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if n.prefix != nil && !strings.HasPrefix(strVal, *n.prefix) {
		violations = append(violations, n.newViolation(strDescs.prefixSite,
			"string.prefix", fmt.Sprintf("does not have prefix `%s`", *n.prefix),
			val, protoreflect.ValueOfString(*n.prefix)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if n.suffix != nil && !strings.HasSuffix(strVal, *n.suffix) {
		violations = append(violations, n.newViolation(strDescs.suffixSite,
			"string.suffix", fmt.Sprintf("does not have suffix `%s`", *n.suffix),
			val, protoreflect.ValueOfString(*n.suffix)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if n.contains != nil && !strings.Contains(strVal, *n.contains) {
		violations = append(violations, n.newViolation(strDescs.containsSite,
			"string.contains", fmt.Sprintf("does not contain substring `%s`", *n.contains),
			val, protoreflect.ValueOfString(*n.contains)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if n.notContains != nil && strings.Contains(strVal, *n.notContains) {
		violations = append(violations, n.newViolation(strDescs.notContainsSite,
			"string.not_contains", fmt.Sprintf("contains substring `%s`", *n.notContains),
			val, protoreflect.ValueOfString(*n.notContains)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if len(n.inVals) > 0 && !slices.Contains(n.inVals, strVal) {
		violations = append(violations, n.newViolation(strDescs.inSite,
			"string.in", "must be in list "+formatStringList(n.inVals),
			val, protoreflect.ValueOfString(strVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	if len(n.notInVals) > 0 && slices.Contains(n.notInVals, strVal) {
		violations = append(violations, n.newViolation(strDescs.notInSite,
			"string.not_in", "must not be in list "+formatStringList(n.notInVals),
			val, protoreflect.ValueOfString(strVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations[:1]}
		}
	}

	//nolint:nestif // there are levels of nested ifs, but it's not hard to follow.
	if n.wellKnownRule != nil {
		if vs := n.checkWellKnown(strVal, val); len(vs) > 0 {
			violations = append(violations, vs...)
			if cfg.failFast {
				return &ValidationError{Violations: violations[:1]}
			}
		}
	} else if n.knownRegex != 0 {
		if vs := n.checkKnownRegex(strVal, val); len(vs) > 0 {
			violations = append(violations, vs...)
			if cfg.failFast {
				return &ValidationError{Violations: violations[:1]}
			}
		}
	}

	if len(violations) > 0 {
		return &ValidationError{
			Violations: violations,
		}
	}
	return nil
}

func (n nativeStringEval) checkWellKnown(strVal string, val protoreflect.Value) []*Violation {
	rule := n.wellKnownRule
	if rule.emptySite.ruleID != nil && strVal == "" {
		return []*Violation{n.newViolation(rule.emptySite,
			"", "",
			val, protoreflect.ValueOfString(strVal))}
	}
	if !rule.validate(strVal) {
		return []*Violation{n.newViolation(rule.site,
			"", "",
			val, protoreflect.ValueOfString(strVal))}
	}
	return nil
}

func (n nativeStringEval) checkKnownRegex(strVal string, val protoreflect.Value) []*Violation {
	// check if strict is set (it is on by default)
	// if not, just validate against the loose rule (^[^\u0000\u000A\u000D]+$)
	// if yes, check whether this is a name or value and use the correct strict rule
	// ^:?[0-9a-zA-Z!#$%&\\'*+-.^_|~\\x60]+$ for name
	// ^[^\u0000-\u0008\u000A-\u001F\u007F]*$ for value
	var matcher *regexp.Regexp
	var rule string
	var msg string
	switch n.knownRegex {
	case validate.KnownRegex_KNOWN_REGEX_HTTP_HEADER_NAME:
		if strVal == "" {
			return []*Violation{n.newViolation(strDescs.wellKnownRegexSite,
				"string.well_known_regex.header_name_empty", "value is empty, which is not a valid HTTP header name",
				val, protoreflect.ValueOfString(strVal))}
		}
		matcher = headerNameRegexp
		rule = "string.well_known_regex.header_name"
		msg = "must be a valid HTTP header name"
	case validate.KnownRegex_KNOWN_REGEX_HTTP_HEADER_VALUE:
		matcher = headerValueRegexp
		rule = "string.well_known_regex.header_value"
		msg = "must be a valid HTTP header value"
	default:
		return nil // should never happen, but just in case
	}
	if !n.strict {
		matcher = looseRegexp
	}
	if !matcher.MatchString(strVal) {
		return []*Violation{n.newViolation(strDescs.wellKnownRegexSite,
			rule, msg,
			val, protoreflect.ValueOfString(strVal))}
	}
	return nil
}

// it would be worse to unify this and evaluateLength than it is to leave them as
// very similar bits of code
//
//nolint:dupl
func (n nativeStringEval) evaluateByteLength(byteCount uint64, val protoreflect.Value) []*Violation {
	var out []*Violation
	if n.exactBytes != nil && byteCount != *n.exactBytes {
		out = append(out, n.newViolation(strDescs.lenBytesSite,
			"string.len_bytes", fmt.Sprintf("must be %d bytes", *n.exactBytes),
			val, protoreflect.ValueOfUint64(*n.exactBytes)))
	}
	if n.minBytes != nil && byteCount < *n.minBytes {
		out = append(out, n.newViolation(strDescs.minBytesSite,
			"string.min_bytes", fmt.Sprintf("must be at least %d bytes", *n.minBytes),
			val, protoreflect.ValueOfUint64(*n.minBytes)))
	}
	if n.maxBytes != nil && byteCount > *n.maxBytes {
		out = append(out, n.newViolation(strDescs.maxBytesSite,
			"string.max_bytes", fmt.Sprintf("must be at most %d bytes", *n.maxBytes),
			val, protoreflect.ValueOfUint64(*n.maxBytes)))
	}
	return out
}

// it would be worse to unify this and evaluateByteLength than it is to leave them as
// very similar bits of code
//
//nolint:dupl
func (n nativeStringEval) evaluateLength(runeCount uint64, val protoreflect.Value) []*Violation {
	var out []*Violation

	if n.exactLen != nil && runeCount != *n.exactLen {
		out = append(out, n.newViolation(strDescs.lenSite,
			"string.len", fmt.Sprintf("must be %d characters", *n.exactLen),
			val, protoreflect.ValueOfUint64(*n.exactLen)))
	}
	if n.minLen != nil && runeCount < *n.minLen {
		out = append(out, n.newViolation(strDescs.minLenSite,
			"string.min_len", fmt.Sprintf("must be at least %d characters", *n.minLen),
			val, protoreflect.ValueOfUint64(*n.minLen)))
	}
	if n.maxLen != nil && runeCount > *n.maxLen {
		out = append(out, n.newViolation(strDescs.maxLenSite,
			"string.max_len", fmt.Sprintf("must be at most %d characters", *n.maxLen),
			val, protoreflect.ValueOfUint64(*n.maxLen)))
	}
	return out
}

func (n nativeStringEval) Tautology() bool {
	return false
}

var _ evaluator = nativeStringEval{}

var errUnsupportedWellKnown = errors.New("unsupported well-known string constraint")

// parseStringWellKnown maps a StringRules well-known oneof to a
// *stringWellKnownRule (for format constraints) or a KnownRegex +
// strict flag (for well-known regex constraints).
func parseStringWellKnown(rules *validate.StringRules) (
	*stringWellKnownRule, validate.KnownRegex, bool, error,
) {
	switch {
	case rules.GetEmail():
		return &stringRuleEmail, 0, false, nil
	case rules.GetHostname():
		return &stringRuleHostname, 0, false, nil
	case rules.GetIp():
		return &stringRuleIP, 0, false, nil
	case rules.GetIpv4():
		return &stringRuleIPv4, 0, false, nil
	case rules.GetIpv6():
		return &stringRuleIPv6, 0, false, nil
	case rules.GetUri():
		return &stringRuleURI, 0, false, nil
	case rules.GetUriRef():
		return &stringRuleURIRef, 0, false, nil
	case rules.GetAddress():
		return &stringRuleAddress, 0, false, nil
	case rules.GetUuid():
		return &stringRuleUUID, 0, false, nil
	case rules.GetTuuid():
		return &stringRuleTUUID, 0, false, nil
	case rules.GetIpWithPrefixlen():
		return &stringRuleIPPrefixLen, 0, false, nil
	case rules.GetIpv4WithPrefixlen():
		return &stringRuleIPv4PrefixLen, 0, false, nil
	case rules.GetIpv6WithPrefixlen():
		return &stringRuleIPv6PrefixLen, 0, false, nil
	case rules.GetIpPrefix():
		return &stringRuleIPPrefix, 0, false, nil
	case rules.GetIpv4Prefix():
		return &stringRuleIPv4Prefix, 0, false, nil
	case rules.GetIpv6Prefix():
		return &stringRuleIPv6Prefix, 0, false, nil
	case rules.GetUlid():
		return &stringRuleULID, 0, false, nil
	case rules.GetHostAndPort():
		return &stringRuleHostAndPort, 0, false, nil
	case rules.GetWellKnownRegex() != validate.KnownRegex_KNOWN_REGEX_UNSPECIFIED:
		knownRegex := rules.GetWellKnownRegex()
		// strict is on by default or if it is explicitly set to true
		strict := !rules.HasStrict() || rules.GetStrict()
		// intentionally doesn't return a *stringWellKnownRule, because well known regex is a special case
		return nil, knownRegex, strict, nil
	default:
		return nil, 0, false, errUnsupportedWellKnown
	}
}

// formatStringList formats a []string as [a, b] to match CEL's
// string list formatting (no quoting around elements).
func formatStringList(vals []string) string {
	return "[" + strings.Join(vals, ", ") + "]"
}
