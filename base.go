// Copyright 2023-2026 Buf Technologies, Inc.
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
	"slices"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// base is a common struct used by all field evaluators. It holds
// some common information used across all field evaluators.
type base struct {
	// Descriptor is the FieldDescriptor targeted by this evaluator, or nil if
	// there is none.
	Descriptor protoreflect.FieldDescriptor

	// FieldPathElement is the field path element that pertains to this evaluator, or
	// nil if there is none.
	FieldPathElement *validate.FieldPathElement

	// RulePrefix is a static prefix this evaluator should add to the rule path
	// of violations.
	RulePrefix *validate.FieldPath
}

func newBase(valEval *value) base {
	return base{
		Descriptor:       valEval.Descriptor,
		FieldPathElement: fieldPathElement(valEval.Descriptor),
		RulePrefix:       valEval.NestedRule,
	}
}

func (b *base) fieldPath() *validate.FieldPath {
	if b.FieldPathElement == nil {
		return nil
	}
	return validate.FieldPath_builder{
		Elements: []*validate.FieldPathElement{
			b.FieldPathElement,
		},
	}.Build()
}

func (b *base) rulePath(suffix *validate.FieldPath) *validate.FieldPath {
	return prefixRulePath(b.RulePrefix, suffix)
}

// ruleSite is a compile-time pre-built bundle for a single rule site: the
// 2-element rule path suffix and the leaf descriptor. Using ruleSite with
// newViolationAt avoids re-allocating FieldPathElement proto messages on
// every violation (each FieldPathElement rebuild allocates ~4 sub-objects).
type ruleSite struct {
	// pathElements is the rule path suffix, e.g.
	// [FieldRules.int32 element, Int32Rules.gt element].
	//
	// Safe to share across violations: only the containing *FieldPath is
	// rebuilt per violation, which is what updateViolationPaths mutates.
	pathElements []*validate.FieldPathElement
	// desc is the leaf rule field descriptor (e.g. Int32Rules.gt); it is
	// stored on the returned *Violation's RuleDescriptor field.
	desc protoreflect.FieldDescriptor

	// if there are constant values for ruleID or message, specify them once so they can be reused
	ruleID  *string
	message *string
}

// makeRuleSite pre-builds a ruleSite for a rule at compile time.
func makeRuleSite(ruleDesc, desc protoreflect.FieldDescriptor) ruleSite {
	return ruleSite{
		pathElements: []*validate.FieldPathElement{
			fieldPathElement(ruleDesc),
			fieldPathElement(desc),
		},
		desc: desc,
	}
}

// makeRuleSiteWithID pre-builds a ruleSite for a rule at compile time.
func makeRuleSiteWithID(ruleDesc, desc protoreflect.FieldDescriptor, ruleID string) ruleSite {
	return ruleSite{
		pathElements: []*validate.FieldPathElement{
			fieldPathElement(ruleDesc),
			fieldPathElement(desc),
		},
		desc:   desc,
		ruleID: proto.String(ruleID),
	}
}

// makeRuleSiteWithID pre-builds a ruleSite for a rule at compile time.
func makeRuleSiteWithIDAndMessage(ruleDesc, desc protoreflect.FieldDescriptor, ruleID string, message string) ruleSite {
	return ruleSite{
		pathElements: []*validate.FieldPathElement{
			fieldPathElement(ruleDesc),
			fieldPathElement(desc),
		},
		desc:    desc,
		ruleID:  proto.String(ruleID),
		message: proto.String(message),
	}
}

// newViolation constructs a Violation.
// ruleDesc is the top-level rule descriptor (e.g., FieldRules.int32),
// desc is the specific constraint descriptor (e.g., Int32Rules.gt).
//
// it uses pre-built path elements in site instead of rebuilding them each call.
func (b *base) newViolation(
	site ruleSite,
	ruleID string,
	message string,
	fieldValue protoreflect.Value,
	ruleValue protoreflect.Value,
) *Violation {
	ruleIDPtr := site.ruleID
	if ruleIDPtr == nil {
		ruleIDPtr = proto.String(ruleID)
	}
	messagePtr := site.message
	if messagePtr == nil {
		messagePtr = proto.String(message)
	}
	return &Violation{
		Proto: validate.Violation_builder{
			Field: b.fieldPath(),
			Rule: b.rulePath(validate.FieldPath_builder{
				Elements: site.pathElements,
			}.Build()),
			RuleId:  ruleIDPtr,
			Message: messagePtr,
		}.Build(),
		FieldValue:      fieldValue,
		FieldDescriptor: b.Descriptor,
		RuleValue:       ruleValue,
		RuleDescriptor:  site.desc,
	}
}

func prefixRulePath(prefix *validate.FieldPath, suffix *validate.FieldPath) *validate.FieldPath {
	if len(prefix.GetElements()) > 0 {
		return validate.FieldPath_builder{
			Elements: slices.Concat(prefix.GetElements(), suffix.GetElements()),
		}.Build()
	}
	return suffix
}
