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
	"slices"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// numericValue is the set of Go types that back protobuf numeric field kinds.
type numericValue interface {
	~int32 | ~int64 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Per-kind builder wrappers. Each handles nil check and type-specific
// concerns before delegating to the generic builder.

func tryBuildNativeInt32Rules(base base, rules *validate.Int32Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, int32Config)
}

func tryBuildNativeSint32Rules(base base, rules *validate.SInt32Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, sint32Config)
}

func tryBuildNativeSfixed32Rules(base base, rules *validate.SFixed32Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, sfixed32Config)
}

func tryBuildNativeInt64Rules(base base, rules *validate.Int64Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, int64Config)
}

func tryBuildNativeSint64Rules(base base, rules *validate.SInt64Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, sint64Config)
}

func tryBuildNativeSfixed64Rules(base base, rules *validate.SFixed64Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, sfixed64Config)
}

func tryBuildNativeUint32Rules(base base, rules *validate.UInt32Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, uint32Config)
}

func tryBuildNativeFixed32Rules(base base, rules *validate.Fixed32Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, fixed32Config)
}

func tryBuildNativeUint64Rules(base base, rules *validate.UInt64Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, uint64Config)
}

func tryBuildNativeFixed64Rules(base base, rules *validate.Fixed64Rules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, fixed64Config)
}

func tryBuildNativeFloatRules(base base, rules *validate.FloatRules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, floatConfig)
}

func tryBuildNativeDoubleRules(base base, rules *validate.DoubleRules) evaluator {
	if rules == nil {
		return nil
	}
	return tryBuildNativeNumericRules(base, rules, doubleConfig)
}

// tryBuildNativeNumericRules attempts to build a native Go evaluator for
// numeric rules. Returns nil if the rules can't be handled natively,
// including cases with unknown fields (custom predefined extensions).
func tryBuildNativeNumericRules[T numericValue, R numericRules[T]](
	base base,
	rules R,
	config numericTypeConfig[T],
) evaluator {
	// Bail out if the rules message has unknown fields, which indicate
	// custom predefined extensions that we can't handle natively.
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}

	hasRule := false

	var lowerValue T
	lower := lowerBoundNone
	switch {
	case rules.HasGt():
		lower = lowerBoundGt
		lowerValue = rules.GetGt()
		hasRule = true
	case rules.HasGte():
		lower = lowerBoundGte
		lowerValue = rules.GetGte()
		hasRule = true
	}

	var upperValue T
	upper := upperBoundNone
	switch {
	case rules.HasLt():
		upper = upperBoundLt
		upperValue = rules.GetLt()
		hasRule = true
	case rules.HasLte():
		upper = upperBoundLte
		upperValue = rules.GetLte()
		hasRule = true
	}

	var constVal *T
	if rules.HasConst() {
		constVal = ptr(rules.GetConst())
		hasRule = true
	}

	var inVals []T
	if inVals = rules.GetIn(); len(inVals) > 0 {
		hasRule = true
	}

	var notInVals []T
	if notInVals = rules.GetNotIn(); len(notInVals) > 0 {
		hasRule = true
	}

	type finiteInterface interface {
		HasFinite() bool
		GetFinite() bool
	}

	finite := false
	if fi, ok := (any)(rules).(finiteInterface); ok && fi.HasFinite() {
		finite = fi.GetFinite()
		hasRule = true
	}

	if !hasRule {
		return nil
	}

	return nativeNumericCompare[T]{
		base:      base,
		config:    config,
		lo:        lowerValue,
		lower:     lower,
		hi:        upperValue,
		upper:     upper,
		constVal:  constVal,
		inVals:    inVals,
		notInVals: notInVals,
		finite:    finite,
	}
}

// numericRules is satisfied by all generated numeric rules types
// (Int32Rules, Int64Rules, UInt32Rules, etc.).
//
//nolint:interfacebloat
type numericRules[T numericValue] interface {
	HasGt() bool
	GetGt() T
	HasGte() bool
	GetGte() T
	HasLt() bool
	GetLt() T
	HasLte() bool
	GetLte() T
	HasConst() bool
	GetConst() T
	GetIn() []T
	GetNotIn() []T
	ProtoReflect() protoreflect.Message
}

// numericDescriptors bundles the field descriptors for a single numeric
// rules type (e.g., Int32Rules). Used to construct violation rule paths.
type numericDescriptors struct {
	ruleDesc   protoreflect.FieldDescriptor // FieldRules.{type} (e.g., "int32")
	gtDesc     protoreflect.FieldDescriptor
	gteDesc    protoreflect.FieldDescriptor
	ltDesc     protoreflect.FieldDescriptor
	lteDesc    protoreflect.FieldDescriptor
	constDesc  protoreflect.FieldDescriptor
	inDesc     protoreflect.FieldDescriptor
	notInDesc  protoreflect.FieldDescriptor
	finiteDesc protoreflect.FieldDescriptor
}

