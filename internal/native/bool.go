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
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buildBoolRules builds native evaluators for bool rules.
func buildBoolRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.BoolRules,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators
	rulesDesc := rules.ProtoReflect().Descriptor()

	// const
	if rules.HasConst() {
		constVal := rules.GetConst()
		evaluators = append(evaluators, &boolConstEval{
			value:    constVal,
			ruleInfo: makeRuleInfo(fieldDesc, rulesDesc, "const", "bool.const", fmt.Sprintf("value must equal %v", constVal)),
		})
	}

	return evaluators
}

// Bool evaluators

type boolConstEval struct {
	value    bool
	ruleInfo RuleInfo
}

func (e *boolConstEval) Eval(val protoreflect.Value, _ bool) ([]*validate.Violation, error) {
	if val.Bool() != e.value {
		return []*validate.Violation{e.ruleInfo.NewViolation("")}, nil
	}
	return nil, nil
}

func (e *boolConstEval) Tautology() bool { return false }

func (e *boolConstEval) RuleDescriptor() protoreflect.FieldDescriptor { return e.ruleInfo.Descriptor() }
