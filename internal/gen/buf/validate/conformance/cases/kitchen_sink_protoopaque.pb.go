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
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/kitchen_sink.proto

//go:build protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComplexTestEnum int32

const (
	ComplexTestEnum_COMPLEX_TEST_ENUM_UNSPECIFIED ComplexTestEnum = 0
	ComplexTestEnum_COMPLEX_TEST_ENUM_ONE         ComplexTestEnum = 1
	ComplexTestEnum_COMPLEX_TEST_ENUM_TWO         ComplexTestEnum = 2
)

// Enum value maps for ComplexTestEnum.
var (
	ComplexTestEnum_name = map[int32]string{
		0: "COMPLEX_TEST_ENUM_UNSPECIFIED",
		1: "COMPLEX_TEST_ENUM_ONE",
		2: "COMPLEX_TEST_ENUM_TWO",
	}
	ComplexTestEnum_value = map[string]int32{
		"COMPLEX_TEST_ENUM_UNSPECIFIED": 0,
		"COMPLEX_TEST_ENUM_ONE":         1,
		"COMPLEX_TEST_ENUM_TWO":         2,
	}
)

func (x ComplexTestEnum) Enum() *ComplexTestEnum {
	p := new(ComplexTestEnum)
	*p = x
	return p
}

func (x ComplexTestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComplexTestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes[0].Descriptor()
}

func (ComplexTestEnum) Type() protoreflect.EnumType {
	return &file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes[0]
}

