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
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/required_field_proto2.proto

//go:build !protoopaque

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

type RequiredProto2ScalarOptional struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *string                `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2ScalarOptional) Reset() {
	*x = RequiredProto2ScalarOptional{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2ScalarOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2ScalarOptional) ProtoMessage() {}

func (x *RequiredProto2ScalarOptional) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2ScalarOptional) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

func (x *RequiredProto2ScalarOptional) SetVal(v string) {
	x.Val = &v
}

func (x *RequiredProto2ScalarOptional) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2ScalarOptional) ClearVal() {
	x.Val = nil
}

type RequiredProto2ScalarOptional_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *string
}

func (b0 RequiredProto2ScalarOptional_builder) Build() *RequiredProto2ScalarOptional {
	m0 := &RequiredProto2ScalarOptional{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2ScalarOptionalDefault struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *string                `protobuf:"bytes,1,opt,name=val,def=foo" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Default values for RequiredProto2ScalarOptionalDefault fields.
const (
	Default_RequiredProto2ScalarOptionalDefault_Val = string("foo")
)

func (x *RequiredProto2ScalarOptionalDefault) Reset() {
	*x = RequiredProto2ScalarOptionalDefault{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2ScalarOptionalDefault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2ScalarOptionalDefault) ProtoMessage() {}

func (x *RequiredProto2ScalarOptionalDefault) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2ScalarOptionalDefault) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return Default_RequiredProto2ScalarOptionalDefault_Val
}

func (x *RequiredProto2ScalarOptionalDefault) SetVal(v string) {
	x.Val = &v
}

func (x *RequiredProto2ScalarOptionalDefault) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2ScalarOptionalDefault) ClearVal() {
	x.Val = nil
}

type RequiredProto2ScalarOptionalDefault_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *string
}

func (b0 RequiredProto2ScalarOptionalDefault_builder) Build() *RequiredProto2ScalarOptionalDefault {
	m0 := &RequiredProto2ScalarOptionalDefault{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2ScalarRequired struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *string                `protobuf:"bytes,1,req,name=val" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2ScalarRequired) Reset() {
	*x = RequiredProto2ScalarRequired{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2ScalarRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2ScalarRequired) ProtoMessage() {}

func (x *RequiredProto2ScalarRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2ScalarRequired) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

func (x *RequiredProto2ScalarRequired) SetVal(v string) {
	x.Val = &v
}

func (x *RequiredProto2ScalarRequired) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2ScalarRequired) ClearVal() {
	x.Val = nil
}

type RequiredProto2ScalarRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *string
}

func (b0 RequiredProto2ScalarRequired_builder) Build() *RequiredProto2ScalarRequired {
	m0 := &RequiredProto2ScalarRequired{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2Message struct {
	state         protoimpl.MessageState     `protogen:"hybrid.v1"`
	Val           *RequiredProto2Message_Msg `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2Message) Reset() {
	*x = RequiredProto2Message{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2Message) ProtoMessage() {}

func (x *RequiredProto2Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2Message) GetVal() *RequiredProto2Message_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *RequiredProto2Message) SetVal(v *RequiredProto2Message_Msg) {
	x.Val = v
}

func (x *RequiredProto2Message) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2Message) ClearVal() {
	x.Val = nil
}

type RequiredProto2Message_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *RequiredProto2Message_Msg
}

