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
	"testing"
	"time"

	pvcel "buf.build/go/protovalidate/cel"
	pb "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/sourcecontextpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				pb.HasMsgExprs_builder{X: 2, Y: 43}.Build(),
				false,
			},
			{
				pb.HasMsgExprs_builder{X: 9, Y: 8}.Build(),
				true,
			},
		}

		for _, test := range tests {
			err := val.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestValidator_ValidateGlobal(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				pb.HasMsgExprs_builder{X: 2, Y: 43}.Build(),
				false,
			},
			{
				pb.HasMsgExprs_builder{X: 9, Y: 8}.Build(),
				true,
			},
		}

		for _, test := range tests {
			err := Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestGlobalValidator(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				pb.HasMsgExprs_builder{X: 2, Y: 43}.Build(),
				false,
			},
			{
				pb.HasMsgExprs_builder{X: 9, Y: 8}.Build(),
				true,
			},
		}

		for _, test := range tests {
			err := GlobalValidator.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestRecursive(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)

	selfRec := pb.SelfRecursive_builder{X: 123, Turtle: pb.SelfRecursive_builder{X: 456}.Build()}.Build()
	err = val.Validate(selfRec)
	require.NoError(t, err)

	loopRec := pb.LoopRecursiveA_builder{B: &pb.LoopRecursiveB{}}.Build()
	err = val.Validate(loopRec)
	require.NoError(t, err)
}

func TestValidator_ValidateOneof(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	oneofMessage := pb.MsgHasOneof_builder{X: proto.String("foo")}.Build()
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = pb.MsgHasOneof_builder{Y: proto.Int32(42)}.Build()
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = pb.MsgHasOneof_builder{Msg: pb.HasMsgExprs_builder{X: 4, Y: 50}.Build()}.Build()
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{}
	err = val.Validate(oneofMessage)
	assert.Error(t, err)
}

func TestValidator_ValidateRepeatedFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	repeatMessage := pb.MsgHasRepeated_builder{
		X: []float32{1, 2, 3},
		Y: []string{"foo", "bar"},
		Z: []*pb.HasMsgExprs{
			pb.HasMsgExprs_builder{
				X: 4,
				Y: 55,
			}.Build(),
			pb.HasMsgExprs_builder{
				X: 4,
				Y: 60,
			}.Build(),
		},
	}.Build()
	err = val.Validate(repeatMessage)
	require.NoError(t, err)
}

func TestValidator_ValidateMapFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	mapMessage := pb.MsgHasMap_builder{
		Int32Map:   map[int32]int32{-1: 1, 2: 2},
		StringMap:  map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		MessageMap: map[int64]*pb.LoopRecursiveA{0: nil},
	}.Build()
	err = val.Validate(mapMessage)
	require.Error(t, err)
}

