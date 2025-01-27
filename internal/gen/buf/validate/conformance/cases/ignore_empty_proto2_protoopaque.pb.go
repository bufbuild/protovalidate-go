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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/ignore_empty_proto2.proto

//go:build protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IgnoreEmptyProto2ScalarOptional struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val         int32                  `protobuf:"varint,1,opt,name=val"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2ScalarOptional) Reset() {
	*x = IgnoreEmptyProto2ScalarOptional{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2ScalarOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2ScalarOptional) ProtoMessage() {}

func (x *IgnoreEmptyProto2ScalarOptional) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2ScalarOptional) GetVal() int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return 0
}

func (x *IgnoreEmptyProto2ScalarOptional) SetVal(v int32) {
	x.xxx_hidden_Val = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *IgnoreEmptyProto2ScalarOptional) HasVal() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *IgnoreEmptyProto2ScalarOptional) ClearVal() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Val = 0
}

type IgnoreEmptyProto2ScalarOptional_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *int32
}

func (b0 IgnoreEmptyProto2ScalarOptional_builder) Build() *IgnoreEmptyProto2ScalarOptional {
	m0 := &IgnoreEmptyProto2ScalarOptional{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Val = *b.Val
	}
	return m0
}

type IgnoreEmptyProto2ScalarOptionalWithDefault struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val         int32                  `protobuf:"varint,1,opt,name=val,def=42"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

// Default values for IgnoreEmptyProto2ScalarOptionalWithDefault fields.
const (
	Default_IgnoreEmptyProto2ScalarOptionalWithDefault_Val = int32(42)
)

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) Reset() {
	*x = IgnoreEmptyProto2ScalarOptionalWithDefault{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2ScalarOptionalWithDefault) ProtoMessage() {}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) GetVal() int32 {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_Val
		}
	}
	return Default_IgnoreEmptyProto2ScalarOptionalWithDefault_Val
}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) SetVal(v int32) {
	x.xxx_hidden_Val = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) HasVal() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *IgnoreEmptyProto2ScalarOptionalWithDefault) ClearVal() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
}

type IgnoreEmptyProto2ScalarOptionalWithDefault_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *int32
}

func (b0 IgnoreEmptyProto2ScalarOptionalWithDefault_builder) Build() *IgnoreEmptyProto2ScalarOptionalWithDefault {
	m0 := &IgnoreEmptyProto2ScalarOptionalWithDefault{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Val = *b.Val
	}
	return m0
}

type IgnoreEmptyProto2ScalarRequired struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val         int32                  `protobuf:"varint,1,req,name=val"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2ScalarRequired) Reset() {
	*x = IgnoreEmptyProto2ScalarRequired{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2ScalarRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2ScalarRequired) ProtoMessage() {}

func (x *IgnoreEmptyProto2ScalarRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2ScalarRequired) GetVal() int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return 0
}

func (x *IgnoreEmptyProto2ScalarRequired) SetVal(v int32) {
	x.xxx_hidden_Val = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *IgnoreEmptyProto2ScalarRequired) HasVal() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *IgnoreEmptyProto2ScalarRequired) ClearVal() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Val = 0
}

type IgnoreEmptyProto2ScalarRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *int32
}

func (b0 IgnoreEmptyProto2ScalarRequired_builder) Build() *IgnoreEmptyProto2ScalarRequired {
	m0 := &IgnoreEmptyProto2ScalarRequired{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Val = *b.Val
	}
	return m0
}

