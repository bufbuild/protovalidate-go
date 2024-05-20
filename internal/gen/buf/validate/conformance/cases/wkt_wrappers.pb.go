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
// source: buf/validate/conformance/cases/wkt_wrappers.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WrapperNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperNone) Reset() {
	*x = WrapperNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperNone) ProtoMessage() {}

func (x *WrapperNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperNone.ProtoReflect.Descriptor instead.
func (*WrapperNone) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{0}
}

func (x *WrapperNone) GetVal() *wrapperspb.Int32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperFloat) Reset() {
	*x = WrapperFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperFloat) ProtoMessage() {}

func (x *WrapperFloat) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperFloat.ProtoReflect.Descriptor instead.
func (*WrapperFloat) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{1}
}

func (x *WrapperFloat) GetVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperDouble struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperDouble) Reset() {
	*x = WrapperDouble{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperDouble) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperDouble) ProtoMessage() {}

func (x *WrapperDouble) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperDouble.ProtoReflect.Descriptor instead.
func (*WrapperDouble) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{2}
}

func (x *WrapperDouble) GetVal() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperInt64) Reset() {
	*x = WrapperInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperInt64) ProtoMessage() {}

func (x *WrapperInt64) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperInt64.ProtoReflect.Descriptor instead.
func (*WrapperInt64) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{3}
}

func (x *WrapperInt64) GetVal() *wrapperspb.Int64Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperInt32) Reset() {
	*x = WrapperInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperInt32) ProtoMessage() {}

func (x *WrapperInt32) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperInt32.ProtoReflect.Descriptor instead.
func (*WrapperInt32) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{4}
}

func (x *WrapperInt32) GetVal() *wrapperspb.Int32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperUInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperUInt64) Reset() {
	*x = WrapperUInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperUInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperUInt64) ProtoMessage() {}

func (x *WrapperUInt64) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperUInt64.ProtoReflect.Descriptor instead.
func (*WrapperUInt64) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{5}
}

func (x *WrapperUInt64) GetVal() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperUInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperUInt32) Reset() {
	*x = WrapperUInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperUInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperUInt32) ProtoMessage() {}

func (x *WrapperUInt32) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperUInt32.ProtoReflect.Descriptor instead.
func (*WrapperUInt32) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{6}
}

func (x *WrapperUInt32) GetVal() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperBool) Reset() {
	*x = WrapperBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperBool) ProtoMessage() {}

func (x *WrapperBool) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperBool.ProtoReflect.Descriptor instead.
func (*WrapperBool) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{7}
}

func (x *WrapperBool) GetVal() *wrapperspb.BoolValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperString) Reset() {
	*x = WrapperString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperString) ProtoMessage() {}

func (x *WrapperString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperString.ProtoReflect.Descriptor instead.
func (*WrapperString) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{8}
}

func (x *WrapperString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperBytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.BytesValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperBytes) Reset() {
	*x = WrapperBytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperBytes) ProtoMessage() {}

func (x *WrapperBytes) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperBytes.ProtoReflect.Descriptor instead.
func (*WrapperBytes) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{9}
}

func (x *WrapperBytes) GetVal() *wrapperspb.BytesValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperRequiredString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperRequiredString) Reset() {
	*x = WrapperRequiredString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperRequiredString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredString) ProtoMessage() {}

func (x *WrapperRequiredString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperRequiredString.ProtoReflect.Descriptor instead.
func (*WrapperRequiredString) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{10}
}

func (x *WrapperRequiredString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperRequiredEmptyString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperRequiredEmptyString) Reset() {
	*x = WrapperRequiredEmptyString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperRequiredEmptyString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredEmptyString) ProtoMessage() {}

func (x *WrapperRequiredEmptyString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperRequiredEmptyString.ProtoReflect.Descriptor instead.
func (*WrapperRequiredEmptyString) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{11}
}

func (x *WrapperRequiredEmptyString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperOptionalUuidString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperOptionalUuidString) Reset() {
	*x = WrapperOptionalUuidString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperOptionalUuidString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperOptionalUuidString) ProtoMessage() {}

func (x *WrapperOptionalUuidString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperOptionalUuidString.ProtoReflect.Descriptor instead.
func (*WrapperOptionalUuidString) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{12}
}

func (x *WrapperOptionalUuidString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

type WrapperRequiredFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *WrapperRequiredFloat) Reset() {
	*x = WrapperRequiredFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperRequiredFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredFloat) ProtoMessage() {}

func (x *WrapperRequiredFloat) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperRequiredFloat.ProtoReflect.Descriptor instead.
func (*WrapperRequiredFloat) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP(), []int{13}
}

func (x *WrapperRequiredFloat) GetVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_buf_validate_conformance_cases_wkt_wrappers_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x77, 0x6b, 0x74, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x0b, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x65, 0x12,
	0x2d, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x49,
	0x0a, 0x0c, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x39,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x0a, 0x05, 0x25,
	0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x4f, 0x0a, 0x0d, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x46, 0x0a, 0x0c, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x36, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x46, 0x0a, 0x0c, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x36, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x48, 0x0a, 0x0d, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x37, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x48, 0x0a, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x37, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x44,
	0x0a, 0x0b, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x35, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x4b, 0x0a, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x42, 0x03, 0x62, 0x61, 0x72, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x46, 0x0a, 0x0c, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x7a, 0x02, 0x10, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x56, 0x0a, 0x15, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0xba,
	0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x58, 0x0a, 0x1a, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x3a, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8,
	0x01, 0x01, 0x72, 0x02, 0x0a, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x55, 0x0a, 0x19, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x75,
	0x69, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x54, 0x0a, 0x14, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x3c, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x0a, 0x05, 0x25, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0xa4, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x10, 0x57, 0x6b, 0x74, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e,
	0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02,
	0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75,
	0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescData = file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc
)

