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
	"os"
	"strings"
	"testing"

	pb "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/rodaine/protogofakeit"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func BenchmarkScalar(b *testing.B) {
	benchSuccess(b, &pb.BenchScalar{})
}

func BenchmarkRepeated(b *testing.B) {
	b.Run("Scalar", func(b *testing.B) {
		benchSuccess(b, &pb.BenchRepeatedScalar{})
	})
	b.Run("Message", func(b *testing.B) {
		benchSuccess(b, &pb.BenchRepeatedMessage{})
	})
	b.Run("Unique", func(b *testing.B) {
		b.Run("Scalar", func(b *testing.B) {
			benchSuccess(b, &pb.BenchRepeatedScalarUnique{})
		})
		b.Run("Bytes", func(b *testing.B) {
			benchSuccess(b, &pb.BenchRepeatedBytesUnique{})
		})
	})
}

func BenchmarkMap(b *testing.B) {
	benchSuccess(b, &pb.BenchMap{})
}

func BenchmarkComplexSchema(b *testing.B) {
	benchSuccess(b, &pb.BenchComplexSchema{})
}

func BenchmarkInt32GT(b *testing.B) {
	benchSuccess(b, &pb.BenchGT{})
}

func BenchmarkTestByteMatching(b *testing.B) {
	benchSuccess(b, &pb.TestByteMatching{})
}

func BenchmarkStringMatching(b *testing.B) {
	benchSuccess(b, &pb.StringMatching{})
}

func BenchmarkWrapperTesting(b *testing.B) {
	msg := pb.WrapperTesting_builder{
		I32: &wrapperspb.Int32Value{Value: 11},
		D:   &wrapperspb.DoubleValue{Value: 11},
		F:   &wrapperspb.FloatValue{Value: 11},
		I64: &wrapperspb.Int64Value{Value: 11},
		U64: &wrapperspb.UInt64Value{Value: 11},
		U32: &wrapperspb.UInt32Value{Value: 11},
		B:   &wrapperspb.BoolValue{Value: true},
		S:   &wrapperspb.StringValue{Value: "hello"},
		Bs:  &wrapperspb.BytesValue{Value: []byte("hello")},
	}.Build()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	val, err := New(options...)
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = val.Validate(msg)
	}
}

func TestInt32GT(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.BenchGT{})
}

func TestComplexSchema(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.BenchComplexSchema{})
}

func TestMap(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.BenchMap{})
}

func TestScalar(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.BenchScalar{})
}

func TestByteMatching(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.TestByteMatching{})
}

func TestStringMatching(t *testing.T) {
	t.Parallel()
	testSuccess(t, &pb.StringMatching{})
}

func TestWrapperTesting(t *testing.T) {
	t.Parallel()
	msg := pb.WrapperTesting_builder{
		I32: &wrapperspb.Int32Value{Value: 11},
		D:   &wrapperspb.DoubleValue{Value: 11},
		F:   &wrapperspb.FloatValue{Value: 11},
		I64: &wrapperspb.Int64Value{Value: 11},
		U64: &wrapperspb.UInt64Value{Value: 11},
		U32: &wrapperspb.UInt32Value{Value: 11},
		B:   &wrapperspb.BoolValue{Value: true},
		S:   &wrapperspb.StringValue{Value: "hello"},
		Bs:  &wrapperspb.BytesValue{Value: []byte("hello")},
	}.Build()
	val, err := New(WithMessages(msg), WithDisableLazy())
	require.NoError(t, err)
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestRepeated(t *testing.T) {
	t.Parallel()
	t.Run("Scalar", func(t *testing.T) {
		t.Parallel()
		testSuccess(t, &pb.BenchRepeatedScalar{})
	})
	t.Run("Message", func(t *testing.T) {
		t.Parallel()
		testSuccess(t, &pb.BenchRepeatedMessage{})
	})
	t.Run("Unique", func(t *testing.T) {
		t.Parallel()
		t.Run("Scalar", func(t *testing.T) {
			t.Parallel()
			testSuccess(t, &pb.BenchRepeatedScalarUnique{})
		})
		t.Run("Bytes", func(t *testing.T) {
			t.Parallel()
			testSuccess(t, &pb.BenchRepeatedBytesUnique{})
		})
	})
}

func BenchmarkCompile(b *testing.B) {
	// Measures compile-time allocations for complex schemas
	msg := &pb.BenchComplexSchema{}
	b.ReportAllocs()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	for b.Loop() {
		_, _ = New(options...)
	}
}

func BenchmarkCompileInt32GT(b *testing.B) {
	// Measures compile-time allocations for complex schemas
	msg := &pb.BenchGT{}
	b.ReportAllocs()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	for b.Loop() {
		_, _ = New(options...)
	}
}

func BenchmarkMultiRuleError(b *testing.B) {
	msg := pb.MultiRule_builder{Many: 1}.Build()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	val, err := New(options...)
	require.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = val.Validate(msg)
	}
}

func BenchmarkMultiRuleNoError(b *testing.B) {
	msg := pb.MultiRule_builder{Many: 10}.Build()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	val, err := New(options...)
	require.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = val.Validate(msg)
	}
}

