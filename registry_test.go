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

	pvcel "buf.build/go/protovalidate/cel"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

// newTestEnv creates a new CEL environment with the protovalidate library
// and registry attached, similar to how the validator creates its env.
func newTestEnv(t testing.TB) *cel.Env {
	t.Helper()
	reg, err := newRegistry()
	require.NoError(t, err)
	env, err := cel.NewEnv(
		cel.CustomTypeProvider(reg),
		cel.CustomTypeAdapter(reg),
		cel.Lib(pvcel.NewLibrary()),
	)
	require.NoError(t, err)
	return env
}

func TestRegistryTypeShadowing(t *testing.T) {
	t.Parallel()

	// Create base environment with protovalidate library and registry
	baseEnv := newTestEnv(t)

	// Extend with first dynamic message (id: string)
	env1, err := extendEnv(baseEnv,
		cel.Types(newRegistryTestMessageType(t, descriptorpb.FieldDescriptorProto_TYPE_STRING).New().Interface()),
	)
	require.NoError(t, err)

	// Compile and run expression with string value
	ast, issues := env1.Compile(`test.Duplicate{id: "abc"}`)
	require.NoError(t, issues.Err())
	prog, err := env1.Program(ast)
	require.NoError(t, err)
	_, _, err = prog.Eval(cel.NoVars())
	require.NoError(t, err)

	// Extend base (not env1) with second dynamic message (id: int32)
	env2, err := extendEnv(baseEnv,
		cel.Types(newRegistryTestMessageType(t, descriptorpb.FieldDescriptorProto_TYPE_INT32).New().Interface()),
	)
	require.NoError(t, err)

	// Compile expression with string value - should fail because env2 expects int32
	_, issues = env2.Compile(`test.Duplicate{id: "abc"}`)
	require.Error(t, issues.Err(), "env2 should expect int32, not string")

	// Verify env2 can compile with int32 value
	ast, issues = env2.Compile(`test.Duplicate{id: 123}`)
	require.NoError(t, issues.Err())
	prog, err = env2.Program(ast)
	require.NoError(t, err)
	_, _, err = prog.Eval(cel.NoVars())
	require.NoError(t, err)
}

func TestRegistryDeepChain(t *testing.T) {
	t.Parallel()

	const depth = 10

	baseEnv := newTestEnv(t)

	// Create a chain of extended environments, each registering a unique type
	envs := make([]*cel.Env, depth)
	envs[0] = baseEnv
	for i := 1; i < depth; i++ {
		msgType := newRegistryTestMessageTypeWithName(t, fmt.Sprintf("test.Level%d", i), "Message",
			descriptorpb.FieldDescriptorProto_TYPE_INT32)
		var err error
		envs[i], err = extendEnv(envs[i-1], cel.Types(msgType.New().Interface()))
		require.NoError(t, err)
	}

	// The deepest environment should be able to see all types from the chain
	deepestEnv := envs[depth-1]
	for i := 1; i < depth; i++ {
		expr := fmt.Sprintf(`test.Level%d.Message{value: %d}`, i, i)
		ast, issues := deepestEnv.Compile(expr)
		require.NoError(t, issues.Err(), "level %d", i)
		prog, err := deepestEnv.Program(ast)
		require.NoError(t, err, "level %d", i)
		_, _, err = prog.Eval(cel.NoVars())
		require.NoError(t, err, "level %d", i)
	}

	// Middle environment should only see types from its ancestors
	middleEnv := envs[depth/2]
	for i := 1; i <= depth/2; i++ {
		expr := fmt.Sprintf(`test.Level%d.Message{value: %d}`, i, i)
		_, issues := middleEnv.Compile(expr)
		require.NoError(t, issues.Err(), "middle env should see level %d", i)
	}
	for i := depth/2 + 1; i < depth; i++ {
		expr := fmt.Sprintf(`test.Level%d.Message{value: %d}`, i, i)
		_, issues := middleEnv.Compile(expr)
		require.Error(t, issues.Err(), "middle env should NOT see level %d", i)
	}
}

func TestRegistryConcurrentAccess(t *testing.T) {
	t.Parallel()

	baseEnv := newTestEnv(t)

	const numGoroutines = 50
	const numOpsPerGoroutine = 100

	// Create all message types and environments before spawning goroutines
	// to avoid require assertions in goroutines
	envs := make([]*cel.Env, numGoroutines)
	for i := range numGoroutines {
		msgType := newRegistryTestMessageTypeWithName(t, fmt.Sprintf("test.Concurrent%d", i), "Message",
			descriptorpb.FieldDescriptorProto_TYPE_STRING)
		var err error
		envs[i], err = extendEnv(baseEnv, cel.Types(msgType.New().Interface()))
		require.NoError(t, err)
	}

	var wg sync.WaitGroup
	for i := range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			env := envs[i]

			// Repeatedly compile and evaluate expressions
			for range numOpsPerGoroutine {
				expr := fmt.Sprintf(`test.Concurrent%d.Message{value: "test"}`, i)
				ast, issues := env.Compile(expr)
				assert.NoError(t, issues.Err())
				if issues.Err() != nil {
					continue
				}
				prog, err := env.Program(ast)
				assert.NoError(t, err)
				if err != nil {
					continue
				}
				_, _, err = prog.Eval(cel.NoVars())
				assert.NoError(t, err)
			}
		}()
	}
	wg.Wait()
}

func TestRegistryCoreTypeIdents(t *testing.T) {
	t.Parallel()

	// Verify that a root registry exposes core CEL type identifiers
	// (string, int, bool, etc.) so that expressions like
	// type(x) == string can compile. See #304.
	env := newTestEnv(t)
	ast, issues := env.Compile(`type("hello") == string`)
	require.NoError(t, issues.Err())
	prog, err := env.Program(ast)
	require.NoError(t, err)
	out, _, err := prog.Eval(cel.NoVars())
	require.NoError(t, err)
	require.Equal(t, true, out.Value())
}