type IgnoreEmptyProto2Message struct {
	state          protoimpl.MessageState        `protogen:"opaque.v1"`
	xxx_hidden_Val *IgnoreEmptyProto2Message_Msg `protobuf:"bytes,1,opt,name=val"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2Message) Reset() {
	*x = IgnoreEmptyProto2Message{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2Message) ProtoMessage() {}

func (x *IgnoreEmptyProto2Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2Message) GetVal() *IgnoreEmptyProto2Message_Msg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto2Message) SetVal(v *IgnoreEmptyProto2Message_Msg) {
	x.xxx_hidden_Val = v
}

func (x *IgnoreEmptyProto2Message) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *IgnoreEmptyProto2Message) ClearVal() {
	x.xxx_hidden_Val = nil
}

type IgnoreEmptyProto2Message_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *IgnoreEmptyProto2Message_Msg
}

func (b0 IgnoreEmptyProto2Message_builder) Build() *IgnoreEmptyProto2Message {
	m0 := &IgnoreEmptyProto2Message{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto2Oneof struct {
	state         protoimpl.MessageState     `protogen:"opaque.v1"`
	xxx_hidden_O  isIgnoreEmptyProto2Oneof_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2Oneof) Reset() {
	*x = IgnoreEmptyProto2Oneof{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2Oneof) ProtoMessage() {}

func (x *IgnoreEmptyProto2Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2Oneof) GetVal() int32 {
	if x != nil {
		if x, ok := x.xxx_hidden_O.(*ignoreEmptyProto2Oneof_Val); ok {
			return x.Val
		}
	}
	return 0
}

func (x *IgnoreEmptyProto2Oneof) SetVal(v int32) {
	x.xxx_hidden_O = &ignoreEmptyProto2Oneof_Val{v}
}

func (x *IgnoreEmptyProto2Oneof) HasO() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_O != nil
}

func (x *IgnoreEmptyProto2Oneof) HasVal() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_O.(*ignoreEmptyProto2Oneof_Val)
	return ok
}

func (x *IgnoreEmptyProto2Oneof) ClearO() {
	x.xxx_hidden_O = nil
}

func (x *IgnoreEmptyProto2Oneof) ClearVal() {
	if _, ok := x.xxx_hidden_O.(*ignoreEmptyProto2Oneof_Val); ok {
		x.xxx_hidden_O = nil
	}
}

const IgnoreEmptyProto2Oneof_O_not_set_case case_IgnoreEmptyProto2Oneof_O = 0
const IgnoreEmptyProto2Oneof_Val_case case_IgnoreEmptyProto2Oneof_O = 1

func (x *IgnoreEmptyProto2Oneof) WhichO() case_IgnoreEmptyProto2Oneof_O {
	if x == nil {
		return IgnoreEmptyProto2Oneof_O_not_set_case
	}
	switch x.xxx_hidden_O.(type) {
	case *ignoreEmptyProto2Oneof_Val:
		return IgnoreEmptyProto2Oneof_Val_case
	default:
		return IgnoreEmptyProto2Oneof_O_not_set_case
	}
}

type IgnoreEmptyProto2Oneof_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_O:
	Val *int32
	// -- end of xxx_hidden_O
}

func (b0 IgnoreEmptyProto2Oneof_builder) Build() *IgnoreEmptyProto2Oneof {
	m0 := &IgnoreEmptyProto2Oneof{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		x.xxx_hidden_O = &ignoreEmptyProto2Oneof_Val{*b.Val}
	}
	return m0
}

type case_IgnoreEmptyProto2Oneof_O protoreflect.FieldNumber

func (x case_IgnoreEmptyProto2Oneof_O) String() string {
	md := file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[4].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isIgnoreEmptyProto2Oneof_O interface {
	isIgnoreEmptyProto2Oneof_O()
}

type ignoreEmptyProto2Oneof_Val struct {
	Val int32 `protobuf:"varint,1,opt,name=val,oneof"`
}

func (*ignoreEmptyProto2Oneof_Val) isIgnoreEmptyProto2Oneof_O() {}

type IgnoreEmptyProto2Repeated struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []int32                `protobuf:"varint,1,rep,name=val"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2Repeated) Reset() {
	*x = IgnoreEmptyProto2Repeated{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2Repeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2Repeated) ProtoMessage() {}

func (x *IgnoreEmptyProto2Repeated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2Repeated) GetVal() []int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto2Repeated) SetVal(v []int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto2Repeated_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []int32
}

func (b0 IgnoreEmptyProto2Repeated_builder) Build() *IgnoreEmptyProto2Repeated {
	m0 := &IgnoreEmptyProto2Repeated{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto2Map struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val map[int32]int32        `protobuf:"bytes,1,rep,name=val" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2Map) Reset() {
	*x = IgnoreEmptyProto2Map{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2Map) ProtoMessage() {}

func (x *IgnoreEmptyProto2Map) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2Map) GetVal() map[int32]int32 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *IgnoreEmptyProto2Map) SetVal(v map[int32]int32) {
	x.xxx_hidden_Val = v
}

type IgnoreEmptyProto2Map_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[int32]int32
}

func (b0 IgnoreEmptyProto2Map_builder) Build() *IgnoreEmptyProto2Map {
	m0 := &IgnoreEmptyProto2Map{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type IgnoreEmptyProto2Message_Msg struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val         *string                `protobuf:"bytes,1,opt,name=val"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *IgnoreEmptyProto2Message_Msg) Reset() {
	*x = IgnoreEmptyProto2Message_Msg{}
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgnoreEmptyProto2Message_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto2Message_Msg) ProtoMessage() {}

func (x *IgnoreEmptyProto2Message_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IgnoreEmptyProto2Message_Msg) GetVal() string {
	if x != nil {
		if x.xxx_hidden_Val != nil {
			return *x.xxx_hidden_Val
		}
		return ""
	}
	return ""
}

func (x *IgnoreEmptyProto2Message_Msg) SetVal(v string) {
	x.xxx_hidden_Val = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *IgnoreEmptyProto2Message_Msg) HasVal() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *IgnoreEmptyProto2Message_Msg) ClearVal() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Val = nil
}

