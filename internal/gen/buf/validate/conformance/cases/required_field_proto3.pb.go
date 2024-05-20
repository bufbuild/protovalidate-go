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
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/required_field_proto3.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequiredProto3Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RequiredProto3Scalar) Reset() {
	*x = RequiredProto3Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3Scalar) ProtoMessage() {}

func (x *RequiredProto3Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3Scalar.ProtoReflect.Descriptor instead.
func (*RequiredProto3Scalar) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{0}
}

func (x *RequiredProto3Scalar) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type RequiredProto3OptionalScalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *string `protobuf:"bytes,1,opt,name=val,proto3,oneof" json:"val,omitempty"`
}

func (x *RequiredProto3OptionalScalar) Reset() {
	*x = RequiredProto3OptionalScalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3OptionalScalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3OptionalScalar) ProtoMessage() {}

func (x *RequiredProto3OptionalScalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3OptionalScalar.ProtoReflect.Descriptor instead.
func (*RequiredProto3OptionalScalar) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{1}
}

func (x *RequiredProto3OptionalScalar) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

type RequiredProto3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *RequiredProto3Message_Msg `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RequiredProto3Message) Reset() {
	*x = RequiredProto3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3Message) ProtoMessage() {}

func (x *RequiredProto3Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3Message.ProtoReflect.Descriptor instead.
func (*RequiredProto3Message) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{2}
}

func (x *RequiredProto3Message) GetVal() *RequiredProto3Message_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

type RequiredProto3OneOf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Val:
	//
	//	*RequiredProto3OneOf_A
	//	*RequiredProto3OneOf_B
	Val isRequiredProto3OneOf_Val `protobuf_oneof:"val"`
}

func (x *RequiredProto3OneOf) Reset() {
	*x = RequiredProto3OneOf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3OneOf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3OneOf) ProtoMessage() {}

func (x *RequiredProto3OneOf) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3OneOf.ProtoReflect.Descriptor instead.
func (*RequiredProto3OneOf) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{3}
}

func (m *RequiredProto3OneOf) GetVal() isRequiredProto3OneOf_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (x *RequiredProto3OneOf) GetA() string {
	if x, ok := x.GetVal().(*RequiredProto3OneOf_A); ok {
		return x.A
	}
	return ""
}

func (x *RequiredProto3OneOf) GetB() string {
	if x, ok := x.GetVal().(*RequiredProto3OneOf_B); ok {
		return x.B
	}
	return ""
}

type isRequiredProto3OneOf_Val interface {
	isRequiredProto3OneOf_Val()
}

type RequiredProto3OneOf_A struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3,oneof"`
}

type RequiredProto3OneOf_B struct {
	B string `protobuf:"bytes,2,opt,name=b,proto3,oneof"`
}

func (*RequiredProto3OneOf_A) isRequiredProto3OneOf_Val() {}

func (*RequiredProto3OneOf_B) isRequiredProto3OneOf_Val() {}

type RequiredProto3Repeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *RequiredProto3Repeated) Reset() {
	*x = RequiredProto3Repeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3Repeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3Repeated) ProtoMessage() {}

func (x *RequiredProto3Repeated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3Repeated.ProtoReflect.Descriptor instead.
func (*RequiredProto3Repeated) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{4}
}

func (x *RequiredProto3Repeated) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

type RequiredProto3Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequiredProto3Map) Reset() {
	*x = RequiredProto3Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3Map) ProtoMessage() {}

func (x *RequiredProto3Map) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3Map.ProtoReflect.Descriptor instead.
func (*RequiredProto3Map) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{5}
}

func (x *RequiredProto3Map) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type RequiredProto3Message_Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RequiredProto3Message_Msg) Reset() {
	*x = RequiredProto3Message_Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequiredProto3Message_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto3Message_Msg) ProtoMessage() {}

func (x *RequiredProto3Message_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequiredProto3Message_Msg.ProtoReflect.Descriptor instead.
func (*RequiredProto3Message_Msg) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RequiredProto3Message_Msg) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_buf_validate_conformance_cases_required_field_proto3_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_required_field_proto3_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x12, 0x18, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x45, 0x0a, 0x1c, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x48, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76,
	0x61, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4d, 0x73, 0x67, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x1a, 0x17, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x44, 0x0a, 0x13, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x6e, 0x65, 0x4f,
	0x66, 0x12, 0x16, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x62, 0x42, 0x05, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x12, 0x54, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xac, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x18, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67,
	0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04,
	0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescData = file_buf_validate_conformance_cases_required_field_proto3_proto_rawDesc
)

func file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_required_field_proto3_proto_rawDescData
}

var file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_buf_validate_conformance_cases_required_field_proto3_proto_goTypes = []interface{}{
	(*RequiredProto3Scalar)(nil),         // 0: buf.validate.conformance.cases.RequiredProto3Scalar
	(*RequiredProto3OptionalScalar)(nil), // 1: buf.validate.conformance.cases.RequiredProto3OptionalScalar
	(*RequiredProto3Message)(nil),        // 2: buf.validate.conformance.cases.RequiredProto3Message
	(*RequiredProto3OneOf)(nil),          // 3: buf.validate.conformance.cases.RequiredProto3OneOf
	(*RequiredProto3Repeated)(nil),       // 4: buf.validate.conformance.cases.RequiredProto3Repeated
	(*RequiredProto3Map)(nil),            // 5: buf.validate.conformance.cases.RequiredProto3Map
	(*RequiredProto3Message_Msg)(nil),    // 6: buf.validate.conformance.cases.RequiredProto3Message.Msg
	nil,                                  // 7: buf.validate.conformance.cases.RequiredProto3Map.ValEntry
}
var file_buf_validate_conformance_cases_required_field_proto3_proto_depIdxs = []int32{
	6, // 0: buf.validate.conformance.cases.RequiredProto3Message.val:type_name -> buf.validate.conformance.cases.RequiredProto3Message.Msg
	7, // 1: buf.validate.conformance.cases.RequiredProto3Map.val:type_name -> buf.validate.conformance.cases.RequiredProto3Map.ValEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_required_field_proto3_proto_init() }
func file_buf_validate_conformance_cases_required_field_proto3_proto_init() {
	if File_buf_validate_conformance_cases_required_field_proto3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3Scalar); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3OptionalScalar); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3Message); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3OneOf); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3Repeated); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3Map); i {
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
		file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequiredProto3Message_Msg); i {
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
	file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RequiredProto3OneOf_A)(nil),
		(*RequiredProto3OneOf_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_required_field_proto3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_required_field_proto3_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_required_field_proto3_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_required_field_proto3_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_required_field_proto3_proto = out.File
	file_buf_validate_conformance_cases_required_field_proto3_proto_rawDesc = nil
	file_buf_validate_conformance_cases_required_field_proto3_proto_goTypes = nil
	file_buf_validate_conformance_cases_required_field_proto3_proto_depIdxs = nil
}
