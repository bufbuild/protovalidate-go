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

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Violation represents a single instance where a validation rule was not met.
// It provides information about the field that caused the violation, the
// specific unfulfilled constraint, and a human-readable error message.
type Violation interface {
	// GetFieldValue returns the value of the specific field that failed
	// validation. If there was no value, this will return an invalid value.
	GetFieldValue() protoreflect.Value

	// GetRuleValue returns the value of the rule that specified the failed
	// constraint. Not all constraints have a value; only standard and
	// predefined constraints have rule values. In violations caused by other
	// kinds of constraints, like custom contraints, this will return an invalid
	// value.
	GetRuleValue() protoreflect.Value

	// ToProto converts this violation into its proto.Message form.
	ToProto() *validate.Violation
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
		violations.Violations[i] = violation.ToProto()
	}
	return violations
}

func (err *ValidationError) Error() string {
	bldr := &strings.Builder{}
	bldr.WriteString("validation error:")
	for _, violation := range err.Violations {
		violation := violation.ToProto()
		bldr.WriteString("\n - ")
		if fieldPath := FieldPathString(violation.GetField().GetElements()); fieldPath != "" {
			bldr.WriteString(fieldPath)
			bldr.WriteString(": ")
		}
		_, _ = fmt.Fprintf(bldr, "%s [%s]",
			violation.GetMessage(),
			violation.GetConstraintId())
	}
	return bldr.String()
}

// ViolationData is a simple implementation of Violation.
type ViolationData struct {
	Field        []*validate.FieldPathElement
	Rule         []*validate.FieldPathElement
	FieldValue   protoreflect.Value
	RuleValue    protoreflect.Value
	ConstraintID string
	Message      string
	ForKey       bool
}

func (v *ViolationData) GetFieldValue() protoreflect.Value {
	return v.FieldValue
}

func (v *ViolationData) GetRuleValue() protoreflect.Value {
	return v.RuleValue
}

func (v *ViolationData) ToProto() *validate.Violation {
	var fieldPathString *string
	if len(v.Field) > 0 {
		fieldPathString = proto.String(FieldPathString(v.Field))
	}
	var forKey *bool
	if v.ForKey {
		forKey = proto.Bool(true)
	}
	return &validate.Violation{
		Field:        fieldPathProto(v.Field),
		Rule:         fieldPathProto(v.Rule),
		FieldPath:    fieldPathString,
		ConstraintId: proto.String(v.ConstraintID),
		Message:      proto.String(v.Message),
		ForKey:       forKey,
	}
}

var _ Violation = &ViolationData{}

func fieldPathProto(elements []*validate.FieldPathElement) *validate.FieldPath {
	if len(elements) == 0 {
		return nil
	}
	return &validate.FieldPath{
		Elements: elements,
	}
}
