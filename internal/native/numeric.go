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

// buildSignedRules builds native evaluators for signed integer rules (int32).
func buildSignedRules[T int32 | int64](
	fieldDesc protoreflect.FieldDescriptor,
	rules any,
	_ string,
) Evaluators {
	if rules == nil {
		return nil
	}

	// Type switch to handle different rule types
	switch typedRules := rules.(type) {
	case *validate.Int32Rules:
		return buildInt32RulesImpl(fieldDesc, typedRules)
	case *validate.Int64Rules:
		return buildInt64RulesImpl(fieldDesc, typedRules)
	case *validate.SInt32Rules:
		return buildSInt32RulesImpl(fieldDesc, typedRules)
	case *validate.SInt64Rules:
		return buildSInt64RulesImpl(fieldDesc, typedRules)
	case *validate.SFixed32Rules:
		return buildSFixed32RulesImpl(fieldDesc, typedRules)
	case *validate.SFixed64Rules:
		return buildSFixed64RulesImpl(fieldDesc, typedRules)
	}
	return nil
}

// buildUnsignedRules builds native evaluators for unsigned integer rules.
func buildUnsignedRules[T uint32 | uint64](
	fieldDesc protoreflect.FieldDescriptor,
	rules any,
	_ string,
) Evaluators {
	if rules == nil {
		return nil
	}

	switch typedRules := rules.(type) {
	case *validate.UInt32Rules:
		return buildUInt32RulesImpl(fieldDesc, typedRules)
	case *validate.UInt64Rules:
		return buildUInt64RulesImpl(fieldDesc, typedRules)
	case *validate.Fixed32Rules:
		return buildFixed32RulesImpl(fieldDesc, typedRules)
	case *validate.Fixed64Rules:
		return buildFixed64RulesImpl(fieldDesc, typedRules)
	}
	return nil
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildInt32RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.Int32Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int32ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "int32.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int32LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "int32.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int32LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "int32.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int32GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "int32.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int32GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "int32.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "int32.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "int32.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildInt64RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.Int64Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int64ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "int64.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int64LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "int64.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int64LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "int64.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int64GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "int64.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int64GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "int64.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "int64.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "int64.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildSInt32RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.SInt32Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int32ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "sint32.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int32LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "sint32.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int32LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "sint32.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int32GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "sint32.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int32GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "sint32.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "sint32.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "sint32.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildSInt64RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.SInt64Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int64ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "sint64.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int64LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "sint64.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int64LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "sint64.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int64GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "sint64.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int64GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "sint64.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "sint64.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "sint64.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildSFixed32RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.SFixed32Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int32ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "sfixed32.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int32LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "sfixed32.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int32LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "sfixed32.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int32GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "sfixed32.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int32GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "sfixed32.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "sfixed32.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int32NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "sfixed32.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildSFixed64RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.SFixed64Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &int64ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "sfixed64.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &int64LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "sfixed64.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &int64LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "sfixed64.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &int64GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "sfixed64.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &int64GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "sfixed64.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[int64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "sfixed64.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[int64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &int64NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "sfixed64.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildUInt32RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.UInt32Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &uint32ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "uint32.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &uint32LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "uint32.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &uint32LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "uint32.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &uint32GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "uint32.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &uint32GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "uint32.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[uint32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint32InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "uint32.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[uint32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint32NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "uint32.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildUInt64RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.UInt64Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &uint64ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "uint64.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &uint64LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "uint64.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &uint64LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "uint64.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &uint64GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "uint64.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &uint64GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "uint64.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[uint64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint64InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "uint64.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[uint64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint64NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "uint64.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildFixed32RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.Fixed32Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &uint32ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "fixed32.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &uint32LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "fixed32.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &uint32LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "fixed32.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &uint32GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "fixed32.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &uint32GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "fixed32.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[uint32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint32InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "fixed32.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[uint32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint32NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "fixed32.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

//nolint:dupl // integer rule implementations share structure but differ by type
func buildFixed64RulesImpl(fieldDesc protoreflect.FieldDescriptor, rules *validate.Fixed64Rules) Evaluators {
	if rules == nil {
		return nil
	}
	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &uint64ConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "fixed64.const", fmt.Sprintf("value must equal %d", constVal)),
		})
	}
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &uint64LtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "fixed64.lt", fmt.Sprintf("value must be less than %d", ltVal)),
		})
	}
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &uint64LteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "fixed64.lte", fmt.Sprintf("value must be less than or equal to %d", lteVal)),
		})
	}
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &uint64GtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "fixed64.gt", fmt.Sprintf("value must be greater than %d", gtVal)),
		})
	}
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &uint64GteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "fixed64.gte", fmt.Sprintf("value must be greater than or equal to %d", gteVal)),
		})
	}
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[uint64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint64InEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "fixed64.in", "value must be in list"),
		})
	}
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[uint64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &uint64NotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "fixed64.not_in", "value must not be in list"),
		})
	}
	return evaluators
}

