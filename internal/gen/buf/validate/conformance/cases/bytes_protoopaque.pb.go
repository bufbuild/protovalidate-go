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
// source: buf/validate/conformance/cases/bytes.proto

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

type BytesNone struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesNone) Reset() {
	*x = BytesNone{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesNone) ProtoMessage() {}

func (x *BytesNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesNone) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesNone) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesNone_builder) Build() *BytesNone {
	m0 := &BytesNone{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesConst struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesConst) Reset() {
	*x = BytesConst{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesConst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesConst) ProtoMessage() {}

func (x *BytesConst) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesConst) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesConst) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesConst_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesConst_builder) Build() *BytesConst {
	m0 := &BytesConst{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesIn struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesIn) Reset() {
	*x = BytesIn{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesIn) ProtoMessage() {}

func (x *BytesIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesIn) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesIn) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesIn_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesIn_builder) Build() *BytesIn {
	m0 := &BytesIn{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesNotIn struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesNotIn) Reset() {
	*x = BytesNotIn{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesNotIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesNotIn) ProtoMessage() {}

func (x *BytesNotIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesNotIn) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesNotIn) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesNotIn_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesNotIn_builder) Build() *BytesNotIn {
	m0 := &BytesNotIn{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesLen struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesLen) Reset() {
	*x = BytesLen{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesLen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesLen) ProtoMessage() {}

func (x *BytesLen) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesLen) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesLen) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesLen_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesLen_builder) Build() *BytesLen {
	m0 := &BytesLen{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesMinLen struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesMinLen) Reset() {
	*x = BytesMinLen{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesMinLen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesMinLen) ProtoMessage() {}

func (x *BytesMinLen) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesMinLen) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesMinLen) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesMinLen_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesMinLen_builder) Build() *BytesMinLen {
	m0 := &BytesMinLen{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesMaxLen struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesMaxLen) Reset() {
	*x = BytesMaxLen{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesMaxLen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesMaxLen) ProtoMessage() {}

func (x *BytesMaxLen) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesMaxLen) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesMaxLen) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesMaxLen_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesMaxLen_builder) Build() *BytesMaxLen {
	m0 := &BytesMaxLen{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesMinMaxLen struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesMinMaxLen) Reset() {
	*x = BytesMinMaxLen{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesMinMaxLen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesMinMaxLen) ProtoMessage() {}

func (x *BytesMinMaxLen) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesMinMaxLen) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesMinMaxLen) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesMinMaxLen_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesMinMaxLen_builder) Build() *BytesMinMaxLen {
	m0 := &BytesMinMaxLen{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesEqualMinMaxLen struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesEqualMinMaxLen) Reset() {
	*x = BytesEqualMinMaxLen{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesEqualMinMaxLen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesEqualMinMaxLen) ProtoMessage() {}

func (x *BytesEqualMinMaxLen) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesEqualMinMaxLen) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesEqualMinMaxLen) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesEqualMinMaxLen_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesEqualMinMaxLen_builder) Build() *BytesEqualMinMaxLen {
	m0 := &BytesEqualMinMaxLen{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesPattern struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesPattern) Reset() {
	*x = BytesPattern{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesPattern) ProtoMessage() {}

func (x *BytesPattern) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesPattern) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesPattern) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesPattern_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesPattern_builder) Build() *BytesPattern {
	m0 := &BytesPattern{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesPrefix struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesPrefix) Reset() {
	*x = BytesPrefix{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesPrefix) ProtoMessage() {}

func (x *BytesPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesPrefix) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesPrefix) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesPrefix_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesPrefix_builder) Build() *BytesPrefix {
	m0 := &BytesPrefix{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesContains struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesContains) Reset() {
	*x = BytesContains{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesContains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesContains) ProtoMessage() {}

func (x *BytesContains) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesContains) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesContains) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesContains_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesContains_builder) Build() *BytesContains {
	m0 := &BytesContains{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesSuffix struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesSuffix) Reset() {
	*x = BytesSuffix{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesSuffix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesSuffix) ProtoMessage() {}

func (x *BytesSuffix) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesSuffix) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesSuffix) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesSuffix_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesSuffix_builder) Build() *BytesSuffix {
	m0 := &BytesSuffix{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesIP struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesIP) Reset() {
	*x = BytesIP{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesIP) ProtoMessage() {}

func (x *BytesIP) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesIP) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesIP) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesIP_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesIP_builder) Build() *BytesIP {
	m0 := &BytesIP{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesNotIP struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesNotIP) Reset() {
	*x = BytesNotIP{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesNotIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesNotIP) ProtoMessage() {}

func (x *BytesNotIP) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesNotIP) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesNotIP) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesNotIP_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesNotIP_builder) Build() *BytesNotIP {
	m0 := &BytesNotIP{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesIPv4 struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesIPv4) Reset() {
	*x = BytesIPv4{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesIPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesIPv4) ProtoMessage() {}

func (x *BytesIPv4) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesIPv4) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesIPv4) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesIPv4_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesIPv4_builder) Build() *BytesIPv4 {
	m0 := &BytesIPv4{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesNotIPv4 struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesNotIPv4) Reset() {
	*x = BytesNotIPv4{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesNotIPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesNotIPv4) ProtoMessage() {}

func (x *BytesNotIPv4) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesNotIPv4) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesNotIPv4) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesNotIPv4_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesNotIPv4_builder) Build() *BytesNotIPv4 {
	m0 := &BytesNotIPv4{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesIPv6 struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesIPv6) Reset() {
	*x = BytesIPv6{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesIPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesIPv6) ProtoMessage() {}

func (x *BytesIPv6) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesIPv6) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesIPv6) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesIPv6_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesIPv6_builder) Build() *BytesIPv6 {
	m0 := &BytesIPv6{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesNotIPv6 struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesNotIPv6) Reset() {
	*x = BytesNotIPv6{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesNotIPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesNotIPv6) ProtoMessage() {}

func (x *BytesNotIPv6) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesNotIPv6) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesNotIPv6) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesNotIPv6_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesNotIPv6_builder) Build() *BytesNotIPv6 {
	m0 := &BytesNotIPv6{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesIPv6Ignore struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesIPv6Ignore) Reset() {
	*x = BytesIPv6Ignore{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[19]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesIPv6Ignore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesIPv6Ignore) ProtoMessage() {}

func (x *BytesIPv6Ignore) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[19]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesIPv6Ignore) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesIPv6Ignore) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesIPv6Ignore_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesIPv6Ignore_builder) Build() *BytesIPv6Ignore {
	m0 := &BytesIPv6Ignore{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type BytesExample struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val []byte                 `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BytesExample) Reset() {
	*x = BytesExample{}
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[20]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BytesExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesExample) ProtoMessage() {}

func (x *BytesExample) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bytes_proto_msgTypes[20]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BytesExample) GetVal() []byte {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *BytesExample) SetVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Val = v
}

type BytesExample_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val []byte
}

func (b0 BytesExample_builder) Build() *BytesExample {
	m0 := &BytesExample{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_bytes_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_bytes_proto_rawDesc = "" +
	"\n" +
	"*buf/validate/conformance/cases/bytes.proto\x12\x1ebuf.validate.conformance.cases\x1a\x1bbuf/validate/validate.proto\"\x1d\n" +
	"\tBytesNone\x12\x10\n" +
	"\x03val\x18\x01 \x01(\fR\x03val\"*\n" +
	"\n" +
	"BytesConst\x12\x1c\n" +
	"\x03val\x18\x01 \x01(\fB\n" +
	"\xbaH\az\x05\n" +
	"\x03fooR\x03val\",\n" +
	"\aBytesIn\x12!\n" +
	"\x03val\x18\x01 \x01(\fB\x0f\xbaH\fz\n" +
	"B\x03barB\x03bazR\x03val\"1\n" +
	"\n" +
	"BytesNotIn\x12#\n" +
	"\x03val\x18\x01 \x01(\fB\x11\xbaH\x0ez\fJ\x04fizzJ\x04buzzR\x03val\"%\n" +
	"\bBytesLen\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02h\x03R\x03val\"(\n" +
	"\vBytesMinLen\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02\x10\x03R\x03val\"(\n" +
	"\vBytesMaxLen\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02\x18\x05R\x03val\"-\n" +
	"\x0eBytesMinMaxLen\x12\x1b\n" +
	"\x03val\x18\x01 \x01(\fB\t\xbaH\x06z\x04\x10\x03\x18\x05R\x03val\"2\n" +
	"\x13BytesEqualMinMaxLen\x12\x1b\n" +
	"\x03val\x18\x01 \x01(\fB\t\xbaH\x06z\x04\x10\x05\x18\x05R\x03val\"7\n" +
	"\fBytesPattern\x12'\n" +
	"\x03val\x18\x01 \x01(\fB\x15\xbaH\x12z\x10\"\x0e^[\\x00-\\x7F]+$R\x03val\")\n" +
	"\vBytesPrefix\x12\x1a\n" +
	"\x03val\x18\x01 \x01(\fB\b\xbaH\x05z\x03*\x01\x99R\x03val\"-\n" +
	"\rBytesContains\x12\x1c\n" +
	"\x03val\x18\x01 \x01(\fB\n" +
	"\xbaH\az\x05:\x03barR\x03val\",\n" +
	"\vBytesSuffix\x12\x1d\n" +
	"\x03val\x18\x01 \x01(\fB\v\xbaH\bz\x062\x04buzzR\x03val\"$\n" +
	"\aBytesIP\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02P\x01R\x03val\"'\n" +
	"\n" +
	"BytesNotIP\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02P\x00R\x03val\"&\n" +
	"\tBytesIPv4\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02X\x01R\x03val\")\n" +
	"\fBytesNotIPv4\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02X\x00R\x03val\"&\n" +
	"\tBytesIPv6\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02`\x01R\x03val\")\n" +
	"\fBytesNotIPv6\x12\x19\n" +
	"\x03val\x18\x01 \x01(\fB\a\xbaH\x04z\x02`\x00R\x03val\"/\n" +
	"\x0fBytesIPv6Ignore\x12\x1c\n" +
	"\x03val\x18\x01 \x01(\fB\n" +
	"\xbaH\a\xd8\x01\x01z\x02`\x01R\x03val\"*\n" +
	"\fBytesExample\x12\x1a\n" +
	"\x03val\x18\x01 \x01(\fB\b\xbaH\x05z\x03r\x01\x99R\x03valB\x94\x02\n" +
	"\"com.buf.validate.conformance.casesB\n" +
	"BytesProtoP\x01ZFbuf.build/go/protovalidate/internal/gen/buf/validate/conformance/cases\xa2\x02\x04BVCC\xaa\x02\x1eBuf.Validate.Conformance.Cases\xca\x02\x1eBuf\\Validate\\Conformance\\Cases\xe2\x02*Buf\\Validate\\Conformance\\Cases\\GPBMetadata\xea\x02!Buf::Validate::Conformance::Casesb\x06proto3"

var file_buf_validate_conformance_cases_bytes_proto_msgTypes = make([]protoimpl.MessageInfo, 21)
var file_buf_validate_conformance_cases_bytes_proto_goTypes = []any{
	(*BytesNone)(nil),           // 0: buf.validate.conformance.cases.BytesNone
	(*BytesConst)(nil),          // 1: buf.validate.conformance.cases.BytesConst
	(*BytesIn)(nil),             // 2: buf.validate.conformance.cases.BytesIn
	(*BytesNotIn)(nil),          // 3: buf.validate.conformance.cases.BytesNotIn
	(*BytesLen)(nil),            // 4: buf.validate.conformance.cases.BytesLen
	(*BytesMinLen)(nil),         // 5: buf.validate.conformance.cases.BytesMinLen
	(*BytesMaxLen)(nil),         // 6: buf.validate.conformance.cases.BytesMaxLen
	(*BytesMinMaxLen)(nil),      // 7: buf.validate.conformance.cases.BytesMinMaxLen
	(*BytesEqualMinMaxLen)(nil), // 8: buf.validate.conformance.cases.BytesEqualMinMaxLen
	(*BytesPattern)(nil),        // 9: buf.validate.conformance.cases.BytesPattern
	(*BytesPrefix)(nil),         // 10: buf.validate.conformance.cases.BytesPrefix
	(*BytesContains)(nil),       // 11: buf.validate.conformance.cases.BytesContains
	(*BytesSuffix)(nil),         // 12: buf.validate.conformance.cases.BytesSuffix
	(*BytesIP)(nil),             // 13: buf.validate.conformance.cases.BytesIP
	(*BytesNotIP)(nil),          // 14: buf.validate.conformance.cases.BytesNotIP
	(*BytesIPv4)(nil),           // 15: buf.validate.conformance.cases.BytesIPv4
	(*BytesNotIPv4)(nil),        // 16: buf.validate.conformance.cases.BytesNotIPv4
	(*BytesIPv6)(nil),           // 17: buf.validate.conformance.cases.BytesIPv6
	(*BytesNotIPv6)(nil),        // 18: buf.validate.conformance.cases.BytesNotIPv6
	(*BytesIPv6Ignore)(nil),     // 19: buf.validate.conformance.cases.BytesIPv6Ignore
	(*BytesExample)(nil),        // 20: buf.validate.conformance.cases.BytesExample
}
var file_buf_validate_conformance_cases_bytes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_bytes_proto_init() }
func file_buf_validate_conformance_cases_bytes_proto_init() {
	if File_buf_validate_conformance_cases_bytes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_bytes_proto_rawDesc), len(file_buf_validate_conformance_cases_bytes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   21,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_bytes_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_bytes_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_bytes_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_bytes_proto = out.File
	file_buf_validate_conformance_cases_bytes_proto_goTypes = nil
	file_buf_validate_conformance_cases_bytes_proto_depIdxs = nil
}
