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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/wkt_wrappers.proto

//go:build !protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WrapperNone struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperNone) Reset() {
	*x = WrapperNone{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperNone) ProtoMessage() {}

func (x *WrapperNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperNone) GetVal() *wrapperspb.Int32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperNone) SetVal(v *wrapperspb.Int32Value) {
	x.Val = v
}

func (x *WrapperNone) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperNone) ClearVal() {
	x.Val = nil
}

type WrapperNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.Int32Value
}

func (b0 WrapperNone_builder) Build() *WrapperNone {
	m0 := &WrapperNone{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperFloat struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperFloat) Reset() {
	*x = WrapperFloat{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperFloat) ProtoMessage() {}

func (x *WrapperFloat) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperFloat) GetVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperFloat) SetVal(v *wrapperspb.FloatValue) {
	x.Val = v
}

func (x *WrapperFloat) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperFloat) ClearVal() {
	x.Val = nil
}

type WrapperFloat_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.FloatValue
}

func (b0 WrapperFloat_builder) Build() *WrapperFloat {
	m0 := &WrapperFloat{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperDouble struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.DoubleValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperDouble) Reset() {
	*x = WrapperDouble{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperDouble) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperDouble) ProtoMessage() {}

func (x *WrapperDouble) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperDouble) GetVal() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperDouble) SetVal(v *wrapperspb.DoubleValue) {
	x.Val = v
}

func (x *WrapperDouble) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperDouble) ClearVal() {
	x.Val = nil
}

type WrapperDouble_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.DoubleValue
}

func (b0 WrapperDouble_builder) Build() *WrapperDouble {
	m0 := &WrapperDouble{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperInt64 struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperInt64) Reset() {
	*x = WrapperInt64{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperInt64) ProtoMessage() {}

func (x *WrapperInt64) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperInt64) GetVal() *wrapperspb.Int64Value {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperInt64) SetVal(v *wrapperspb.Int64Value) {
	x.Val = v
}

func (x *WrapperInt64) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperInt64) ClearVal() {
	x.Val = nil
}

type WrapperInt64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.Int64Value
}

func (b0 WrapperInt64_builder) Build() *WrapperInt64 {
	m0 := &WrapperInt64{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperInt32 struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperInt32) Reset() {
	*x = WrapperInt32{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperInt32) ProtoMessage() {}

func (x *WrapperInt32) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperInt32) GetVal() *wrapperspb.Int32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperInt32) SetVal(v *wrapperspb.Int32Value) {
	x.Val = v
}

func (x *WrapperInt32) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperInt32) ClearVal() {
	x.Val = nil
}

type WrapperInt32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.Int32Value
}

func (b0 WrapperInt32_builder) Build() *WrapperInt32 {
	m0 := &WrapperInt32{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperUInt64 struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperUInt64) Reset() {
	*x = WrapperUInt64{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperUInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperUInt64) ProtoMessage() {}

func (x *WrapperUInt64) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperUInt64) GetVal() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperUInt64) SetVal(v *wrapperspb.UInt64Value) {
	x.Val = v
}

func (x *WrapperUInt64) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperUInt64) ClearVal() {
	x.Val = nil
}

type WrapperUInt64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.UInt64Value
}

func (b0 WrapperUInt64_builder) Build() *WrapperUInt64 {
	m0 := &WrapperUInt64{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperUInt32 struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperUInt32) Reset() {
	*x = WrapperUInt32{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperUInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperUInt32) ProtoMessage() {}

func (x *WrapperUInt32) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperUInt32) GetVal() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperUInt32) SetVal(v *wrapperspb.UInt32Value) {
	x.Val = v
}

func (x *WrapperUInt32) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperUInt32) ClearVal() {
	x.Val = nil
}

type WrapperUInt32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.UInt32Value
}

