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
	"errors"
	"slices"
	"strconv"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Merge is a utility to resolve and combine Errors resulting from
// evaluation. If ok is false, execution of validation should stop (either due
// to failFast or the result is not a ValidationErrors).
//
//nolint:errorlint
func Merge(dst, src error, failFast bool) (ok bool, err error) {
	if src == nil {
		return true, dst
	}

	srcValErrs, ok := src.(*ValidationError)
	if !ok {
		return false, src
	}

	if dst == nil {
		return !(failFast && len(srcValErrs.Violations) > 0), src
	}

	dstValErrs, ok := dst.(*ValidationError)
	if !ok {
		// what should we do here?
		return false, dst
	}

	dstValErrs.Violations = append(dstValErrs.Violations, srcValErrs.Violations...)
	return !(failFast && len(dstValErrs.Violations) > 0), dst
}

func FieldPathElement(field protoreflect.FieldDescriptor) *validate.FieldPathElement {
	return &validate.FieldPathElement{
		FieldNumber: proto.Int32(int32(field.Number())),
		FieldName:   proto.String(field.TextName()),
		FieldType:   descriptorpb.FieldDescriptorProto_Type(field.Kind()).Enum(),
	}
}

// AppendFieldPath appends an element to the end of each field path in err.
// As an exception, if skipSubscript is true, any field paths ending in a
// subscript element will not have a suffix element appended to them.
//
// Note that this function is ordinarily used to append field paths in reverse
// order, as the stack bubbles up through the evaluators. Then, at the end, the
// path is reversed.
func AppendFieldPath(err error, suffix *validate.FieldPathElement, skipSubscript bool) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			// Special case: Here we skip appending if the last element had a
			// subscript. This is a weird special case that makes it
			// significantly simpler to handle reverse-constructing paths with
			// maps and slices.
			if elements := violation.Proto.GetField().GetElements(); skipSubscript &&
				len(elements) > 0 && elements[len(elements)-1].Subscript != nil {
				continue
			}
			if violation.Proto.GetField() == nil {
				violation.Proto.Field = &validate.FieldPath{}
			}
			violation.Proto.Field.Elements = append(violation.Proto.Field.Elements, suffix)
		}
	}
}

// PrependRulePath prepends some elements to the beginning of each rule path in
// err. Note that unlike field paths, the rule path is not appended in reverse
// order. This is because rule paths are very often fixed, simple paths, so it
// is better to avoid the copy instead if possible. This prepend is only used in
// the error case for nested rules (repeated.items, map.keys, map.values.)
func PrependRulePath(err error, prefix []*validate.FieldPathElement) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			if violation.Proto.GetRule() == nil {
				violation.Proto.Rule = &validate.FieldPath{}
			}
			violation.Proto.Rule.Elements = append(
				append([]*validate.FieldPathElement{}, prefix...),
				violation.Proto.GetRule().GetElements()...,
			)
		}
	}
}

// ReverseFieldPaths reverses all field paths in the error.
func ReverseFieldPaths(err error) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			if violation.Proto.GetField() != nil {
				slices.Reverse(violation.Proto.GetField().GetElements())
			}
		}
	}
}

// PopulateFieldPathStrings populates the field path strings in the error.
func PopulateFieldPathStrings(err error) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			if violation.Proto.GetField() != nil {
				//nolint:staticcheck // Intentional use of deprecated field
				violation.Proto.FieldPath = proto.String(FieldPathString(violation.Proto.GetField().GetElements()))
			}
		}
	}
}

// FieldPathString takes a FieldPath and encodes it to a string-based dotted
// field path.
func FieldPathString(path []*validate.FieldPathElement) string {
	var result strings.Builder
	for i, element := range path {
		if i > 0 {
			result.WriteByte('.')
		}
		result.WriteString(element.GetFieldName())
		subscript := element.GetSubscript()
		if subscript == nil {
			continue
		}
		result.WriteByte('[')
		switch value := subscript.(type) {
		case *validate.FieldPathElement_Index:
			result.WriteString(strconv.FormatUint(value.Index, 10))
		case *validate.FieldPathElement_BoolKey:
			result.WriteString(strconv.FormatBool(value.BoolKey))
		case *validate.FieldPathElement_IntKey:
			result.WriteString(strconv.FormatInt(value.IntKey, 10))
		case *validate.FieldPathElement_UintKey:
			result.WriteString(strconv.FormatUint(value.UintKey, 10))
		case *validate.FieldPathElement_StringKey:
			result.WriteString(strconv.Quote(value.StringKey))
		}
		result.WriteByte(']')
	}
	return result.String()
}

func MarkForKey(err error) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			violation.Proto.ForKey = proto.Bool(true)
		}
	}
}