func (b0 RequiredProto2Message_builder) Build() *RequiredProto2Message {
	m0 := &RequiredProto2Message{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2Oneof struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to Val:
	//
	//	*RequiredProto2Oneof_A
	//	*RequiredProto2Oneof_B
	Val           isRequiredProto2Oneof_Val `protobuf_oneof:"val"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2Oneof) Reset() {
	*x = RequiredProto2Oneof{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2Oneof) ProtoMessage() {}

func (x *RequiredProto2Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2Oneof) GetVal() isRequiredProto2Oneof_Val {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *RequiredProto2Oneof) GetA() string {
	if x != nil {
		if x, ok := x.Val.(*RequiredProto2Oneof_A); ok {
			return x.A
		}
	}
	return ""
}

func (x *RequiredProto2Oneof) GetB() string {
	if x != nil {
		if x, ok := x.Val.(*RequiredProto2Oneof_B); ok {
			return x.B
		}
	}
	return ""
}

func (x *RequiredProto2Oneof) SetA(v string) {
	x.Val = &RequiredProto2Oneof_A{v}
}

func (x *RequiredProto2Oneof) SetB(v string) {
	x.Val = &RequiredProto2Oneof_B{v}
}

func (x *RequiredProto2Oneof) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2Oneof) HasA() bool {
	if x == nil {
		return false
	}
	_, ok := x.Val.(*RequiredProto2Oneof_A)
	return ok
}

func (x *RequiredProto2Oneof) HasB() bool {
	if x == nil {
		return false
	}
	_, ok := x.Val.(*RequiredProto2Oneof_B)
	return ok
}

func (x *RequiredProto2Oneof) ClearVal() {
	x.Val = nil
}

func (x *RequiredProto2Oneof) ClearA() {
	if _, ok := x.Val.(*RequiredProto2Oneof_A); ok {
		x.Val = nil
	}
}

func (x *RequiredProto2Oneof) ClearB() {
	if _, ok := x.Val.(*RequiredProto2Oneof_B); ok {
		x.Val = nil
	}
}

const RequiredProto2Oneof_Val_not_set_case case_RequiredProto2Oneof_Val = 0
const RequiredProto2Oneof_A_case case_RequiredProto2Oneof_Val = 1
const RequiredProto2Oneof_B_case case_RequiredProto2Oneof_Val = 2

func (x *RequiredProto2Oneof) WhichVal() case_RequiredProto2Oneof_Val {
	if x == nil {
		return RequiredProto2Oneof_Val_not_set_case
	}
	switch x.Val.(type) {
	case *RequiredProto2Oneof_A:
		return RequiredProto2Oneof_A_case
	case *RequiredProto2Oneof_B:
		return RequiredProto2Oneof_B_case
	default:
		return RequiredProto2Oneof_Val_not_set_case
	}
}

type RequiredProto2Oneof_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof Val:
	A *string
	B *string
	// -- end of Val
}

func (b0 RequiredProto2Oneof_builder) Build() *RequiredProto2Oneof {
	m0 := &RequiredProto2Oneof{}
	b, x := &b0, m0
	_, _ = b, x
	if b.A != nil {
		x.Val = &RequiredProto2Oneof_A{*b.A}
	}
	if b.B != nil {
		x.Val = &RequiredProto2Oneof_B{*b.B}
	}
	return m0
}

type case_RequiredProto2Oneof_Val protoreflect.FieldNumber

func (x case_RequiredProto2Oneof_Val) String() string {
	md := file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[4].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isRequiredProto2Oneof_Val interface {
	isRequiredProto2Oneof_Val()
}

type RequiredProto2Oneof_A struct {
	A string `protobuf:"bytes,1,opt,name=a,oneof"`
}

type RequiredProto2Oneof_B struct {
	B string `protobuf:"bytes,2,opt,name=b,oneof"`
}

func (*RequiredProto2Oneof_A) isRequiredProto2Oneof_Val() {}

func (*RequiredProto2Oneof_B) isRequiredProto2Oneof_Val() {}

type RequiredProto2Repeated struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           []string               `protobuf:"bytes,1,rep,name=val" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2Repeated) Reset() {
	*x = RequiredProto2Repeated{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2Repeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2Repeated) ProtoMessage() {}

func (x *RequiredProto2Repeated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2Repeated) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *RequiredProto2Repeated) SetVal(v []string) {
	x.Val = v
}

type RequiredProto2Repeated_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []string
}

