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
	"buf.build/go/protovalidate/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// stringDescriptors bundles the field descriptors for StringRules.
type stringDescriptors struct {
	ruleDesc           protoreflect.FieldDescriptor // FieldRules.string
	constDesc          protoreflect.FieldDescriptor
	lenDesc            protoreflect.FieldDescriptor
	minLenDesc         protoreflect.FieldDescriptor
	maxLenDesc         protoreflect.FieldDescriptor
	lenBytesDesc       protoreflect.FieldDescriptor
	minBytesDesc       protoreflect.FieldDescriptor
	maxBytesDesc       protoreflect.FieldDescriptor
	patternDesc        protoreflect.FieldDescriptor
	prefixDesc         protoreflect.FieldDescriptor
	suffixDesc         protoreflect.FieldDescriptor
	containsDesc       protoreflect.FieldDescriptor
	notContainsDesc    protoreflect.FieldDescriptor
	inDesc             protoreflect.FieldDescriptor
	notInDesc          protoreflect.FieldDescriptor
	hostNameDesc       protoreflect.FieldDescriptor
	hostAndPortDesc    protoreflect.FieldDescriptor
	emailDesc          protoreflect.FieldDescriptor
	ipDesc             protoreflect.FieldDescriptor
	ipv4Desc           protoreflect.FieldDescriptor
	ipv6Desc           protoreflect.FieldDescriptor
	uriDesc            protoreflect.FieldDescriptor
	uriRefDesc         protoreflect.FieldDescriptor
	addressDesc        protoreflect.FieldDescriptor
	uuidDesc           protoreflect.FieldDescriptor
	tuuidDesc          protoreflect.FieldDescriptor
	ipPrefixLenDesc    protoreflect.FieldDescriptor
	ipv4PrefixLenDesc  protoreflect.FieldDescriptor
	ipv6PrefixLenDesc  protoreflect.FieldDescriptor
	ipPrefixDesc       protoreflect.FieldDescriptor
	ipv4PrefixDesc     protoreflect.FieldDescriptor
	ipv6PrefixDesc     protoreflect.FieldDescriptor
	ulidDesc           protoreflect.FieldDescriptor
	wellKnownRegexDesc protoreflect.FieldDescriptor
	strictDesc         protoreflect.FieldDescriptor
}

func makeStringDescriptors() stringDescriptors {
	rulesDesc := (*validate.StringRules)(nil).ProtoReflect().Descriptor()
	return stringDescriptors{
		ruleDesc:           fieldRulesDesc.Fields().ByName("string"),
		constDesc:          rulesDesc.Fields().ByName("const"),
		lenDesc:            rulesDesc.Fields().ByName("len"),
		minLenDesc:         rulesDesc.Fields().ByName("min_len"),
		maxLenDesc:         rulesDesc.Fields().ByName("max_len"),
		lenBytesDesc:       rulesDesc.Fields().ByName("len_bytes"),
		minBytesDesc:       rulesDesc.Fields().ByName("min_bytes"),
		maxBytesDesc:       rulesDesc.Fields().ByName("max_bytes"),
		patternDesc:        rulesDesc.Fields().ByName("pattern"),
		prefixDesc:         rulesDesc.Fields().ByName("prefix"),
		suffixDesc:         rulesDesc.Fields().ByName("suffix"),
		containsDesc:       rulesDesc.Fields().ByName("contains"),
		notContainsDesc:    rulesDesc.Fields().ByName("not_contains"),
		inDesc:             rulesDesc.Fields().ByName("in"),
		notInDesc:          rulesDesc.Fields().ByName("not_in"),
		emailDesc:          rulesDesc.Fields().ByName("email"),
		hostNameDesc:       rulesDesc.Fields().ByName("hostname"),
		ipDesc:             rulesDesc.Fields().ByName("ip"),
		ipv4Desc:           rulesDesc.Fields().ByName("ipv4"),
		ipv6Desc:           rulesDesc.Fields().ByName("ipv6"),
		uriDesc:            rulesDesc.Fields().ByName("uri"),
		uriRefDesc:         rulesDesc.Fields().ByName("uri_ref"),
		addressDesc:        rulesDesc.Fields().ByName("address"),
		uuidDesc:           rulesDesc.Fields().ByName("uuid"),
		tuuidDesc:          rulesDesc.Fields().ByName("tuuid"),
		ipPrefixLenDesc:    rulesDesc.Fields().ByName("ip_with_prefixlen"),
		ipv4PrefixLenDesc:  rulesDesc.Fields().ByName("ipv4_with_prefixlen"),
		ipv6PrefixLenDesc:  rulesDesc.Fields().ByName("ipv6_with_prefixlen"),
		ipPrefixDesc:       rulesDesc.Fields().ByName("ip_prefix"),
		ipv4PrefixDesc:     rulesDesc.Fields().ByName("ipv4_prefix"),
		ipv6PrefixDesc:     rulesDesc.Fields().ByName("ipv6_prefix"),
		hostAndPortDesc:    rulesDesc.Fields().ByName("host_and_port"),
		ulidDesc:           rulesDesc.Fields().ByName("ulid"),
		wellKnownRegexDesc: rulesDesc.Fields().ByName("well_known_regex"),
		strictDesc:         rulesDesc.Fields().ByName("strict"),
	}
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
	desc        protoreflect.FieldDescriptor
	ruleID      string // e.g. "string.ip"
	emptyRuleID string // e.g. "string.ip_empty"; empty means skip the empty check
	mainMsg     string // e.g. "must be a valid IP address"
	emptyMsg    string // e.g. "value is empty, which is not a valid IP address"
	validate    func(string) bool
}

