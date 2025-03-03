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

package protovalidate

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	repeatedRuleDescriptor      = (&validate.FieldConstraints{}).ProtoReflect().Descriptor().Fields().ByName("repeated")
	repeatedItemsRuleDescriptor = (&validate.RepeatedRules{}).ProtoReflect().Descriptor().Fields().ByName("items")
	repeatedItemsRulePath       = &validate.FieldPath{
		Elements: []*validate.FieldPathElement{
			fieldPathElement(repeatedRuleDescriptor),
			fieldPathElement(repeatedItemsRuleDescriptor),
		},
	}
)

// listItems performs validation on the elements of a repeated field.
type listItems struct {
	base

	// ItemConstraints are checked on every item of the list
	ItemConstraints value
}

func newListItems(valEval *value) listItems {
	return listItems{
		base:            newBase(valEval),
		ItemConstraints: value{NestedRule: repeatedItemsRulePath},
	}
}

func (r listItems) Evaluate(msg protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	list := val.List()
	var ok bool
	var err error
	for i := 0; i < list.Len(); i++ {
		itemErr := r.ItemConstraints.EvaluateField(msg, list.Get(i), cfg, true)
		if itemErr != nil {
			updateViolationPaths(itemErr, &validate.FieldPathElement{
				FieldNumber: proto.Int32(r.base.FieldPathElement.GetFieldNumber()),
				FieldType:   r.base.FieldPathElement.GetFieldType().Enum(),
				FieldName:   proto.String(r.base.FieldPathElement.GetFieldName()),
				Subscript:   &validate.FieldPathElement_Index{Index: uint64(i)}, //nolint:gosec // indices are guaranteed to be non-negative
			}, r.base.RulePrefix.GetElements())
		}
		if ok, err = mergeViolations(err, itemErr, cfg); !ok {
			return err
		}
	}
	return err
}

func (r listItems) Tautology() bool {
	return r.ItemConstraints.Tautology()
}

var _ evaluator = listItems{}
