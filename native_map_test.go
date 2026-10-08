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
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	examplev1 "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNativeMapMinPairs(t *testing.T) {
	t.Parallel()
	eval := buildNativeMap(t, validate.MapRules_builder{MinPairs: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	// 2 entries passes
	m := newStringStringMap(t, map[string]string{"a": "1", "b": "2"})
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfMap(m), &validationConfig{}))

	// 1 entry fails
	m = newStringStringMap(t, map[string]string{"a": "1"})
	err := eval.Evaluate(nil, protoreflect.ValueOfMap(m), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "map.min_pairs", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "map must be at least 2 entries", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeMapMaxPairs(t *testing.T) {
	t.Parallel()
	eval := buildNativeMap(t, validate.MapRules_builder{MaxPairs: proto.Uint64(2)}.Build())
	require.NotNil(t, eval)

	// 2 entries passes
	m := newStringStringMap(t, map[string]string{"a": "1", "b": "2"})
	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfMap(m), &validationConfig{}))

	// 3 entries fails
	m = newStringStringMap(t, map[string]string{"a": "1", "b": "2", "c": "3"})
	err := eval.Evaluate(nil, protoreflect.ValueOfMap(m), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "map.max_pairs", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "map must be at most 2 entries", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeMapPairViolations(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		minPairs uint64
		maxPairs uint64
		size     int
		wantIDs  []string
	}{
		{"conflicting_bounds/empty", 2, 0, 0, []string{"map.min_pairs"}},
		{"conflicting_bounds/both", 2, 0, 1, []string{"map.min_pairs", "map.max_pairs"}},
		{"conflicting_bounds/minimum", 2, 0, 2, []string{"map.max_pairs"}},
		{"normal_bounds/below", 1, 2, 0, []string{"map.min_pairs"}},
		{"normal_bounds/minimum", 1, 2, 1, nil},
		{"normal_bounds/maximum", 1, 2, 2, nil},
		{"normal_bounds/above", 1, 2, 3, []string{"map.max_pairs"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			msgType := newDynamicMapMessageType(t, "test.map", "PairViolations",
				descriptorpb.FieldDescriptorProto_TYPE_STRING,
				descriptorpb.FieldDescriptorProto_TYPE_STRING,
				validate.FieldRules_builder{
					Map: validate.MapRules_builder{
						MinPairs: new(test.minPairs),
						MaxPairs: new(test.maxPairs),
					}.Build(),
				}.Build(),
			)
			msg := msgType.New()
			entries := msg.Mutable(msg.Descriptor().Fields().ByName("entries")).Map()
			for _, key := range []string{"a", "b", "c"}[:test.size] {
				entries.Set(protoreflect.ValueOfString(key).MapKey(), protoreflect.ValueOfString(key))
			}

			modes := []struct {
				name              string
				validatorOptions  []ValidatorOption
				validationOptions []ValidationOption
				failFast          bool
			}{
				{name: "all_violations"},
				{name: "validator_fail_fast", validatorOptions: []ValidatorOption{WithFailFast()}, failFast: true},
				{name: "validation_fail_fast", validationOptions: []ValidationOption{WithFailFast()}, failFast: true},
			}
			for _, mode := range modes {
				t.Run(mode.name, func(t *testing.T) {
					nativeVal, err := New(mode.validatorOptions...)
					require.NoError(t, err)
					celVal, err := New(append([]ValidatorOption{WithDisableNativeRules()}, mode.validatorOptions...)...)
					require.NoError(t, err)
					nativeErr := nativeVal.Validate(msg.Interface(), mode.validationOptions...)
					celErr := celVal.Validate(msg.Interface(), mode.validationOptions...)
					if len(test.wantIDs) == 0 {
						require.NoError(t, nativeErr)
						require.NoError(t, celErr)
						return
					}
					wantIDs := test.wantIDs
					if mode.failFast {
						wantIDs = wantIDs[:1]
					}
					assert.Equal(t, wantIDs, violationRuleIDs(t, nativeErr))
					assert.Equal(t, wantIDs, violationRuleIDs(t, celErr))
					var nativeValErr, celValErr *ValidationError
					require.ErrorAs(t, nativeErr, &nativeValErr)
					require.ErrorAs(t, celErr, &celValErr)
					assert.True(t, proto.Equal(celValErr.ToProto(), nativeValErr.ToProto()))
				})
			}
		})
	}
}

func TestTryNativeMapRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.MapRules
	}{
		{"nil_rules", nil},
		{"empty_rules", validate.MapRules_builder{}.Build()},
		{"keys_only", validate.MapRules_builder{
			Keys: validate.FieldRules_builder{
				String: validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build(),
			}.Build(),
		}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryNativeMapRules(base{}, tt.rules))
		})
	}
}

func TestNativeMapTautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeMap(t, validate.MapRules_builder{MinPairs: proto.Uint64(1)}.Build())
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}

// --- Helpers ---

