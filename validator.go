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
	"context"
	"fmt"
	"sync"

	pvcel "buf.build/go/protovalidate/cel"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	getGlobalValidator = sync.OnceValues(func() (ContextValidator, error) { return New() })

	// GlobalValidator provides access to the global Validator instance that is
	// used by the [Validate] function. This is intended to be used by libraries
	// that use protovalidate. This Validator can be used as a default when the
	// user does not specify a Validator instance to use.
	//
	// Using the global Validator instance (either through [Validator] or via
	// GlobalValidator) will result in lower memory usage than using multiple
	// Validator instances, because each Validator instance has its own caches.
	GlobalValidator ContextValidator = globalValidator{}
)

// Validator performs validation on any proto.Message values. The Validator is
// safe for concurrent use.
type Validator interface {
	// Validate checks that message satisfies its rules. Rules are
	// defined within the Protobuf file as options from the buf.validate
	// package. An error is returned if the rules are violated
	// (ValidationError), the evaluation logic for the message cannot be built
	// (CompilationError), or there is a type error when attempting to evaluate
	// a CEL expression associated with the message (RuntimeError).
	Validate(msg proto.Message, options ...ValidationOption) error
}

// ContextValidator is a [Validator] that additionally supports context-aware
// validation. The value returned by [New] and the [GlobalValidator] both
// implement it. Callers that hold only a [Validator] can type-assert it to
// ContextValidator, or use the package-level [ValidateContext] function.
//
// ContextValidator is a separate interface (rather than a method on Validator)
// to preserve backward compatibility for external implementations of Validator.
type ContextValidator interface {
	Validator

	// ValidateContext behaves like Validate, but stops early if ctx is
	// cancelled or its deadline is exceeded, returning the context error
	// (context.Canceled or context.DeadlineExceeded), detectable with
	// errors.Is.
	//
	// Cancellation is observed at message, repeated and map boundaries, before
	// each CEL rule, and inside CEL comprehension macros (all, exists, map,
	// filter): the loops in which an expensive expression spends its time. A
	// CEL expression containing no comprehension is not interruptible once it
	// begins; see [WithCELInterruptCheckFrequency].
	//
	// Passing a context that can never be cancelled, such as context.Background,
	// costs nothing: evaluation then takes exactly the same path as Validate.
	ValidateContext(ctx context.Context, msg proto.Message, options ...ValidationOption) error
}

// New creates a ContextValidator with the given options. An error may occur in
// setting up the CEL execution environment if the configuration is invalid. See
// the individual ValidatorOption for how they impact the fallibility of New.
//
// The returned value implements [ContextValidator], and therefore also
// [Validator]; assigning it to a Validator remains valid.
func New(options ...ValidatorOption) (ContextValidator, error) {
	cfg := config{
		extensionTypeResolver:   protoregistry.GlobalTypes,
		nowFn:                   timestamppb.Now,
		interruptCheckFrequency: pvcel.DefaultInterruptCheckFrequency,
	}
	for _, opt := range options {
		opt.applyToValidator(&cfg)
	}

	reg, err := newRegistry()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct CEL type registry: %w", err)
	}
	env, err := cel.NewEnv(
		cel.CustomTypeProvider(reg),
		cel.CustomTypeAdapter(reg),
		cel.Lib(pvcel.NewLibrary(
			pvcel.WithInterruptCheckFrequency(cfg.interruptCheckFrequency),
		)),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct CEL environment: %w", err)
	}

	bldr := newBuilder(
		env,
		cfg.disableLazy,
		cfg.extensionTypeResolver,
		cfg.allowUnknownFields,
		cfg.disableNativeRules,
		cfg.desc...,
	)

	baseCfg := &validationConfig{
		failFast: cfg.failFast,
		filter:   nopFilter{},
		nowFn:    cfg.nowFn,
	}
	cancellableCfg := baseCfg.clone()
	cancellableCfg.cancellable = true

	return &validator{
		builder:        bldr,
		cfg:            baseCfg,
		cancellableCfg: cancellableCfg,
	}, nil
}

type validator struct {
	builder *builder
	// cfg and cancellableCfg are identical but for the cancellable flag. Both
	// are built once and treated as immutable, so a ValidateContext call with no
	// ValidationOptions can select one without allocating a copy.
	cfg            *validationConfig
	cancellableCfg *validationConfig
}

func (v *validator) Validate(
	msg proto.Message,
	options ...ValidationOption,
) error {
	// Pass cancellable=false without consulting the context: a background
	// context can never be cancelled, and probing it would add interface calls
	// to the hot path.
	return v.validate(context.Background(), false, msg, options...)
}

func (v *validator) ValidateContext(
	ctx context.Context,
	msg proto.Message,
	options ...ValidationOption,
) error {
	// Check the context before anything else (including a nil msg) so that a
	// cancelled context always yields its error, as documented.
	if err := ctx.Err(); err != nil {
		return err
	}
	return v.validate(ctx, ctx.Done() != nil, msg, options...)
}

func (v *validator) validate(
	ctx context.Context,
	cancellable bool,
	msg proto.Message,
	options ...ValidationOption,
) error {
	if msg == nil {
		return nil
	}
	cfg := v.cfg
	if cancellable {
		cfg = v.cancellableCfg
	}
	if len(options) > 0 {
		cfg = cfg.clone()
		for _, opt := range options {
			opt.applyToValidation(cfg)
		}
	}
	refl := msg.ProtoReflect()
	eval := v.builder.Load(refl.Descriptor())
	err := eval.EvaluateMessageContext(ctx, refl, cfg)
	finalizeViolationPaths(err)
	return err
}

// Validate uses a global instance of Validator constructed with no ValidatorOptions and
// calls its Validate function. For the vast majority of validation cases, using this global
// function is safe and acceptable. If you need to provide i.e. a custom
// ExtensionTypeResolver, you'll need to construct a Validator.
func Validate(msg proto.Message, options ...ValidationOption) error {
	return ValidateContext(context.Background(), msg, options...)
}

// ValidateContext is the context-aware counterpart to Validate, using the
// global Validator instance. See [ContextValidator.ValidateContext].
func ValidateContext(ctx context.Context, msg proto.Message, options ...ValidationOption) error {
	globalValidator, err := getGlobalValidator()
	if err != nil {
		return err
	}
	return globalValidator.ValidateContext(ctx, msg, options...)
}

type config struct {
	failFast                bool
	disableLazy             bool
	desc                    []protoreflect.MessageDescriptor
	extensionTypeResolver   protoregistry.ExtensionTypeResolver
	allowUnknownFields      bool
	nowFn                   func() *timestamppb.Timestamp
	disableNativeRules      bool
	interruptCheckFrequency uint
}

type validationConfig struct {
	failFast bool
	filter   Filter
	nowFn    func() *timestamppb.Timestamp
	// cancellable reports whether the context passed to ValidateContext can
	// ever be cancelled (ctx.Done() != nil). When false, evaluation skips every
	// per-node context check and cel-go's ContextEval, so the Validate path
	// costs exactly what it did before context support existed.
	cancellable bool
}

func (cfg *validationConfig) clone() *validationConfig {
	clonedCfg := *cfg
	return &clonedCfg
}

type globalValidator struct{}

func (globalValidator) Validate(msg proto.Message, options ...ValidationOption) error {
	return Validate(msg, options...)
}

func (globalValidator) ValidateContext(ctx context.Context, msg proto.Message, options ...ValidationOption) error {
	return ValidateContext(ctx, msg, options...)
}