func (x ComplexTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type ComplexTestMsg struct {
	state                 protoimpl.MessageState    `protogen:"opaque.v1"`
	xxx_hidden_Const      string                    `protobuf:"bytes,1,opt,name=const,proto3" json:"const,omitempty"`
	xxx_hidden_Nested     *ComplexTestMsg           `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
	xxx_hidden_IntConst   int32                     `protobuf:"varint,3,opt,name=int_const,json=intConst,proto3" json:"int_const,omitempty"`
	xxx_hidden_BoolConst  bool                      `protobuf:"varint,4,opt,name=bool_const,json=boolConst,proto3" json:"bool_const,omitempty"`
	xxx_hidden_FloatVal   *wrapperspb.FloatValue    `protobuf:"bytes,5,opt,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	xxx_hidden_DurVal     *durationpb.Duration      `protobuf:"bytes,6,opt,name=dur_val,json=durVal,proto3" json:"dur_val,omitempty"`
	xxx_hidden_TsVal      *timestamppb.Timestamp    `protobuf:"bytes,7,opt,name=ts_val,json=tsVal,proto3" json:"ts_val,omitempty"`
	xxx_hidden_Another    *ComplexTestMsg           `protobuf:"bytes,8,opt,name=another,proto3" json:"another,omitempty"`
	xxx_hidden_FloatConst float32                   `protobuf:"fixed32,9,opt,name=float_const,json=floatConst,proto3" json:"float_const,omitempty"`
	xxx_hidden_DoubleIn   float64                   `protobuf:"fixed64,10,opt,name=double_in,json=doubleIn,proto3" json:"double_in,omitempty"`
	xxx_hidden_EnumConst  ComplexTestEnum           `protobuf:"varint,11,opt,name=enum_const,json=enumConst,proto3,enum=buf.validate.conformance.cases.ComplexTestEnum" json:"enum_const,omitempty"`
	xxx_hidden_AnyVal     *anypb.Any                `protobuf:"bytes,12,opt,name=any_val,json=anyVal,proto3" json:"any_val,omitempty"`
	xxx_hidden_RepTsVal   *[]*timestamppb.Timestamp `protobuf:"bytes,13,rep,name=rep_ts_val,json=repTsVal,proto3" json:"rep_ts_val,omitempty"`
	xxx_hidden_MapVal     map[int32]string          `protobuf:"bytes,14,rep,name=map_val,json=mapVal,proto3" json:"map_val,omitempty" protobuf_key:"zigzag32,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_BytesVal   []byte                    `protobuf:"bytes,15,opt,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
	xxx_hidden_O          isComplexTestMsg_O        `protobuf_oneof:"o"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ComplexTestMsg) Reset() {
	*x = ComplexTestMsg{}
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplexTestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexTestMsg) ProtoMessage() {}

func (x *ComplexTestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ComplexTestMsg) GetConst() string {
	if x != nil {
		return x.xxx_hidden_Const
	}
	return ""
}

func (x *ComplexTestMsg) GetNested() *ComplexTestMsg {
	if x != nil {
		return x.xxx_hidden_Nested
	}
	return nil
}

func (x *ComplexTestMsg) GetIntConst() int32 {
	if x != nil {
		return x.xxx_hidden_IntConst
	}
	return 0
}

func (x *ComplexTestMsg) GetBoolConst() bool {
	if x != nil {
		return x.xxx_hidden_BoolConst
	}
	return false
}

func (x *ComplexTestMsg) GetFloatVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.xxx_hidden_FloatVal
	}
	return nil
}

func (x *ComplexTestMsg) GetDurVal() *durationpb.Duration {
	if x != nil {
		return x.xxx_hidden_DurVal
	}
	return nil
}

func (x *ComplexTestMsg) GetTsVal() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_TsVal
	}
	return nil
}

func (x *ComplexTestMsg) GetAnother() *ComplexTestMsg {
	if x != nil {
		return x.xxx_hidden_Another
	}
	return nil
}

func (x *ComplexTestMsg) GetFloatConst() float32 {
	if x != nil {
		return x.xxx_hidden_FloatConst
	}
	return 0
}

func (x *ComplexTestMsg) GetDoubleIn() float64 {
	if x != nil {
		return x.xxx_hidden_DoubleIn
	}
	return 0
}

func (x *ComplexTestMsg) GetEnumConst() ComplexTestEnum {
	if x != nil {
		return x.xxx_hidden_EnumConst
	}
	return ComplexTestEnum_COMPLEX_TEST_ENUM_UNSPECIFIED
}

func (x *ComplexTestMsg) GetAnyVal() *anypb.Any {
	if x != nil {
		return x.xxx_hidden_AnyVal
	}
	return nil
}

func (x *ComplexTestMsg) GetRepTsVal() []*timestamppb.Timestamp {
	if x != nil {
		if x.xxx_hidden_RepTsVal != nil {
			return *x.xxx_hidden_RepTsVal
		}
	}
	return nil
}

func (x *ComplexTestMsg) GetMapVal() map[int32]string {
	if x != nil {
		return x.xxx_hidden_MapVal
	}
	return nil
}

func (x *ComplexTestMsg) GetBytesVal() []byte {
	if x != nil {
		return x.xxx_hidden_BytesVal
	}
	return nil
}

func (x *ComplexTestMsg) GetX() string {
	if x != nil {
		if x, ok := x.xxx_hidden_O.(*complexTestMsg_X); ok {
			return x.X
		}
	}
	return ""
}

func (x *ComplexTestMsg) GetY() int32 {
	if x != nil {
		if x, ok := x.xxx_hidden_O.(*complexTestMsg_Y); ok {
			return x.Y
		}
	}
	return 0
}

func (x *ComplexTestMsg) SetConst(v string) {
	x.xxx_hidden_Const = v
}

func (x *ComplexTestMsg) SetNested(v *ComplexTestMsg) {
	x.xxx_hidden_Nested = v
}

func (x *ComplexTestMsg) SetIntConst(v int32) {
	x.xxx_hidden_IntConst = v
}

func (x *ComplexTestMsg) SetBoolConst(v bool) {
	x.xxx_hidden_BoolConst = v
}

func (x *ComplexTestMsg) SetFloatVal(v *wrapperspb.FloatValue) {
	x.xxx_hidden_FloatVal = v
}

func (x *ComplexTestMsg) SetDurVal(v *durationpb.Duration) {
	x.xxx_hidden_DurVal = v
}

func (x *ComplexTestMsg) SetTsVal(v *timestamppb.Timestamp) {
	x.xxx_hidden_TsVal = v
}

func (x *ComplexTestMsg) SetAnother(v *ComplexTestMsg) {
	x.xxx_hidden_Another = v
}

func (x *ComplexTestMsg) SetFloatConst(v float32) {
	x.xxx_hidden_FloatConst = v
}

func (x *ComplexTestMsg) SetDoubleIn(v float64) {
	x.xxx_hidden_DoubleIn = v
}

func (x *ComplexTestMsg) SetEnumConst(v ComplexTestEnum) {
	x.xxx_hidden_EnumConst = v
}

func (x *ComplexTestMsg) SetAnyVal(v *anypb.Any) {
	x.xxx_hidden_AnyVal = v
}

func (x *ComplexTestMsg) SetRepTsVal(v []*timestamppb.Timestamp) {
	x.xxx_hidden_RepTsVal = &v
}

func (x *ComplexTestMsg) SetMapVal(v map[int32]string) {
	x.xxx_hidden_MapVal = v
}

func (x *ComplexTestMsg) SetBytesVal(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_BytesVal = v
}

func (x *ComplexTestMsg) SetX(v string) {
	x.xxx_hidden_O = &complexTestMsg_X{v}
}

func (x *ComplexTestMsg) SetY(v int32) {
	x.xxx_hidden_O = &complexTestMsg_Y{v}
}

func (x *ComplexTestMsg) HasNested() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Nested != nil
}

func (x *ComplexTestMsg) HasFloatVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_FloatVal != nil
}

func (x *ComplexTestMsg) HasDurVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DurVal != nil
}

func (x *ComplexTestMsg) HasTsVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_TsVal != nil
}

func (x *ComplexTestMsg) HasAnother() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Another != nil
}

func (x *ComplexTestMsg) HasAnyVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_AnyVal != nil
}

func (x *ComplexTestMsg) HasO() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_O != nil
}

func (x *ComplexTestMsg) HasX() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_O.(*complexTestMsg_X)
	return ok
}

func (x *ComplexTestMsg) HasY() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_O.(*complexTestMsg_Y)
	return ok
}

func (x *ComplexTestMsg) ClearNested() {
	x.xxx_hidden_Nested = nil
}

func (x *ComplexTestMsg) ClearFloatVal() {
	x.xxx_hidden_FloatVal = nil
}

func (x *ComplexTestMsg) ClearDurVal() {
	x.xxx_hidden_DurVal = nil
}

func (x *ComplexTestMsg) ClearTsVal() {
	x.xxx_hidden_TsVal = nil
}

func (x *ComplexTestMsg) ClearAnother() {
	x.xxx_hidden_Another = nil
}

func (x *ComplexTestMsg) ClearAnyVal() {
	x.xxx_hidden_AnyVal = nil
}

func (x *ComplexTestMsg) ClearO() {
	x.xxx_hidden_O = nil
}

func (x *ComplexTestMsg) ClearX() {
	if _, ok := x.xxx_hidden_O.(*complexTestMsg_X); ok {
		x.xxx_hidden_O = nil
	}
}

func (x *ComplexTestMsg) ClearY() {
	if _, ok := x.xxx_hidden_O.(*complexTestMsg_Y); ok {
		x.xxx_hidden_O = nil
	}
}

const ComplexTestMsg_O_not_set_case case_ComplexTestMsg_O = 0
const ComplexTestMsg_X_case case_ComplexTestMsg_O = 16
const ComplexTestMsg_Y_case case_ComplexTestMsg_O = 17

func (x *ComplexTestMsg) WhichO() case_ComplexTestMsg_O {
	if x == nil {
		return ComplexTestMsg_O_not_set_case
	}
	switch x.xxx_hidden_O.(type) {
	case *complexTestMsg_X:
		return ComplexTestMsg_X_case
	case *complexTestMsg_Y:
		return ComplexTestMsg_Y_case
	default:
		return ComplexTestMsg_O_not_set_case
	}
}

type ComplexTestMsg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Const      string
	Nested     *ComplexTestMsg
	IntConst   int32
	BoolConst  bool
	FloatVal   *wrapperspb.FloatValue
	DurVal     *durationpb.Duration
	TsVal      *timestamppb.Timestamp
	Another    *ComplexTestMsg
	FloatConst float32
	DoubleIn   float64
	EnumConst  ComplexTestEnum
	AnyVal     *anypb.Any
	RepTsVal   []*timestamppb.Timestamp
	MapVal     map[int32]string
	BytesVal   []byte
	// Fields of oneof xxx_hidden_O:
	X *string
	Y *int32
	// -- end of xxx_hidden_O
}

func (b0 ComplexTestMsg_builder) Build() *ComplexTestMsg {
	m0 := &ComplexTestMsg{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Const = b.Const
	x.xxx_hidden_Nested = b.Nested
	x.xxx_hidden_IntConst = b.IntConst
	x.xxx_hidden_BoolConst = b.BoolConst
	x.xxx_hidden_FloatVal = b.FloatVal
	x.xxx_hidden_DurVal = b.DurVal
	x.xxx_hidden_TsVal = b.TsVal
	x.xxx_hidden_Another = b.Another
	x.xxx_hidden_FloatConst = b.FloatConst
	x.xxx_hidden_DoubleIn = b.DoubleIn
	x.xxx_hidden_EnumConst = b.EnumConst
	x.xxx_hidden_AnyVal = b.AnyVal
	x.xxx_hidden_RepTsVal = &b.RepTsVal
	x.xxx_hidden_MapVal = b.MapVal
	x.xxx_hidden_BytesVal = b.BytesVal
	if b.X != nil {
		x.xxx_hidden_O = &complexTestMsg_X{*b.X}
	}
	if b.Y != nil {
		x.xxx_hidden_O = &complexTestMsg_Y{*b.Y}
	}
	return m0
}

type case_ComplexTestMsg_O protoreflect.FieldNumber

func (x case_ComplexTestMsg_O) String() string {
	md := file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isComplexTestMsg_O interface {
	isComplexTestMsg_O()
}

type complexTestMsg_X struct {
	X string `protobuf:"bytes,16,opt,name=x,proto3,oneof"`
}

type complexTestMsg_Y struct {
	Y int32 `protobuf:"varint,17,opt,name=y,proto3,oneof"`
}

func (*complexTestMsg_X) isComplexTestMsg_O() {}

func (*complexTestMsg_Y) isComplexTestMsg_O() {}

type KitchenSinkMessage struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *ComplexTestMsg        `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *KitchenSinkMessage) Reset() {
	*x = KitchenSinkMessage{}
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KitchenSinkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KitchenSinkMessage) ProtoMessage() {}

func (x *KitchenSinkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *KitchenSinkMessage) GetVal() *ComplexTestMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *KitchenSinkMessage) SetVal(v *ComplexTestMsg) {
	x.xxx_hidden_Val = v
}

