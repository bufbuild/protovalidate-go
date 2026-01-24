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
	"math"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buildFloatRules builds native evaluators for float32 rules.
//
//nolint:dupl // float32 and float64 rules are structurally similar but use different types
func buildFloatRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FloatRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// const
	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &floatConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "float.const", fmt.Sprintf("value must equal %v", constVal)),
		})
	}

	// lt
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &floatLtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "float.lt", fmt.Sprintf("value must be less than %v", ltVal)),
		})
	}

	// lte
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &floatLteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "float.lte", fmt.Sprintf("value must be less than or equal to %v", lteVal)),
		})
	}

	// gt
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &floatGtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "float.gt", fmt.Sprintf("value must be greater than %v", gtVal)),
		})
	}

	// gte
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &floatGteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "float.gte", fmt.Sprintf("value must be greater than or equal to %v", gteVal)),
		})
	}

	// in
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[float32]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &floatInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "float.in", "value must be in list"),
		})
	}

	// not_in
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[float32]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &floatNotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "float.not_in", "value must not be in list"),
		})
	}

	// finite
	if rules.GetFinite() {
		evaluators = append(evaluators, &floatFiniteEval{
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "finite", "float.finite", "value must be finite"),
		})
	}

	return evaluators
}

// buildDoubleRules builds native evaluators for float64 rules.
//
//nolint:dupl // float32 and float64 rules are structurally similar but use different types
func buildDoubleRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.DoubleRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// const
	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &doubleConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "double.const", fmt.Sprintf("value must equal %v", constVal)),
		})
	}

	// lt
	if rules.HasLt() {
		ltVal := rules.GetLt()
		evaluators = append(evaluators, &doubleLtEval{
			threshold: ltVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lt", "double.lt", fmt.Sprintf("value must be less than %v", ltVal)),
		})
	}

	// lte
	if rules.HasLte() {
		lteVal := rules.GetLte()
		evaluators = append(evaluators, &doubleLteEval{
			threshold: lteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "lte", "double.lte", fmt.Sprintf("value must be less than or equal to %v", lteVal)),
		})
	}

	// gt
	if rules.HasGt() {
		gtVal := rules.GetGt()
		evaluators = append(evaluators, &doubleGtEval{
			threshold: gtVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gt", "double.gt", fmt.Sprintf("value must be greater than %v", gtVal)),
		})
	}

	// gte
	if rules.HasGte() {
		gteVal := rules.GetGte()
		evaluators = append(evaluators, &doubleGteEval{
			threshold: gteVal,
			ruleInfo:  makeRuleInfo(fieldDesc, rulesDesc, "gte", "double.gte", fmt.Sprintf("value must be greater than or equal to %v", gteVal)),
		})
	}

	// in
	if inVals := rules.GetIn(); len(inVals) > 0 {
		set := make(map[float64]struct{}, len(inVals))
		for _, v := range inVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &doubleInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "in", "double.in", "value must be in list"),
		})
	}

	// not_in
	if notInVals := rules.GetNotIn(); len(notInVals) > 0 {
		set := make(map[float64]struct{}, len(notInVals))
		for _, v := range notInVals {
			set[v] = struct{}{}
		}
		evaluators = append(evaluators, &doubleNotInEval{
			set:      set,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "not_in", "double.not_in", "value must not be in list"),
		})
	}

	// finite
	if rules.GetFinite() {
		evaluators = append(evaluators, &doubleFiniteEval{
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "finite", "double.finite", "value must be finite"),
		})
	}

	return evaluators
}

// Float evaluators

type floatConstEval struct {
	value    float32
	ruleInfo RuleInfo
}

func (e *floatConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if float32(val.Float()) != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatConstEval) Tautology() bool { return false }

func (e *floatConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type floatLtEval struct {
	threshold float32
	ruleInfo  RuleInfo
}

func (e *floatLtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if float32(val.Float()) >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatLtEval) Tautology() bool { return false }

func (e *floatLtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type floatLteEval struct {
	threshold float32
	ruleInfo  RuleInfo
}

func (e *floatLteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if float32(val.Float()) > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatLteEval) Tautology() bool { return false }

func (e *floatLteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type floatGtEval struct {
	threshold float32
	ruleInfo  RuleInfo
}

func (e *floatGtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if float32(val.Float()) <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatGtEval) Tautology() bool { return false }

func (e *floatGtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type floatGteEval struct {
	threshold float32
	ruleInfo  RuleInfo
}

func (e *floatGteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if float32(val.Float()) < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatGteEval) Tautology() bool { return false }

func (e *floatGteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type floatInEval struct {
	set      map[float32]struct{}
	ruleInfo RuleInfo
}

func (e *floatInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[float32(val.Float())]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatInEval) Tautology() bool { return false }

func (e *floatInEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type floatNotInEval struct {
	set      map[float32]struct{}
	ruleInfo RuleInfo
}

func (e *floatNotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[float32(val.Float())]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatNotInEval) Tautology() bool { return false }

func (e *floatNotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type floatFiniteEval struct {
	ruleInfo RuleInfo
}

func (e *floatFiniteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	f := float32(val.Float())
	if math.IsInf(float64(f), 0) || math.IsNaN(float64(f)) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *floatFiniteEval) Tautology() bool { return false }

func (e *floatFiniteEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

// Double evaluators

type doubleConstEval struct {
	value    float64
	ruleInfo RuleInfo
}

func (e *doubleConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Float() != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleConstEval) Tautology() bool { return false }

func (e *doubleConstEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type doubleLtEval struct {
	threshold float64
	ruleInfo  RuleInfo
}

func (e *doubleLtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Float() >= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleLtEval) Tautology() bool { return false }

func (e *doubleLtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type doubleLteEval struct {
	threshold float64
	ruleInfo  RuleInfo
}

func (e *doubleLteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Float() > e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleLteEval) Tautology() bool { return false }

func (e *doubleLteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type doubleGtEval struct {
	threshold float64
	ruleInfo  RuleInfo
}

func (e *doubleGtEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Float() <= e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleGtEval) Tautology() bool { return false }

func (e *doubleGtEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type doubleGteEval struct {
	threshold float64
	ruleInfo  RuleInfo
}

func (e *doubleGteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Float() < e.threshold {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleGteEval) Tautology() bool { return false }

func (e *doubleGteEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type doubleInEval struct {
	set      map[float64]struct{}
	ruleInfo RuleInfo
}

func (e *doubleInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Float()]; !ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleInEval) Tautology() bool { return false }

func (e *doubleInEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }

type doubleNotInEval struct {
	set      map[float64]struct{}
	ruleInfo RuleInfo
}

func (e *doubleNotInEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if _, ok := e.set[val.Float()]; ok {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleNotInEval) Tautology() bool { return false }

func (e *doubleNotInEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}

type doubleFiniteEval struct {
	ruleInfo RuleInfo
}

func (e *doubleFiniteEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	f := val.Float()
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *doubleFiniteEval) Tautology() bool { return false }

func (e *doubleFiniteEval) RuleDescriptor() protoreflect.FieldDescriptor {
	return e.ruleInfo.Descriptor()
}
