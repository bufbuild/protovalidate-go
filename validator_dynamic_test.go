// Copyright 2023-2025 Buf Technologies, Inc.
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
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestValidator_DynamicMessageTypeShadowing(t *testing.T) {
	t.Parallel()

	stringMsgType := newDynamicMessageType(t, "test.shadowing", "Message", &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("value"),
		Number: proto.Int32(1),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Options: fieldOpts(validate.FieldRules_builder{
			String: validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build(),
		}.Build()),
	})
	int32MsgType := newDynamicMessageType(t, "test.shadowing", "Message", &descriptorpb.FieldDescriptorProto{
		Name:   proto.String("value"),
		Number: proto.Int32(1),
		Type:   descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
		Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		Options: fieldOpts(validate.FieldRules_builder{
			Int32: validate.Int32Rules_builder{Gt: proto.Int32(0)}.Build(),
		}.Build()),
	})

	require.Equal(t,
		stringMsgType.Descriptor().FullName(),
		int32MsgType.Descriptor().FullName(),
		"both types should have the same fully-qualified name")
	require.NotEqual(t,
		stringMsgType.Descriptor().Fields().ByName("value").Kind(),
		int32MsgType.Descriptor().Fields().ByName("value").Kind(),
		"types should have different field kinds")

	val, err := New()
	require.NoError(t, err)

	stringMsg := stringMsgType.New()
	stringMsg.Set(stringMsgType.Descriptor().Fields().ByName("value"), protoreflect.ValueOfString("hello"))

	int32Msg := int32MsgType.New()
	int32Msg.Set(int32MsgType.Descriptor().Fields().ByName("value"), protoreflect.ValueOfInt32(42))

	require.NoError(t, val.Validate(stringMsg.Interface()), "valid string")
	require.NoError(t, val.Validate(int32Msg.Interface()), "valid int32")
	require.NoError(t, val.Validate(int32Msg.Interface()), "valid int32 re-validation")
	require.NoError(t, val.Validate(stringMsg.Interface()), "valid string re-validation")

	emptyStringMsg := stringMsgType.New()
	require.Error(t, val.Validate(emptyStringMsg.Interface()), "empty string violates min_len=1")

	zeroInt32Msg := int32MsgType.New()
	require.Error(t, val.Validate(zeroInt32Msg.Interface()), "zero violates gt=0")
}

func TestValidator_ParallelDynamicMessageRegistration(t *testing.T) {
	t.Parallel()

	const numTypes = 50

	val, err := New()
	require.NoError(t, err)

	types := make([]protoreflect.MessageType, numTypes)
	for i := range numTypes {
		var fieldType descriptorpb.FieldDescriptorProto_Type
		var opts *descriptorpb.FieldOptions
		if i%2 == 0 {
			fieldType = descriptorpb.FieldDescriptorProto_TYPE_STRING
			opts = fieldOpts(validate.FieldRules_builder{
				String: validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build(),
			}.Build())
		} else {
			fieldType = descriptorpb.FieldDescriptorProto_TYPE_INT32
			opts = fieldOpts(validate.FieldRules_builder{
				Int32: validate.Int32Rules_builder{Gte: proto.Int32(0)}.Build(),
			}.Build())
		}
		types[i] = newDynamicMessageType(t, fmt.Sprintf("test.parallel%d", i), "Message", &descriptorpb.FieldDescriptorProto{
			Name:    proto.String("value"),
			Number:  proto.Int32(1),
			Type:    fieldType.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: opts,
		})
	}

	var wg sync.WaitGroup
	for i, msgType := range types {
		wg.Add(1)
		go func() {
			defer wg.Done()

			msg := msgType.New()
			fieldDesc := msgType.Descriptor().Fields().ByName("value")
			if fieldDesc.Kind() == protoreflect.StringKind {
				msg.Set(fieldDesc, protoreflect.ValueOfString(fmt.Sprintf("value%d", i)))
			} else {
				msg.Set(fieldDesc, protoreflect.ValueOfInt32(int32(i)))
			}

			for range 10 {
				assert.NoError(t, val.Validate(msg.Interface()), "type %d", i)
			}
		}()
	}
	wg.Wait()
}

