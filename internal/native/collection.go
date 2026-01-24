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

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buildMapSizeRules builds native evaluators for map size rules.
func buildMapSizeRules(
	_ protoreflect.FieldDescriptor,
	rules *validate.MapRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// min_pairs
	if rules.HasMinPairs() {
		minPairs := rules.GetMinPairs()
		evaluators = append(evaluators, &mapMinPairsEval{
			minPairs: minPairs,
			ruleInfo: makeCollectionRuleInfo(rulesDesc, "map", "min_pairs", "map.min_pairs", fmt.Sprintf("map must contain at least %d entries", minPairs)),
		})
	}

	// max_pairs
	if rules.HasMaxPairs() {
		maxPairs := rules.GetMaxPairs()
		evaluators = append(evaluators, &mapMaxPairsEval{
			maxPairs: maxPairs,
			ruleInfo: makeCollectionRuleInfo(rulesDesc, "map", "max_pairs", "map.max_pairs", fmt.Sprintf("map must contain at most %d entries", maxPairs)),
		})
	}

	return evaluators
}

// buildRepeatedSizeRules builds native evaluators for repeated field size rules.
func buildRepeatedSizeRules(
	_ protoreflect.FieldDescriptor,
	rules *validate.RepeatedRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// min_items
	if rules.HasMinItems() {
		minItems := rules.GetMinItems()
		evaluators = append(evaluators, &repeatedMinItemsEval{
			minItems: minItems,
			ruleInfo: makeCollectionRuleInfo(rulesDesc, "repeated", "min_items", "repeated.min_items", fmt.Sprintf("repeated field must contain at least %d items", minItems)),
		})
	}

	// max_items
	if rules.HasMaxItems() {
		maxItems := rules.GetMaxItems()
		evaluators = append(evaluators, &repeatedMaxItemsEval{
			maxItems: maxItems,
			ruleInfo: makeCollectionRuleInfo(rulesDesc, "repeated", "max_items", "repeated.max_items", fmt.Sprintf("repeated field must contain at most %d items", maxItems)),
		})
	}

	return evaluators
}

// makeCollectionRuleInfo creates a RuleInfo for collection rules (map or repeated).
func makeCollectionRuleInfo(
	rulesDesc protoreflect.MessageDescriptor,
	collectionType string,
	ruleName string,
	ruleID string,
	message string,
) RuleInfo {
	ruleFieldDesc := rulesDesc.Fields().ByName(protoreflect.Name(ruleName))
	fieldRulesDesc := (*validate.FieldRules)(nil).ProtoReflect().Descriptor()
	collectionFieldDesc := fieldRulesDesc.Fields().ByName(protoreflect.Name(collectionType))

	elements := []*validate.FieldPathElement{
		validate.FieldPathElement_builder{
			FieldNumber: proto.Int32(int32(collectionFieldDesc.Number())),
			FieldName:   proto.String(string(collectionFieldDesc.Name())),
		}.Build(),
	}
	if ruleFieldDesc != nil {
		elements = append(elements, validate.FieldPathElement_builder{
			FieldNumber: proto.Int32(int32(ruleFieldDesc.Number())),
			FieldName:   proto.String(string(ruleFieldDesc.Name())),
		}.Build())
	}

	return RuleInfo{
		RuleID:         ruleID,
		Message:        message,
		RulePath:       validate.FieldPath_builder{Elements: elements}.Build(),
		RuleDescriptor: ruleFieldDesc,
	}
}

// Map evaluators

type mapMinPairsEval struct {
	minPairs uint64
	ruleInfo RuleInfo
}

func (e *mapMinPairsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(val.Map().Len()) < e.minPairs {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *mapMinPairsEval) Tautology() bool { return e.minPairs == 0 }

func (e *mapMinPairsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type mapMaxPairsEval struct {
	maxPairs uint64
	ruleInfo RuleInfo
}

func (e *mapMaxPairsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(val.Map().Len()) > e.maxPairs {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *mapMaxPairsEval) Tautology() bool { return false }

func (e *mapMaxPairsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

// Repeated evaluators

type repeatedMinItemsEval struct {
	minItems uint64
	ruleInfo RuleInfo
}

func (e *repeatedMinItemsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(val.List().Len()) < e.minItems {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *repeatedMinItemsEval) Tautology() bool { return e.minItems == 0 }

func (e *repeatedMinItemsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type repeatedMaxItemsEval struct {
	maxItems uint64
	ruleInfo RuleInfo
}

func (e *repeatedMaxItemsEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint64(val.List().Len()) > e.maxItems {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *repeatedMaxItemsEval) Tautology() bool { return false }

func (e *repeatedMaxItemsEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}
