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
	"fmt"
	"slices"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	enumConstDesc = (*validate.EnumRules)(nil).ProtoReflect().Descriptor().Fields().ByName("const")
	enumInDesc    = (*validate.EnumRules)(nil).ProtoReflect().Descriptor().Fields().ByName("in")
	enumNotInDesc = (*validate.EnumRules)(nil).ProtoReflect().Descriptor().Fields().ByName("not_in")
)

// tryBuildNativeEnumRules attempts to build a native Go evaluator for
// enum const/in/not_in rules. Returns nil if the rules can't be handled
// natively. Note: defined_only is handled separately in enum.go.
func tryBuildNativeEnumRules(base base, rules *validate.EnumRules) evaluator {
	if rules == nil {
		return nil
	}
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	var constVal *int32
	if rules.HasConst() {
		constVal = ptr(rules.GetConst())
		hasRule = true
	}

	var inVals []int32
	if inVals = rules.GetIn(); len(inVals) > 0 {
		hasRule = true
	}

	var notInVals []int32
	if notInVals = rules.GetNotIn(); len(notInVals) > 0 {
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeEnumEval{
		base:      base,
		constVal:  constVal,
		inVals:    inVals,
		notInVals: notInVals,
	}
}

var _ evaluator = nativeEnumEval{}

// nativeEnumEval is a native Go evaluator for enum const/in/not_in rules.
type nativeEnumEval struct {
	base
	constVal  *int32
	inVals    []int32
	notInVals []int32
}

type enumProcessor func(n nativeEnumEval, val protoreflect.Value, enumVal int32) *Violation

//nolint:gochecknoglobals // slice of all the processors that are used, value never modified, effectively immutable
var enumProcessors = []enumProcessor{
	// const
	func(n nativeEnumEval, val protoreflect.Value, enumVal int32) *Violation {
		if n.constVal != nil && enumVal != *n.constVal {
			return n.newViolation(enumRuleDescriptor, enumConstDesc,
				"enum.const", fmt.Sprintf("must equal %d", *n.constVal),
				val, protoreflect.ValueOfInt32(*n.constVal))
		}
		return nil
	},
	// in
	func(n nativeEnumEval, val protoreflect.Value, enumVal int32) *Violation {
		if len(n.inVals) > 0 && !slices.Contains(n.inVals, enumVal) {
			return n.newViolation(enumRuleDescriptor, enumInDesc,
				"enum.in", "must be in list "+formatList(n.inVals),
				val, protoreflect.ValueOfInt32(enumVal))
		}
		return nil
	},
	// not_in
	func(n nativeEnumEval, val protoreflect.Value, enumVal int32) *Violation {
		if len(n.notInVals) > 0 && slices.Contains(n.notInVals, enumVal) {
			return n.newViolation(enumRuleDescriptor, enumNotInDesc,
				"enum.not_in", "must not be in list "+formatList(n.notInVals),
				val, protoreflect.ValueOfInt32(enumVal))
		}
		return nil
	},
}

func (n nativeEnumEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	enumVal := int32(val.Enum())

	var violations []*Violation

	for _, enumProcessor := range enumProcessors {
		violation := enumProcessor(n, val, enumVal)
		if violation != nil {
			violations = append(violations, violation)
			if cfg.failFast {
				break
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

func (n nativeEnumEval) Tautology() bool {
	return false
}
