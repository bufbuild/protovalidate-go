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

package evaluator

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go/internal/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

//nolint:gochecknoglobals
var enumDefinedOnlyRulePath = []*validate.FieldPathElement{
	{FieldName: proto.String("enum"), FieldNumber: proto.Int32(16), FieldType: descriptorpb.FieldDescriptorProto_Type(11).Enum()},
	{FieldName: proto.String("defined_only"), FieldNumber: proto.Int32(2), FieldType: descriptorpb.FieldDescriptorProto_Type(8).Enum()},
}

// definedEnum is an evaluator that checks an enum value being a member of
// the defined values exclusively. This check is handled outside CEL as enums
// are completely type erased to integers.
type definedEnum struct {
	// ValueDescriptors captures all the defined values for this enum
	ValueDescriptors protoreflect.EnumValueDescriptors
}

func (d definedEnum) Evaluate(val protoreflect.Value, _ bool) error {
	if d.ValueDescriptors.ByNumber(val.Enum()) == nil {
		return &errors.ValidationError{Violations: []errors.Violation{&errors.ViolationData{
			Rule:         enumDefinedOnlyRulePath,
			FieldValue:   val,
			RuleValue:    protoreflect.ValueOfBool(true),
			ConstraintID: "enum.defined_only",
			Message:      "value must be one of the defined enum values",
		}}}
	}
	return nil
}

func (d definedEnum) Tautology() bool {
	return false
}

var _ evaluator = definedEnum{}
