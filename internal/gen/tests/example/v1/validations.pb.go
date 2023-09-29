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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tests/example/v1/validations.proto

package examplev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	apipb "google.golang.org/protobuf/types/known/apipb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HasMsgExprs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *HasMsgExprs) Reset() {
	*x = HasMsgExprs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasMsgExprs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasMsgExprs) ProtoMessage() {}

func (x *HasMsgExprs) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasMsgExprs.ProtoReflect.Descriptor instead.
func (*HasMsgExprs) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{0}
}

func (x *HasMsgExprs) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *HasMsgExprs) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type SelfRecursive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X      int32          `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Turtle *SelfRecursive `protobuf:"bytes,2,opt,name=turtle,proto3" json:"turtle,omitempty"`
}

func (x *SelfRecursive) Reset() {
	*x = SelfRecursive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfRecursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfRecursive) ProtoMessage() {}

func (x *SelfRecursive) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfRecursive.ProtoReflect.Descriptor instead.
func (*SelfRecursive) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{1}
}

func (x *SelfRecursive) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SelfRecursive) GetTurtle() *SelfRecursive {
	if x != nil {
		return x.Turtle
	}
	return nil
}

type LoopRecursiveA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B *LoopRecursiveB `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *LoopRecursiveA) Reset() {
	*x = LoopRecursiveA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoopRecursiveA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoopRecursiveA) ProtoMessage() {}

func (x *LoopRecursiveA) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoopRecursiveA.ProtoReflect.Descriptor instead.
func (*LoopRecursiveA) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{2}
}

func (x *LoopRecursiveA) GetB() *LoopRecursiveB {
	if x != nil {
		return x.B
	}
	return nil
}

type LoopRecursiveB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *LoopRecursiveA `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *LoopRecursiveB) Reset() {
	*x = LoopRecursiveB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoopRecursiveB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoopRecursiveB) ProtoMessage() {}

func (x *LoopRecursiveB) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoopRecursiveB.ProtoReflect.Descriptor instead.
func (*LoopRecursiveB) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{3}
}

func (x *LoopRecursiveB) GetA() *LoopRecursiveA {
	if x != nil {
		return x.A
	}
	return nil
}

type MsgHasOneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to O:
	//
	//	*MsgHasOneof_X
	//	*MsgHasOneof_Y
	//	*MsgHasOneof_Msg
	O isMsgHasOneof_O `protobuf_oneof:"o"`
}

func (x *MsgHasOneof) Reset() {
	*x = MsgHasOneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHasOneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHasOneof) ProtoMessage() {}

func (x *MsgHasOneof) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHasOneof.ProtoReflect.Descriptor instead.
func (*MsgHasOneof) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{4}
}

func (m *MsgHasOneof) GetO() isMsgHasOneof_O {
	if m != nil {
		return m.O
	}
	return nil
}

func (x *MsgHasOneof) GetX() string {
	if x, ok := x.GetO().(*MsgHasOneof_X); ok {
		return x.X
	}
	return ""
}

func (x *MsgHasOneof) GetY() int32 {
	if x, ok := x.GetO().(*MsgHasOneof_Y); ok {
		return x.Y
	}
	return 0
}

func (x *MsgHasOneof) GetMsg() *HasMsgExprs {
	if x, ok := x.GetO().(*MsgHasOneof_Msg); ok {
		return x.Msg
	}
	return nil
}

type isMsgHasOneof_O interface {
	isMsgHasOneof_O()
}

type MsgHasOneof_X struct {
	X string `protobuf:"bytes,1,opt,name=x,proto3,oneof"`
}

type MsgHasOneof_Y struct {
	Y int32 `protobuf:"varint,2,opt,name=y,proto3,oneof"`
}

type MsgHasOneof_Msg struct {
	Msg *HasMsgExprs `protobuf:"bytes,3,opt,name=msg,proto3,oneof"`
}

func (*MsgHasOneof_X) isMsgHasOneof_O() {}

func (*MsgHasOneof_Y) isMsgHasOneof_O() {}

func (*MsgHasOneof_Msg) isMsgHasOneof_O() {}

type MsgHasRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X []float32      `protobuf:"fixed32,1,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y []string       `protobuf:"bytes,2,rep,name=y,proto3" json:"y,omitempty"`
	Z []*HasMsgExprs `protobuf:"bytes,3,rep,name=z,proto3" json:"z,omitempty"`
}

func (x *MsgHasRepeated) Reset() {
	*x = MsgHasRepeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHasRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHasRepeated) ProtoMessage() {}