func testSuccess(t *testing.T, msg proto.Message) {
	faker := protogofakeit.New(gofakeit.New(1))
	require.NoError(t, faker.FakeProto(msg))
	val, err := New(WithMessages(msg), WithDisableLazy())
	require.NoError(t, err)
	err = val.Validate(msg)
	require.NoError(t, err)
}

func benchSuccess(b *testing.B, msg proto.Message) {
	faker := protogofakeit.New(gofakeit.New(1))
	require.NoError(b, faker.FakeProto(msg))
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	val, err := New(options...)
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_ = val.Validate(msg)
	}
}

// newCrossReference builds a BenchCrossReference with n entries and n rules of
// n references each. Every reference resolves to the *last* entry, so the
// exists() macro scans the whole entry list and no short-circuit truncates the
// work: evaluating the message-level rule costs Theta(n^3) comprehension
// iterations. The message is valid, so validation runs to completion.
func newCrossReference(n int) *pb.BenchCrossReference {
	entries := make([]*pb.BenchCrossReference_Entry, n)
	for i := range n {
		entries[i] = pb.BenchCrossReference_Entry_builder{
			Key: pb.BenchCrossReference_Key_builder{Kind: "k", Name: fmt.Sprintf("n-%d", i)}.Build(),
		}.Build()
	}
	last := pb.BenchCrossReference_Key_builder{Kind: "k", Name: fmt.Sprintf("n-%d", n-1)}.Build()
	rules := make([]*pb.BenchCrossReference_Rule, n)
	for i := range n {
		refs := make([]*pb.BenchCrossReference_Key, n)
		for j := range n {
			refs[j] = last
		}
		rules[i] = pb.BenchCrossReference_Rule_builder{Refs: refs}.Build()
	}
	return pb.BenchCrossReference_builder{Entries: entries, Rules: rules}.Build()
}

// benchContextVariants compares the three evaluation paths on one message:
//
//   - Validate:    the legacy entry point, which delegates with context.Background()
//   - Background:  ValidateContext with a context that can never be cancelled
//   - Cancellable: ValidateContext with a real cancellable context
//
// Validate and Background must match the pre-context baseline: because
// context.Background().Done() is nil, evalContext skips cel-go's ContextEval
// and its per-evaluation activation wrapping. Only Cancellable pays for
// cancellation support.
func benchContextVariants(b *testing.B, msg proto.Message) {
	b.Helper()
	options := []ValidatorOption{WithMessages(msg), WithDisableLazy()}
	if strings.EqualFold(os.Getenv("DISABLE_NATIVE_RULES"), "true") {
		options = append(options, WithDisableNativeRules())
	}
	val, err := New(options...)
	require.NoError(b, err)

	b.Run("Validate", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = val.Validate(msg)
		}
	})
	b.Run("Background", func(b *testing.B) {
		ctx := context.Background()
		b.ReportAllocs()
		for b.Loop() {
			_ = val.ValidateContext(ctx, msg)
		}
	})
	b.Run("Cancellable", func(b *testing.B) {
		ctx := b.Context() // cancelled when the benchmark ends; Done() is non-nil
		b.ReportAllocs()
		for b.Loop() {
			_ = val.ValidateContext(ctx, msg)
		}
	})
}

// BenchmarkContextOverhead quantifies what context support costs across message
// shapes exercising the different cancellation checkpoints.
func BenchmarkContextOverhead(b *testing.B) {
	faked := func(msg proto.Message) proto.Message {
		require.NoError(b, protogofakeit.New(gofakeit.New(1)).FakeProto(msg))
		return msg
	}
	b.Run("Scalar", func(b *testing.B) {
		benchContextVariants(b, faked(&pb.BenchScalar{}))
	})
	b.Run("RepeatedMessage", func(b *testing.B) {
		benchContextVariants(b, faked(&pb.BenchRepeatedMessage{}))
	})
	b.Run("Map", func(b *testing.B) {
		benchContextVariants(b, faked(&pb.BenchMap{}))
	})
	b.Run("ComplexSchema", func(b *testing.B) {
		benchContextVariants(b, faked(&pb.BenchComplexSchema{}))
	})
}

// BenchmarkInterruptCheckFrequency shows that the interrupt check frequency does
// not meaningfully affect throughput. cel-go calls checkInterrupt on every
// comprehension iteration regardless of the frequency; the frequency only gates
// a non-blocking channel poll. Raising it therefore trades cancellation latency
// away for no throughput gain, which is why the default is 1.
func BenchmarkInterruptCheckFrequency(b *testing.B) {
	msg := newCrossReference(8)
	for _, frequency := range []uint{0, 1, 10, 100} {
		b.Run(fmt.Sprintf("freq=%d", frequency), func(b *testing.B) {
			val, err := New(WithMessages(msg), WithDisableLazy(),
				WithCELInterruptCheckFrequency(frequency))
			require.NoError(b, err)
			ctx := b.Context()
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = val.ValidateContext(ctx, msg)
			}
		})
	}
}
