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
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
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

	// FieldValue is the value of the specific field that failed validation.
	FieldValue protoreflect.Value

	// ForKey` indicates whether the violation was caused by a map key, rather
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