// makeRuleInfo creates a RuleInfo from field and rule descriptors.
func makeRuleInfo(
	fieldDesc protoreflect.FieldDescriptor,
	rulesDesc protoreflect.MessageDescriptor,
	ruleName string,
	ruleID string,
	message string,
) RuleInfo {
	ruleFieldDesc := rulesDesc.Fields().ByName(protoreflect.Name(ruleName))
	return RuleInfo{
		RuleID:         ruleID,
		Message:        message,
		RulePath:       makeRulePath(fieldDesc, rulesDesc, ruleFieldDesc),
		RuleDescriptor: ruleFieldDesc,
	}
}

// makeRulePath creates a FieldPath for a rule.
func makeRulePath(
	_ protoreflect.FieldDescriptor,
	rulesDesc protoreflect.MessageDescriptor,
	ruleFieldDesc protoreflect.FieldDescriptor,
) *validate.FieldPath {
	// Find the actual field descriptor for this type in FieldRules
	fieldRulesDesc := (*validate.FieldRules)(nil).ProtoReflect().Descriptor()
	var typeFieldDesc protoreflect.FieldDescriptor
	typeOneof := fieldRulesDesc.Oneofs().ByName("type")
	if typeOneof != nil {
		for i := range typeOneof.Fields().Len() {
			field := typeOneof.Fields().Get(i)
			if field.Message() != nil && field.Message().FullName() == rulesDesc.FullName() {
				typeFieldDesc = field
				break
			}
		}
	}

	elements := []*validate.FieldPathElement{}
	if typeFieldDesc != nil {
		elements = append(elements, validate.FieldPathElement_builder{
			FieldNumber: proto.Int32(int32(typeFieldDesc.Number())),
			FieldName:   proto.String(string(typeFieldDesc.Name())),
		}.Build())
	}
	if ruleFieldDesc != nil {
		elements = append(elements, validate.FieldPathElement_builder{
			FieldNumber: proto.Int32(int32(ruleFieldDesc.Number())),
			FieldName:   proto.String(string(ruleFieldDesc.Name())),
		}.Build())
	}

	return validate.FieldPath_builder{Elements: elements}.Build()
}

// Int32 evaluators

type int32ConstEval struct {
	value    int32
	ruleInfo RuleInfo
}

func (e *int32ConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if int32(val.Int()) != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32ConstEval) Tautology() bool { return false }