//nolint:gochecknoglobals
var (
	stringRuleEmail = stringWellKnownRule{
		desc:        strDescs.emailDesc,
		ruleID:      "string.email",
		emptyRuleID: "string.email_empty",
		mainMsg:     "must be a valid email address",
		emptyMsg:    "value is empty, which is not a valid email address",
		validate:    cel.IsEmail,
	}
	stringRuleHostname = stringWellKnownRule{
		desc:        strDescs.hostNameDesc,
		ruleID:      "string.hostname",
		emptyRuleID: "string.hostname_empty",
		mainMsg:     "must be a valid hostname",
		emptyMsg:    "value is empty, which is not a valid hostname",
		validate:    cel.IsHostname,
	}
	stringRuleIP = stringWellKnownRule{
		desc:        strDescs.ipDesc,
		ruleID:      "string.ip",
		emptyRuleID: "string.ip_empty",
		mainMsg:     "must be a valid IP address",
		emptyMsg:    "value is empty, which is not a valid IP address",
		validate:    func(s string) bool { return cel.IsIP(s, 0) },
	}
	stringRuleIPv4 = stringWellKnownRule{
		desc:        strDescs.ipv4Desc,
		ruleID:      "string.ipv4",
		emptyRuleID: "string.ipv4_empty",
		mainMsg:     "must be a valid IPv4 address",
		emptyMsg:    "value is empty, which is not a valid IPv4 address",
		validate:    func(s string) bool { return cel.IsIP(s, 4) },
	}
	stringRuleIPv6 = stringWellKnownRule{
		desc:        strDescs.ipv6Desc,
		ruleID:      "string.ipv6",
		emptyRuleID: "string.ipv6_empty",
		mainMsg:     "must be a valid IPv6 address",
		emptyMsg:    "value is empty, which is not a valid IPv6 address",
		validate:    func(s string) bool { return cel.IsIP(s, 6) },
	}
	stringRuleURI = stringWellKnownRule{
		desc:        strDescs.uriDesc,
		ruleID:      "string.uri",
		emptyRuleID: "string.uri_empty",
		mainMsg:     "must be a valid URI",
		emptyMsg:    "value is empty, which is not a valid URI",
		validate:    cel.IsURI,
	}
	stringRuleURIRef = stringWellKnownRule{
		desc:     strDescs.uriRefDesc,
		ruleID:   "string.uri_ref",
		mainMsg:  "must be a valid URI Reference",
		validate: cel.IsURIRef,
	}
	stringRuleAddress = stringWellKnownRule{
		desc:        strDescs.addressDesc,
		ruleID:      "string.address",
		emptyRuleID: "string.address_empty",
		mainMsg:     "must be a valid hostname, or ip address",
		emptyMsg:    "value is empty, which is not a valid hostname, or ip address",
		validate:    func(s string) bool { return cel.IsHostname(s) || cel.IsIP(s, 0) },
	}
	stringRuleUUID = stringWellKnownRule{
		desc:        strDescs.uuidDesc,
		ruleID:      "string.uuid",
		emptyRuleID: "string.uuid_empty",
		mainMsg:     "must be a valid UUID",
		emptyMsg:    "value is empty, which is not a valid UUID",
		validate:    uuidRegexp.MatchString,
	}
	stringRuleTUUID = stringWellKnownRule{
		desc:        strDescs.tuuidDesc,
		ruleID:      "string.tuuid",
		emptyRuleID: "string.tuuid_empty",
		mainMsg:     "must be a valid trimmed UUID",
		emptyMsg:    "value is empty, which is not a valid trimmed UUID",
		validate:    tuuidRegexp.MatchString,
	}
	stringRuleIPPrefixLen = stringWellKnownRule{
		desc:        strDescs.ipPrefixLenDesc,
		ruleID:      "string.ip_with_prefixlen",
		emptyRuleID: "string.ip_with_prefixlen_empty",
		mainMsg:     "must be a valid IP prefix",
		emptyMsg:    "value is empty, which is not a valid IP prefix",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 0, false) },
	}
	stringRuleIPv4PrefixLen = stringWellKnownRule{
		desc:        strDescs.ipv4PrefixLenDesc,
		ruleID:      "string.ipv4_with_prefixlen",
		emptyRuleID: "string.ipv4_with_prefixlen_empty",
		mainMsg:     "must be a valid IPv4 address with prefix length",
		emptyMsg:    "value is empty, which is not a valid IPv4 address with prefix length",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 4, false) },
	}
	stringRuleIPv6PrefixLen = stringWellKnownRule{
		desc:        strDescs.ipv6PrefixLenDesc,
		ruleID:      "string.ipv6_with_prefixlen",
		emptyRuleID: "string.ipv6_with_prefixlen_empty",
		mainMsg:     "must be a valid IPv6 address with prefix length",
		emptyMsg:    "value is empty, which is not a valid IPv6 address with prefix length",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 6, false) },
	}
	stringRuleIPPrefix = stringWellKnownRule{
		desc:        strDescs.ipPrefixDesc,
		ruleID:      "string.ip_prefix",
		emptyRuleID: "string.ip_prefix_empty",
		mainMsg:     "must be a valid IP prefix",
		emptyMsg:    "value is empty, which is not a valid IP prefix",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 0, true) },
	}
	stringRuleIPv4Prefix = stringWellKnownRule{
		desc:        strDescs.ipv4PrefixDesc,
		ruleID:      "string.ipv4_prefix",
		emptyRuleID: "string.ipv4_prefix_empty",
		mainMsg:     "must be a valid IPv4 prefix",
		emptyMsg:    "value is empty, which is not a valid IPv4 prefix",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 4, true) },
	}
	stringRuleIPv6Prefix = stringWellKnownRule{
		desc:        strDescs.ipv6PrefixDesc,
		ruleID:      "string.ipv6_prefix",
		emptyRuleID: "string.ipv6_prefix_empty",
		mainMsg:     "must be a valid IPv6 prefix",
		emptyMsg:    "value is empty, which is not a valid IPv6 prefix",
		validate:    func(s string) bool { return cel.IsIPPrefix(s, 6, true) },
	}
	stringRuleHostAndPort = stringWellKnownRule{
		desc:        strDescs.hostAndPortDesc,
		ruleID:      "string.host_and_port",
		emptyRuleID: "string.host_and_port_empty",
		mainMsg:     "must be a valid host (hostname or IP address) and port pair",
		emptyMsg:    "value is empty, which is not a valid host and port pair",
		validate:    func(s string) bool { return cel.IsHostAndPort(s, true) },
	}
	stringRuleULID = stringWellKnownRule{
		desc:        strDescs.ulidDesc,
		ruleID:      "string.ulid",
		emptyRuleID: "string.ulid_empty",
		mainMsg:     "must be a valid ULID",
		emptyMsg:    "value is empty, which is not a valid ULID",
		validate:    ulidRegexp.MatchString,
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

func (n nativeStringEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, _ *validationConfig) error {
	strVal := val.String()

	var runeCount uint64
	if n.exactLen != nil || n.minLen != nil || n.maxLen != nil {
		runeCount = uint64(utf8.RuneCountInString(strVal)) //nolint:gosec
		err := n.evaluateLength(runeCount, val)
		if err != nil {
			return err
		}
	}

	var byteCount uint64
	if n.exactBytes != nil || n.minBytes != nil || n.maxBytes != nil {
		byteCount = uint64(len(strVal))
		err := n.evaluateByteLength(byteCount, val)
		if err != nil {
			return err
		}
	}

	// const
	if n.constVal != nil && strVal != *n.constVal {
		return n.newViolation(strDescs.ruleDesc, strDescs.constDesc,
			"string.const", fmt.Sprintf("must equal `%s`", *n.constVal),
			val, protoreflect.ValueOfString(*n.constVal))
	}

	// pattern
	if n.pattern != nil && !n.pattern.MatchString(strVal) {
		return n.newViolation(strDescs.ruleDesc, strDescs.patternDesc,
			"string.pattern", fmt.Sprintf("does not match regex pattern `%s`", n.patternStr),
			val, protoreflect.ValueOfString(n.patternStr))
	}

	// prefix
	if n.prefix != nil && !strings.HasPrefix(strVal, *n.prefix) {
		return n.newViolation(strDescs.ruleDesc, strDescs.prefixDesc,
			"string.prefix", fmt.Sprintf("does not have prefix `%s`", *n.prefix),
			val, protoreflect.ValueOfString(*n.prefix))
	}

	// suffix
	if n.suffix != nil && !strings.HasSuffix(strVal, *n.suffix) {
		return n.newViolation(strDescs.ruleDesc, strDescs.suffixDesc,
			"string.suffix", fmt.Sprintf("does not have suffix `%s`", *n.suffix),
			val, protoreflect.ValueOfString(*n.suffix))
	}

	// contains
	if n.contains != nil && !strings.Contains(strVal, *n.contains) {
		return n.newViolation(strDescs.ruleDesc, strDescs.containsDesc,
			"string.contains", fmt.Sprintf("does not contain substring `%s`", *n.contains),
			val, protoreflect.ValueOfString(*n.contains))
	}

	// not_contains
	if n.notContains != nil && strings.Contains(strVal, *n.notContains) {
		return n.newViolation(strDescs.ruleDesc, strDescs.notContainsDesc,
			"string.not_contains", fmt.Sprintf("value contains substring `%s`", *n.notContains),
			val, protoreflect.ValueOfString(*n.notContains))
	}

	// in
	if len(n.inVals) > 0 && !slices.Contains(n.inVals, strVal) {
		return n.newViolation(strDescs.ruleDesc, strDescs.inDesc,
			"string.in", "must be in list "+formatStringList(n.inVals),
			val, protoreflect.ValueOfString(strVal))
	}

	// not_in
	if len(n.notInVals) > 0 && slices.Contains(n.notInVals, strVal) {
		return n.newViolation(strDescs.ruleDesc, strDescs.notInDesc,
			"string.not_in", "must not be in list "+formatStringList(n.notInVals),
			val, protoreflect.ValueOfString(strVal))
	}

	if n.wellKnownRule != nil {
		if err := n.checkWellKnown(strVal, val); err != nil {
			return err
		}
	} else if n.knownRegex != 0 {
		if err := n.checkKnownRegex(strVal, val); err != nil {
			return err
		}
	}

	return nil
}

func (n nativeStringEval) checkWellKnown(strVal string, val protoreflect.Value) error {
	rule := n.wellKnownRule
	if rule.emptyRuleID != "" && strVal == "" {
		return n.newViolation(strDescs.ruleDesc, rule.desc,
			rule.emptyRuleID, rule.emptyMsg,
			val, protoreflect.ValueOfString(strVal))
	}
	if !rule.validate(strVal) {
		return n.newViolation(strDescs.ruleDesc, rule.desc,
			rule.ruleID, rule.mainMsg,
			val, protoreflect.ValueOfString(strVal))
	}
	return nil
}

func (n nativeStringEval) checkKnownRegex(strVal string, val protoreflect.Value) error {
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
			return n.newViolation(strDescs.ruleDesc, strDescs.wellKnownRegexDesc,
				"string.well_known_regex.header_name_empty", "value is empty, which is not a valid HTTP header name",
				val, protoreflect.ValueOfString(strVal))
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
		return n.newViolation(strDescs.ruleDesc, strDescs.wellKnownRegexDesc,
			rule, msg,
			val, protoreflect.ValueOfString(strVal))
	}
	return nil
}

