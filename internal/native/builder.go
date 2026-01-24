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

package native

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Builder creates native evaluators from field rules.
type Builder struct{}

// NewBuilder creates a new native evaluator builder.
func NewBuilder() *Builder {
	return &Builder{}
}

// Build attempts to create native evaluators for the given field and rules.
// Returns nil if the rules cannot be handled natively (requires CEL fallback).
// The forItems parameter indicates if we're building rules for repeated items.
func (b *Builder) Build(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
	forItems bool,
) Evaluators {
	if rules == nil {
		return nil
	}

	var evaluators Evaluators

	// Build evaluators based on field type
	switch {
	case fieldDesc.IsMap() && !forItems:
		evaluators = b.buildMapRules(fieldDesc, rules.GetMap())
	case fieldDesc.IsList() && !forItems:
		evaluators = b.buildRepeatedRules(fieldDesc, rules.GetRepeated())
	default:
		evaluators = b.buildScalarRules(fieldDesc, rules)
	}

	if len(evaluators) == 0 {
		return nil
	}
	return evaluators
}

// buildScalarRules builds native evaluators for scalar field types.
func (b *Builder) buildScalarRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	var evaluators Evaluators

	switch fieldDesc.Kind() {
	// Signed integers
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		evaluators = b.buildInt32Rules(fieldDesc, rules)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		evaluators = b.buildInt64Rules(fieldDesc, rules)

	// Unsigned integers
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		evaluators = b.buildUint32Rules(fieldDesc, rules)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		evaluators = b.buildUint64Rules(fieldDesc, rules)

	// Floating point
	case protoreflect.FloatKind:
		evaluators = b.buildFloatRules(fieldDesc, rules)
	case protoreflect.DoubleKind:
		evaluators = b.buildDoubleRules(fieldDesc, rules)

	// Boolean
	case protoreflect.BoolKind:
		evaluators = b.buildBoolRules(fieldDesc, rules)

	// String and Bytes
	case protoreflect.StringKind:
		evaluators = b.buildStringRules(fieldDesc, rules)
	case protoreflect.BytesKind:
		evaluators = b.buildBytesRules(fieldDesc, rules)
	}

	return evaluators
}

// buildMapRules builds native evaluators for map field rules.
func (b *Builder) buildMapRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.MapRules,
) Evaluators {
	if rules == nil {
		return nil
	}
	return buildMapSizeRules(fieldDesc, rules)
}

// buildRepeatedRules builds native evaluators for repeated field rules.
func (b *Builder) buildRepeatedRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.RepeatedRules,
) Evaluators {
	if rules == nil {
		return nil
	}
	return buildRepeatedSizeRules(fieldDesc, rules)
}

// Placeholder methods for each type - these delegate to type-specific files

func (b *Builder) buildInt32Rules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildSignedRules[int32](fieldDesc, rules.GetInt32(), "int32")
}

func (b *Builder) buildInt64Rules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildSignedRules[int64](fieldDesc, rules.GetInt64(), "int64")
}

func (b *Builder) buildUint32Rules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildUnsignedRules[uint32](fieldDesc, rules.GetUint32(), "uint32")
}

func (b *Builder) buildUint64Rules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildUnsignedRules[uint64](fieldDesc, rules.GetUint64(), "uint64")
}

func (b *Builder) buildFloatRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildFloatRules(fieldDesc, rules.GetFloat())
}

func (b *Builder) buildDoubleRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildDoubleRules(fieldDesc, rules.GetDouble())
}

func (b *Builder) buildBoolRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildBoolRules(fieldDesc, rules.GetBool())
}

func (b *Builder) buildStringRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildStringRulesFromRules(fieldDesc, rules.GetString())
}

func (b *Builder) buildBytesRules(
	fieldDesc protoreflect.FieldDescriptor,
	rules *validate.FieldRules,
) Evaluators {
	return buildBytesRulesFromRules(fieldDesc, rules.GetBytes())
}
