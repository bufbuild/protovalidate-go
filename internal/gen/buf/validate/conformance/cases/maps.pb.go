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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/maps.proto

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

type MapNone struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[uint32]bool        `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapNone) Reset() {
	*x = MapNone{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNone) ProtoMessage() {}

func (x *MapNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapNone) GetVal() map[uint32]bool {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapNone) SetVal(v map[uint32]bool) {
	x.Val = v
}

type MapNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[uint32]bool
}

func (b0 MapNone_builder) Build() *MapNone {
	m0 := &MapNone{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapMin struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[int32]float32      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapMin) Reset() {
	*x = MapMin{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapMin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMin) ProtoMessage() {}

func (x *MapMin) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapMin) GetVal() map[int32]float32 {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapMin) SetVal(v map[int32]float32) {
	x.Val = v
}

type MapMin_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[int32]float32
}

func (b0 MapMin_builder) Build() *MapMin {
	m0 := &MapMin{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapMax struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[int64]float64      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapMax) Reset() {
	*x = MapMax{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapMax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMax) ProtoMessage() {}

func (x *MapMax) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapMax) GetVal() map[int64]float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapMax) SetVal(v map[int64]float64) {
	x.Val = v
}

type MapMax_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[int64]float64
}

func (b0 MapMax_builder) Build() *MapMax {
	m0 := &MapMax{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapMinMax struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[string]bool        `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapMinMax) Reset() {
	*x = MapMinMax{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapMinMax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMinMax) ProtoMessage() {}

func (x *MapMinMax) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapMinMax) GetVal() map[string]bool {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapMinMax) SetVal(v map[string]bool) {
	x.Val = v
}

type MapMinMax_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]bool
}

func (b0 MapMinMax_builder) Build() *MapMinMax {
	m0 := &MapMinMax{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapExact struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[uint64]string      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapExact) Reset() {
	*x = MapExact{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapExact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapExact) ProtoMessage() {}

func (x *MapExact) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapExact) GetVal() map[uint64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapExact) SetVal(v map[uint64]string) {
	x.Val = v
}

type MapExact_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[uint64]string
}

func (b0 MapExact_builder) Build() *MapExact {
	m0 := &MapExact{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapKeys struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[int64]string       `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapKeys) Reset() {
	*x = MapKeys{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapKeys) ProtoMessage() {}

func (x *MapKeys) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapKeys) GetVal() map[int64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapKeys) SetVal(v map[int64]string) {
	x.Val = v
}

type MapKeys_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[int64]string
}

func (b0 MapKeys_builder) Build() *MapKeys {
	m0 := &MapKeys{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapValues struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[string]string      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapValues) Reset() {
	*x = MapValues{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValues) ProtoMessage() {}

func (x *MapValues) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapValues) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapValues) SetVal(v map[string]string) {
	x.Val = v
}

type MapValues_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]string
}

func (b0 MapValues_builder) Build() *MapValues {
	m0 := &MapValues{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapKeysPattern struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[string]string      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapKeysPattern) Reset() {
	*x = MapKeysPattern{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapKeysPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapKeysPattern) ProtoMessage() {}

func (x *MapKeysPattern) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapKeysPattern) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapKeysPattern) SetVal(v map[string]string) {
	x.Val = v
}

type MapKeysPattern_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]string
}

func (b0 MapKeysPattern_builder) Build() *MapKeysPattern {
	m0 := &MapKeysPattern{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapValuesPattern struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[string]string      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapValuesPattern) Reset() {
	*x = MapValuesPattern{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapValuesPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValuesPattern) ProtoMessage() {}

func (x *MapValuesPattern) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapValuesPattern) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapValuesPattern) SetVal(v map[string]string) {
	x.Val = v
}

type MapValuesPattern_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[string]string
}

func (b0 MapValuesPattern_builder) Build() *MapValuesPattern {
	m0 := &MapValuesPattern{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapRecursive struct {
	state         protoimpl.MessageState       `protogen:"hybrid.v1"`
	Val           map[uint32]*MapRecursive_Msg `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapRecursive) Reset() {
	*x = MapRecursive{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapRecursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRecursive) ProtoMessage() {}

func (x *MapRecursive) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapRecursive) GetVal() map[uint32]*MapRecursive_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapRecursive) SetVal(v map[uint32]*MapRecursive_Msg) {
	x.Val = v
}

type MapRecursive_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[uint32]*MapRecursive_Msg
}