func (b0 WrapperUInt32_builder) Build() *WrapperUInt32 {
	m0 := &WrapperUInt32{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperBool struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.BoolValue  `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperBool) Reset() {
	*x = WrapperBool{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperBool) ProtoMessage() {}

func (x *WrapperBool) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperBool) GetVal() *wrapperspb.BoolValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperBool) SetVal(v *wrapperspb.BoolValue) {
	x.Val = v
}

func (x *WrapperBool) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperBool) ClearVal() {
	x.Val = nil
}

type WrapperBool_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.BoolValue
}

func (b0 WrapperBool_builder) Build() *WrapperBool {
	m0 := &WrapperBool{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperString struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperString) Reset() {
	*x = WrapperString{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperString) ProtoMessage() {}

func (x *WrapperString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperString) SetVal(v *wrapperspb.StringValue) {
	x.Val = v
}

func (x *WrapperString) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperString) ClearVal() {
	x.Val = nil
}

type WrapperString_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.StringValue
}

func (b0 WrapperString_builder) Build() *WrapperString {
	m0 := &WrapperString{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperBytes struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.BytesValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperBytes) Reset() {
	*x = WrapperBytes{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperBytes) ProtoMessage() {}

func (x *WrapperBytes) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperBytes) GetVal() *wrapperspb.BytesValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperBytes) SetVal(v *wrapperspb.BytesValue) {
	x.Val = v
}

func (x *WrapperBytes) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperBytes) ClearVal() {
	x.Val = nil
}

type WrapperBytes_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.BytesValue
}

func (b0 WrapperBytes_builder) Build() *WrapperBytes {
	m0 := &WrapperBytes{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperRequiredString struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperRequiredString) Reset() {
	*x = WrapperRequiredString{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperRequiredString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredString) ProtoMessage() {}

func (x *WrapperRequiredString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperRequiredString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperRequiredString) SetVal(v *wrapperspb.StringValue) {
	x.Val = v
}

func (x *WrapperRequiredString) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperRequiredString) ClearVal() {
	x.Val = nil
}

type WrapperRequiredString_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.StringValue
}

func (b0 WrapperRequiredString_builder) Build() *WrapperRequiredString {
	m0 := &WrapperRequiredString{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperRequiredEmptyString struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperRequiredEmptyString) Reset() {
	*x = WrapperRequiredEmptyString{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperRequiredEmptyString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredEmptyString) ProtoMessage() {}

func (x *WrapperRequiredEmptyString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperRequiredEmptyString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperRequiredEmptyString) SetVal(v *wrapperspb.StringValue) {
	x.Val = v
}

func (x *WrapperRequiredEmptyString) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperRequiredEmptyString) ClearVal() {
	x.Val = nil
}

type WrapperRequiredEmptyString_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.StringValue
}

func (b0 WrapperRequiredEmptyString_builder) Build() *WrapperRequiredEmptyString {
	m0 := &WrapperRequiredEmptyString{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperOptionalUuidString struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Val           *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperOptionalUuidString) Reset() {
	*x = WrapperOptionalUuidString{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperOptionalUuidString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperOptionalUuidString) ProtoMessage() {}

func (x *WrapperOptionalUuidString) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperOptionalUuidString) GetVal() *wrapperspb.StringValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperOptionalUuidString) SetVal(v *wrapperspb.StringValue) {
	x.Val = v
}

func (x *WrapperOptionalUuidString) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperOptionalUuidString) ClearVal() {
	x.Val = nil
}

type WrapperOptionalUuidString_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.StringValue
}

func (b0 WrapperOptionalUuidString_builder) Build() *WrapperOptionalUuidString {
	m0 := &WrapperOptionalUuidString{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type WrapperRequiredFloat struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *wrapperspb.FloatValue `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperRequiredFloat) Reset() {
	*x = WrapperRequiredFloat{}
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperRequiredFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperRequiredFloat) ProtoMessage() {}

func (x *WrapperRequiredFloat) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *WrapperRequiredFloat) GetVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *WrapperRequiredFloat) SetVal(v *wrapperspb.FloatValue) {
	x.Val = v
}

