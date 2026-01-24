// Copyright 2023-2025 Buf Technologies, Inc.
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

package native

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buildStringRulesFromRules builds native evaluators for string rules.
func buildStringRulesFromRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.StringRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// const
	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &stringConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "string.const", fmt.Sprintf("value must equal %q", constVal)),
		})
	}

	// len
	if rules.HasLen() {
		lenVal := rules.GetLen()
		evaluators = append(evaluators, &stringLenEval{
			length:   lenVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "len", "string.len", fmt.Sprintf("value length must be %d characters", lenVal)),
		})
	}

	// min_len
	if rules.HasMinLen() {
		minLen := rules.GetMinLen()
		evaluators = append(evaluators, &stringMinLenEval{
			minLen:   minLen,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "min_len", "string.min_len", fmt.Sprintf("value length must be at least %d characters", minLen)),
		})
	}

	// max_len
	if rules.HasMaxLen() {
		maxLen := rules.GetMaxLen()
		evaluators = append(evaluators, &stringMaxLenEval{
			maxLen:   maxLen,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "max_len", "string.max_len", fmt.Sprintf("value length must be at most %d characters", maxLen)),
		})
	}

	// len_bytes
	if rules.HasLenBytes() {
		lenBytes := rules.GetLenBytes()
		evaluators = append(evaluators, &stringLenBytesEval{
			length:   lenBytes,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "len_bytes", "string.len_bytes", fmt.Sprintf("value length must be %d bytes", lenBytes)),
		})
	}

	// min_bytes
	if rules.HasMinBytes() {
		minBytes := rules.GetMinBytes()
		evaluators = append(evaluators, &stringMinBytesEval{
			minBytes: minBytes,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "min_bytes", "string.min_bytes", fmt.Sprintf("value length must be at least %d bytes", minBytes)),
		})
	}

	// max_bytes
	if rules.HasMaxBytes() {
		maxBytes := rules.GetMaxBytes()
		evaluators = append(evaluators, &stringMaxBytesEval{
			maxBytes: maxBytes,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "max_bytes", "string.max_bytes", fmt.Sprintf("value length must be at most %d bytes", maxBytes)),
		})
	}

	// prefix
	if rules.HasPrefix() {
		prefix := rules.GetPrefix()
		evaluators = append(evaluators, &stringPrefixEval{
			prefix:   prefix,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "prefix", "string.prefix", fmt.Sprintf("value does not have prefix %q", prefix)),
		})
	}

	// suffix
	if rules.HasSuffix() {
		suffix := rules.GetSuffix()
		evaluators = append(evaluators, &stringSuffixEval{
			suffix:   suffix,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "suffix", "string.suffix", fmt.Sprintf("value does not have suffix %q", suffix)),
		})
	}

	// contains
	if rules.HasContains() {
		contains := rules.GetContains()
		evaluators = append(evaluators, &stringContainsEval{
			substr:   contains,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "contains", "string.contains", fmt.Sprintf("value does not contain %q", contains)),
		})
	}

	// not_contains
	if rules.HasNotContains() {
		notContains := rules.GetNotContains()
		evaluators = append(evaluators, &stringNotContainsEval{
			substr:   notContains,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_contains", "string.not_contains", fmt.Sprintf("value contains %q", notContains)),
		})
	}

	// in
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[string]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &stringInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "string.in", "value must be in list"),
		})
	}

	// not_in
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[string]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &stringNotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "string.not_in", "value must not be in list"),
		})
	}

	return evaluators
}

// String evaluators

type stringConstEval struct {
	value    string
	ruleInfo RuleInfo
}

func (e *stringConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.String() != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringConstEval) Tautology() bool { return false }

func (e *stringConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringLenEval struct {
	length   uint64
	ruleInfo RuleInfo
}

func (e *stringLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(utf8.RuneCountInString(val.String())) != e.length {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringLenEval) Tautology() bool { return false }

func (e *stringLenEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type stringMinLenEval struct {
	minLen   uint64
	ruleInfo RuleInfo
}

func (e *stringMinLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(utf8.RuneCountInString(val.String())) < e.minLen {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringMinLenEval) Tautology() bool { return e.minLen == 0 }

func (e *stringMinLenEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringMaxLenEval struct {
	maxLen   uint64
	ruleInfo RuleInfo
}

func (e *stringMaxLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(utf8.RuneCountInString(val.String())) > e.maxLen {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringMaxLenEval) Tautology() bool { return false }

func (e *stringMaxLenEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringLenBytesEval struct {
	length   uint64
	ruleInfo RuleInfo
}

func (e *stringLenBytesEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.String())) != e.length {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringLenBytesEval) Tautology() bool { return false }

func (e *stringLenBytesEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringMinBytesEval struct {
	minBytes uint64
	ruleInfo RuleInfo
}

func (e *stringMinBytesEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.String())) < e.minBytes {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringMinBytesEval) Tautology() bool { return e.minBytes == 0 }

func (e *stringMinBytesEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringMaxBytesEval struct {
	maxBytes uint64
	ruleInfo RuleInfo
}

func (e *stringMaxBytesEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.String())) > e.maxBytes {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringMaxBytesEval) Tautology() bool { return false }

func (e *stringMaxBytesEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringPrefixEval struct {
	prefix   string
	ruleInfo RuleInfo
}

func (e *stringPrefixEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !strings.HasPrefix(val.String(), e.prefix) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringPrefixEval) Tautology() bool { return e.prefix == "" }

func (e *stringPrefixEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringSuffixEval struct {
	suffix   string
	ruleInfo RuleInfo
}

func (e *stringSuffixEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !strings.HasSuffix(val.String(), e.suffix) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringSuffixEval) Tautology() bool { return e.suffix == "" }

func (e *stringSuffixEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringContainsEval struct {
	substr   string
	ruleInfo RuleInfo
}

func (e *stringContainsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !strings.Contains(val.String(), e.substr) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringContainsEval) Tautology() bool { return e.substr == "" }

func (e *stringContainsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringNotContainsEval struct {
	substr   string
	ruleInfo RuleInfo
}

func (e *stringNotContainsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if strings.Contains(val.String(), e.substr) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringNotContainsEval) Tautology() bool { return false }

func (e *stringNotContainsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type stringInEval struct {
	set      map[string]struct{}
	ruleInfo RuleInfo
}

func (e *stringInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.String()]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringInEval) Tautology() bool { return false }

func (e *stringInEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type stringNotInEval struct {
	set      map[string]struct{}
	ruleInfo RuleInfo
}

func (e *stringNotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.String()]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *stringNotInEval) Tautology() bool { return len(e.set) == 0 }

func (e *stringNotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}
