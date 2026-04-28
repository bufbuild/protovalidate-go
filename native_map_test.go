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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
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
		Name:     proto.String("entries"),
		Number:   proto.Int32(1),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String("." + pkg + "." + name + "." + mapEntryName),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
	}
	if rules != nil {
		field.Options = fieldOpts(rules)
	}

	file := &descriptorpb.FileDescriptorProto{
		Name:    proto.String(pkg + "." + name + ".proto"),
		Package: proto.String(pkg),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"buf/validate/validate.proto",
		},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:  proto.String(name),
			Field: []*descriptorpb.FieldDescriptorProto{field},
			NestedType: []*descriptorpb.DescriptorProto{{
				Name: proto.String(mapEntryName),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("key"),
						Number: proto.Int32(1),
						Type:   keyType.Enum(),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					},
					{
						Name:   proto.String("value"),
						Number: proto.Int32(2),
						Type:   valueType.Enum(),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					},
				},
				Options: &descriptorpb.MessageOptions{
					MapEntry: proto.Bool(true),
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
