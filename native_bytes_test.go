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
	"errors"
	"fmt"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	examplev1 "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestNativeBytes(t *testing.T) {
	t.Parallel()
	msg := examplev1.BenchTestBytes_builder{
		B1: []byte{'\x32', '\x33'},
		B:  []byte{'\x03', '\x04'},
	}.Build()
	// with native rules off
	val, err := New(WithMessages(msg), WithDisableLazy(), WithDisableNativeRules())
	if err != nil {
		t.Fatalf("native off: expected no error, got %v", err)
	}
	err = val.Validate(msg)
	if err == nil {
		t.Errorf("native off: expected error, got nil")
	} else {
		x := err.Error()
		if diff := cmp.Diff(x, `validation errors:
 - b1: must not be in list [23, 45, 67]
 - b: must be in list [23, 45, 67]`); diff != "" {
			t.Error("native off, difference in error: " + diff)
		}
	}

	// with native rules on
	val, err = New(WithMessages(msg), WithDisableLazy())
	if err != nil {
		t.Fatalf("native on: expected no error, got %v", err)
	}
	err = val.Validate(msg)
	if err == nil {
		t.Errorf("native on: expected error, got nil")
	} else {
		x := err.Error()
		if diff := cmp.Diff(x, `validation errors:
 - b1: must not be in list [23, 45, 67]
 - b: must be in list [23, 45, 67]`); diff != "" {
			t.Error("native on, difference in error: " + diff)
		}
	}
}
func TestNativeBytesConst(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Const: []byte{0x01, 0x02}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01, 0x02}), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x03}), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "bytes.const", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must be 0102", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeBytesLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Len: proto.Uint64(4)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3, 4}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3}), &validationConfig{}))
}

func TestNativeBytesMinLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{MinLen: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1}), &validationConfig{}))
}

func TestNativeBytesMaxLen(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{MaxLen: proto.Uint64(3)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3, 4}), &validationConfig{}))
}

func TestNativeBytesPattern(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Pattern: proto.String("^[a-zA-Z0-9]+$")}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte("abc123")), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte("abc 123")), &validationConfig{}))
}

func TestNativeBytesPrefix(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Prefix: []byte{0x01, 0x02}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01, 0x02, 0x03}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x03, 0x02, 0x01}), &validationConfig{}))
}

func TestNativeBytesSuffix(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Suffix: []byte{0x03}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01, 0x02, 0x03}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x03, 0x02, 0x01}), &validationConfig{}))
}

func TestNativeBytesContains(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Contains: []byte{0x02}}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01, 0x02, 0x03}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01, 0x03}), &validationConfig{}))
}

func TestNativeBytesIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{
		In: [][]byte{{0x01}, {0x02}},
	}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x03}), &validationConfig{}))
}

func TestNativeBytesNotIn(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{
		NotIn: [][]byte{{0x00}},
	}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x01}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{0x00}), &validationConfig{}))
}

func TestNativeBytesIP(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Ip: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	// valid IPv4 (4 bytes)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{127, 0, 0, 1}), &validationConfig{}))
	// valid IPv6 (16 bytes)
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes(make([]byte, 16)), &validationConfig{}))

	// empty → empty-specific message
	err := eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{}), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "bytes.ip_empty", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "value is empty, which is not a valid IP address", valErr.Violations[0].Proto.GetMessage())

	// wrong length → main message
	err = eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3}), &validationConfig{})
	require.Error(t, err)
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "bytes.ip", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must be a valid IP address", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeBytesIPv4(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Ipv4: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{10, 0, 0, 1}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes(make([]byte, 16)), &validationConfig{}))
}

func TestNativeBytesIPv6(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Ipv6: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes(make([]byte, 16)), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{}), &validationConfig{}))
	require.Error(t, eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3, 4}), &validationConfig{}))
}

func TestNativeBytesUUID(t *testing.T) {
	t.Parallel()
	eval := buildNativeBytes(t, validate.BytesRules_builder{Uuid: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBytes(make([]byte, 16)), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{}), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "bytes.uuid_empty", valErr.Violations[0].Proto.GetRuleId())

	err = eval.Evaluate(nil, protoreflect.ValueOfBytes([]byte{1, 2, 3}), &validationConfig{})
	require.Error(t, err)
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "bytes.uuid", valErr.Violations[0].Proto.GetRuleId())
}

// TestNativeBytesBroken runs a test twice, once with CEL, once with native, making sure that when
// we have two rules broken, we get two violations.
func TestNativeBytesBroken(t *testing.T) {
	t.Parallel()
	// should fail, want same message from both native and CEL
	msg := examplev1.TestByteBroken_builder{
		Broken: []byte("greetings and salutations"),
	}.Build()
	for _, d := range []bool{false, true} {
		t.Run(fmt.Sprintf("disableNativeRules=%t", d), func(t *testing.T) {
			t.Parallel()
			options := []ValidatorOption{WithDisableLazy(), WithMessageDescriptors(msg.ProtoReflect().Descriptor())}
			if d {
				options = append(options, WithDisableNativeRules())
			}
			validator, err := New(options...)
			require.NoError(t, err)
			err = validator.Validate(msg)
			require.Error(t, err)
			var valErr *ValidationError
			if errors.As(err, &valErr) {
				if len(valErr.Violations) != 2 {
					t.Errorf("expected 2 violations, got %d: %v", len(valErr.Violations), valErr)
				}
			} else {
				t.Error("expected a validation error")
			}
		})
	}
}

func TestTryBuildNativeBytesRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.BytesRules
	}{
		{"nil_rules", nil},
		{"empty_rules", validate.BytesRules_builder{}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryBuildNativeBytesRules(base{}, tt.rules))
		})
	}
}

func TestNativeBytes_EndToEnd(t *testing.T) {
	t.Parallel()
	msgType := newDynamicMessageType(t, "test.native", "BytesMsg", &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("value"),
		Number: proto.Int32(1),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_BYTES.Enum(),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Options: fieldOpts(validate.FieldRules_builder{
			Bytes: validate.BytesRules_builder{MinLen: proto.Uint64(2)}.Build(),
		}.Build()),
	})

	validator, err := New(WithDisableLazy(), WithMessageDescriptors(msgType.Descriptor()))
	require.NoError(t, err)

	passing := dynamicpb.NewMessage(msgType.Descriptor())
	passing.Set(msgType.Descriptor().Fields().ByName("value"), protoreflect.ValueOfBytes([]byte{0x01, 0x02}))
	require.NoError(t, validator.Validate(passing))

	failing := dynamicpb.NewMessage(msgType.Descriptor())
	failing.Set(msgType.Descriptor().Fields().ByName("value"), protoreflect.ValueOfBytes([]byte{0x01}))
	err = validator.Validate(failing)
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "bytes.min_len", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must be at least 2 bytes", valErr.Violations[0].Proto.GetMessage())
}

func buildNativeBytes(t testing.TB, rules *validate.BytesRules) evaluator {
	t.Helper()
	fdesc := newFieldDescriptor(t, descriptorpb.FieldDescriptorProto_TYPE_BYTES,
		descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum())
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryBuildNativeBytesRules(b, rules)
}
