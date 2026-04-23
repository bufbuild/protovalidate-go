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
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// tryBuildNativeBytesRules attempts to build a native Go evaluator for
// bytes rules. Returns nil if the rules can't be handled natively.
func tryBuildNativeBytesRules(base base, rules *validate.BytesRules) evaluator {
	if rules == nil {
		return nil
	}
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	// Detect well-known format constraint (ip, ipv4, ipv6, uuid).
	// Check both presence and value — setting ip=false means no check.
	var wellKnown *bytesWellKnown
	switch {
	case rules.GetIp():
		wellKnown = &bytesWellKnownIP
		hasRule = true
	case rules.GetIpv4():
		wellKnown = &bytesWellKnownIPv4
		hasRule = true
	case rules.GetIpv6():
		wellKnown = &bytesWellKnownIPv6
		hasRule = true
	case rules.GetUuid():
		wellKnown = &bytesWellKnownUUID
		hasRule = true
	}

	var constVal []byte
	var hasConst bool
	if rules.HasConst() {
		constVal = rules.GetConst()
		hasConst = true
		hasRule = true
	}

	var exactLen *uint64
	if rules.HasLen() {
		exactLen = ptr(rules.GetLen())
		hasRule = true
	}

	var minLen uint64
	if rules.HasMinLen() {
		minLen = rules.GetMinLen()
		hasRule = true
	}

	var maxLen uint64 = math.MaxUint64
	if rules.HasMaxLen() {
		maxLen = rules.GetMaxLen()
		hasRule = true
	}

	var compiledPattern *regexp.Regexp
	var patternStr string
	if rules.HasPattern() {
		patternStr = rules.GetPattern()
		var err error
		compiledPattern, err = regexp.Compile(patternStr)
		if err != nil {
			return nil // bail to CEL
		}
		hasRule = true
	}

	var prefix []byte
	var hasPrefix bool
	if rules.HasPrefix() {
		prefix = rules.GetPrefix()
		hasPrefix = true
		hasRule = true
	}

	var suffix []byte
	var hasSuffix bool
	if rules.HasSuffix() {
		suffix = rules.GetSuffix()
		hasSuffix = true
		hasRule = true
	}

	var contains []byte
	var hasContains bool
	if rules.HasContains() {
		contains = rules.GetContains()
		hasContains = true
		hasRule = true
	}

	var inVals [][]byte
	if inVals = rules.GetIn(); len(inVals) > 0 {
		hasRule = true
	}

	var notInVals [][]byte
	if notInVals = rules.GetNotIn(); len(notInVals) > 0 {
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeBytesEval{
		base:        base,
		constVal:    constVal,
		hasConst:    hasConst,
		exactLen:    exactLen,
		minLen:      minLen,
		maxLen:      maxLen,
		pattern:     compiledPattern,
		patternStr:  patternStr,
		prefix:      prefix,
		hasPrefix:   hasPrefix,
		suffix:      suffix,
		hasSuffix:   hasSuffix,
		contains:    contains,
		hasContains: hasContains,
		inVals:      inVals,
		notInVals:   notInVals,
		wellKnown:   wellKnown,
	}
}

// bytesWellKnown identifies which well-known bytes format constraint is active.
type bytesWellKnown struct {
	desc                protoreflect.FieldDescriptor
	ruleID, emptyRuleID string
	mainMsg, emptyMsg   string
	validSizes          []int
}

var (
	//nolint:gochecknoglobals
	bytesWellKnownIP = bytesWellKnown{
		desc:        bytesDescs.ipDesc,
		ruleID:      "bytes.ip",
		emptyRuleID: "bytes.ip_empty",
		mainMsg:     "must be a valid IP address",
		emptyMsg:    "value is empty, which is not a valid IP address",
		validSizes:  []int{4, 16},
	}
	//nolint:gochecknoglobals
	bytesWellKnownIPv4 = bytesWellKnown{
		desc:        bytesDescs.ipv4Desc,
		ruleID:      "bytes.ipv4",
		emptyRuleID: "bytes.ipv4_empty",
		mainMsg:     "must be a valid IPv4 address",
		emptyMsg:    "value is empty, which is not a valid IPv4 address",
		validSizes:  []int{4},
	}
	//nolint:gochecknoglobals
	bytesWellKnownIPv6 = bytesWellKnown{
		desc:        bytesDescs.ipv6Desc,
		ruleID:      "bytes.ipv6",
		emptyRuleID: "bytes.ipv6_empty",
		mainMsg:     "must be a valid IPv6 address",
		emptyMsg:    "value is empty, which is not a valid IPv6 address",
		validSizes:  []int{16},
	}
	//nolint:gochecknoglobals
	bytesWellKnownUUID = bytesWellKnown{
		desc:        bytesDescs.uuidDesc,
		ruleID:      "bytes.uuid",
		emptyRuleID: "bytes.uuid_empty",
		mainMsg:     "must be a valid UUID",
		emptyMsg:    "value is empty, which is not a valid UUID",
		validSizes:  []int{16},
	}
)

// bytesDescriptors bundles the field descriptors for BytesRules.
type bytesDescriptors struct {
	ruleDesc     protoreflect.FieldDescriptor
	constDesc    protoreflect.FieldDescriptor
	lenDesc      protoreflect.FieldDescriptor
	minLenDesc   protoreflect.FieldDescriptor
	maxLenDesc   protoreflect.FieldDescriptor
	patternDesc  protoreflect.FieldDescriptor
	prefixDesc   protoreflect.FieldDescriptor
	suffixDesc   protoreflect.FieldDescriptor
	containsDesc protoreflect.FieldDescriptor
	inDesc       protoreflect.FieldDescriptor
	notInDesc    protoreflect.FieldDescriptor
	ipDesc       protoreflect.FieldDescriptor
	ipv4Desc     protoreflect.FieldDescriptor
	ipv6Desc     protoreflect.FieldDescriptor
	uuidDesc     protoreflect.FieldDescriptor
}

func makeBytesDescriptors() bytesDescriptors {
	rulesDesc := (*validate.BytesRules)(nil).ProtoReflect().Descriptor()
	return bytesDescriptors{
		ruleDesc:     fieldRulesDesc.Fields().ByName("bytes"),
		constDesc:    rulesDesc.Fields().ByName("const"),
		lenDesc:      rulesDesc.Fields().ByName("len"),
		minLenDesc:   rulesDesc.Fields().ByName("min_len"),
		maxLenDesc:   rulesDesc.Fields().ByName("max_len"),
		patternDesc:  rulesDesc.Fields().ByName("pattern"),
		prefixDesc:   rulesDesc.Fields().ByName("prefix"),
		suffixDesc:   rulesDesc.Fields().ByName("suffix"),
		containsDesc: rulesDesc.Fields().ByName("contains"),
		inDesc:       rulesDesc.Fields().ByName("in"),
		notInDesc:    rulesDesc.Fields().ByName("not_in"),
		ipDesc:       rulesDesc.Fields().ByName("ip"),
		ipv4Desc:     rulesDesc.Fields().ByName("ipv4"),
		ipv6Desc:     rulesDesc.Fields().ByName("ipv6"),
		uuidDesc:     rulesDesc.Fields().ByName("uuid"),
	}
}

//nolint:gochecknoglobals
var bytesDescs = makeBytesDescriptors()

var _ evaluator = nativeBytesEval{}

// nativeBytesEval is a native Go evaluator for bytes rules.
type nativeBytesEval struct {
	base
	constVal    []byte
	hasConst    bool
	exactLen    *uint64
	minLen      uint64
	maxLen      uint64
	pattern     *regexp.Regexp
	patternStr  string
	prefix      []byte
	hasPrefix   bool
	suffix      []byte
	hasSuffix   bool
	contains    []byte
	hasContains bool
	inVals      [][]byte
	notInVals   [][]byte
	wellKnown   *bytesWellKnown
}

var errNotUTF8 = errors.New("must be valid UTF-8 to apply regexp")

//nolint:gocyclo // this code has nested ifs but it's not hard to follow.
func (n nativeBytesEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	bytesVal := val.Bytes()
	byteLen := uint64(len(bytesVal))
	var violations []*Violation

	if n.hasConst && !bytes.Equal(bytesVal, n.constVal) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.constDesc,
			"bytes.const", fmt.Sprintf("must be %x", n.constVal),
			val, protoreflect.ValueOfBytes(n.constVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.exactLen != nil && byteLen != *n.exactLen {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.lenDesc,
			"bytes.len", fmt.Sprintf("must be %d bytes", *n.exactLen),
			val, protoreflect.ValueOfUint64(*n.exactLen)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if byteLen < n.minLen {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.minLenDesc,
			"bytes.min_len", fmt.Sprintf("must be at least %d bytes", n.minLen),
			val, protoreflect.ValueOfUint64(n.minLen)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if byteLen > n.maxLen {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.maxLenDesc,
			"bytes.max_len", fmt.Sprintf("must be at most %d bytes", n.maxLen),
			val, protoreflect.ValueOfUint64(n.maxLen)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.pattern != nil {
		if !utf8.Valid(bytesVal) {
			// the bytes.pattern rule requires the value to be UTF-8. Surface
			// this as a RuntimeError to match CEL behavior / conformance tests.
			return &RuntimeError{cause: errNotUTF8}
		}
		if !n.pattern.MatchString(string(bytesVal)) {
			violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.patternDesc,
				"bytes.pattern", fmt.Sprintf("must match regex pattern `%s`", n.patternStr),
				val, protoreflect.ValueOfString(n.patternStr)))
			if cfg.failFast {
				return &ValidationError{Violations: violations}
			}
		}
	}

	if n.hasPrefix && !bytes.HasPrefix(bytesVal, n.prefix) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.prefixDesc,
			"bytes.prefix", fmt.Sprintf("does not have prefix %x", n.prefix),
			val, protoreflect.ValueOfBytes(n.prefix)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.hasSuffix && !bytes.HasSuffix(bytesVal, n.suffix) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.suffixDesc,
			"bytes.suffix", fmt.Sprintf("does not have suffix %x", n.suffix),
			val, protoreflect.ValueOfBytes(n.suffix)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.hasContains && !bytes.Contains(bytesVal, n.contains) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.containsDesc,
			"bytes.contains", fmt.Sprintf("does not contain %x", n.contains),
			val, protoreflect.ValueOfBytes(n.contains)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if len(n.inVals) > 0 && !slices.ContainsFunc(n.inVals, func(v []byte) bool { return bytes.Equal(v, bytesVal) }) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.inDesc,
			"bytes.in", "must be in list "+formatBytesList(n.inVals),
			val, protoreflect.ValueOfBytes(bytesVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if len(n.notInVals) > 0 && slices.ContainsFunc(n.notInVals, func(v []byte) bool { return bytes.Equal(v, bytesVal) }) {
		violations = append(violations, n.newViolation(bytesDescs.ruleDesc, bytesDescs.notInDesc,
			"bytes.not_in", "must not be in list "+formatBytesList(n.notInVals),
			val, protoreflect.ValueOfBytes(bytesVal)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.wellKnown != nil {
		if v := n.evaluateWellKnown(bytesVal, val); v != nil {
			violations = append(violations, v)
			if cfg.failFast {
				return &ValidationError{Violations: violations}
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

func (n nativeBytesEval) evaluateWellKnown(bytesVal []byte, val protoreflect.Value) *Violation {
	size := len(bytesVal)
	wellKnown := n.wellKnown

	if size == 0 {
		return n.newViolation(bytesDescs.ruleDesc, wellKnown.desc,
			wellKnown.emptyRuleID, wellKnown.emptyMsg,
			val, protoreflect.ValueOfBool(true))
	}

	if slices.Contains(wellKnown.validSizes, size) {
		return nil
	}

	return n.newViolation(bytesDescs.ruleDesc, wellKnown.desc,
		wellKnown.ruleID, wellKnown.mainMsg,
		val, protoreflect.ValueOfBool(true))
}

func (n nativeBytesEval) Tautology() bool {
	return false
}

// formatBytesList formats a [][]byte to match CEL's list formatting.
func formatBytesList(vals [][]byte) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		// this is what CEL does for a byte slice; displays it as a string
		parts[i] = string(v)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