func (e *int32ConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type int32LtEval struct {
	threshold int32
	ruleInfo  RuleInfo
}

func (e *int32LtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if int32(val.Int()) >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32LtEval) Tautology() bool { return false }

func (e *int32LtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int32LteEval struct {
	threshold int32
	ruleInfo  RuleInfo
}

func (e *int32LteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if int32(val.Int()) > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32LteEval) Tautology() bool { return false }

func (e *int32LteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int32GtEval struct {
	threshold int32
	ruleInfo  RuleInfo
}

func (e *int32GtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if int32(val.Int()) <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32GtEval) Tautology() bool { return false }

func (e *int32GtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int32GteEval struct {
	threshold int32
	ruleInfo  RuleInfo
}

func (e *int32GteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if int32(val.Int()) < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32GteEval) Tautology() bool { return false }

func (e *int32GteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int32InEval struct {
	set      map[int32]struct{}
	ruleInfo RuleInfo
}

func (e *int32InEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[int32(val.Int())]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32InEval) Tautology() bool { return false }

func (e *int32InEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int32NotInEval struct {
	set      map[int32]struct{}
	ruleInfo RuleInfo
}

func (e *int32NotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[int32(val.Int())]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int32NotInEval) Tautology() bool { return false }

func (e *int32NotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

// Int64 evaluators

type int64ConstEval struct {
	value    int64
	ruleInfo RuleInfo
}

func (e *int64ConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Int() != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64ConstEval) Tautology() bool { return false }

func (e *int64ConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type int64LtEval struct {
	threshold int64
	ruleInfo  RuleInfo
}

func (e *int64LtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Int() >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64LtEval) Tautology() bool { return false }

func (e *int64LtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int64LteEval struct {
	threshold int64
	ruleInfo  RuleInfo
}

func (e *int64LteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Int() > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64LteEval) Tautology() bool { return false }

func (e *int64LteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int64GtEval struct {
	threshold int64
	ruleInfo  RuleInfo
}

func (e *int64GtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Int() <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64GtEval) Tautology() bool { return false }

func (e *int64GtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int64GteEval struct {
	threshold int64
	ruleInfo  RuleInfo
}

func (e *int64GteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Int() < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64GteEval) Tautology() bool { return false }

func (e *int64GteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int64InEval struct {
	set      map[int64]struct{}
	ruleInfo RuleInfo
}

func (e *int64InEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Int()]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64InEval) Tautology() bool { return false }

func (e *int64InEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type int64NotInEval struct {
	set      map[int64]struct{}
	ruleInfo RuleInfo
}

func (e *int64NotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Int()]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *int64NotInEval) Tautology() bool { return false }

func (e *int64NotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

// Uint32 evaluators

type uint32ConstEval struct {
	value    uint32
	ruleInfo RuleInfo
}

func (e *uint32ConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint32(val.Uint()) != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32ConstEval) Tautology() bool { return false }

func (e *uint32ConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type uint32LtEval struct {
	threshold uint32
	ruleInfo  RuleInfo
}

func (e *uint32LtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint32(val.Uint()) >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32LtEval) Tautology() bool { return false }

func (e *uint32LtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint32LteEval struct {
	threshold uint32
	ruleInfo  RuleInfo
}

func (e *uint32LteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint32(val.Uint()) > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32LteEval) Tautology() bool { return false }

func (e *uint32LteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint32GtEval struct {
	threshold uint32
	ruleInfo  RuleInfo
}

func (e *uint32GtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint32(val.Uint()) <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32GtEval) Tautology() bool { return false }

func (e *uint32GtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint32GteEval struct {
	threshold uint32
	ruleInfo  RuleInfo
}

func (e *uint32GteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if uint32(val.Uint()) < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32GteEval) Tautology() bool { return false }

func (e *uint32GteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint32InEval struct {
	set      map[uint32]struct{}
	ruleInfo RuleInfo
}

func (e *uint32InEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[uint32(val.Uint())]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32InEval) Tautology() bool { return false }

func (e *uint32InEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint32NotInEval struct {
	set      map[uint32]struct{}
	ruleInfo RuleInfo
}

func (e *uint32NotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[uint32(val.Uint())]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint32NotInEval) Tautology() bool { return false }

func (e *uint32NotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

// Uint64 evaluators

type uint64ConstEval struct {
	value    uint64
	ruleInfo RuleInfo
}

func (e *uint64ConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Uint() != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64ConstEval) Tautology() bool { return false }

func (e *uint64ConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type uint64LtEval struct {
	threshold uint64
	ruleInfo  RuleInfo
}

func (e *uint64LtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Uint() >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64LtEval) Tautology() bool { return false }

func (e *uint64LtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint64LteEval struct {
	threshold uint64
	ruleInfo  RuleInfo
}

func (e *uint64LteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Uint() > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64LteEval) Tautology() bool { return false }

func (e *uint64LteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint64GtEval struct {
	threshold uint64
	ruleInfo  RuleInfo
}

func (e *uint64GtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Uint() <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64GtEval) Tautology() bool { return false }

func (e *uint64GtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint64GteEval struct {
	threshold uint64
	ruleInfo  RuleInfo
}

func (e *uint64GteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Uint() < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64GteEval) Tautology() bool { return false }

func (e *uint64GteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint64InEval struct {
	set      map[uint64]struct{}
	ruleInfo RuleInfo
}

func (e *uint64InEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Uint()]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64InEval) Tautology() bool { return false }

func (e *uint64InEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type uint64NotInEval struct {
	set      map[uint64]struct{}
	ruleInfo RuleInfo
}

func (e *uint64NotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Uint()]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *uint64NotInEval) Tautology() bool { return false }

func (e *uint64NotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}