func (b0 MapRecursive_builder) Build() *MapRecursive {
	m0 := &MapRecursive{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MapExactIgnore struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           map[uint64]string      `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapExactIgnore) Reset() {
	*x = MapExactIgnore{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapExactIgnore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapExactIgnore) ProtoMessage() {}

func (x *MapExactIgnore) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapExactIgnore) GetVal() map[uint64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *MapExactIgnore) SetVal(v map[uint64]string) {
	x.Val = v
}

type MapExactIgnore_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val map[uint64]string
}

func (b0 MapExactIgnore_builder) Build() *MapExactIgnore {
	m0 := &MapExactIgnore{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type MultipleMaps struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	First         map[uint32]string      `protobuf:"bytes,1,rep,name=first,proto3" json:"first,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Second        map[int32]bool         `protobuf:"bytes,2,rep,name=second,proto3" json:"second,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Third         map[int32]bool         `protobuf:"bytes,3,rep,name=third,proto3" json:"third,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MultipleMaps) Reset() {
	*x = MultipleMaps{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultipleMaps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleMaps) ProtoMessage() {}

func (x *MultipleMaps) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MultipleMaps) GetFirst() map[uint32]string {
	if x != nil {
		return x.First
	}
	return nil
}

func (x *MultipleMaps) GetSecond() map[int32]bool {
	if x != nil {
		return x.Second
	}
	return nil
}

func (x *MultipleMaps) GetThird() map[int32]bool {
	if x != nil {
		return x.Third
	}
	return nil
}

func (x *MultipleMaps) SetFirst(v map[uint32]string) {
	x.First = v
}

func (x *MultipleMaps) SetSecond(v map[int32]bool) {
	x.Second = v
}

func (x *MultipleMaps) SetThird(v map[int32]bool) {
	x.Third = v
}

type MultipleMaps_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	First  map[uint32]string
	Second map[int32]bool
	Third  map[int32]bool
}

func (b0 MultipleMaps_builder) Build() *MultipleMaps {
	m0 := &MultipleMaps{}
	b, x := &b0, m0
	_, _ = b, x
	x.First = b.First
	x.Second = b.Second
	x.Third = b.Third
	return m0
}

type MapRecursive_Msg struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           string                 `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapRecursive_Msg) Reset() {
	*x = MapRecursive_Msg{}
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[22]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapRecursive_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRecursive_Msg) ProtoMessage() {}

func (x *MapRecursive_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_maps_proto_msgTypes[22]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapRecursive_Msg) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *MapRecursive_Msg) SetVal(v string) {
	x.Val = v
}

type MapRecursive_Msg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val string
}