func TestValidator_SharedNestedDynamicMessage(t *testing.T) {
	t.Parallel()

	registry := new(protoregistry.Files)
	require.NoError(t, registry.RegisterFile(validate.File_buf_validate_validate_proto))

	nestedFilePB := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.shared.nested.proto"),
		Package: proto.String("test.shared"),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"buf/validate/validate.proto",
		},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Nested"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   proto.String("value"),
				Number: proto.Int32(1),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				Options: fieldOpts(validate.FieldRules_builder{
					String: validate.StringRules_builder{MinLen: proto.Uint64(1)}.Build(),
				}.Build()),
			}},
		}},
	}
	nestedFile, err := protodesc.FileOptions{}.New(nestedFilePB, registry)
	require.NoError(t, err)
	require.NoError(t, registry.RegisterFile(nestedFile))

	nestedField := &descriptorpb.FieldDescriptorProto{
		Name:     proto.String("nested"),
		Number:   proto.Int32(1),
		Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
		TypeName: proto.String(".test.shared.Nested"),
		Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
	}
	parentAType := newDynamicMessageTypeWithRegistry(t, registry, "test.shared", "ParentA", nestedField)
	parentBType := newDynamicMessageTypeWithRegistry(t, registry, "test.shared", "ParentB", nestedField)
	sharedNestedDesc := parentAType.Descriptor().Fields().ByName("nested").Message()

	require.NotEqual(t,
		parentAType.Descriptor().FullName(),
		parentBType.Descriptor().FullName(),
		"parent types should have different names")
	require.Equal(t,
		sharedNestedDesc.FullName(),
		parentBType.Descriptor().Fields().ByName("nested").Message().FullName(),
		"both parents should reference the same nested type")

	val, err := New()
	require.NoError(t, err)

	nestedA := dynamicpb.NewMessage(sharedNestedDesc)
	nestedA.Set(sharedNestedDesc.Fields().ByName("value"), protoreflect.ValueOfString("nested-value-a"))

	nestedB := dynamicpb.NewMessage(sharedNestedDesc)
	nestedB.Set(sharedNestedDesc.Fields().ByName("value"), protoreflect.ValueOfString("nested-value-b"))

	parentA := parentAType.New()
	parentA.Set(parentAType.Descriptor().Fields().ByName("nested"), protoreflect.ValueOfMessage(nestedA))

	parentB := parentBType.New()
	parentB.Set(parentBType.Descriptor().Fields().ByName("nested"), protoreflect.ValueOfMessage(nestedB))

	require.NoError(t, val.Validate(parentA.Interface()), "valid parentA")
	require.NoError(t, val.Validate(parentB.Interface()), "valid parentB")
	require.NoError(t, val.Validate(parentB.Interface()), "valid parentB re-validation")
	require.NoError(t, val.Validate(parentA.Interface()), "valid parentA re-validation")

	emptyNested := dynamicpb.NewMessage(sharedNestedDesc)
	parentWithEmptyNested := parentAType.New()
	parentWithEmptyNested.Set(parentAType.Descriptor().Fields().ByName("nested"), protoreflect.ValueOfMessage(emptyNested))
	require.Error(t, val.Validate(parentWithEmptyNested.Interface()), "empty nested value violates min_len=1")
}

func newDynamicMessageType(
	t testing.TB,
	pkg, name string,
	field *descriptorpb.FieldDescriptorProto,
) protoreflect.MessageType {
	t.Helper()

	registry := new(protoregistry.Files)
	require.NoError(t, registry.RegisterFile(validate.File_buf_validate_validate_proto))

	return newDynamicMessageTypeWithRegistry(t, registry, pkg, name, field)
}

func newDynamicMessageTypeWithRegistry(
	t testing.TB,
	registry *protoregistry.Files,
	pkg, name string,
	field *descriptorpb.FieldDescriptorProto,
) protoreflect.MessageType {
	t.Helper()

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
		}},
	}

	registry.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if fd.Path() != "buf/validate/validate.proto" {
			file.Dependency = append(file.Dependency, fd.Path())
		}
		return true
	})

	fd, err := protodesc.FileOptions{}.New(file, registry)
	require.NoError(t, err)

	desc := fd.Messages().ByName(protoreflect.Name(name))
	require.NotNil(t, desc)

	return dynamicpb.NewMessageType(desc)
}

func fieldOpts(rules *validate.FieldRules) *descriptorpb.FieldOptions {
	opts := &descriptorpb.FieldOptions{}
	proto.SetExtension(opts, validate.E_Field, rules)
	return opts
}
