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

package protovalidate

import (
	"buf.build/go/protovalidate/internal/native"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// nativeEval wraps native.Evaluators to implement the evaluator interface.
type nativeEval struct {
	base
	evaluators native.Evaluators
}

func (n nativeEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	violations, err := n.evaluators.Eval(val, cfg.failFast)
	if err != nil {
		return &RuntimeError{cause: err}
	}
	if len(violations) == 0 {
		return nil
	}

	// Convert native violations to ValidationError
	validationErr := &ValidationError{
		Violations: make([]*Violation, 0, len(violations)),
	}
	for _, v := range violations {
		violation := &Violation{
			Proto:           v,
			FieldValue:      val,
			FieldDescriptor: n.Descriptor,
		}
		violation.Proto.SetField(n.fieldPath())
		violation.Proto.SetRule(n.rulePath(violation.Proto.GetRule()))
		validationErr.Violations = append(validationErr.Violations, violation)
		if cfg.failFast {
			break
		}
	}
	return validationErr
}

func (n nativeEval) Tautology() bool {
	return n.evaluators.Tautology()
}

var _ evaluator = nativeEval{}