func TestValidator_Validate_TransitiveFieldRules(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.TransitiveFieldRule_builder{
		Mask: &fieldmaskpb.FieldMask{Paths: []string{"foo", "bar"}},
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_MultipleStepsTransitiveFieldRules(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.MultipleStepsTransitiveFieldRules_builder{
		Api: &apipb.Api{
			SourceContext: &sourcecontextpb.SourceContext{
				FileName: "path/file",
			},
		},
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_FieldOfTypeAny(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	simple := pb.Simple_builder{S: "foo"}.Build()
	anyFromSimple, err := anypb.New(simple)
	require.NoError(t, err)
	msg := pb.FieldOfTypeAny_builder{
		Any: anyFromSimple,
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_CelMapOnARepeated(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.CelMapOnARepeated_builder{Values: []*pb.CelMapOnARepeated_Value{
		pb.CelMapOnARepeated_Value_builder{Name: "foo"}.Build(),
		pb.CelMapOnARepeated_Value_builder{Name: "bar"}.Build(),
		pb.CelMapOnARepeated_Value_builder{Name: "baz"}.Build(),
	}}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.SetValues(append(msg.GetValues(), pb.CelMapOnARepeated_Value_builder{Name: "foo"}.Build()))
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
}

func TestValidator_Validate_RepeatedItemCel(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.RepeatedItemCel_builder{Paths: []string{"foo"}}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.SetPaths(append(msg.GetPaths(), " bar"))
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "paths.no_space", valErr.Violations[0].Proto.GetRuleId())
	pathsFd := msg.ProtoReflect().Descriptor().Fields().ByName("paths")
	err = val.Validate(msg, WithFilter(FilterFunc(func(m protoreflect.Message, d protoreflect.Descriptor) bool {
		return !(m.Interface() == msg && d == pathsFd)
	})))
	require.NoError(t, err)
}

func TestValidator_Validate_Filter(t *testing.T) {
	t.Parallel()

	t.Run("FilterField", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.Person{}
		err = val.Validate(msg)
		valErr := &ValidationError{}
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 3)
		idFd := msg.ProtoReflect().Descriptor().Fields().ByName("id")
		err = val.Validate(msg, WithFilter(FilterFunc(func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
			return d == idFd
		})))
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 1)
	})

	t.Run("FilterInvalid", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.InvalidRules{}
		err = val.Validate(msg)
		require.Error(t, err)
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, _ protoreflect.Descriptor) bool {
				return false
			},
		)))
		require.NoError(t, err)
	})

	t.Run("FilterNested", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.NestedRules_builder{
			Field:         &pb.AllRuleTypes{},
			RepeatedField: []*pb.AllRuleTypes{{}},
			MapField:      map[string]*pb.AllRuleTypes{"test": {}},
		}.Build()
		var descs []string
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				descs = append(descs, string(d.FullName()))
				return false
			},
		)))
		require.Equal(t, []string{
			"tests.example.v1.NestedRules",
			"tests.example.v1.NestedRules.required_oneof",
			"tests.example.v1.NestedRules.field",
			"tests.example.v1.NestedRules.field2",
			"tests.example.v1.NestedRules.repeated_field",
			"tests.example.v1.NestedRules.map_field",
		}, descs)
		require.NoError(t, err)
		descs = []string{}
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				descs = append(descs, string(d.FullName()))
				return true
			},
		)))
		require.Equal(t, []string{
			"tests.example.v1.NestedRules",
			"tests.example.v1.NestedRules.required_oneof",
			"tests.example.v1.NestedRules.field",
			"tests.example.v1.AllRuleTypes",
			"tests.example.v1.AllRuleTypes.required_oneof",
			"tests.example.v1.AllRuleTypes.field",
			"tests.example.v1.NestedRules.field2",
			"tests.example.v1.NestedRules.repeated_field",
			"tests.example.v1.AllRuleTypes",
			"tests.example.v1.AllRuleTypes.required_oneof",
			"tests.example.v1.AllRuleTypes.field",
			"tests.example.v1.NestedRules.map_field",
			"tests.example.v1.AllRuleTypes",
			"tests.example.v1.AllRuleTypes.required_oneof",
			"tests.example.v1.AllRuleTypes.field",
		}, descs)
		require.Error(t, err)
	})

	t.Run("FilterIncludeCompilationError", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.MixedValidInvalidRules_builder{
			StringFieldBoolRule: "foo",
			ValidStringRule:     "bar",
		}.Build()
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				return d == msg.ProtoReflect().Descriptor().Fields().Get(0)
			},
		)))
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})

	t.Run("FilterExcludeCompilationError", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.MixedValidInvalidRules_builder{
			ValidStringRule:     "bar",
			StringFieldBoolRule: "foo",
		}.Build()
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				return d == msg.ProtoReflect().Descriptor().Fields().Get(1)
			},
		)))
		require.Error(t, err)
		compErr := &CompilationError{}
		require.NotErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 1)
	})
}

func TestValidator_ValidateCompilationError(t *testing.T) {
	t.Parallel()

	t.Run("CompilationErrorNoViolations", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.MismatchRules{}
		err = val.Validate(msg)
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})

	t.Run("CompilationErrorWithViolations", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.MixedValidInvalidRules_builder{
			StringFieldBoolRule: "foo",
			ValidStringRule:     "bar",
		}.Build()
		err = val.Validate(msg)
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})
}

func TestValidator_WithDisableLazy(t *testing.T) {
	t.Parallel()

	t.Run("no_evaluator_available", func(t *testing.T) {
		t.Parallel()
		val, err := New(
			WithDisableLazy(),
		)
		require.NoError(t, err)
		msg := &pb.Simple{}
		err = val.Validate(msg)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		require.ErrorContains(t, err, "no evaluator available for tests.example.v1.Simple")
	})
}