func file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDescData
}

var file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_buf_validate_conformance_cases_wkt_wrappers_proto_goTypes = []interface{}{
	(*WrapperNone)(nil),                // 0: buf.validate.conformance.cases.WrapperNone
	(*WrapperFloat)(nil),               // 1: buf.validate.conformance.cases.WrapperFloat
	(*WrapperDouble)(nil),              // 2: buf.validate.conformance.cases.WrapperDouble
	(*WrapperInt64)(nil),               // 3: buf.validate.conformance.cases.WrapperInt64
	(*WrapperInt32)(nil),               // 4: buf.validate.conformance.cases.WrapperInt32
	(*WrapperUInt64)(nil),              // 5: buf.validate.conformance.cases.WrapperUInt64
	(*WrapperUInt32)(nil),              // 6: buf.validate.conformance.cases.WrapperUInt32
	(*WrapperBool)(nil),                // 7: buf.validate.conformance.cases.WrapperBool
	(*WrapperString)(nil),              // 8: buf.validate.conformance.cases.WrapperString
	(*WrapperBytes)(nil),               // 9: buf.validate.conformance.cases.WrapperBytes
	(*WrapperRequiredString)(nil),      // 10: buf.validate.conformance.cases.WrapperRequiredString
	(*WrapperRequiredEmptyString)(nil), // 11: buf.validate.conformance.cases.WrapperRequiredEmptyString
	(*WrapperOptionalUuidString)(nil),  // 12: buf.validate.conformance.cases.WrapperOptionalUuidString
	(*WrapperRequiredFloat)(nil),       // 13: buf.validate.conformance.cases.WrapperRequiredFloat
	(*wrapperspb.Int32Value)(nil),      // 14: google.protobuf.Int32Value
	(*wrapperspb.FloatValue)(nil),      // 15: google.protobuf.FloatValue
	(*wrapperspb.DoubleValue)(nil),     // 16: google.protobuf.DoubleValue
	(*wrapperspb.Int64Value)(nil),      // 17: google.protobuf.Int64Value
	(*wrapperspb.UInt64Value)(nil),     // 18: google.protobuf.UInt64Value
	(*wrapperspb.UInt32Value)(nil),     // 19: google.protobuf.UInt32Value
	(*wrapperspb.BoolValue)(nil),       // 20: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),     // 21: google.protobuf.StringValue
	(*wrapperspb.BytesValue)(nil),      // 22: google.protobuf.BytesValue
}
var file_buf_validate_conformance_cases_wkt_wrappers_proto_depIdxs = []int32{
	14, // 0: buf.validate.conformance.cases.WrapperNone.val:type_name -> google.protobuf.Int32Value
	15, // 1: buf.validate.conformance.cases.WrapperFloat.val:type_name -> google.protobuf.FloatValue
	16, // 2: buf.validate.conformance.cases.WrapperDouble.val:type_name -> google.protobuf.DoubleValue
	17, // 3: buf.validate.conformance.cases.WrapperInt64.val:type_name -> google.protobuf.Int64Value
	14, // 4: buf.validate.conformance.cases.WrapperInt32.val:type_name -> google.protobuf.Int32Value
	18, // 5: buf.validate.conformance.cases.WrapperUInt64.val:type_name -> google.protobuf.UInt64Value
	19, // 6: buf.validate.conformance.cases.WrapperUInt32.val:type_name -> google.protobuf.UInt32Value
	20, // 7: buf.validate.conformance.cases.WrapperBool.val:type_name -> google.protobuf.BoolValue
	21, // 8: buf.validate.conformance.cases.WrapperString.val:type_name -> google.protobuf.StringValue
	22, // 9: buf.validate.conformance.cases.WrapperBytes.val:type_name -> google.protobuf.BytesValue
	21, // 10: buf.validate.conformance.cases.WrapperRequiredString.val:type_name -> google.protobuf.StringValue
	21, // 11: buf.validate.conformance.cases.WrapperRequiredEmptyString.val:type_name -> google.protobuf.StringValue
	21, // 12: buf.validate.conformance.cases.WrapperOptionalUuidString.val:type_name -> google.protobuf.StringValue
	15, // 13: buf.validate.conformance.cases.WrapperRequiredFloat.val:type_name -> google.protobuf.FloatValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_wkt_wrappers_proto_init() }
func file_buf_validate_conformance_cases_wkt_wrappers_proto_init() {
	if File_buf_validate_conformance_cases_wkt_wrappers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperNone); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperFloat); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperDouble); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperInt64); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperInt32); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperUInt64); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperUInt32); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperBool); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperString); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperBytes); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperRequiredString); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperRequiredEmptyString); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperOptionalUuidString); i {
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
		file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperRequiredFloat); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_wkt_wrappers_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_wkt_wrappers_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_wkt_wrappers_proto = out.File
	file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc = nil
	file_buf_validate_conformance_cases_wkt_wrappers_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_wrappers_proto_depIdxs = nil
}