func (b0 RequiredProto2Repeated_builder) Build() *RequiredProto2Repeated {
	m0 := &RequiredProto2Repeated{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2Map struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[string]string      `protobuf:"bytes,1,rep,name=val" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2Map) Reset() {
	*x = RequiredProto2Map{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2Map) ProtoMessage() {}

func (x *RequiredProto2Map) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2Map) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *RequiredProto2Map) SetVal(v map[string]string) {
	x.Val = v
}

type RequiredProto2Map_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]string
}

func (b0 RequiredProto2Map_builder) Build() *RequiredProto2Map {
	m0 := &RequiredProto2Map{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type RequiredProto2Message_Msg struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *string                `protobuf:"bytes,1,opt,name=val" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequiredProto2Message_Msg) Reset() {
	*x = RequiredProto2Message_Msg{}
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequiredProto2Message_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequiredProto2Message_Msg) ProtoMessage() {}

func (x *RequiredProto2Message_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RequiredProto2Message_Msg) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

func (x *RequiredProto2Message_Msg) SetVal(v string) {
	x.Val = &v
}

func (x *RequiredProto2Message_Msg) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *RequiredProto2Message_Msg) ClearVal() {
	x.Val = nil
}

type RequiredProto2Message_Msg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *string
}

func (b0 RequiredProto2Message_Msg_builder) Build() *RequiredProto2Message_Msg {
	m0 := &RequiredProto2Message_Msg{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_required_field_proto2_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_required_field_proto2_proto_rawDesc = string([]byte{
	0x0a, 0x3a, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x1c, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x44, 0x0a, 0x23, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x03, 0x66, 0x6f, 0x6f, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x38, 0x0a, 0x1c, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x1a, 0x17, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x44, 0x0a, 0x13, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4f, 0x6e, 0x65,
	0x6f, 0x66, 0x12, 0x16, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01, 0x62, 0x42, 0x05, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x61, 0x70, 0x12, 0x54, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xac, 0x02, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x42, 0x18, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d,
	0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02,
	0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73,
})

var file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_buf_validate_conformance_cases_required_field_proto2_proto_goTypes = []any{
	(*RequiredProto2ScalarOptional)(nil),        // 0: buf.validate.conformance.cases.RequiredProto2ScalarOptional
	(*RequiredProto2ScalarOptionalDefault)(nil), // 1: buf.validate.conformance.cases.RequiredProto2ScalarOptionalDefault
	(*RequiredProto2ScalarRequired)(nil),        // 2: buf.validate.conformance.cases.RequiredProto2ScalarRequired
	(*RequiredProto2Message)(nil),               // 3: buf.validate.conformance.cases.RequiredProto2Message
	(*RequiredProto2Oneof)(nil),                 // 4: buf.validate.conformance.cases.RequiredProto2Oneof
	(*RequiredProto2Repeated)(nil),              // 5: buf.validate.conformance.cases.RequiredProto2Repeated
	(*RequiredProto2Map)(nil),                   // 6: buf.validate.conformance.cases.RequiredProto2Map
	(*RequiredProto2Message_Msg)(nil),           // 7: buf.validate.conformance.cases.RequiredProto2Message.Msg
	nil,                                         // 8: buf.validate.conformance.cases.RequiredProto2Map.ValEntry
}
var file_buf_validate_conformance_cases_required_field_proto2_proto_depIdxs = []int32{
	7, // 0: buf.validate.conformance.cases.RequiredProto2Message.val:type_name -> buf.validate.conformance.cases.RequiredProto2Message.Msg
	8, // 1: buf.validate.conformance.cases.RequiredProto2Map.val:type_name -> buf.validate.conformance.cases.RequiredProto2Map.ValEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_required_field_proto2_proto_init() }
func file_buf_validate_conformance_cases_required_field_proto2_proto_init() {
	if File_buf_validate_conformance_cases_required_field_proto2_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes[4].OneofWrappers = []any{
		(*RequiredProto2Oneof_A)(nil),
		(*RequiredProto2Oneof_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_required_field_proto2_proto_rawDesc), len(file_buf_validate_conformance_cases_required_field_proto2_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_required_field_proto2_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_required_field_proto2_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_required_field_proto2_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_required_field_proto2_proto = out.File
	file_buf_validate_conformance_cases_required_field_proto2_proto_goTypes = nil
	file_buf_validate_conformance_cases_required_field_proto2_proto_depIdxs = nil
}
