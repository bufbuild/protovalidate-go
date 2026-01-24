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

// Package native provides native Go evaluators for standard validation rules.
// These evaluators bypass the CEL interpreter for common operations, reducing
// memory allocations and improving performance.
package native

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Evaluator performs validation directly on protoreflect.Value without CEL.
// Implementations are expected to be immutable and safe for concurrent use.
type Evaluator interface {
	// Eval validates the value, returning violations if invalid.
	// If failFast is true, evaluation stops at the first violation.
	Eval(val protoreflect.Value, failFast bool) ([]*validate.Violation, error)

	// Tautology returns true if this evaluator always succeeds.
	Tautology() bool

	// RuleDescriptor returns the proto field descriptor for the rule this
	// evaluator handles. This is used to skip CEL compilation for the same rule.
	RuleDescriptor() protoreflect.FieldDescriptor
}

// Evaluators is a slice of Evaluator that can be evaluated together.
type Evaluators []Evaluator

// Eval evaluates all evaluators and merges violations.
func (e Evaluators) Eval(val protoreflect.Value, failFast bool) ([]*validate.Violation, error) {
	var violations []*validate.Violation
	for _, eval := range e {
		vs, err := eval.Eval(val, failFast)
		if err != nil {
			return violations, err
		}
		violations = append(violations, vs...)
		if failFast && len(violations) > 0 {
			return violations, nil
		}
	}
	return violations, nil
}

// Tautology returns true if all evaluators are tautologies.
func (e Evaluators) Tautology() bool {
	for _, eval := range e {
		if !eval.Tautology() {
			return false
		}
	}
	return true
}

// HandledRules returns a set of rule field descriptors that these evaluators handle.
// This is used to skip CEL compilation for rules that are handled natively.
func (e Evaluators) HandledRules() map[protoreflect.FieldDescriptor]struct{} {
	if len(e) == 0 {
		return nil
	}
	result := make(map[protoreflect.FieldDescriptor]struct{}, len(e))
	for _, eval := range e {
		if desc := eval.RuleDescriptor(); desc != nil {
			result[desc] = struct{}{}
		}
	}
	return result
}

// RuleInfo contains metadata about a rule for violation reporting.
type RuleInfo struct {
	// RuleID is the ID of the rule (e.g., "int32.gt").
	RuleID string
	// Message is the default violation message.
	Message string
	// RulePath is the path to the rule in the field rules.
	RulePath *validate.FieldPath
	// RuleValue is the value of the rule for violation reporting.
	RuleValue protoreflect.Value
	// RuleDescriptor is the descriptor of the rule field.
	RuleDescriptor protoreflect.FieldDescriptor
}

// NewViolation creates a new violation from rule info.
func (r RuleInfo) NewViolation(message string) *validate.Violation {
	if message == "" {
		message = r.Message
	}
	return validate.Violation_builder{
		Rule:    r.RulePath,
		RuleId:  proto.String(r.RuleID),
		Message: proto.String(message),
	}.Build()
}

// Descriptor returns the rule field descriptor.
func (r RuleInfo) Descriptor() protoreflect.FieldDescriptor {
	return r.RuleDescriptor
}