func buildNativeMap(t testing.TB, rules *validate.MapRules) evaluator {
	t.Helper()
	fdesc := newMapFieldDescriptor(t)
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryNativeMapRules(b, rules)
}

func newMapFieldDescriptor(t testing.TB) protoreflect.FieldDescriptor {
	t.Helper()
	msgType := newDynamicMapMessageType(t, "test.map", "MapTestMsg",
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		nil,
	)
	return msgType.Descriptor().Fields().ByName("entries")
}

func newStringStringMap(t testing.TB, entries map[string]string) protoreflect.Map {
	t.Helper()
	msgType := newDynamicMapMessageType(t, "test.map", "MapHelper",
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_STRING,
		nil,
	)
	msg := dynamicpb.NewMessage(msgType.Descriptor())
	fd := msgType.Descriptor().Fields().ByName("entries")
	mapField := msg.NewField(fd)
	for k, v := range entries {
		mapField.Map().Set(
			protoreflect.ValueOfString(k).MapKey(),
			protoreflect.ValueOfString(v),
		)
	}
	msg.Set(fd, mapField)
	return msg.Get(fd).Map()
}

// newDynamicMapMessageType creates a dynamic message type with a map field.
func newDynamicMapMessageType(
	t testing.TB,
	pkg, name string,
	keyType, valueType descriptorpb.FieldDescriptorProto_Type,
	rules *validate.FieldRules,
) protoreflect.MessageType {
	t.Helper()

	mapEntryName := "EntriesEntry"
	field := &descriptorpb.FieldDescriptorProto{
		Name:     new("entries"),
		Number:   proto.Int32(1),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: new("." + pkg + "." + name + "." + mapEntryName),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
	}
	if rules != nil {
		field.Options = fieldOpts(rules)
	}

	file := &descriptorpb.FileDescriptorProto{
		Name:    new(pkg + "." + name + ".proto"),
		Package: new(pkg),
		Syntax:  new("proto3"),
		Dependency: []string{
			"buf/validate/validate.proto",
		},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:  new(name),
			Field: []*descriptorpb.FieldDescriptorProto{field},
			NestedType: []*descriptorpb.DescriptorProto{{
				Name: new(mapEntryName),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   new("key"),
						Number: proto.Int32(1),
						Type:   keyType.Enum(),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					},
					{
						Name:   new("value"),
						Number: proto.Int32(2),
						Type:   valueType.Enum(),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					},
				},
				Options: &descriptorpb.MessageOptions{
					MapEntry: new(true),
				},
			}},
		}},
	}

	registry := newRegistryWithValidateProto(t)
	fd, err := protodesc.FileOptions{}.New(file, registry)
	require.NoError(t, err)

	desc := fd.Messages().ByName(protoreflect.Name(name))
	require.NotNil(t, desc)

	return dynamicpb.NewMessageType(desc)
}

func TestNativeMapWrapperValues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		msg      proto.Message
		wantRule string
		wantPath string
	}{
		{
			name: "int32/invalid",
			msg: examplev1.MapWrapperValues_builder{
				Val: map[string]*wrapperspb.Int32Value{"key": wrapperspb.Int32(5)},
			}.Build(),
			wantRule: "int32.gt",
			wantPath: `val["key"]`,
		},
		{
			name: "int32/valid",
			msg: examplev1.MapWrapperValues_builder{
				Val: map[string]*wrapperspb.Int32Value{"key": wrapperspb.Int32(100)},
			}.Build(),
		},
		{
			name: "string/invalid",
			msg: examplev1.MapStringWrapperValues_builder{
				Val: map[string]*wrapperspb.StringValue{"k": wrapperspb.String("a")},
			}.Build(),
			wantRule: "string.min_len",
			wantPath: `val["k"]`,
		},
		{
			name: "string/valid",
			msg: examplev1.MapStringWrapperValues_builder{
				Val: map[string]*wrapperspb.StringValue{"k": wrapperspb.String("abc")},
			}.Build(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			nativeVal, err := New()
			require.NoError(t, err)
			celVal, err := New(WithDisableNativeRules())
			require.NoError(t, err)
			nativeErr := nativeVal.Validate(test.msg)
			celErr := celVal.Validate(test.msg)
			if test.wantRule == "" {
				require.NoError(t, nativeErr)
				require.NoError(t, celErr)
				return
			}
			var nativeValErr, celValErr *ValidationError
			require.ErrorAs(t, nativeErr, &nativeValErr)
			require.ErrorAs(t, celErr, &celValErr)
			require.Len(t, nativeValErr.Violations, 1)
			assert.Equal(t, test.wantRule, nativeValErr.Violations[0].Proto.GetRuleId())
			assert.Equal(t, test.wantPath, FieldPathString(nativeValErr.Violations[0].Proto.GetField()))
			assert.True(t, proto.Equal(celValErr.ToProto(), nativeValErr.ToProto()))
		})
	}
}
