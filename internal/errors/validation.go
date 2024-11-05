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
)

// Violation represents a single instance where a validation rule was not met.
// It provides information about the field that caused the violation, the
// specific unfulfilled constraint, and a human-readable error message.
type Violation struct {
	// FieldPath is an identifier that points to the specific field
	// that failed the validation.  This could be a nested field, in which case
	// the path will include all the parent fields leading to the actual field
	// that caused the violation.
	FieldPath FieldPath

	// RulePath is a machine-readable identifier that points to the specific
	// constraint rule that failed validation. This will be a nested field
	// starting from the FieldConstraints of the field that failed validation.
	// This value is only set for standard or predefined rules on fields.
	RulePath FieldPath

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

// ToProto converts this error into its proto.Message form.
func (err *ValidationError) ToProto() *validate.Violations {
	violations := &validate.Violations{
		Violations: make([]*validate.Violation, len(err.Violations)),
	}
	for i, violation := range err.Violations {
		var fieldPath *validate.FieldPath
		if len(violation.FieldPath.path) > 0 {
			fieldPath = violation.FieldPath.ToProto()
		}
		var rulePath *validate.FieldPath
		if len(violation.RulePath.path) > 0 {
			rulePath = violation.RulePath.ToProto()
		}
		var fieldPathString *string
		if fieldPath != nil {
			fieldPathString = proto.String(FieldPathString(fieldPath.GetElements()))
		}
		var forKey *bool
		if violation.ForKey {
			forKey = proto.Bool(true)
		}
		violations.Violations[i] = &validate.Violation{
			FieldPathString: fieldPathString,
			FieldPath:       fieldPath,
			RulePath:        rulePath,
			ConstraintId:    proto.String(violation.ConstraintID),
			Message:         proto.String(violation.Message),
			ForKey:          forKey,
		}
	}
	return violations
}

func (err *ValidationError) Error() string {
	bldr := &strings.Builder{}
	bldr.WriteString("validation error:")
	for _, violation := range err.Violations {
		bldr.WriteString("\n - ")
		if fieldPath := violation.FieldPath.String(); fieldPath != "" {
			bldr.WriteString(fieldPath)
			bldr.WriteString(": ")
		}
		_, _ = fmt.Fprintf(bldr, "%s [%s]",
			violation.Message,
			violation.ConstraintID)
	}
	return bldr.String()
}

// A FieldPath specifies a nested field inside of a protobuf message.
type FieldPath struct {
	path []*validate.FieldPathElement
}

func NewFieldPath(elements []*validate.FieldPathElement) FieldPath {
	return FieldPath{path: elements}
}

func (f FieldPath) ToProto() *validate.FieldPath {
	return &validate.FieldPath{
		Elements: f.path,
	}
}

func (f FieldPath) String() string {
	return FieldPathString(f.path)
}

// EqualViolations returns true if the underlying violations are equal.
func EqualViolations(a, b []Violation) bool {
	return slices.EqualFunc(a, b, EqualViolation)
}

// EqualViolation returns true if the underlying violations are equal.
func EqualViolation(a, b Violation) bool {
	return (a.ConstraintID == b.ConstraintID &&
		a.Message == b.Message &&
		a.ForKey == b.ForKey &&
		proto.Equal(a.FieldPath.ToProto(), b.FieldPath.ToProto()) &&
		proto.Equal(a.RulePath.ToProto(), b.RulePath.ToProto()))
}
