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
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	repeatedMinItemsDesc = (*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("min_items")
	repeatedMaxItemsDesc = (*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("max_items")
	repeatedUniqueDesc   = (*validate.RepeatedRules)(nil).ProtoReflect().Descriptor().Fields().ByName("unique")
)

// nativeRepeatedEval is a native Go evaluator for repeated list-level rules
// (min_items, max_items, unique). Item-level rules are handled separately
// by the listItems evaluator in repeated.go.
type nativeRepeatedEval struct {
	base
	minItems *uint64
	maxItems *uint64
	unique   bool
}

func (n nativeRepeatedEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, _ *validationConfig) error {
	list := val.List()
	size := uint64(list.Len()) //nolint:gosec

	// min_items
	if n.minItems != nil && size < *n.minItems {
		return n.newViolation(repeatedFieldRulesDesc, repeatedMinItemsDesc,
			"repeated.min_items",
			fmt.Sprintf("must contain at least %d item(s)", *n.minItems),
			val, protoreflect.ValueOfUint64(*n.minItems))
	}

	// max_items
	if n.maxItems != nil && size > *n.maxItems {
		return n.newViolation(repeatedFieldRulesDesc, repeatedMaxItemsDesc,
			"repeated.max_items",
			fmt.Sprintf("must contain no more than %d item(s)", *n.maxItems),
			val, protoreflect.ValueOfUint64(*n.maxItems))
	}

	// unique
	if n.unique && !isUnique(list) {
		return n.newViolation(repeatedFieldRulesDesc, repeatedUniqueDesc,
			"repeated.unique",
			"repeated value must contain unique items",
			val, protoreflect.ValueOfBool(true))
	}

	return nil
}

// isUnique checks whether all elements in the list are distinct.
func isUnique(list protoreflect.List) bool {
	length := list.Len()
	if length <= 1 {
		return true
	}
	// type-specific maps avoid any-boxing allocations
	switch list.Get(0).Interface().(type) {
	case int32:
		return isUniqueTyped[int32](list, length)
	case int64:
		return isUniqueTyped[int64](list, length)
	case uint32:
		return isUniqueTyped[uint32](list, length)
	case uint64:
		return isUniqueTyped[uint64](list, length)
	case float32:
		return isUniqueTyped[float32](list, length)
	case float64:
		return isUniqueTyped[float64](list, length)
	case string:
		return isUniqueTyped[string](list, length)
	case bool:
		return isUniqueTyped[bool](list, length)
	case protoreflect.EnumNumber:
		return isUniqueTyped[protoreflect.EnumNumber](list, length)
	case []byte:
		return isUniqueBytes(list, length)
	default:
		return false // message, list, and map types are not supported, only enum and scalars
	}
}

func isUniqueTyped[T comparable](list protoreflect.List, length int) bool {
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

func isUniqueBytes(list protoreflect.List, length int) bool {
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

var _ evaluator = nativeRepeatedEval{}

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

	var minItems *uint64
	if rules.HasMinItems() {
		minItems = ptr(rules.GetMinItems())
		hasRule = true
	}

	var maxItems *uint64
	if rules.HasMaxItems() {
		maxItems = ptr(rules.GetMaxItems())
		hasRule = true
	}

	unique := false
	if rules.GetUnique() {
		unique = true
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeRepeatedEval{
		base:     base,
		minItems: minItems,
		maxItems: maxItems,
		unique:   unique,
	}
}