func TestRegistryNativeToValueDelegation(t *testing.T) {
	t.Parallel()

	// Create a chain: root -> child -> grandchild
	root, err := newRegistry()
	require.NoError(t, err)
	child := root.Copy()
	grandchild := child.Copy()

	// Register a message type only in child
	msgType := newRegistryTestMessageTypeWithName(t, "test.delegation", "Message",
		descriptorpb.FieldDescriptorProto_TYPE_STRING)
	require.NoError(t, child.RegisterMessage(msgType.New().Interface()))

	// Create a message instance
	msg := msgType.New()
	msg.Set(msgType.Descriptor().Fields().ByName("value"), protoreflect.ValueOfString("hello"))

	// Grandchild should be able to convert via delegation to child
	result := grandchild.NativeToValue(msg.Interface())
	require.False(t, types.IsError(result), "grandchild should delegate to child")

	// Root should NOT be able to convert (type not registered there)
	result = root.NativeToValue(msg.Interface())
	require.True(t, types.IsError(result), "root should not know about the type")
}

func TestRegistryMessageFirstRegistrationWins(t *testing.T) {
	t.Parallel()

	root, err := newRegistry()
	require.NoError(t, err)
	child := root.Copy()

	// Register a message with certain fields in root
	rootMsgType := newRegistryTestMessageTypeWithFields(t, "test.fields", "Message", map[string]descriptorpb.FieldDescriptorProto_Type{
		"field_a": descriptorpb.FieldDescriptorProto_TYPE_STRING,
		"field_b": descriptorpb.FieldDescriptorProto_TYPE_INT32,
	})
	require.NoError(t, root.RegisterMessage(rootMsgType.New().Interface()))

	// Attempt to register a different version with different fields in child.
	// First-registration-wins semantics: child's registration is skipped
	// because the message is already visible through the parent.
	childMsgType := newRegistryTestMessageTypeWithFields(t, "test.fields", "Message", map[string]descriptorpb.FieldDescriptorProto_Type{
		"field_x": descriptorpb.FieldDescriptorProto_TYPE_BOOL,
		"field_y": descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
		"field_z": descriptorpb.FieldDescriptorProto_TYPE_BYTES,
	})
	require.NoError(t, child.RegisterMessage(childMsgType.New().Interface()))

	// Child sees root's fields (first registration wins, no shadowing for RegisterMessage)
	childFields, found := child.FindStructFieldNames("test.fields.Message")
	require.True(t, found)
	require.ElementsMatch(t, []string{"field_a", "field_b"}, childFields)

	// Root should see original field names
	rootFields, found := root.FindStructFieldNames("test.fields.Message")
	require.True(t, found)
	require.ElementsMatch(t, []string{"field_a", "field_b"}, rootFields)
}

func newRegistryTestMessageType(t testing.TB, fieldType descriptorpb.FieldDescriptorProto_Type) protoreflect.MessageType {
	t.Helper()
	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{{
			Name:    proto.String("test.proto"),
			Package: proto.String("test"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{{
				Name: proto.String("Duplicate"),
				Field: []*descriptorpb.FieldDescriptorProto{{
					Name:     proto.String("id"),
					Number:   proto.Int32(1),
					Type:     fieldType.Enum(),
					JsonName: proto.String("id"),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				}},
			}},
		}},
	})
	require.NoError(t, err)
	desc, err := files.FindDescriptorByName("test.Duplicate")
	require.NoError(t, err)
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	require.True(t, ok)
	return dynamicpb.NewMessageType(msgDesc)
}

func newRegistryTestMessageTypeWithName(t testing.TB, pkg, name string, fieldType descriptorpb.FieldDescriptorProto_Type) protoreflect.MessageType {
	t.Helper()
	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{{
			Name:    proto.String(fmt.Sprintf("%s.%s.proto", pkg, name)),
			Package: proto.String(pkg),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{{
				Name: proto.String(name),
				Field: []*descriptorpb.FieldDescriptorProto{{
					Name:     proto.String("value"),
					Number:   proto.Int32(1),
					Type:     fieldType.Enum(),
					JsonName: proto.String("value"),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				}},
			}},
		}},
	})
	require.NoError(t, err)
	desc, err := files.FindDescriptorByName(protoreflect.FullName(pkg + "." + name))
	require.NoError(t, err)
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	require.True(t, ok)
	return dynamicpb.NewMessageType(msgDesc)
}

func newRegistryTestMessageTypeWithFields(t testing.TB, pkg, name string, fields map[string]descriptorpb.FieldDescriptorProto_Type) protoreflect.MessageType {
	t.Helper()

	fieldDescs := make([]*descriptorpb.FieldDescriptorProto, 0, len(fields))
	fieldNum := int32(1)
	for fieldName, fieldType := range fields {
		fieldDescs = append(fieldDescs, &descriptorpb.FieldDescriptorProto{
			Name:     proto.String(fieldName),
			Number:   proto.Int32(fieldNum),
			Type:     fieldType.Enum(),
			JsonName: proto.String(fieldName),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		})
		fieldNum++
	}

	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{{
			Name:    proto.String(fmt.Sprintf("%s.%s.proto", pkg, name)),
			Package: proto.String(pkg),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{{
				Name:  proto.String(name),
				Field: fieldDescs,
			}},
		}},
	})
	require.NoError(t, err)
	desc, err := files.FindDescriptorByName(protoreflect.FullName(pkg + "." + name))
	require.NoError(t, err)
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	require.True(t, ok)
	return dynamicpb.NewMessageType(msgDesc)
}