type IgnoreEmptyProto2Message_Msg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *string
}

func (b0 IgnoreEmptyProto2Message_Msg_builder) Build() *IgnoreEmptyProto2Message_Msg {
	m0 := &IgnoreEmptyProto2Message_Msg{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Val = b.Val
	}
	return m0
}

var File_buf_validate_conformance_cases_ignore_empty_proto2_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_ignore_empty_proto2_proto_rawDesc = string([]byte{
	0x0a, 0x38, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x1f, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xd0, 0x01, 0x01, 0x1a,
	0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x4e, 0x0a, 0x2a, 0x49, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x3a, 0x02, 0x34, 0x32, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xd0, 0x01, 0x01, 0x1a,
	0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x1f, 0x49, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xd0, 0x01, 0x01,
	0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xc7, 0x01, 0x0a, 0x18, 0x49, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x42, 0x41, 0xba, 0x48, 0x3e, 0xba, 0x01, 0x38, 0x0a, 0x1b, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x06, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x1a,
	0x11, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x61, 0x6c, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x66, 0x6f,
	0x6f, 0x27, 0xd0, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x17, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1e, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xd0,
	0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x03, 0x0a,
	0x01, 0x6f, 0x22, 0x3a, 0x0a, 0x19, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0b, 0xba, 0x48,
	0x08, 0xd0, 0x01, 0x01, 0x92, 0x01, 0x02, 0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xac,
	0x01, 0x0a, 0x14, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x61, 0x70, 0x12, 0x5c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xd0, 0x01, 0x01, 0x9a, 0x01, 0x02, 0x08, 0x03,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xaa, 0x02,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x42, 0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73,
})

var file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_buf_validate_conformance_cases_ignore_empty_proto2_proto_goTypes = []any{
	(*IgnoreEmptyProto2ScalarOptional)(nil),            // 0: buf.validate.conformance.cases.IgnoreEmptyProto2ScalarOptional
	(*IgnoreEmptyProto2ScalarOptionalWithDefault)(nil), // 1: buf.validate.conformance.cases.IgnoreEmptyProto2ScalarOptionalWithDefault
	(*IgnoreEmptyProto2ScalarRequired)(nil),            // 2: buf.validate.conformance.cases.IgnoreEmptyProto2ScalarRequired
	(*IgnoreEmptyProto2Message)(nil),                   // 3: buf.validate.conformance.cases.IgnoreEmptyProto2Message
	(*IgnoreEmptyProto2Oneof)(nil),                     // 4: buf.validate.conformance.cases.IgnoreEmptyProto2Oneof
	(*IgnoreEmptyProto2Repeated)(nil),                  // 5: buf.validate.conformance.cases.IgnoreEmptyProto2Repeated
	(*IgnoreEmptyProto2Map)(nil),                       // 6: buf.validate.conformance.cases.IgnoreEmptyProto2Map
	(*IgnoreEmptyProto2Message_Msg)(nil),               // 7: buf.validate.conformance.cases.IgnoreEmptyProto2Message.Msg
	nil,                                                // 8: buf.validate.conformance.cases.IgnoreEmptyProto2Map.ValEntry
}
var file_buf_validate_conformance_cases_ignore_empty_proto2_proto_depIdxs = []int32{
	7, // 0: buf.validate.conformance.cases.IgnoreEmptyProto2Message.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto2Message.Msg
	8, // 1: buf.validate.conformance.cases.IgnoreEmptyProto2Map.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto2Map.ValEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_ignore_empty_proto2_proto_init() }
func file_buf_validate_conformance_cases_ignore_empty_proto2_proto_init() {
	if File_buf_validate_conformance_cases_ignore_empty_proto2_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes[4].OneofWrappers = []any{
		(*ignoreEmptyProto2Oneof_Val)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_ignore_empty_proto2_proto_rawDesc), len(file_buf_validate_conformance_cases_ignore_empty_proto2_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_ignore_empty_proto2_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_ignore_empty_proto2_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_ignore_empty_proto2_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_ignore_empty_proto2_proto = out.File
	file_buf_validate_conformance_cases_ignore_empty_proto2_proto_goTypes = nil
	file_buf_validate_conformance_cases_ignore_empty_proto2_proto_depIdxs = nil
}
