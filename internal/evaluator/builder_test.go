// Copyright 2023 Buf Technologies, Inc.
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

package evaluator

import (
	"sync"
	"testing"

	"github.com/bufbuild/protovalidate-go/celext"
	pb "github.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestBuildCache(t *testing.T) {
	t.Parallel()

	env, err := celext.DefaultEnv(true)
	require.NoError(t, err, "failed to construct CEL environment")
	bldr := NewBuilder(
		env, false, DefaultResolver{},
	)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		dynamicMsg := dynamicProto{&pb.Person{
			Id:    1234,
			Email: "protovalidate@buf.build",
			Name:  "Protocol Buffer",
		}, int32(i)}
		desc := dynamicMsg.ProtoReflect().Descriptor()
		go func() {
			defer wg.Done()
			eval := bldr.Load(desc)
			assert.NotNil(t, eval)
		}()
	}
	wg.Wait()
}

type dynamicProto struct {
	proto.Message
	salt int32
}

func (d dynamicProto) ProtoReflect() protoreflect.Message {
	return dynamicMessage{d.Message.ProtoReflect(), d.salt}
}

type dynamicMessage struct {
	protoreflect.Message
	salt int32
}

func (d dynamicMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	d.Message.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		return f(dynamicFieldDescriptor{fd, d.salt}, v)
	})
}

func (d dynamicMessage) Has(fd protoreflect.FieldDescriptor) bool {
	return d.Message.Has(unwrapFieldDescriptor(fd))
}
func (d dynamicMessage) Clear(fd protoreflect.FieldDescriptor) {
	d.Message.Clear(unwrapFieldDescriptor(fd))
}
func (d dynamicMessage) Get(fd protoreflect.FieldDescriptor) protoreflect.Value {
	return d.Message.Get(unwrapFieldDescriptor(fd))
}
func (d dynamicMessage) Set(fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	d.Message.Set(unwrapFieldDescriptor(fd), v)
}
func (d dynamicMessage) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	return d.Message.Mutable(unwrapFieldDescriptor(fd))
}
func (d dynamicMessage) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	return d.Message.NewField(unwrapFieldDescriptor(fd))
}

func (d dynamicMessage) Descriptor() protoreflect.MessageDescriptor {
	return dynamicMessageDescriptor{d.Message.Descriptor(), d.salt}
}

type dynamicMessageDescriptor struct {
	protoreflect.MessageDescriptor
	salt int32
}

func (d dynamicMessageDescriptor) Fields() protoreflect.FieldDescriptors {
	return dynamicFieldDescriptors{d.MessageDescriptor.Fields(), d.salt}
}

type dynamicFieldDescriptors struct {
	protoreflect.FieldDescriptors
	salt int32
}

func (d dynamicFieldDescriptors) Get(i int) protoreflect.FieldDescriptor {
	return dynamicFieldDescriptor{d.FieldDescriptors.Get(i), d.salt}
}
func (d dynamicFieldDescriptors) ByName(s protoreflect.Name) protoreflect.FieldDescriptor {
	return dynamicFieldDescriptor{d.FieldDescriptors.ByName(s), d.salt}
}
func (d dynamicFieldDescriptors) ByJSONName(s string) protoreflect.FieldDescriptor {
	return dynamicFieldDescriptor{d.FieldDescriptors.ByJSONName(s), d.salt}
}
func (d dynamicFieldDescriptors) ByTextName(s string) protoreflect.FieldDescriptor {
	return dynamicFieldDescriptor{d.FieldDescriptors.ByTextName(s), d.salt}
}
func (d dynamicFieldDescriptors) ByNumber(n protoreflect.FieldNumber) protoreflect.FieldDescriptor {
	return dynamicFieldDescriptor{d.FieldDescriptors.ByNumber(n), d.salt}
}

type dynamicFieldDescriptor struct {
	protoreflect.FieldDescriptor
	salt int32
}

func unwrapFieldDescriptor(fd protoreflect.FieldDescriptor) protoreflect.FieldDescriptor {
	if d, ok := fd.(dynamicFieldDescriptor); ok {
		return d.FieldDescriptor
	}
	return fd
}