func TestValidator_WithMessages(t *testing.T) {
	t.Parallel()

	t.Run("defers_compile_error", func(t *testing.T) {
		t.Parallel()
		val, err := New(
			WithMessages(&pb.MismatchRules{}),
			WithDisableLazy(), // disable lazy to ensure pre-warmed descriptors are used
		)
		require.NoError(t, err)
		msg := &pb.MismatchRules{}
		err = val.Validate(msg)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		require.ErrorContains(t, err, "expected rule \"buf.validate.FieldRules.string\"")
	})
}

func TestValidator_WithNowFunc_Issue211(t *testing.T) {
	t.Parallel()

	nowFn := func() *timestamppb.Timestamp {
		return timestamppb.New(time.Now().Add(time.Hour))
	}

	msg := pb.Issue211_builder{
		Value: timestamppb.New(time.Now().Add(time.Minute)),
	}.Build()
	val, err := New()
	require.NoError(t, err)
	err = val.Validate(msg)
	require.NoError(t, err)
	err = val.Validate(msg, WithNowFunc(nowFn))
	require.Error(t, err)

	val, err = New(WithNowFunc(nowFn))
	require.NoError(t, err)
	err = val.Validate(msg)
	require.Error(t, err)
	err = val.Validate(msg, WithNowFunc(timestamppb.Now))
	require.NoError(t, err)
}

func TestValidator_Validate_Issue141(t *testing.T) {
	t.Parallel()

	t.Run("FieldWithIssue", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.FieldWithIssue{}
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})

	t.Run("OneTwo", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.OneTwo_builder{
			Field1: pb.F1_builder{
				Field: &pb.FieldWithIssue{},
			}.Build(),
		}.Build()
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})

	t.Run("TwoOne", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := pb.TwoOne_builder{
			Field1: pb.F1_builder{
				Field: &pb.FieldWithIssue{},
			}.Build(),
		}.Build()
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})
}