func (b0 MapRecursive_Msg_builder) Build() *MapRecursive_Msg {
	m0 := &MapRecursive_Msg{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_maps_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_maps_proto_rawDesc = "" +
	"\n" +
	")buf/validate/conformance/cases/maps.proto\x12\x1ebuf.validate.conformance.cases\x1a\x1bbuf/validate/validate.proto\"\x85\x01\n" +
	"\aMapNone\x12B\n" +
	"\x03val\x18\x01 \x03(\v20.buf.validate.conformance.cases.MapNone.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\rR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01\"\x8d\x01\n" +
	"\x06MapMin\x12K\n" +
	"\x03val\x18\x01 \x03(\v2/.buf.validate.conformance.cases.MapMin.ValEntryB\b\xbaH\x05\x9a\x01\x02\b\x02R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x02R\x05value:\x028\x01\"\x8d\x01\n" +
	"\x06MapMax\x12K\n" +
	"\x03val\x18\x01 \x03(\v2/.buf.validate.conformance.cases.MapMax.ValEntryB\b\xbaH\x05\x9a\x01\x02\x10\x03R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x03R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01\"\x95\x01\n" +
	"\tMapMinMax\x12P\n" +
	"\x03val\x18\x01 \x03(\v22.buf.validate.conformance.cases.MapMinMax.ValEntryB\n" +
	"\xbaH\a\x9a\x01\x04\b\x02\x10\x04R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01\"\x93\x01\n" +
	"\bMapExact\x12O\n" +
	"\x03val\x18\x01 \x03(\v21.buf.validate.conformance.cases.MapExact.ValEntryB\n" +
	"\xbaH\a\x9a\x01\x04\b\x03\x10\x03R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x04R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x93\x01\n" +
	"\aMapKeys\x12P\n" +
	"\x03val\x18\x01 \x03(\v20.buf.validate.conformance.cases.MapKeys.ValEntryB\f\xbaH\t\x9a\x01\x06\"\x04B\x02\x10\x00R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x12R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x97\x01\n" +
	"\tMapValues\x12R\n" +
	"\x03val\x18\x01 \x03(\v22.buf.validate.conformance.cases.MapValues.ValEntryB\f\xbaH\t\x9a\x01\x06*\x04r\x02\x10\x03R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xb0\x01\n" +
	"\x0eMapKeysPattern\x12f\n" +
	"\x03val\x18\x01 \x03(\v27.buf.validate.conformance.cases.MapKeysPattern.ValEntryB\x1b\xbaH\x18\x9a\x01\x15\"\x13r\x112\x0f(?i)^[a-z0-9]+$R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xb4\x01\n" +
	"\x10MapValuesPattern\x12h\n" +
	"\x03val\x18\x01 \x03(\v29.buf.validate.conformance.cases.MapValuesPattern.ValEntryB\x1b\xbaH\x18\x9a\x01\x15*\x13r\x112\x0f(?i)^[a-z0-9]+$R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xe3\x01\n" +
	"\fMapRecursive\x12G\n" +
	"\x03val\x18\x01 \x03(\v25.buf.validate.conformance.cases.MapRecursive.ValEntryR\x03val\x1ah\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\rR\x03key\x12F\n" +
	"\x05value\x18\x02 \x01(\v20.buf.validate.conformance.cases.MapRecursive.MsgR\x05value:\x028\x01\x1a \n" +
	"\x03Msg\x12\x19\n" +
	"\x03val\x18\x01 \x01(\tB\a\xbaH\x04r\x02\x10\x03R\x03val\"\xa2\x01\n" +
	"\x0eMapExactIgnore\x12X\n" +
	"\x03val\x18\x01 \x03(\v27.buf.validate.conformance.cases.MapExactIgnore.ValEntryB\r\xbaH\n" +
	"\xd8\x01\x01\x9a\x01\x04\b\x03\x10\x03R\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x04R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xd7\x03\n" +
	"\fMultipleMaps\x12[\n" +
	"\x05first\x18\x01 \x03(\v27.buf.validate.conformance.cases.MultipleMaps.FirstEntryB\f\xbaH\t\x9a\x01\x06\"\x04*\x02 \x00R\x05first\x12^\n" +
	"\x06second\x18\x02 \x03(\v28.buf.validate.conformance.cases.MultipleMaps.SecondEntryB\f\xbaH\t\x9a\x01\x06\"\x04\x1a\x02\x10\x00R\x06second\x12[\n" +
	"\x05third\x18\x03 \x03(\v27.buf.validate.conformance.cases.MultipleMaps.ThirdEntryB\f\xbaH\t\x9a\x01\x06\"\x04\x1a\x02 \x00R\x05third\x1a8\n" +
	"\n" +
	"FirstEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\rR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a9\n" +
	"\vSecondEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01\x1a8\n" +
	"\n" +
	"ThirdEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01B\x9d\x02\n" +
	"\"com.buf.validate.conformance.casesB\tMapsProtoP\x01ZPgithub.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases\xa2\x02\x04BVCC\xaa\x02\x1eBuf.Validate.Conformance.Cases\xca\x02\x1eBuf\\Validate\\Conformance\\Cases\xe2\x02*Buf\\Validate\\Conformance\\Cases\\GPBMetadata\xea\x02!Buf::Validate::Conformance::Casesb\x06proto3"

var file_buf_validate_conformance_cases_maps_proto_msgTypes = make([]protoimpl.MessageInfo, 27)
var file_buf_validate_conformance_cases_maps_proto_goTypes = []any{
	(*MapNone)(nil),          // 0: buf.validate.conformance.cases.MapNone
	(*MapMin)(nil),           // 1: buf.validate.conformance.cases.MapMin
	(*MapMax)(nil),           // 2: buf.validate.conformance.cases.MapMax
	(*MapMinMax)(nil),        // 3: buf.validate.conformance.cases.MapMinMax
	(*MapExact)(nil),         // 4: buf.validate.conformance.cases.MapExact
	(*MapKeys)(nil),          // 5: buf.validate.conformance.cases.MapKeys
	(*MapValues)(nil),        // 6: buf.validate.conformance.cases.MapValues
	(*MapKeysPattern)(nil),   // 7: buf.validate.conformance.cases.MapKeysPattern
	(*MapValuesPattern)(nil), // 8: buf.validate.conformance.cases.MapValuesPattern
	(*MapRecursive)(nil),     // 9: buf.validate.conformance.cases.MapRecursive
	(*MapExactIgnore)(nil),   // 10: buf.validate.conformance.cases.MapExactIgnore
	(*MultipleMaps)(nil),     // 11: buf.validate.conformance.cases.MultipleMaps
	nil,                      // 12: buf.validate.conformance.cases.MapNone.ValEntry
	nil,                      // 13: buf.validate.conformance.cases.MapMin.ValEntry
	nil,                      // 14: buf.validate.conformance.cases.MapMax.ValEntry
	nil,                      // 15: buf.validate.conformance.cases.MapMinMax.ValEntry
	nil,                      // 16: buf.validate.conformance.cases.MapExact.ValEntry
	nil,                      // 17: buf.validate.conformance.cases.MapKeys.ValEntry
	nil,                      // 18: buf.validate.conformance.cases.MapValues.ValEntry
	nil,                      // 19: buf.validate.conformance.cases.MapKeysPattern.ValEntry
	nil,                      // 20: buf.validate.conformance.cases.MapValuesPattern.ValEntry
	nil,                      // 21: buf.validate.conformance.cases.MapRecursive.ValEntry
	(*MapRecursive_Msg)(nil), // 22: buf.validate.conformance.cases.MapRecursive.Msg
	nil,                      // 23: buf.validate.conformance.cases.MapExactIgnore.ValEntry
	nil,                      // 24: buf.validate.conformance.cases.MultipleMaps.FirstEntry
	nil,                      // 25: buf.validate.conformance.cases.MultipleMaps.SecondEntry
	nil,                      // 26: buf.validate.conformance.cases.MultipleMaps.ThirdEntry
}
var file_buf_validate_conformance_cases_maps_proto_depIdxs = []int32{
	12, // 0: buf.validate.conformance.cases.MapNone.val:type_name -> buf.validate.conformance.cases.MapNone.ValEntry
	13, // 1: buf.validate.conformance.cases.MapMin.val:type_name -> buf.validate.conformance.cases.MapMin.ValEntry
	14, // 2: buf.validate.conformance.cases.MapMax.val:type_name -> buf.validate.conformance.cases.MapMax.ValEntry
	15, // 3: buf.validate.conformance.cases.MapMinMax.val:type_name -> buf.validate.conformance.cases.MapMinMax.ValEntry
	16, // 4: buf.validate.conformance.cases.MapExact.val:type_name -> buf.validate.conformance.cases.MapExact.ValEntry
	17, // 5: buf.validate.conformance.cases.MapKeys.val:type_name -> buf.validate.conformance.cases.MapKeys.ValEntry
	18, // 6: buf.validate.conformance.cases.MapValues.val:type_name -> buf.validate.conformance.cases.MapValues.ValEntry
	19, // 7: buf.validate.conformance.cases.MapKeysPattern.val:type_name -> buf.validate.conformance.cases.MapKeysPattern.ValEntry
	20, // 8: buf.validate.conformance.cases.MapValuesPattern.val:type_name -> buf.validate.conformance.cases.MapValuesPattern.ValEntry
	21, // 9: buf.validate.conformance.cases.MapRecursive.val:type_name -> buf.validate.conformance.cases.MapRecursive.ValEntry
	23, // 10: buf.validate.conformance.cases.MapExactIgnore.val:type_name -> buf.validate.conformance.cases.MapExactIgnore.ValEntry
	24, // 11: buf.validate.conformance.cases.MultipleMaps.first:type_name -> buf.validate.conformance.cases.MultipleMaps.FirstEntry
	25, // 12: buf.validate.conformance.cases.MultipleMaps.second:type_name -> buf.validate.conformance.cases.MultipleMaps.SecondEntry
	26, // 13: buf.validate.conformance.cases.MultipleMaps.third:type_name -> buf.validate.conformance.cases.MultipleMaps.ThirdEntry
	22, // 14: buf.validate.conformance.cases.MapRecursive.ValEntry.value:type_name -> buf.validate.conformance.cases.MapRecursive.Msg
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_maps_proto_init() }
func file_buf_validate_conformance_cases_maps_proto_init() {
	if File_buf_validate_conformance_cases_maps_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_maps_proto_rawDesc), len(file_buf_validate_conformance_cases_maps_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   27,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_maps_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_maps_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_maps_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_maps_proto = out.File
	file_buf_validate_conformance_cases_maps_proto_goTypes = nil
	file_buf_validate_conformance_cases_maps_proto_depIdxs = nil
}
