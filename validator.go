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
	"sync"

	pvcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var getGlobalValidator = sync.OnceValues(func() (*Validator, error) { return New() })

// Validator performs validation on any proto.Message values. The Validator is
// safe for concurrent use.
type Validator struct {
	builder  *builder
	failFast bool
}

// New creates a Validator with the given options. An error may occur in setting
// up the CEL execution environment if the configuration is invalid. See the
// individual ValidatorOption for how they impact the fallibility of New.
func New(options ...ValidatorOption) (*Validator, error) {
	cfg := config{
		extensionTypeResolver: protoregistry.GlobalTypes,
	}
	for _, opt := range options {
		opt(&cfg)
	}

	env, err := cel.NewEnv(cel.Lib(pvcel.NewLibrary()))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct CEL environment: %w", err)
	}

	bldr := newBuilder(
		env,
		cfg.disableLazy,
		cfg.extensionTypeResolver,
		cfg.allowUnknownFields,
		cfg.desc...,
	)

	return &Validator{
		failFast: cfg.failFast,
		builder:  bldr,
	}, nil
}

// Validate checks that message satisfies its constraints. Constraints are
// defined within the Protobuf file as options from the buf.validate package.
// An error is returned if the constraints are violated (ValidationError), the
// evaluation logic for the message cannot be built (CompilationError), or
// there is a type error when attempting to evaluate a CEL expression
// associated with the message (RuntimeError).
func (v *Validator) Validate(msg proto.Message) error {
	if msg == nil {
		return nil
	}
	refl := msg.ProtoReflect()
	eval := v.builder.Load(refl.Descriptor())
	err := eval.EvaluateMessage(refl, v.failFast)
	finalizeViolationPaths(err)
	return err
}

// Validate uses a global instance of Validator constructed with no ValidatorOptions and
// calls its Validate function. For the vast majority of validation cases, using this global
// function is safe and acceptable. If you need to provide i.e. a custom
// ExtensionTypeResolver, you'll need to construct a Validator.
func Validate(msg proto.Message) error {
	globalValidator, err := getGlobalValidator()
	if err != nil {
		return err
	}
	return globalValidator.Validate(msg)
}

type config struct {
	failFast              bool
	disableLazy           bool
	desc                  []protoreflect.MessageDescriptor
	extensionTypeResolver protoregistry.ExtensionTypeResolver
	allowUnknownFields    bool
}

// A ValidatorOption modifies the default configuration of a Validator. See the
// individual options for their defaults and affects on the fallibility of
// configuring a Validator.
type ValidatorOption func(*config)

// WithFailFast specifies whether validation should fail on the first constraint
// violation encountered or if all violations should be accumulated. By default,
// all violations are accumulated.
func WithFailFast() ValidatorOption {
	return func(cfg *config) {
		cfg.failFast = true
	}
}

// WithMessages allows warming up the Validator with messages that are
// expected to be validated. Messages included transitively (i.e., fields with
// message values) are automatically handled.
func WithMessages(messages ...proto.Message) ValidatorOption {
	desc := make([]protoreflect.MessageDescriptor, len(messages))
	for i, msg := range messages {
		desc[i] = msg.ProtoReflect().Descriptor()
	}
	return WithMessageDescriptors(desc...)
}

// WithMessageDescriptors allows warming up the Validator with message
// descriptors that are expected to be validated. Messages included transitively
// (i.e., fields with message values) are automatically handled.
func WithMessageDescriptors(descriptors ...protoreflect.MessageDescriptor) ValidatorOption {
	return func(cfg *config) {
		cfg.desc = append(cfg.desc, descriptors...)
	}
}

// WithDisableLazy prevents the Validator from lazily building validation logic
// for a message it has not encountered before. Disabling lazy logic
// additionally eliminates any internal locking as the validator becomes
// read-only.
//
// Note: All expected messages must be provided by WithMessages or
// WithMessageDescriptors during initialization.
func WithDisableLazy() ValidatorOption {
	return func(cfg *config) {
		cfg.disableLazy = true
	}
}

// WithExtensionTypeResolver specifies a resolver to use when reparsing unknown
// extension types. When dealing with dynamic file descriptor sets, passing this
// option will allow extensions to be resolved using a custom resolver.
//
// To ignore unknown extension fields, use the [WithAllowUnknownFields] option.
// Note that this may result in messages being treated as valid even though not
// all constraints are being applied.
func WithExtensionTypeResolver(extensionTypeResolver protoregistry.ExtensionTypeResolver) ValidatorOption {
	return func(c *config) {
		c.extensionTypeResolver = extensionTypeResolver
	}
}

// WithAllowUnknownFields specifies if the presence of unknown field constraints
// should cause compilation to fail with an error. When set to false, an unknown
// field will simply be ignored, which will cause constraints to silently not be
// applied. This condition may occur if a predefined constraint definition isn't
// present in the extension type resolver, or when passing dynamic messages with
// standard constraints defined in a newer version of protovalidate. The default
// value is false, to prevent silently-incorrect validation from occurring.
func WithAllowUnknownFields() ValidatorOption {
	return func(c *config) {
		c.allowUnknownFields = true
	}
}