func (x *MsgHasRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHasRepeated.ProtoReflect.Descriptor instead.
func (*MsgHasRepeated) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{5}
}

func (x *MsgHasRepeated) GetX() []float32 {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *MsgHasRepeated) GetY() []string {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *MsgHasRepeated) GetZ() []*HasMsgExprs {
	if x != nil {
		return x.Z
	}
	return nil
}

type MsgHasMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32Map   map[int32]int32           `protobuf:"bytes,1,rep,name=int32map,proto3" json:"int32map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	StringMap  map[string]string         `protobuf:"bytes,2,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MessageMap map[int64]*LoopRecursiveA `protobuf:"bytes,3,rep,name=message_map,json=messageMap,proto3" json:"message_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgHasMap) Reset() {
	*x = MsgHasMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHasMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHasMap) ProtoMessage() {}

func (x *MsgHasMap) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHasMap.ProtoReflect.Descriptor instead.
func (*MsgHasMap) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{6}
}

func (x *MsgHasMap) GetInt32Map() map[int32]int32 {
	if x != nil {
		return x.Int32Map
	}
	return nil
}

func (x *MsgHasMap) GetStringMap() map[string]string {
	if x != nil {
		return x.StringMap
	}
	return nil
}

func (x *MsgHasMap) GetMessageMap() map[int64]*LoopRecursiveA {
	if x != nil {
		return x.MessageMap
	}
	return nil
}

type TransitiveFieldConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *TransitiveFieldConstraint) Reset() {
	*x = TransitiveFieldConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitiveFieldConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitiveFieldConstraint) ProtoMessage() {}

func (x *TransitiveFieldConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitiveFieldConstraint.ProtoReflect.Descriptor instead.
func (*TransitiveFieldConstraint) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{7}
}

func (x *TransitiveFieldConstraint) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

type MultipleStepsTransitiveFieldConstraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *apipb.Api `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
}

func (x *MultipleStepsTransitiveFieldConstraints) Reset() {
	*x = MultipleStepsTransitiveFieldConstraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_validations_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleStepsTransitiveFieldConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleStepsTransitiveFieldConstraints) ProtoMessage() {}

func (x *MultipleStepsTransitiveFieldConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_validations_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleStepsTransitiveFieldConstraints.ProtoReflect.Descriptor instead.
func (*MultipleStepsTransitiveFieldConstraints) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_validations_proto_rawDescGZIP(), []int{8}
}

func (x *MultipleStepsTransitiveFieldConstraints) GetApi() *apipb.Api {
	if x != nil {
		return x.Api
	}
	return nil
}

var File_tests_example_v1_validations_proto protoreflect.FileDescriptor

var file_tests_example_v1_validations_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x02, 0x0a, 0x0b, 0x48, 0x61, 0x73, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x70, 0x72, 0x73,
	0x12, 0x7f, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x71, 0xba, 0x48, 0x6e,
	0xba, 0x01, 0x27, 0x0a, 0x06, 0x78, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x12, 0x0e, 0x78, 0x20, 0x6d,
	0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x1a, 0x0d, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x25, 0x20, 0x32, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0xba, 0x01, 0x41, 0x0a, 0x0b, 0x78,
	0x5f, 0x63, 0x6f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x33, 0x1a, 0x32, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x25, 0x20, 0x33, 0x20, 0x21, 0x3d, 0x20, 0x30, 0x20, 0x3f, 0x20, 0x27, 0x27, 0x3a, 0x20,
	0x27, 0x78, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x62, 0x79, 0x20, 0x33, 0x27, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x3a,
	0x6f, 0xba, 0x48, 0x6c, 0x1a, 0x30, 0x0a, 0x06, 0x78, 0x5f, 0x6c, 0x74, 0x5f, 0x79, 0x12, 0x15,
	0x78, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x68, 0x61, 0x6e, 0x20, 0x79, 0x1a, 0x0f, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x78, 0x20, 0x3c, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x79, 0x1a, 0x38, 0x0a, 0x07, 0x79, 0x5f, 0x67, 0x74, 0x5f, 0x34,
	0x32, 0x1a, 0x2d, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x79, 0x20, 0x3e, 0x20, 0x34, 0x32, 0x20, 0x3f,
	0x20, 0x27, 0x27, 0x3a, 0x20, 0x27, 0x79, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x34, 0x32, 0x27,
	0x22, 0xfe, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78,
	0x12, 0x8d, 0x01, 0x0a, 0x06, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x42, 0x54, 0xba, 0x48, 0x51, 0xba, 0x01, 0x4e, 0x0a, 0x14, 0x6e, 0x6f, 0x6e, 0x5f,
	0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x62, 0x61, 0x62, 0x79, 0x5f, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65,
	0x12, 0x2a, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x20, 0x74, 0x75, 0x72, 0x74, 0x6c,
	0x65, 0x27, 0x73, 0x20, 0x78, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x7a, 0x65, 0x72, 0x6f, 0x1a, 0x0a, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x78, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x06, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65,
	0x3a, 0x4f, 0xba, 0x48, 0x4c, 0x1a, 0x4a, 0x0a, 0x0e, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f,
	0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e,
	0x74, 0x20, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x1a, 0x17, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x78,
	0x20, 0x21, 0x3d, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x74, 0x75, 0x72, 0x74, 0x6c, 0x65, 0x2e,
	0x78, 0x22, 0x40, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x41, 0x12, 0x2e, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x42,
	0x52, 0x01, 0x62, 0x22, 0x40, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72,
	0x73, 0x69, 0x76, 0x65, 0x42, 0x12, 0x2e, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x41, 0x52, 0x01, 0x61, 0x22, 0xfb, 0x01, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x48, 0x61, 0x73,
	0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1a, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x3a, 0x03, 0x66, 0x6f, 0x6f, 0x48, 0x00, 0x52, 0x01,
	0x78, 0x12, 0x17, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x01, 0x79, 0x12, 0x31, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x4d, 0x73,
	0x67, 0x45, 0x78, 0x70, 0x72, 0x73, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x78, 0xba,
	0x48, 0x75, 0x1a, 0x5c, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x20, 0x78, 0x1a, 0x52, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x78, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x3f, 0x20, 0x27, 0x27, 0x20,
	0x3a, 0x20, 0x0a, 0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x78, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x28, 0x27, 0x66, 0x6f, 0x6f, 0x27, 0x29, 0x20, 0x3f, 0x20, 0x27,
	0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x20, 0x60, 0x66, 0x6f, 0x6f, 0x60, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27,
	0x1a, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x78, 0x74, 0x20, 0x79, 0x1a, 0x0b, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x79, 0x20, 0x3e, 0x3d, 0x20, 0x30, 0x42, 0x0a, 0x0a, 0x01, 0x6f, 0x12, 0x05, 0xba, 0x48,
	0x02, 0x08, 0x01, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x48, 0x61, 0x73, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x02, 0x42, 0x31, 0xba, 0x48, 0x2e, 0x92, 0x01, 0x2b, 0x08, 0x01, 0x10, 0x03, 0x18, 0x01, 0x22,
	0x23, 0xba, 0x01, 0x19, 0x12, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x1a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x0a, 0x05, 0x25,
	0x00, 0x00, 0x00, 0x00, 0x52, 0x01, 0x78, 0x12, 0x16, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x18, 0x01, 0x52, 0x01, 0x79, 0x12,
	0x35, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61,
	0x73, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x70, 0x72, 0x73, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01,
	0x02, 0x10, 0x02, 0x52, 0x01, 0x7a, 0x22, 0xf1, 0x03, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x48, 0x61,
	0x73, 0x4d, 0x61, 0x70, 0x12, 0x5b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x6d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x61, 0x73,
	0x4d, 0x61, 0x70, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x6d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x14, 0xba, 0x48, 0x11, 0x9a, 0x01, 0x0e, 0x08, 0x03, 0x22, 0x04, 0x1a, 0x02, 0x20,
	0x00, 0x2a, 0x04, 0x1a, 0x02, 0x10, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x6d, 0x61,
	0x70, 0x12, 0x53, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x61, 0x73, 0x4d,
	0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x08, 0xba, 0x48, 0x05, 0x9a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x56, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x48, 0x61, 0x73, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xba, 0x48, 0x05, 0x9a, 0x01, 0x02,
	0x08, 0x02, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x3b,
	0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x6d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5f, 0x0a, 0x0f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x41, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x71, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x41, 0xba, 0x48, 0x3e, 0xba, 0x01, 0x3b, 0x0a, 0x0a, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x1c, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x29, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0xce, 0x01, 0x0a, 0x27,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x65, 0x70, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0xa2, 0x01, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x69, 0x42, 0x7a, 0xba, 0x48, 0x77,
	0xba, 0x01, 0x74, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x61, 0x70, 0x69, 0x27, 0x73, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x6e, 0x61, 0x6d,
	0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x22, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x52, 0x03, 0x61, 0x70, 0x69, 0x42, 0xd8, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x10,
	0x54, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x73, 0x74, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_example_v1_validations_proto_rawDescOnce sync.Once
	file_tests_example_v1_validations_proto_rawDescData = file_tests_example_v1_validations_proto_rawDesc
)

func file_tests_example_v1_validations_proto_rawDescGZIP() []byte {
	file_tests_example_v1_validations_proto_rawDescOnce.Do(func() {
		file_tests_example_v1_validations_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_example_v1_validations_proto_rawDescData)
	})
	return file_tests_example_v1_validations_proto_rawDescData
}

var file_tests_example_v1_validations_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tests_example_v1_validations_proto_goTypes = []interface{}{
	(*HasMsgExprs)(nil),                             // 0: tests.example.v1.HasMsgExprs
	(*SelfRecursive)(nil),                           // 1: tests.example.v1.SelfRecursive
	(*LoopRecursiveA)(nil),                          // 2: tests.example.v1.LoopRecursiveA
	(*LoopRecursiveB)(nil),                          // 3: tests.example.v1.LoopRecursiveB
	(*MsgHasOneof)(nil),                             // 4: tests.example.v1.MsgHasOneof
	(*MsgHasRepeated)(nil),                          // 5: tests.example.v1.MsgHasRepeated
	(*MsgHasMap)(nil),                               // 6: tests.example.v1.MsgHasMap
	(*TransitiveFieldConstraint)(nil),               // 7: tests.example.v1.TransitiveFieldConstraint
	(*MultipleStepsTransitiveFieldConstraints)(nil), // 8: tests.example.v1.MultipleStepsTransitiveFieldConstraints
	nil,                           // 9: tests.example.v1.MsgHasMap.Int32mapEntry
	nil,                           // 10: tests.example.v1.MsgHasMap.StringMapEntry
	nil,                           // 11: tests.example.v1.MsgHasMap.MessageMapEntry
	(*fieldmaskpb.FieldMask)(nil), // 12: google.protobuf.FieldMask
	(*apipb.Api)(nil),             // 13: google.protobuf.Api
}
var file_tests_example_v1_validations_proto_depIdxs = []int32{
	1,  // 0: tests.example.v1.SelfRecursive.turtle:type_name -> tests.example.v1.SelfRecursive
	3,  // 1: tests.example.v1.LoopRecursiveA.b:type_name -> tests.example.v1.LoopRecursiveB
	2,  // 2: tests.example.v1.LoopRecursiveB.a:type_name -> tests.example.v1.LoopRecursiveA
	0,  // 3: tests.example.v1.MsgHasOneof.msg:type_name -> tests.example.v1.HasMsgExprs
	0,  // 4: tests.example.v1.MsgHasRepeated.z:type_name -> tests.example.v1.HasMsgExprs
	9,  // 5: tests.example.v1.MsgHasMap.int32map:type_name -> tests.example.v1.MsgHasMap.Int32mapEntry
	10, // 6: tests.example.v1.MsgHasMap.string_map:type_name -> tests.example.v1.MsgHasMap.StringMapEntry
	11, // 7: tests.example.v1.MsgHasMap.message_map:type_name -> tests.example.v1.MsgHasMap.MessageMapEntry
	12, // 8: tests.example.v1.TransitiveFieldConstraint.mask:type_name -> google.protobuf.FieldMask
	13, // 9: tests.example.v1.MultipleStepsTransitiveFieldConstraints.api:type_name -> google.protobuf.Api
	2,  // 10: tests.example.v1.MsgHasMap.MessageMapEntry.value:type_name -> tests.example.v1.LoopRecursiveA
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tests_example_v1_validations_proto_init() }
func file_tests_example_v1_validations_proto_init() {
	if File_tests_example_v1_validations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_example_v1_validations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasMsgExprs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfRecursive); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoopRecursiveA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoopRecursiveB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHasOneof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHasRepeated); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHasMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitiveFieldConstraint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_example_v1_validations_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleStepsTransitiveFieldConstraints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tests_example_v1_validations_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*MsgHasOneof_X)(nil),
		(*MsgHasOneof_Y)(nil),
		(*MsgHasOneof_Msg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_example_v1_validations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_example_v1_validations_proto_goTypes,
		DependencyIndexes: file_tests_example_v1_validations_proto_depIdxs,
		MessageInfos:      file_tests_example_v1_validations_proto_msgTypes,
	}.Build()
	File_tests_example_v1_validations_proto = out.File
	file_tests_example_v1_validations_proto_rawDesc = nil
	file_tests_example_v1_validations_proto_goTypes = nil
	file_tests_example_v1_validations_proto_depIdxs = nil
}
