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
	"math"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	mapMinPairsSite = makeRuleSite(
		mapFieldRulesDesc,
		(*validate.MapRules)(nil).ProtoReflect().Descriptor().Fields().ByName("min_pairs"),
		"map.min_pairs", "",
	)
	mapMaxPairsSite = makeRuleSite(
		mapFieldRulesDesc,
		(*validate.MapRules)(nil).ProtoReflect().Descriptor().Fields().ByName("max_pairs"),
		"map.max_pairs", "",
	)
)

// tryNativeMapRules attempts to build a native Go evaluator for
// map-level rules (min_pairs, max_pairs).
// Returns nil if the rules can't be handled natively.
func tryNativeMapRules(base base, rules *validate.MapRules) evaluator {
	if rules == nil {
		return nil
	}
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	var minPairs uint64
	if rules.HasMinPairs() {
		minPairs = rules.GetMinPairs()
		rules.ProtoReflect().Clear(mapMinPairsSite.desc)
		hasRule = true
	}

	var maxPairs uint64 = math.MaxUint64
	if rules.HasMaxPairs() {
		maxPairs = rules.GetMaxPairs()
		rules.ProtoReflect().Clear(mapMaxPairsSite.desc)
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeMapEval{
		base:     base,
		minPairs: minPairs,
		maxPairs: maxPairs,
	}
}

var _ evaluator = nativeMapEval{}

// nativeMapEval is a native Go evaluator for map-level rules
// (min_pairs, max_pairs). Key/value rules are handled separately
// by the kvPairs evaluator in map.go.
type nativeMapEval struct {
	base
	minPairs uint64
	maxPairs uint64
}

func (n nativeMapEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, _ *validationConfig) error {
	size := uint64(val.Map().Len()) //nolint:gosec // int will never be negative or out of uint64 range

	// min_pairs
	if size < n.minPairs {
		return &ValidationError{Violations: []*Violation{n.newViolation(mapMinPairsSite,
			"map.min_pairs",
			fmt.Sprintf("map must be at least %d entries", n.minPairs),
			val, protoreflect.ValueOfUint64(n.minPairs)),
		}}
	}

	// max_pairs
	if size > n.maxPairs {
		return &ValidationError{Violations: []*Violation{n.newViolation(mapMaxPairsSite,
			"map.max_pairs",
			fmt.Sprintf("map must be at most %d entries", n.maxPairs),
			val, protoreflect.ValueOfUint64(n.maxPairs)),
		}}
	}

	return nil
}

func (n nativeMapEval) Tautology() bool {
	return n.minPairs == 0 && n.maxPairs == math.MaxUint64
}