// it would be worse to unify this and evaluateLength than it is to leave them as
// very similar bits of code
//
//nolint:dupl
func (n nativeStringEval) evaluateByteLength(byteCount uint64, val protoreflect.Value) error {
	if n.exactBytes != nil && byteCount != *n.exactBytes {
		return n.newViolation(strDescs.ruleDesc, strDescs.lenBytesDesc,
			"string.len_bytes", fmt.Sprintf("must be %d bytes", *n.exactBytes),
			val, protoreflect.ValueOfUint64(*n.exactBytes))
	}
	if n.minBytes != nil && byteCount < *n.minBytes {
		return n.newViolation(strDescs.ruleDesc, strDescs.minBytesDesc,
			"string.min_bytes", fmt.Sprintf("must be at least %d bytes", *n.minBytes),
			val, protoreflect.ValueOfUint64(*n.minBytes))
	}
	if n.maxBytes != nil && byteCount > *n.maxBytes {
		return n.newViolation(strDescs.ruleDesc, strDescs.maxBytesDesc,
			"string.max_bytes", fmt.Sprintf("must be at most %d bytes", *n.maxBytes),
			val, protoreflect.ValueOfUint64(*n.maxBytes))
	}
	return nil
}

// it would be worse to unify this and evaluateByteLength than it is to leave them as
// very similar bits of code
//
//nolint:dupl
func (n nativeStringEval) evaluateLength(runeCount uint64, val protoreflect.Value) error {
	if n.exactLen != nil && runeCount != *n.exactLen {
		return n.newViolation(strDescs.ruleDesc, strDescs.lenDesc,
			"string.len", fmt.Sprintf("must be %d characters", *n.exactLen),
			val, protoreflect.ValueOfUint64(*n.exactLen))
	}
	if n.minLen != nil && runeCount < *n.minLen {
		return n.newViolation(strDescs.ruleDesc, strDescs.minLenDesc,
			"string.min_len", fmt.Sprintf("must be at least %d characters", *n.minLen),
			val, protoreflect.ValueOfUint64(*n.minLen))
	}
	if n.maxLen != nil && runeCount > *n.maxLen {
		return n.newViolation(strDescs.ruleDesc, strDescs.maxLenDesc,
			"string.max_len", fmt.Sprintf("must be at most %d characters", *n.maxLen),
			val, protoreflect.ValueOfUint64(*n.maxLen))
	}
	return nil
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
		hasRule = true
	}

	var constVal *string
	if rules.HasConst() {
		constVal = ptr(rules.GetConst())
		hasRule = true
	}

	var exactLen *uint64
	if rules.HasLen() {
		exactLen = ptr(rules.GetLen())
		hasRule = true
	}

	var minLen *uint64
	if rules.HasMinLen() {
		minLen = ptr(rules.GetMinLen())
		hasRule = true
	}

	var maxLen *uint64
	if rules.HasMaxLen() {
		maxLen = ptr(rules.GetMaxLen())
		hasRule = true
	}

	var exactBytes *uint64
	if rules.HasLenBytes() {
		exactBytes = ptr(rules.GetLenBytes())
		hasRule = true
	}

	var minBytes *uint64
	if rules.HasMinBytes() {
		minBytes = ptr(rules.GetMinBytes())
		hasRule = true
	}

	var maxBytes *uint64
	if rules.HasMaxBytes() {
		maxBytes = ptr(rules.GetMaxBytes())
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
		hasRule = true
	}

	var prefix *string
	if rules.HasPrefix() {
		prefix = ptr(rules.GetPrefix())
		hasRule = true
	}

	var suffix *string
	if rules.HasSuffix() {
		suffix = ptr(rules.GetSuffix())
		hasRule = true
	}

	var containsVal *string
	if rules.HasContains() {
		containsVal = ptr(rules.GetContains())
		hasRule = true
	}

	var notContains *string
	if rules.HasNotContains() {
		notContains = ptr(rules.GetNotContains())
		hasRule = true
	}

	var inVals []string
	if inVals = rules.GetIn(); len(inVals) > 0 {
		hasRule = true
	}

	var notInVals []string
	if notInVals = rules.GetNotIn(); len(notInVals) > 0 {
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

// formatStringList formats a []string as [a, b] to match CEL's
// string list formatting (no quoting around elements).
func formatStringList(vals []string) string {
	return "[" + strings.Join(vals, ", ") + "]"
}