func TestValidator_Validate_Issue148(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.Issue148_builder{Test: proto.Int32(1)}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_Issue187(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.Issue187_builder{
		FalseField: proto.Bool(false),
		TrueField:  proto.Bool(true),
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_Issue307(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)

	t.Run("map/valid", func(t *testing.T) {
		t.Parallel()
		name := "child"
		msg := pb.SelfReferentialMap_builder{
			Name: &name,
			Children: map[string]*pb.SelfReferentialMap{
				"a": pb.SelfReferentialMap_builder{Name: &name}.Build(),
				"b": pb.SelfReferentialMap_builder{Name: &name}.Build(),
			},
		}.Build()
		err := val.Validate(msg)
		require.NoError(t, err)
	})

	t.Run("map/invalid", func(t *testing.T) {
		t.Parallel()
		name := "parent"
		msg := pb.SelfReferentialMap_builder{
			Name: &name,
			Children: map[string]*pb.SelfReferentialMap{
				"a": pb.SelfReferentialMap_builder{Name: &name}.Build(),
				"b": pb.SelfReferentialMap_builder{}.Build(), // no name
			},
		}.Build()
		err := val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})

	t.Run("repeated/valid", func(t *testing.T) {
		t.Parallel()
		name := "child"
		msg := pb.SelfReferentialRepeated_builder{
			Name: &name,
			Children: []*pb.SelfReferentialRepeated{
				pb.SelfReferentialRepeated_builder{Name: &name}.Build(),
				pb.SelfReferentialRepeated_builder{Name: &name}.Build(),
			},
		}.Build()
		err := val.Validate(msg)
		require.NoError(t, err)
	})

	t.Run("repeated/invalid", func(t *testing.T) {
		t.Parallel()
		name := "parent"
		msg := pb.SelfReferentialRepeated_builder{
			Name: &name,
			Children: []*pb.SelfReferentialRepeated{
				pb.SelfReferentialRepeated_builder{Name: &name}.Build(),
				pb.SelfReferentialRepeated_builder{}.Build(), // no name
			},
		}.Build()
		err := val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})
}

func TestValidator_Validate_Issue296(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.Issue296_builder{
		Fm: &fieldmaskpb.FieldMask{
			Paths: []string{"a"},
		},
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

// validateOnly implements only the Validator interface. It exists to guard the
// backward-compatibility contract: ValidateContext must stay off Validator, so
// external types implementing only Validate keep satisfying it.
type validateOnly struct{}

func (validateOnly) Validate(proto.Message, ...ValidationOption) error { return nil }

var (
	_ Validator        = validateOnly{}
	_ ContextValidator = (*validator)(nil)
)

func TestValidator_ValidateContext_Cancelled(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = val.ValidateContext(ctx, &pb.Person{})
	assert.ErrorIs(t, err, context.Canceled)
}

// TestValidator_ValidateContext_CancelledNilMessage pins the check order: a
// cancelled context yields its error even when the message is nil, so callers
// using the error as a cancellation signal never miss it.
func TestValidator_ValidateContext_CancelledNilMessage(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = val.ValidateContext(ctx, nil)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestValidator_ValidateContext_DelegatesToValidate(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	// A valid message returns nil through both entry points.
	msg := &pb.Person{Id: 1, Email: "a@b.co", Name: "x", Home: &pb.Coordinates{}}
	assert.Equal(t,
		val.Validate(msg),
		val.ValidateContext(context.Background(), msg),
	)
}

func TestValidateContext_Global(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := ValidateContext(ctx, &pb.Person{})
	assert.ErrorIs(t, err, context.Canceled)
}

// TestValidateContext_InterruptsExpensiveCELExpression covers the motivating
// case: a single CEL rule whose cost grows multiplicatively with message size
// (see BenchCrossReference). Uncancelled it runs for seconds; a deadline must
// cut it short rather than waiting for it to finish.
func TestValidateContext_InterruptsExpensiveCELExpression(t *testing.T) {
	t.Parallel()
	msg := newCrossReference(150) // ~3s of CEL evaluation when uncancelled
	val, err := New(WithMessages(msg), WithDisableLazy())
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	err = val.ValidateContext(ctx, msg)
	elapsed := time.Since(start)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	assert.Less(t, elapsed, time.Second,
		"the deadline must interrupt evaluation, not be noticed after it completes")
}

// TestValidateContext_InterruptCheckFrequency documents why the default
// frequency is 1. cel-go polls ctx.Done() once every `frequency` comprehension
// iterations, so cancellation latency scales with the frequency, while the
// per-iteration bookkeeping happens either way. A coarse frequency therefore
// buys no throughput (see BenchmarkInterruptCheckFrequency) and only makes a
// deadline less meaningful.
func TestValidateContext_InterruptCheckFrequency(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("timing-sensitive; runs a multi-second CEL expression")
	}
	msg := newCrossReference(150)

	latency := func(frequency uint) time.Duration {
		val, err := New(WithMessages(msg), WithDisableLazy(),
			WithCELInterruptCheckFrequency(frequency))
		require.NoError(t, err)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		start := time.Now()
		err = val.ValidateContext(ctx, msg)
		require.ErrorIs(t, err, context.DeadlineExceeded)
		return time.Since(start)
	}

	def := latency(pvcel.DefaultInterruptCheckFrequency)
	coarse := latency(100)

	assert.Less(t, def, 500*time.Millisecond,
		"the default frequency should honor the deadline promptly")
	assert.Greater(t, coarse, 2*def,
		"a coarser frequency should notice the deadline substantially later")
}

// TestWithCELInterruptCheckFrequency_Disabled shows that a frequency of 0 turns
// interrupt checking off: validation still works, but CEL expressions run to
// completion and cannot be cancelled mid-evaluation.
func TestWithCELInterruptCheckFrequency_Disabled(t *testing.T) {
	t.Parallel()
	msg := newCrossReference(8)
	val, err := New(WithMessages(msg), WithDisableLazy(),
		WithCELInterruptCheckFrequency(0))
	require.NoError(t, err)
	require.NoError(t, val.Validate(msg))
	require.NoError(t, val.ValidateContext(context.Background(), msg))
}

func TestValidator_ValidateContext_CancelMidTraversal(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	// Cancel as soon as traversal asks whether to validate anything.
	filter := FilterFunc(func(protoreflect.Message, protoreflect.Descriptor) bool {
		cancel()
		return true
	})
	msg := &pb.BenchRepeatedMessage{X: []*pb.BenchScalar{{}, {}}}
	err = val.ValidateContext(ctx, msg, WithFilter(filter))
	assert.ErrorIs(t, err, context.Canceled)
}
