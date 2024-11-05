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
var repeatedItemsFieldPath = []*validate.FieldPathElement{
	{FieldName: proto.String("repeated"), FieldNumber: proto.Int32(18), FieldType: descriptorpb.FieldDescriptorProto_Type(11).Enum()},
	{FieldName: proto.String("items"), FieldNumber: proto.Int32(4), FieldType: descriptorpb.FieldDescriptorProto_Type(11).Enum()},
}

// listItems performs validation on the elements of a repeated field.
type listItems struct {
	// Descriptor is the FieldDescriptor targeted by this evaluator
	Descriptor protoreflect.FieldDescriptor
	// ItemConstraints are checked on every item of the list
	ItemConstraints value
}

func (r listItems) Evaluate(val protoreflect.Value, failFast bool) error {
	list := val.List()
	var ok bool
	var err error
	for i := 0; i < list.Len(); i++ {
		itemErr := r.ItemConstraints.Evaluate(list.Get(i), failFast)
		if itemErr != nil {
			element := errors.FieldPathElement(r.Descriptor)
			element.Subscript = &validate.FieldPathElement_Index{Index: uint64(i)}
			errors.AppendFieldPath(itemErr, element, false)
		}
		if ok, err = errors.Merge(err, itemErr, failFast); !ok {
			return err
		}
	}
	return err
}

func (r listItems) Tautology() bool {
	return r.ItemConstraints.Tautology()
}

// itemsWrapper wraps the evaluation of nested repeated field rules.
type itemsWrapper struct {
	evaluator
}

func newItemsWrapper(evaluator evaluator) evaluator {
	return itemsWrapper{evaluator}
}

func (e itemsWrapper) Evaluate(val protoreflect.Value, failFast bool) error {
	err := e.evaluator.Evaluate(val, failFast)
	errors.PrependRulePath(err, repeatedItemsFieldPath)
	return err
}

var (
	_ evaluator = listItems{}
	_ evaluator = itemsWrapper{}
	_ wrapper   = newItemsWrapper
)