func (x *KitchenSinkMessage) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *KitchenSinkMessage) ClearVal() {
	x.xxx_hidden_Val = nil
}

type KitchenSinkMessage_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *ComplexTestMsg
}

func (b0 KitchenSinkMessage_builder) Build() *KitchenSinkMessage {
	m0 := &KitchenSinkMessage{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_kitchen_sink_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x08, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12,
	0x21, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0x72, 0x06, 0x0a, 0x04, 0x61, 0x62, 0x63, 0x64, 0x52, 0x05, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x1a, 0x02, 0x08, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xba, 0x48, 0x04, 0x6a, 0x02, 0x08, 0x00, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x0a, 0x05, 0x25,
	0x00, 0x00, 0x00, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x41,
	0x0a, 0x07, 0x64, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8,
	0x01, 0x01, 0xaa, 0x01, 0x04, 0x1a, 0x02, 0x08, 0x11, 0x52, 0x06, 0x64, 0x75, 0x72, 0x56, 0x61,
	0x6c, 0x12, 0x3d, 0x0a, 0x06, 0x74, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0xba,
	0x48, 0x07, 0xb2, 0x01, 0x04, 0x2a, 0x02, 0x08, 0x07, 0x52, 0x05, 0x74, 0x73, 0x56, 0x61, 0x6c,
	0x12, 0x48, 0x0a, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73,
	0x67, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0b, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0x0a, 0x05, 0x15, 0x00, 0x00, 0x00, 0x41, 0x52, 0x0a, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xba, 0x48, 0x14, 0x12,
	0x12, 0x31, 0xb4, 0xc8, 0x76, 0xbe, 0x9f, 0x8c, 0x7c, 0x40, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xc0, 0x5e, 0x40, 0x52, 0x08, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x58, 0x0a,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x75, 0x6d, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x08, 0x02, 0x52, 0x09, 0x65, 0x6e,
	0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x07, 0x61, 0x6e, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x34,
	0xba, 0x48, 0x31, 0xa2, 0x01, 0x2e, 0x12, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x5f, 0x74, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x11, 0xba, 0x48,
	0x0e, 0x92, 0x01, 0x0b, 0x22, 0x09, 0xb2, 0x01, 0x06, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x54, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x61, 0x0a, 0x07, 0x6d, 0x61, 0x70,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61,
	0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0c, 0xba, 0x48, 0x09, 0x9a, 0x01, 0x06, 0x22, 0x04,
	0x3a, 0x02, 0x10, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x09,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x09, 0xba, 0x48, 0x06, 0x7a, 0x04, 0x0a, 0x02, 0x00, 0x99, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x01, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x01, 0x78, 0x12, 0x0e, 0x0a, 0x01, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x01, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x0a, 0x01, 0x6f, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x56, 0x0a, 0x12, 0x4b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x40, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x2a, 0x6a, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x58, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x58, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42,
	0xa4, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42, 0x10, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x42,
	0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43,
	0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c,
	0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a,
	0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes = []any{
	(ComplexTestEnum)(0),          // 0: buf.validate.conformance.cases.ComplexTestEnum
	(*ComplexTestMsg)(nil),        // 1: buf.validate.conformance.cases.ComplexTestMsg
	(*KitchenSinkMessage)(nil),    // 2: buf.validate.conformance.cases.KitchenSinkMessage
	nil,                           // 3: buf.validate.conformance.cases.ComplexTestMsg.MapValEntry
	(*wrapperspb.FloatValue)(nil), // 4: google.protobuf.FloatValue
	(*durationpb.Duration)(nil),   // 5: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
}
var file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs = []int32{
	1,  // 0: buf.validate.conformance.cases.ComplexTestMsg.nested:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	4,  // 1: buf.validate.conformance.cases.ComplexTestMsg.float_val:type_name -> google.protobuf.FloatValue
	5,  // 2: buf.validate.conformance.cases.ComplexTestMsg.dur_val:type_name -> google.protobuf.Duration
	6,  // 3: buf.validate.conformance.cases.ComplexTestMsg.ts_val:type_name -> google.protobuf.Timestamp
	1,  // 4: buf.validate.conformance.cases.ComplexTestMsg.another:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	0,  // 5: buf.validate.conformance.cases.ComplexTestMsg.enum_const:type_name -> buf.validate.conformance.cases.ComplexTestEnum
	7,  // 6: buf.validate.conformance.cases.ComplexTestMsg.any_val:type_name -> google.protobuf.Any
	6,  // 7: buf.validate.conformance.cases.ComplexTestMsg.rep_ts_val:type_name -> google.protobuf.Timestamp
	3,  // 8: buf.validate.conformance.cases.ComplexTestMsg.map_val:type_name -> buf.validate.conformance.cases.ComplexTestMsg.MapValEntry
	1,  // 9: buf.validate.conformance.cases.KitchenSinkMessage.val:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_kitchen_sink_proto_init() }
func file_buf_validate_conformance_cases_kitchen_sink_proto_init() {
	if File_buf_validate_conformance_cases_kitchen_sink_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0].OneofWrappers = []any{
		(*complexTestMsg_X)(nil),
		(*complexTestMsg_Y)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs,
		EnumInfos:         file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes,
		MessageInfos:      file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_kitchen_sink_proto = out.File
	file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc = nil
	file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes = nil
	file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs = nil
}