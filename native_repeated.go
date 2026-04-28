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
	"bytes"
	"fmt"
	"math"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	repeatedMinItemsSite = makeRuleSite(
		repeatedFieldRulesDesc,
		(*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("min_items"),
		"repeated.min_items", "",
	)
	repeatedMaxItemsSite = makeRuleSite(
		repeatedFieldRulesDesc,
		(*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("max_items"),
		"repeated.max_items", "",
	)
	repeatedUniqueSite = makeRuleSite(
		repeatedFieldRulesDesc,
		(*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("unique"),
		"repeated.unique",
		"repeated value must contain unique items",
	)
)

// tryNativeRepeatedRules attempts to build a native Go evaluator for
// repeated list-level rules (min_items, max_items, unique).
// Returns nil if the rules can't be handled natively.
func tryNativeRepeatedRules(base base, rules *validate.RepeatedRules) evaluator {
	if rules == nil {
		return nil
	}
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	var minItems uint64
	if rules.HasMinItems() {
		minItems = rules.GetMinItems()
		rules.ProtoReflect().Clear(repeatedMinItemsSite.desc)
		hasRule = true
	}

	var maxItems uint64 = math.MaxUint64
	if rules.HasMaxItems() {
		maxItems = rules.GetMaxItems()
		rules.ProtoReflect().Clear(repeatedMaxItemsSite.desc)
		hasRule = true
	}

	var uniqueFn uniqueChecker
	if rules.GetUnique() {
		uniqueFn = uniqueCheckerForKind(base.Descriptor)
		if uniqueFn == nil {
			// message/list/map elements can't be checked for uniqueness
			// natively; fall through to CEL.
			return nil
		}
		rules.ProtoReflect().Clear(repeatedUniqueSite.desc)
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeRepeatedEval{
		base:     base,
		minItems: minItems,
		maxItems: maxItems,
		uniqueFn: uniqueFn,
	}
}

// uniqueChecker tests whether all elements in a repeated list are distinct.
// A nil uniqueChecker means the `unique` rule is not active for this field.
type uniqueChecker func(protoreflect.List) bool

// uniqueCheckerForKind returns the concrete uniqueness check for the element
// kind of a repeated field, or nil if the kind isn't supported (message, list,
// map — none of which are valid element kinds for a repeated scalar field
// with `unique`).
func uniqueCheckerForKind(desc protoreflect.FieldDescriptor) uniqueChecker {
	if desc == nil {
		return nil
	}
	switch desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return isUniqueList[int32]
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return isUniqueList[int64]
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return isUniqueList[uint32]
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return isUniqueList[uint64]
	case protoreflect.FloatKind:
		return isUniqueList[float32]
	case protoreflect.DoubleKind:
		return isUniqueList[float64]
	case protoreflect.StringKind:
		return isUniqueList[string]
	case protoreflect.BoolKind:
		return isUniqueList[bool]
	case protoreflect.EnumKind:
		return isUniqueList[protoreflect.EnumNumber]
	case protoreflect.BytesKind:
		return isUniqueBytes
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return nil
	default:
		return nil
	}
}

var _ evaluator = nativeRepeatedEval{}

// nativeRepeatedEval is a native Go evaluator for repeated list-level rules
// (min_items, max_items, unique). Item-level rules are handled separately
// by the listItems evaluator in repeated.go.
type nativeRepeatedEval struct {
	base
	minItems uint64
	maxItems uint64
	// uniqueFn is nil when the `unique` rule is not active. When set, it is
	// specialized for the field's element kind at compile time.
	uniqueFn uniqueChecker
}

func (n nativeRepeatedEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	list := val.List()
	size := uint64(list.Len()) //nolint:gosec // len can't be < 0 and is always within uint64 range
	var violations []*Violation

	if size < n.minItems {
		violations = append(violations, n.newViolation(repeatedMinItemsSite,
			"repeated.min_items",
			fmt.Sprintf("must contain at least %d item(s)", n.minItems),
			val, protoreflect.ValueOfUint64(n.minItems)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if size > n.maxItems {
		violations = append(violations, n.newViolation(repeatedMaxItemsSite,
			"repeated.max_items",
			fmt.Sprintf("must contain no more than %d item(s)", n.maxItems),
			val, protoreflect.ValueOfUint64(n.maxItems)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if n.uniqueFn != nil && !n.uniqueFn(list) {
		violations = append(violations, n.newViolation(repeatedUniqueSite,
			"repeated.unique",
			"repeated value must contain unique items",
			val, protoreflect.ValueOfBool(true)))
		if cfg.failFast {
			return &ValidationError{Violations: violations}
		}
	}

	if len(violations) > 0 {
		return &ValidationError{
			Violations: violations,
		}
	}
	return nil
}

// uniqueLinearThreshold is the list length at and below which uniqueness is
// checked with an O(n²) scan over a stack-allocated array instead of a map.
// For small lists the linear scan is faster and avoids the map allocation;
// at larger sizes the map's O(n) lookup wins.
const uniqueLinearThreshold = 16

// isUniqueList is the generic uniqueness check used for all comparable scalar
// element kinds (the concrete T is bound at compile time via
// uniqueCheckerForKind).
func isUniqueList[T comparable](list protoreflect.List) bool {
	length := list.Len()
	if length <= 1 {
		return true
	}
	if length <= uniqueLinearThreshold {
		var seen = make([]T, length)
		for i := range length {
			key, ok := list.Get(i).Interface().(T)
			if !ok {
				return false
			}
			for j := range i {
				if seen[j] == key {
					return false
				}
			}
			seen[i] = key
		}
		return true
	}

	seen := make(map[T]struct{}, length)
	for i := range length {
		key, ok := list.Get(i).Interface().(T)
		if !ok {
			// should never happen, but just in case
			return false
		}
		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

func isUniqueBytes(list protoreflect.List) bool {
	length := list.Len()
	if length <= 1 {
		return true
	}
	if length <= uniqueLinearThreshold {
		// storing []byte directly avoids the []byte→string allocation the
		// map path needs for a hashable key.
		var seen = make([][]byte, uniqueLinearThreshold)
		for i := range length {
			byteVal := list.Get(i).Bytes()
			for j := range i {
				if bytes.Equal(seen[j], byteVal) {
					return false
				}
			}
			seen[i] = byteVal
		}
		return true
	}

	seen := make(map[string]struct{}, length)
	for i := range length {
		byteVal := list.Get(i).Bytes()
		// []byte is not comparable; convert to string for use as map key.
		// this is the same action performed by CEL in library.uniqueBytes
		key := string(byteVal)
		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}
	return true
}

func (n nativeRepeatedEval) Tautology() bool {
	return false
}