func makeNumericDescriptors(
	fieldName string,
	rulesMsg protoreflect.ProtoMessage,
) numericDescriptors {
	rulesDesc := rulesMsg.ProtoReflect().Descriptor()
	var finiteDesc protoreflect.FieldDescriptor
	if rulesDesc.Name() == "FloatRules" || rulesDesc.Name() == "DoubleRules" {
		finiteDesc = rulesDesc.Fields().ByName("finite")
	}
	return numericDescriptors{
		ruleDesc:   fieldRulesDesc.Fields().ByName(protoreflect.Name(fieldName)),
		gtDesc:     rulesDesc.Fields().ByName("gt"),
		gteDesc:    rulesDesc.Fields().ByName("gte"),
		ltDesc:     rulesDesc.Fields().ByName("lt"),
		lteDesc:    rulesDesc.Fields().ByName("lte"),
		constDesc:  rulesDesc.Fields().ByName("const"),
		inDesc:     rulesDesc.Fields().ByName("in"),
		notInDesc:  rulesDesc.Fields().ByName("not_in"),
		finiteDesc: finiteDesc,
	}
}

// numericTypeConfig holds all type-specific operations and metadata
// for a single proto numeric kind.
type numericTypeConfig[T numericValue] struct {
	typeName      string                     // proto rule prefix: "int32", "sint32", "float", etc.
	descs         numericDescriptors         // descriptor bundle for rule path construction
	extractVal    func(protoreflect.Value) T // val.Int/Uint/Float + cast
	makeRuleVal   func(T) protoreflect.Value // ValueOfInt32, ValueOfFloat32, etc.
	nanFailsRange bool                       // true only for float32, float64
}

//nolint:gochecknoglobals
var (
	int32Descs    = makeNumericDescriptors("int32", (*validate.Int32Rules)(nil))
	sint32Descs   = makeNumericDescriptors("sint32", (*validate.SInt32Rules)(nil))
	sfixed32Descs = makeNumericDescriptors("sfixed32", (*validate.SFixed32Rules)(nil))
	int64Descs    = makeNumericDescriptors("int64", (*validate.Int64Rules)(nil))
	sint64Descs   = makeNumericDescriptors("sint64", (*validate.SInt64Rules)(nil))
	sfixed64Descs = makeNumericDescriptors("sfixed64", (*validate.SFixed64Rules)(nil))
	uint32Descs   = makeNumericDescriptors("uint32", (*validate.UInt32Rules)(nil))
	fixed32Descs  = makeNumericDescriptors("fixed32", (*validate.Fixed32Rules)(nil))
	uint64Descs   = makeNumericDescriptors("uint64", (*validate.UInt64Rules)(nil))
	fixed64Descs  = makeNumericDescriptors("fixed64", (*validate.Fixed64Rules)(nil))
	floatDescs    = makeNumericDescriptors("float", (*validate.FloatRules)(nil))
	doubleDescs   = makeNumericDescriptors("double", (*validate.DoubleRules)(nil))
)

