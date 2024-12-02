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
var (
	anyInRulePath = []*validate.FieldPathElement{
		{FieldName: proto.String("any"), FieldNumber: proto.Int32(20), FieldType: descriptorpb.FieldDescriptorProto_Type(11).Enum()},
		{FieldName: proto.String("in"), FieldNumber: proto.Int32(2), FieldType: descriptorpb.FieldDescriptorProto_Type(9).Enum()},
	}
	anyNotInRulePath = []*validate.FieldPathElement{
		{FieldName: proto.String("any"), FieldNumber: proto.Int32(20), FieldType: descriptorpb.FieldDescriptorProto_Type(11).Enum()},
		{FieldName: proto.String("not_in"), FieldNumber: proto.Int32(3), FieldType: descriptorpb.FieldDescriptorProto_Type(9).Enum()},
	}
)

// anyPB is a specialized evaluator for applying validate.AnyRules to an
// anypb.Any message. This is handled outside CEL which attempts to
// hydrate anyPB's within an expression, breaking evaluation if the type is
// unknown at runtime.
type anyPB struct {
	// TypeURLDescriptor is the descriptor for the TypeURL field
	TypeURLDescriptor protoreflect.FieldDescriptor
	// In specifies which type URLs the value may possess
	In map[string]struct{}
	// NotIn specifies which type URLs the value may not possess
	NotIn map[string]struct{}
	// InValue contains the original `in` rule value.
	InValue protoreflect.Value
	// NotInValue contains the original `not_in` rule value.
	NotInValue protoreflect.Value
}

func (a anyPB) Evaluate(val protoreflect.Value, failFast bool) error {
	typeURL := val.Message().Get(a.TypeURLDescriptor).String()

	err := &errors.ValidationError{}
	if len(a.In) > 0 {
		if _, ok := a.In[typeURL]; !ok {
			err.Violations = append(err.Violations, &errors.Violation{
				Proto: &validate.Violation{
					Rule:         &validate.FieldPath{Elements: anyInRulePath},
					ConstraintId: proto.String("any.in"),
					Message:      proto.String("type URL must be in the allow list"),
				},
				FieldValue: val,
				RuleValue:  a.InValue,
			})
			if failFast {
				return err
			}
		}
	}

	if len(a.NotIn) > 0 {
		if _, ok := a.NotIn[typeURL]; ok {
			err.Violations = append(err.Violations, &errors.Violation{
				Proto: &validate.Violation{
					Rule:         &validate.FieldPath{Elements: anyNotInRulePath},
					ConstraintId: proto.String("any.not_in"),
					Message:      proto.String("type URL must not be in the block list"),
				},
				FieldValue: val,
				RuleValue:  a.NotInValue,
			})
		}
	}

	if len(err.Violations) > 0 {
		return err
	}
	return nil
}

func (a anyPB) Tautology() bool {
	return len(a.In) == 0 && len(a.NotIn) == 0
}

func stringsToSet(ss []string) map[string]struct{} {
	if len(ss) == 0 {
		return nil
	}
	set := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}
	return set
}

var _ evaluator = anyPB{}