func (x *WrapperRequiredFloat) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *WrapperRequiredFloat) ClearVal() {
	x.Val = nil
}

type WrapperRequiredFloat_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *wrapperspb.FloatValue
}

func (b0 WrapperRequiredFloat_builder) Build() *WrapperRequiredFloat {
	m0 := &WrapperRequiredFloat{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_wkt_wrappers_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc = "" +
	"\n" +
	"1buf/validate/conformance/cases/wkt_wrappers.proto\x12\x1ebuf.validate.conformance.cases\x1a\x1bbuf/validate/validate.proto\x1a\x1egoogle/protobuf/wrappers.proto\"<\n" +
	"\vWrapperNone\x12-\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueR\x03val\"I\n" +
	"\fWrapperFloat\x129\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.FloatValueB\n" +
	"\xbaH\a\n" +
	"\x05%\x00\x00\x00\x00R\x03val\"O\n" +
	"\rWrapperDouble\x12>\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.DoubleValueB\x0e\xbaH\v\x12\t!\x00\x00\x00\x00\x00\x00\x00\x00R\x03val\"F\n" +
	"\fWrapperInt64\x126\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xbaH\x04\"\x02 \x00R\x03val\"F\n" +
	"\fWrapperInt32\x126\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.Int32ValueB\a\xbaH\x04\x1a\x02 \x00R\x03val\"H\n" +
	"\rWrapperUInt64\x127\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.UInt64ValueB\a\xbaH\x042\x02 \x00R\x03val\"H\n" +
	"\rWrapperUInt32\x127\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.UInt32ValueB\a\xbaH\x04*\x02 \x00R\x03val\"D\n" +
	"\vWrapperBool\x125\n" +
	"\x03val\x18\x01 \x01(\v2\x1a.google.protobuf.BoolValueB\a\xbaH\x04j\x02\b\x01R\x03val\"K\n" +
	"\rWrapperString\x12:\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\n" +
	"\xbaH\ar\x05B\x03barR\x03val\"F\n" +
	"\fWrapperBytes\x126\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.BytesValueB\a\xbaH\x04z\x02\x10\x03R\x03val\"V\n" +
	"\x15WrapperRequiredString\x12=\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\r\xbaH\n" +
	"\xc8\x01\x01r\x05\n" +
	"\x03barR\x03val\"X\n" +
	"\x1aWrapperRequiredEmptyString\x12:\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\n" +
	"\xbaH\a\xc8\x01\x01r\x02\n" +
	"\x00R\x03val\"X\n" +
	"\x19WrapperOptionalUuidString\x12;\n" +
	"\x03val\x18\x01 \x01(\v2\x1c.google.protobuf.StringValueB\v\xbaH\b\xc8\x01\x00r\x03\xb0\x01\x01R\x03val\"T\n" +
	"\x14WrapperRequiredFloat\x12<\n" +
	"\x03val\x18\x01 \x01(\v2\x1b.google.protobuf.FloatValueB\r\xbaH\n" +
	"\xc8\x01\x01\n" +
	"\x05%\x00\x00\x00\x00R\x03valB\x9a\x02\n" +
	"\"com.buf.validate.conformance.casesB\x10WktWrappersProtoP\x01ZFbuf.build/go/protovalidate/internal/gen/buf/validate/conformance/cases\xa2\x02\x04BVCC\xaa\x02\x1eBuf.Validate.Conformance.Cases\xca\x02\x1eBuf\\Validate\\Conformance\\Cases\xe2\x02*Buf\\Validate\\Conformance\\Cases\\GPBMetadata\xea\x02!Buf::Validate::Conformance::Casesb\x06proto3"

var file_buf_validate_conformance_cases_wkt_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_buf_validate_conformance_cases_wkt_wrappers_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc), len(file_buf_validate_conformance_cases_wkt_wrappers_proto_rawDesc)),
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
	file_buf_validate_conformance_cases_wkt_wrappers_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_wrappers_proto_depIdxs = nil
}