//nolint:gochecknoglobals
var (
	int32Config = numericTypeConfig[int32]{
		typeName:    "int32",
		descs:       int32Descs,
		extractVal:  func(v protoreflect.Value) int32 { return int32(v.Int()) },
		makeRuleVal: protoreflect.ValueOfInt32,
	}
	sint32Config = numericTypeConfig[int32]{
		typeName:    "sint32",
		descs:       sint32Descs,
		extractVal:  func(v protoreflect.Value) int32 { return int32(v.Int()) },
		makeRuleVal: protoreflect.ValueOfInt32,
	}
	sfixed32Config = numericTypeConfig[int32]{
		typeName:    "sfixed32",
		descs:       sfixed32Descs,
		extractVal:  func(v protoreflect.Value) int32 { return int32(v.Int()) },
		makeRuleVal: protoreflect.ValueOfInt32,
	}
	int64Config = numericTypeConfig[int64]{
		typeName:    "int64",
		descs:       int64Descs,
		extractVal:  func(v protoreflect.Value) int64 { return v.Int() },
		makeRuleVal: protoreflect.ValueOfInt64,
	}
	sint64Config = numericTypeConfig[int64]{
		typeName:    "sint64",
		descs:       sint64Descs,
		extractVal:  func(v protoreflect.Value) int64 { return v.Int() },
		makeRuleVal: protoreflect.ValueOfInt64,
	}
	sfixed64Config = numericTypeConfig[int64]{
		typeName:    "sfixed64",
		descs:       sfixed64Descs,
		extractVal:  func(v protoreflect.Value) int64 { return v.Int() },
		makeRuleVal: protoreflect.ValueOfInt64,
	}
	uint32Config = numericTypeConfig[uint32]{
		typeName:    "uint32",
		descs:       uint32Descs,
		extractVal:  func(v protoreflect.Value) uint32 { return uint32(v.Uint()) },
		makeRuleVal: protoreflect.ValueOfUint32,
	}
	fixed32Config = numericTypeConfig[uint32]{
		typeName:    "fixed32",
		descs:       fixed32Descs,
		extractVal:  func(v protoreflect.Value) uint32 { return uint32(v.Uint()) },
		makeRuleVal: protoreflect.ValueOfUint32,
	}
	uint64Config = numericTypeConfig[uint64]{
		typeName:    "uint64",
		descs:       uint64Descs,
		extractVal:  func(v protoreflect.Value) uint64 { return v.Uint() },
		makeRuleVal: protoreflect.ValueOfUint64,
	}
	fixed64Config = numericTypeConfig[uint64]{
		typeName:    "fixed64",
		descs:       fixed64Descs,
		extractVal:  func(v protoreflect.Value) uint64 { return v.Uint() },
		makeRuleVal: protoreflect.ValueOfUint64,
	}
	floatConfig = numericTypeConfig[float32]{
		typeName:      "float",
		descs:         floatDescs,
		extractVal:    func(v protoreflect.Value) float32 { return float32(v.Float()) },
		makeRuleVal:   protoreflect.ValueOfFloat32,
		nanFailsRange: true,
	}
	doubleConfig = numericTypeConfig[float64]{
		typeName:      "double",
		descs:         doubleDescs,
		extractVal:    func(v protoreflect.Value) float64 { return v.Float() },
		makeRuleVal:   protoreflect.ValueOfFloat64,
		nanFailsRange: true,
	}
)

// lowerBound describes which lower bound constraint is active.
type lowerBound int

const (
	lowerBoundNone lowerBound = iota
	// lowerBoundGte is an inclusive lower bound (>=).
	lowerBoundGte
	// lowerBoundGt is an exclusive lower bound (>).
	lowerBoundGt
)

// upperBound describes which upper bound constraint is active.
type upperBound int

const (
	upperBoundNone upperBound = iota
	upperBoundLt
	upperBoundLte
)

// nativeNumericCompare is a native Go evaluator for numeric gt/gte/lt/lte/
// const/in/not_in rules. It replaces CEL evaluation with direct Go comparisons.
type nativeNumericCompare[T numericValue] struct {
	base
	config    numericTypeConfig[T]
	lo        T          // lower bound value (gt or gte threshold)
	lower     lowerBound // gt (exclusive) or gte (inclusive)
	hi        T          // upper bound value (lt or lte threshold)
	upper     upperBound // none, lt, or lte
	constVal  *T         // constant value for comparison
	inVals    []T        // slice of values for IN comparison
	notInVals []T        // slice of values for NOT_IN comparison
	finite    bool       // true if the value is finite (not NaN or Infinity)
}

// belowLo reports whether v violates the lower bound.
func (n nativeNumericCompare[T]) belowLo(v T) bool {
	if n.lower == lowerBoundGt {
		return v <= n.lo
	}
	return v < n.lo
}

// aboveHi reports whether v violates the upper bound.
func (n nativeNumericCompare[T]) aboveHi(v T) bool {
	if n.upper == upperBoundLt {
		return v >= n.hi
	}
	return v > n.hi
}

// isNormalRange reports whether lo and hi form a normal (non-exclusive) range.
func (n nativeNumericCompare[T]) isNormalRange() bool {
	return n.hi >= n.lo
}

func (n nativeNumericCompare[T]) loDesc() protoreflect.FieldDescriptor {
	if n.lower == lowerBoundGt {
		return n.config.descs.gtDesc
	}
	return n.config.descs.gteDesc
}

func (n nativeNumericCompare[T]) hiDesc() protoreflect.FieldDescriptor {
	if n.upper == upperBoundLt {
		return n.config.descs.ltDesc
	}
	return n.config.descs.lteDesc
}

func (n nativeNumericCompare[T]) gtRulePrefix() string {
	if n.lower == lowerBoundGt {
		return n.config.typeName + ".gt"
	}
	return n.config.typeName + ".gte"
}

func (n nativeNumericCompare[T]) ltRulePrefix() string {
	if n.upper == upperBoundLt {
		return n.config.typeName + ".lt"
	}
	return n.config.typeName + ".lte"
}

func (n nativeNumericCompare[T]) gtltRule() string {
	if n.lower != lowerBoundNone {
		prefix := n.gtRulePrefix()
		switch n.upper {
		case upperBoundLt:
			prefix += "_lt"
			if !n.isNormalRange() {
				prefix += "_exclusive"
			}
		case upperBoundLte:
			prefix += "_lte"
			if !n.isNormalRange() {
				prefix += "_exclusive"
			}
		}
		return prefix
	}
	return n.ltRulePrefix()
}

