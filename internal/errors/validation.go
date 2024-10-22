// Copyright 2023-2024 Buf Technologies, Inc.
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

package errors

import (
	"fmt"
	"slices"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Violation represents a single instance where a validation rule was not met.
// It provides information about the field that caused the violation, the
// specific unfulfilled constraint, and a human-readable error message.
type Violation struct {
	// FieldPath is a machine-readable identifier that points to the specific
	// field that failed the validation.  This could be a nested field, in which
	// case the path will include all the parent fields leading to the actual
	// field that caused the violation.
	FieldPath string

	// RulePath is a machine-readable identifier that points to the specific
	// constraint rule that failed validation. This will be a nested field
	// starting from the FieldConstraints of the field that failed validation.
	// This value is only present for standard or predefined rules on fields.
	RulePath string

	// FieldValue is the value of the specific field that failed validation.
	FieldValue protoreflect.Value

	// ForKey indicates whether the violation was caused by a map key, rather
	// than a value.
	ForKey bool

	// ConstraintID is the unique identifier of the constraint that was not
	// fulfilled.
	ConstraintID string

	// RuleValue is the value of the rule that specified the failed constraint.
	// Only constraints that have a corresponding rule are set (i.e.: standard
	// constraints and predefined constraints). In other cases, such as custom
	// constraints, this will be an invalid value.
	RuleValue protoreflect.Value

	// Message is a human-readable error message that describes the nature of
	// the violation. This can be the default error message from the violated
	// constraint, or it can be a custom message that gives more context about
	// the violation.
	Message string
}

// A ValidationError is returned if one or more constraint violations were
// detected.
type ValidationError struct {
	Violations []Violation
}

// FromProto converts the proto.Message form of the error back into native form.
func FromProto(
	registry protoregistry.ExtensionTypeResolver,
	message proto.Message,
	violations *validate.Violations,
) (*ValidationError, error) {
	valErr := &ValidationError{
		Violations: make([]Violation, len(violations.GetViolations())),
	}
	for i, violation := range violations.GetViolations() {
		valErr.Violations[i] = Violation{
			FieldPath:    violation.GetFieldPath(),
			RulePath:     violation.GetRulePath(),
			ConstraintID: violation.GetConstraintId(),
			Message:      violation.GetMessage(),
			ForKey:       violation.GetForKey(),
		}
		if valErr.Violations[i].FieldPath == "" {
			continue
		}
		fieldValue, descriptor, err := getFieldValue(registry, message, violation.GetFieldPath())
		if err != nil {
			return nil, err
		}
		valErr.Violations[i].FieldValue = fieldValue
		if valErr.Violations[i].RulePath == "" {
			continue
		}
		ruleValue, _, err := getFieldValue(registry, descriptor.Options(), valErr.Violations[i].RulePath)
		if err != nil {
			return nil, err
		}
		valErr.Violations[i].RuleValue = ruleValue
	}
	return valErr, nil
}

// ToProto converts this error into its proto.Message form.
func (err *ValidationError) ToProto() *validate.Violations {
	violations := &validate.Violations{
		Violations: make([]*validate.Violation, len(err.Violations)),
	}
	for i, violation := range err.Violations {
		violations.Violations[i] = &validate.Violation{
			FieldPath:    proto.String(violation.FieldPath),
			RulePath:     proto.String(violation.RulePath),
			ConstraintId: proto.String(violation.ConstraintID),
			Message:      proto.String(violation.Message),
			ForKey:       proto.Bool(violation.ForKey),
		}
	}
	return violations
}

func (err *ValidationError) Error() string {
	bldr := &strings.Builder{}
	bldr.WriteString("validation error:")
	for _, violation := range err.Violations {
		bldr.WriteString("\n - ")
		if fieldPath := violation.FieldPath; fieldPath != "" {
			bldr.WriteString(fieldPath)
			bldr.WriteString(": ")
		}
		_, _ = fmt.Fprintf(bldr, "%s [%s]",
			violation.Message,
			violation.ConstraintID)
	}
	return bldr.String()
}

// PrefixFieldPaths prepends to the provided prefix to the error's internal
// field paths.
func PrefixFieldPaths(err *ValidationError, format string, args ...any) {
	prefix := fmt.Sprintf(format, args...)
	for i := range err.Violations {
		violation := &err.Violations[i]
		switch {
		case violation.FieldPath == "": // no existing field path
			violation.FieldPath = prefix
		case violation.FieldPath[0] == '[': // field is a map/list
			violation.FieldPath = prefix + violation.FieldPath
		default: // any other field
			violation.FieldPath = fmt.Sprintf("%s.%s", prefix, violation.FieldPath)
		}
	}
}

// PrefixRulePaths prepends to the provided prefix to the error's internal
// rule paths.
func PrefixRulePaths(err *ValidationError, format string, args ...any) {
	prefix := fmt.Sprintf(format, args...)
	for i := range err.Violations {
		violation := &err.Violations[i]
		switch {
		case violation.RulePath == "": // no existing rule path
			violation.RulePath = prefix
		case violation.RulePath[0] == '[': // rule is a map/list
			violation.RulePath = prefix + violation.RulePath
		default: // any other rule
			violation.RulePath = fmt.Sprintf("%s.%s", prefix, violation.RulePath)
		}
	}
}

// EqualViolations returns true if the underlying violations are equal.
func EqualViolations(a, b []Violation) bool {
	return slices.EqualFunc(a, b, EqualViolation)
}

// EqualViolation returns true if the underlying violations are equal.
func EqualViolation(a, b Violation) bool {
	return a.FieldPath == b.FieldPath && a.ConstraintID == b.ConstraintID && a.Message == b.Message && a.ForKey == b.ForKey
}
