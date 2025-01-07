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
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pvcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/bufbuild/protovalidate-go/internal/extensions"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// cache is a build-through cache to computed standard constraints.
type cache struct {
	cache map[protoreflect.FieldDescriptor]astSet
}

// newCache constructs a new build-through cache for the standard constraints.
func newCache() cache {
	return cache{
		cache: map[protoreflect.FieldDescriptor]astSet{},
	}
}

// Build creates the standard constraints for the given field. If forItems is
// true, the constraints for repeated list items is built instead of the
// constraints on the list itself.
func (c *cache) Build(
	env *cel.Env,
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	extensionTypeResolver protoregistry.ExtensionTypeResolver,
	allowUnknownFields bool,
	forItems bool,
) (set programSet, err error) {
	constraints, setOneof, done, err := c.resolveConstraints(
		fieldDesc,
		fieldConstraints,
		forItems,
	)
	if done {
		return nil, err
	}

	if err = reparseUnrecognized(extensionTypeResolver, constraints); err != nil {
		return nil, &CompilationError{cause: fmt.Errorf("error reparsing message: %w", err)}
	}
	if !allowUnknownFields && len(constraints.GetUnknown()) > 0 {
		return nil, &CompilationError{cause: fmt.Errorf("unknown constraints in %s; see protovalidate.WithExtensionTypeResolver", constraints.Descriptor().FullName())}
	}

	env, err = c.prepareEnvironment(env, fieldDesc, constraints, forItems)
	if err != nil {
		return nil, err
	}

	var asts astSet
	constraints.Range(func(desc protoreflect.FieldDescriptor, rule protoreflect.Value) bool {
		fieldEnv, compileErr := env.Extend(
			cel.Constant(
				"rule",
				pvcel.ProtoFieldToCELType(desc, true, false),
				pvcel.ProtoFieldToCELValue(desc, rule, false),
			),
		)
		if compileErr != nil {
			err = compileErr
			return false
		}
		precomputedASTs, compileErr := c.loadOrCompileStandardConstraint(fieldEnv, setOneof, desc)
		if compileErr != nil {
			err = compileErr
			return false
		}
		precomputedASTs.SetRuleValue(rule, desc)
		asts = asts.Merge(precomputedASTs)
		return true
	})
	if err != nil {
		return nil, err
	}

	rulesGlobal := cel.Globals(&variable{Name: "rules", Val: constraints.Interface()})
	set, err = asts.ReduceResiduals(rulesGlobal)
	return set, err
}

// resolveConstraints extracts the standard constraints for the specified field. An
// error is returned if the wrong constraints are applied to a field (typically
// if there is a type-mismatch). The done result is true if an error is returned
// or if there are now standard constraints to apply to this field.
func (c *cache) resolveConstraints(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
) (rules protoreflect.Message, fieldRule protoreflect.FieldDescriptor, done bool, err error) {
	constraints := fieldConstraints.ProtoReflect()
	setOneof := constraints.WhichOneof(fieldConstraintsOneofDesc)
	if setOneof == nil {
		return nil, nil, true, nil
	}
	expected, ok := c.getExpectedConstraintDescriptor(fieldDesc, forItems)
	if ok && setOneof.FullName() != expected.FullName() {
		return nil, nil, true, &CompilationError{cause: fmt.Errorf(
			"expected constraint %q, got %q on field %q",
			expected.FullName(),
			setOneof.FullName(),
			fieldDesc.FullName(),
		)}
	}
	if !ok || !constraints.Has(setOneof) {
		return nil, nil, true, nil
	}
	rules = constraints.Get(setOneof).Message()
	return rules, setOneof, false, nil
}

// prepareEnvironment prepares the environment for compiling standard constraint
// expressions.
func (c *cache) prepareEnvironment(
	env *cel.Env,
	fieldDesc protoreflect.FieldDescriptor,
	rules protoreflect.Message,
	forItems bool,
) (*cel.Env, error) {
	env, err := env.Extend(
		cel.Types(rules.Interface()),
		cel.Variable("this", pvcel.ProtoFieldToCELType(fieldDesc, true, forItems)),
		cel.Variable("rules",
			cel.ObjectType(string(rules.Descriptor().FullName()))),
	)
	if err != nil {
		return nil, &CompilationError{cause: fmt.Errorf(
			"failed to extend base environment: %w", err)}
	}
	return env, nil
}

// loadOrCompileStandardConstraint loads the precompiled ASTs for the
// specified constraint field from the Cache if present or precomputes them
// otherwise. The result may be empty if the constraint does not have associated
// CEL expressions.
func (c *cache) loadOrCompileStandardConstraint(
	env *cel.Env,
	setOneOf protoreflect.FieldDescriptor,
	constraintFieldDesc protoreflect.FieldDescriptor,
) (set astSet, err error) {
	if cachedConstraint, ok := c.cache[constraintFieldDesc]; ok {
		return cachedConstraint, nil
	}
	exprs := expressions{
		Constraints: extensions.Resolve[*validate.PredefinedConstraints](
			constraintFieldDesc.Options(),
			validate.E_Predefined,
		).GetCel(),
		RulePath: []*validate.FieldPathElement{
			fieldPathElement(setOneOf),
			fieldPathElement(constraintFieldDesc),
		},
	}
	set, err = compileASTs(exprs, env)
	if err != nil {
		return set, &CompilationError{cause: fmt.Errorf(
			"failed to compile standard constraint %q: %w",
			constraintFieldDesc.FullName(), err)}
	}
	c.cache[constraintFieldDesc] = set
	return set, nil
}

// getExpectedConstraintDescriptor produces the field descriptor from the
// validate.FieldConstraints 'type' oneof that matches the provided target
// field descriptor. If ok is false, the field does not expect any standard
// constraints.
func (c *cache) getExpectedConstraintDescriptor(
	targetFieldDesc protoreflect.FieldDescriptor,
	forItems bool,
) (expected protoreflect.FieldDescriptor, ok bool) {
	switch {
	case targetFieldDesc.IsMap():
		return mapFieldConstraintsDesc, true
	case targetFieldDesc.IsList() && !forItems:
		return repeatedFieldConstraintsDesc, true
	case targetFieldDesc.Kind() == protoreflect.MessageKind,
		targetFieldDesc.Kind() == protoreflect.GroupKind:
		expected, ok = expectedWKTConstraints[targetFieldDesc.Message().FullName()]
		return expected, ok
	default:
		expected, ok = expectedStandardConstraints[targetFieldDesc.Kind()]
		return expected, ok
	}
}

func reparseUnrecognized(
	extensionTypeResolver protoregistry.ExtensionTypeResolver,
	reflectMessage protoreflect.Message,
) error {
	if unknown := reflectMessage.GetUnknown(); len(unknown) > 0 {
		reflectMessage.SetUnknown(nil)
		options := proto.UnmarshalOptions{
			Resolver: extensionTypeResolver,
			Merge:    true,
		}
		if err := options.Unmarshal(unknown, reflectMessage.Interface()); err != nil {
			return err
		}
	}
	return nil
}
