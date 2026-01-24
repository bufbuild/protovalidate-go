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
	"bytes"
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buildBytesRulesFromRules builds native evaluators for bytes rules.
func buildBytesRulesFromRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.BytesRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// const
	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &bytesConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "bytes.const", fmt.Sprintf("value must equal %x", constVal)),
		})
	}

	// len
	if rules.HasLen() {
		lenVal := rules.GetLen()
		evaluators = append(evaluators, &bytesLenEval{
			length:   lenVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "len", "bytes.len", fmt.Sprintf("value length must be %d bytes", lenVal)),
		})
	}

	// min_len
	if rules.HasMinLen() {
		minLen := rules.GetMinLen()
		evaluators = append(evaluators, &bytesMinLenEval{
			minLen:   minLen,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "min_len", "bytes.min_len", fmt.Sprintf("value length must be at least %d bytes", minLen)),
		})
	}

	// max_len
	if rules.HasMaxLen() {
		maxLen := rules.GetMaxLen()
		evaluators = append(evaluators, &bytesMaxLenEval{
			maxLen:   maxLen,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "max_len", "bytes.max_len", fmt.Sprintf("value length must be at most %d bytes", maxLen)),
		})
	}

	// prefix
	if rules.HasPrefix() {
		prefix := rules.GetPrefix()
		evaluators = append(evaluators, &bytesPrefixEval{
			prefix:   prefix,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "prefix", "bytes.prefix", fmt.Sprintf("value does not have prefix %x", prefix)),
		})
	}

	// suffix
	if rules.HasSuffix() {
		suffix := rules.GetSuffix()
		evaluators = append(evaluators, &bytesSuffixEval{
			suffix:   suffix,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "suffix", "bytes.suffix", fmt.Sprintf("value does not have suffix %x", suffix)),
		})
	}

	// contains
	if rules.HasContains() {
		contains := rules.GetContains()
		evaluators = append(evaluators, &bytesContainsEval{
			substr:   contains,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "contains", "bytes.contains", fmt.Sprintf("value does not contain %x", contains)),
		})
	}

	// in
	if inVals := rules.GetIn(); len(inVals) > 0 {
		evaluators = append(evaluators, &bytesInEval{
			set:      inVals,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "bytes.in", "value must be in list"),
		})
	}

	// not_in
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		evaluators = append(evaluators, &bytesNotInEval{
			set:      notInVals,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "bytes.not_in", "value must not be in list"),
		})
	}

	return evaluators
}

// Bytes evaluators

type bytesConstEval struct {
	value    []byte
	ruleInfo RuleInfo
}

func (e *bytesConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !bytes.Equal(val.Bytes(), e.value) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesConstEval) Tautology() bool { return false }

func (e *bytesConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesLenEval struct {
	length   uint64
	ruleInfo RuleInfo
}

func (e *bytesLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.Bytes())) != e.length {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesLenEval) Tautology() bool { return false }

func (e *bytesLenEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type bytesMinLenEval struct {
	minLen   uint64
	ruleInfo RuleInfo
}

func (e *bytesMinLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.Bytes())) < e.minLen {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesMinLenEval) Tautology() bool { return e.minLen == 0 }

func (e *bytesMinLenEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesMaxLenEval struct {
	maxLen   uint64
	ruleInfo RuleInfo
}

func (e *bytesMaxLenEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(len(val.Bytes())) > e.maxLen {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesMaxLenEval) Tautology() bool { return false }

func (e *bytesMaxLenEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesPrefixEval struct {
	prefix   []byte
	ruleInfo RuleInfo
}

func (e *bytesPrefixEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !bytes.HasPrefix(val.Bytes(), e.prefix) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesPrefixEval) Tautology() bool { return len(e.prefix) == 0 }

func (e *bytesPrefixEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesSuffixEval struct {
	suffix   []byte
	ruleInfo RuleInfo
}

func (e *bytesSuffixEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !bytes.HasSuffix(val.Bytes(), e.suffix) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesSuffixEval) Tautology() bool { return len(e.suffix) == 0 }

func (e *bytesSuffixEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesContainsEval struct {
	substr   []byte
	ruleInfo RuleInfo
}

func (e *bytesContainsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if !bytes.Contains(val.Bytes(), e.substr) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *bytesContainsEval) Tautology() bool { return len(e.substr) == 0 }

func (e *bytesContainsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type bytesInEval struct {
	set      [][]byte
	ruleInfo RuleInfo
}

func (e *bytesInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	v := val.Bytes()
	for _, b := range e.set {
		if bytes.Equal(v, b) {
			return nil, nil
		}
	}
	return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
}

func (e *bytesInEval) Tautology() bool { return false }

func (e *bytesInEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type bytesNotInEval struct {
	set      [][]byte
	ruleInfo RuleInfo
}

func (e *bytesNotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	v := val.Bytes()
	for _, b := range e.set {
		if bytes.Equal(v, b) {
			return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
		}
	}
	return nil, nil
}

func (e *bytesNotInEval) Tautology() bool { return len(e.set) == 0 }

func (e *bytesNotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}