func (n nativeNumericCompare[T]) loMessage() string {
	if n.lower == lowerBoundGt {
		return fmt.Sprintf("greater than %v", n.lo)
	}
	return fmt.Sprintf("greater than or equal to %v", n.lo)
}

func (n nativeNumericCompare[T]) hiMessage() string {
	if n.upper == upperBoundLt {
		return fmt.Sprintf("less than %v", n.hi)
	}
	return fmt.Sprintf("less than or equal to %v", n.hi)
}

func (n nativeNumericCompare[T]) conjunction() string {
	if n.isNormalRange() {
		return "and"
	}
	return "or"
}

func (n nativeNumericCompare[T]) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	valT := n.config.extractVal(val)
	var violations []*Violation

	// for numerics, since it is using generics, we need to instantiate inside the method. We can't have a slice of
	// uninstantiated generic functions in Go. Because it's inside the method, we'll just capture the parameters instead
	// of passing them in.
	type numericProcessor[T numericValue] func() *Violation

	var numericProcessors = []numericProcessor[T]{
		// const support
		func() *Violation {
			if n.constVal != nil && valT != *n.constVal {
				return n.newViolation(n.config.descs.ruleDesc, n.config.descs.constDesc,
					n.config.typeName+".const",
					fmt.Sprintf("must equal %v", *n.constVal),
					val, n.config.makeRuleVal(*n.constVal))
			}
			return nil
		},
		// in support
		func() *Violation {
			if len(n.inVals) > 0 && !slices.Contains(n.inVals, valT) {
				return n.newViolation(n.config.descs.ruleDesc, n.config.descs.inDesc,
					n.config.typeName+".in",
					"must be in list "+formatList(n.inVals),
					val, n.config.makeRuleVal(valT))
			}
			return nil
		},
		// not in support
		func() *Violation {
			if len(n.notInVals) > 0 && slices.Contains(n.notInVals, valT) {
				return n.newViolation(n.config.descs.ruleDesc, n.config.descs.notInDesc,
					n.config.typeName+".not_in",
					"must not be in list "+formatList(n.notInVals),
					val, n.config.makeRuleVal(valT))
			}
			return nil
		},
		// finite support
		func() *Violation {
			if n.finite && (math.IsNaN(float64(valT)) || math.IsInf(float64(valT), 0)) {
				return n.newViolation(n.config.descs.ruleDesc, n.config.descs.finiteDesc,
					n.config.typeName+".finite",
					"must be finite",
					val, n.config.makeRuleVal(valT))
			}
			return nil
		},
		// range support
		func() *Violation {
			if n.lower == lowerBoundNone && n.upper == upperBoundNone {
				return nil
			}

			// For float/double, NaN fails all range checks (matches CEL behavior).
			isNaN := n.config.nanFailsRange && math.IsNaN(float64(valT))

			switch {
			case n.lower == lowerBoundNone:
				if isNaN || n.aboveHi(valT) {
					return n.newViolation(n.config.descs.ruleDesc, n.hiDesc(),
						n.gtltRule(), "must be "+n.hiMessage(),
						val, n.config.makeRuleVal(n.hi))
				}
			case n.upper == upperBoundNone:
				if isNaN || n.belowLo(valT) {
					return n.newViolation(n.config.descs.ruleDesc, n.loDesc(),
						n.gtltRule(), "must be "+n.loMessage(),
						val, n.config.makeRuleVal(n.lo))
				}
			default:
				var failure bool
				if n.isNormalRange() {
					failure = isNaN || n.aboveHi(valT) || n.belowLo(valT)
				} else {
					failure = isNaN || (n.aboveHi(valT) && n.belowLo(valT))
				}
				if failure {
					return n.newViolation(n.config.descs.ruleDesc, n.loDesc(),
						n.gtltRule(),
						fmt.Sprintf("must be %s %s %s", n.loMessage(), n.conjunction(), n.hiMessage()),
						val, n.config.makeRuleVal(n.lo))
				}
			}
			return nil
		},
	}

	for _, numericProcessor := range numericProcessors {
		violation := numericProcessor()
		if violation != nil {
			violations = append(violations, violation)
			if cfg.failFast {
				break
			}
		}
	}

	if len(violations) > 0 {
		return &ValidationError{
			Violations: violations,
		}
	}
	return nil
}

func (n nativeNumericCompare[T]) Tautology() bool {
	return false
}

var _ evaluator = nativeNumericCompare[int32]{}

func ptr[T any](v T) *T { return &v }

// formatList formats a slice as "list [val1, val2]" to match CEL message format.
func formatList[T any](vals []T) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%v", v)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
