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
// source: buf/validate/conformance/cases/oneofs.proto

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

type TestOneofMsg struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           bool                   `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestOneofMsg) Reset() {
	*x = TestOneofMsg{}
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestOneofMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOneofMsg) ProtoMessage() {}

func (x *TestOneofMsg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TestOneofMsg) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

func (x *TestOneofMsg) SetVal(v bool) {
	x.Val = v
}

type TestOneofMsg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val bool
}

func (b0 TestOneofMsg_builder) Build() *TestOneofMsg {
	m0 := &TestOneofMsg{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type OneofNone struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to O:
	//
	//	*OneofNone_X
	//	*OneofNone_Y
	O             isOneofNone_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OneofNone) Reset() {
	*x = OneofNone{}
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OneofNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofNone) ProtoMessage() {}

func (x *OneofNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OneofNone) GetO() isOneofNone_O {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *OneofNone) GetX() string {
	if x != nil {
		if x, ok := x.O.(*OneofNone_X); ok {
			return x.X
		}
	}
	return ""
}

func (x *OneofNone) GetY() int32 {
	if x != nil {
		if x, ok := x.O.(*OneofNone_Y); ok {
			return x.Y
		}
	}
	return 0
}

func (x *OneofNone) SetX(v string) {
	x.O = &OneofNone_X{v}
}

func (x *OneofNone) SetY(v int32) {
	x.O = &OneofNone_Y{v}
}

func (x *OneofNone) HasO() bool {
	if x == nil {
		return false
	}
	return x.O != nil
}

func (x *OneofNone) HasX() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofNone_X)
	return ok
}

func (x *OneofNone) HasY() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofNone_Y)
	return ok
}

func (x *OneofNone) ClearO() {
	x.O = nil
}

func (x *OneofNone) ClearX() {
	if _, ok := x.O.(*OneofNone_X); ok {
		x.O = nil
	}
}

func (x *OneofNone) ClearY() {
	if _, ok := x.O.(*OneofNone_Y); ok {
		x.O = nil
	}
}

const OneofNone_O_not_set_case case_OneofNone_O = 0
const OneofNone_X_case case_OneofNone_O = 1
const OneofNone_Y_case case_OneofNone_O = 2

func (x *OneofNone) WhichO() case_OneofNone_O {
	if x == nil {
		return OneofNone_O_not_set_case
	}
	switch x.O.(type) {
	case *OneofNone_X:
		return OneofNone_X_case
	case *OneofNone_Y:
		return OneofNone_Y_case
	default:
		return OneofNone_O_not_set_case
	}
}

type OneofNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof O:
	X *string
	Y *int32
	// -- end of O
}

func (b0 OneofNone_builder) Build() *OneofNone {
	m0 := &OneofNone{}
	b, x := &b0, m0
	_, _ = b, x
	if b.X != nil {
		x.O = &OneofNone_X{*b.X}
	}
	if b.Y != nil {
		x.O = &OneofNone_Y{*b.Y}
	}
	return m0
}

type case_OneofNone_O protoreflect.FieldNumber

func (x case_OneofNone_O) String() string {
	md := file_buf_validate_conformance_cases_oneofs_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isOneofNone_O interface {
	isOneofNone_O()
}

type OneofNone_X struct {
	X string `protobuf:"bytes,1,opt,name=x,proto3,oneof"`
}

type OneofNone_Y struct {
	Y int32 `protobuf:"varint,2,opt,name=y,proto3,oneof"`
}

func (*OneofNone_X) isOneofNone_O() {}

func (*OneofNone_Y) isOneofNone_O() {}

type Oneof struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to O:
	//
	//	*Oneof_X
	//	*Oneof_Y
	//	*Oneof_Z
	O             isOneof_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Oneof) Reset() {
	*x = Oneof{}
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oneof) ProtoMessage() {}

func (x *Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Oneof) GetO() isOneof_O {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *Oneof) GetX() string {
	if x != nil {
		if x, ok := x.O.(*Oneof_X); ok {
			return x.X
		}
	}
	return ""
}

func (x *Oneof) GetY() int32 {
	if x != nil {
		if x, ok := x.O.(*Oneof_Y); ok {
			return x.Y
		}
	}
	return 0
}

func (x *Oneof) GetZ() *TestOneofMsg {
	if x != nil {
		if x, ok := x.O.(*Oneof_Z); ok {
			return x.Z
		}
	}
	return nil
}

func (x *Oneof) SetX(v string) {
	x.O = &Oneof_X{v}
}

func (x *Oneof) SetY(v int32) {
	x.O = &Oneof_Y{v}
}

func (x *Oneof) SetZ(v *TestOneofMsg) {
	if v == nil {
		x.O = nil
		return
	}
	x.O = &Oneof_Z{v}
}

func (x *Oneof) HasO() bool {
	if x == nil {
		return false
	}
	return x.O != nil
}

func (x *Oneof) HasX() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*Oneof_X)
	return ok
}

func (x *Oneof) HasY() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*Oneof_Y)
	return ok
}

func (x *Oneof) HasZ() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*Oneof_Z)
	return ok
}

func (x *Oneof) ClearO() {
	x.O = nil
}

func (x *Oneof) ClearX() {
	if _, ok := x.O.(*Oneof_X); ok {
		x.O = nil
	}
}

func (x *Oneof) ClearY() {
	if _, ok := x.O.(*Oneof_Y); ok {
		x.O = nil
	}
}

func (x *Oneof) ClearZ() {
	if _, ok := x.O.(*Oneof_Z); ok {
		x.O = nil
	}
}

const Oneof_O_not_set_case case_Oneof_O = 0
const Oneof_X_case case_Oneof_O = 1
const Oneof_Y_case case_Oneof_O = 2
const Oneof_Z_case case_Oneof_O = 3

func (x *Oneof) WhichO() case_Oneof_O {
	if x == nil {
		return Oneof_O_not_set_case
	}
	switch x.O.(type) {
	case *Oneof_X:
		return Oneof_X_case
	case *Oneof_Y:
		return Oneof_Y_case
	case *Oneof_Z:
		return Oneof_Z_case
	default:
		return Oneof_O_not_set_case
	}
}

type Oneof_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof O:
	X *string
	Y *int32
	Z *TestOneofMsg
	// -- end of O
}

func (b0 Oneof_builder) Build() *Oneof {
	m0 := &Oneof{}
	b, x := &b0, m0
	_, _ = b, x
	if b.X != nil {
		x.O = &Oneof_X{*b.X}
	}
	if b.Y != nil {
		x.O = &Oneof_Y{*b.Y}
	}
	if b.Z != nil {
		x.O = &Oneof_Z{b.Z}
	}
	return m0
}

type case_Oneof_O protoreflect.FieldNumber

func (x case_Oneof_O) String() string {
	md := file_buf_validate_conformance_cases_oneofs_proto_msgTypes[2].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isOneof_O interface {
	isOneof_O()
}

type Oneof_X struct {
	X string `protobuf:"bytes,1,opt,name=x,proto3,oneof"`
}

type Oneof_Y struct {
	Y int32 `protobuf:"varint,2,opt,name=y,proto3,oneof"`
}

type Oneof_Z struct {
	Z *TestOneofMsg `protobuf:"bytes,3,opt,name=z,proto3,oneof"`
}

func (*Oneof_X) isOneof_O() {}

func (*Oneof_Y) isOneof_O() {}

func (*Oneof_Z) isOneof_O() {}

type OneofRequired struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to O:
	//
	//	*OneofRequired_X
	//	*OneofRequired_Y
	//	*OneofRequired_NameWithUnderscores
	//	*OneofRequired_UnderAnd_1Number
	O             isOneofRequired_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OneofRequired) Reset() {
	*x = OneofRequired{}
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OneofRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofRequired) ProtoMessage() {}

func (x *OneofRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OneofRequired) GetO() isOneofRequired_O {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *OneofRequired) GetX() string {
	if x != nil {
		if x, ok := x.O.(*OneofRequired_X); ok {
			return x.X
		}
	}
	return ""
}

func (x *OneofRequired) GetY() int32 {
	if x != nil {
		if x, ok := x.O.(*OneofRequired_Y); ok {
			return x.Y
		}
	}
	return 0
}

func (x *OneofRequired) GetNameWithUnderscores() int32 {
	if x != nil {
		if x, ok := x.O.(*OneofRequired_NameWithUnderscores); ok {
			return x.NameWithUnderscores
		}
	}
	return 0
}

func (x *OneofRequired) GetUnderAnd_1Number() int32 {
	if x != nil {
		if x, ok := x.O.(*OneofRequired_UnderAnd_1Number); ok {
			return x.UnderAnd_1Number
		}
	}
	return 0
}

func (x *OneofRequired) SetX(v string) {
	x.O = &OneofRequired_X{v}
}

func (x *OneofRequired) SetY(v int32) {
	x.O = &OneofRequired_Y{v}
}

func (x *OneofRequired) SetNameWithUnderscores(v int32) {
	x.O = &OneofRequired_NameWithUnderscores{v}
}

func (x *OneofRequired) SetUnderAnd_1Number(v int32) {
	x.O = &OneofRequired_UnderAnd_1Number{v}
}

func (x *OneofRequired) HasO() bool {
	if x == nil {
		return false
	}
	return x.O != nil
}

func (x *OneofRequired) HasX() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequired_X)
	return ok
}

func (x *OneofRequired) HasY() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequired_Y)
	return ok
}

func (x *OneofRequired) HasNameWithUnderscores() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequired_NameWithUnderscores)
	return ok
}

func (x *OneofRequired) HasUnderAnd_1Number() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequired_UnderAnd_1Number)
	return ok
}

func (x *OneofRequired) ClearO() {
	x.O = nil
}

func (x *OneofRequired) ClearX() {
	if _, ok := x.O.(*OneofRequired_X); ok {
		x.O = nil
	}
}

func (x *OneofRequired) ClearY() {
	if _, ok := x.O.(*OneofRequired_Y); ok {
		x.O = nil
	}
}

func (x *OneofRequired) ClearNameWithUnderscores() {
	if _, ok := x.O.(*OneofRequired_NameWithUnderscores); ok {
		x.O = nil
	}
}

func (x *OneofRequired) ClearUnderAnd_1Number() {
	if _, ok := x.O.(*OneofRequired_UnderAnd_1Number); ok {
		x.O = nil
	}
}

const OneofRequired_O_not_set_case case_OneofRequired_O = 0
const OneofRequired_X_case case_OneofRequired_O = 1
const OneofRequired_Y_case case_OneofRequired_O = 2
const OneofRequired_NameWithUnderscores_case case_OneofRequired_O = 3
const OneofRequired_UnderAnd_1Number_case case_OneofRequired_O = 4

func (x *OneofRequired) WhichO() case_OneofRequired_O {
	if x == nil {
		return OneofRequired_O_not_set_case
	}
	switch x.O.(type) {
	case *OneofRequired_X:
		return OneofRequired_X_case
	case *OneofRequired_Y:
		return OneofRequired_Y_case
	case *OneofRequired_NameWithUnderscores:
		return OneofRequired_NameWithUnderscores_case
	case *OneofRequired_UnderAnd_1Number:
		return OneofRequired_UnderAnd_1Number_case
	default:
		return OneofRequired_O_not_set_case
	}
}

type OneofRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof O:
	X                   *string
	Y                   *int32
	NameWithUnderscores *int32
	UnderAnd_1Number    *int32
	// -- end of O
}

func (b0 OneofRequired_builder) Build() *OneofRequired {
	m0 := &OneofRequired{}
	b, x := &b0, m0
	_, _ = b, x
	if b.X != nil {
		x.O = &OneofRequired_X{*b.X}
	}
	if b.Y != nil {
		x.O = &OneofRequired_Y{*b.Y}
	}
	if b.NameWithUnderscores != nil {
		x.O = &OneofRequired_NameWithUnderscores{*b.NameWithUnderscores}
	}
	if b.UnderAnd_1Number != nil {
		x.O = &OneofRequired_UnderAnd_1Number{*b.UnderAnd_1Number}
	}
	return m0
}

type case_OneofRequired_O protoreflect.FieldNumber

func (x case_OneofRequired_O) String() string {
	md := file_buf_validate_conformance_cases_oneofs_proto_msgTypes[3].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isOneofRequired_O interface {
	isOneofRequired_O()
}

type OneofRequired_X struct {
	X string `protobuf:"bytes,1,opt,name=x,proto3,oneof"`
}

type OneofRequired_Y struct {
	Y int32 `protobuf:"varint,2,opt,name=y,proto3,oneof"`
}

type OneofRequired_NameWithUnderscores struct {
	NameWithUnderscores int32 `protobuf:"varint,3,opt,name=name_with_underscores,json=nameWithUnderscores,proto3,oneof"`
}

type OneofRequired_UnderAnd_1Number struct {
	UnderAnd_1Number int32 `protobuf:"varint,4,opt,name=under_and_1_number,json=underAnd1Number,proto3,oneof"`
}

func (*OneofRequired_X) isOneofRequired_O() {}

func (*OneofRequired_Y) isOneofRequired_O() {}

func (*OneofRequired_NameWithUnderscores) isOneofRequired_O() {}

func (*OneofRequired_UnderAnd_1Number) isOneofRequired_O() {}

type OneofRequiredWithRequiredField struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to O:
	//
	//	*OneofRequiredWithRequiredField_A
	//	*OneofRequiredWithRequiredField_B
	O             isOneofRequiredWithRequiredField_O `protobuf_oneof:"o"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OneofRequiredWithRequiredField) Reset() {
	*x = OneofRequiredWithRequiredField{}
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OneofRequiredWithRequiredField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofRequiredWithRequiredField) ProtoMessage() {}

func (x *OneofRequiredWithRequiredField) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_oneofs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OneofRequiredWithRequiredField) GetO() isOneofRequiredWithRequiredField_O {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *OneofRequiredWithRequiredField) GetA() string {
	if x != nil {
		if x, ok := x.O.(*OneofRequiredWithRequiredField_A); ok {
			return x.A
		}
	}
	return ""
}

func (x *OneofRequiredWithRequiredField) GetB() string {
	if x != nil {
		if x, ok := x.O.(*OneofRequiredWithRequiredField_B); ok {
			return x.B
		}
	}
	return ""
}

func (x *OneofRequiredWithRequiredField) SetA(v string) {
	x.O = &OneofRequiredWithRequiredField_A{v}
}

func (x *OneofRequiredWithRequiredField) SetB(v string) {
	x.O = &OneofRequiredWithRequiredField_B{v}
}

func (x *OneofRequiredWithRequiredField) HasO() bool {
	if x == nil {
		return false
	}
	return x.O != nil
}

func (x *OneofRequiredWithRequiredField) HasA() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequiredWithRequiredField_A)
	return ok
}

func (x *OneofRequiredWithRequiredField) HasB() bool {
	if x == nil {
		return false
	}
	_, ok := x.O.(*OneofRequiredWithRequiredField_B)
	return ok
}

func (x *OneofRequiredWithRequiredField) ClearO() {
	x.O = nil
}

func (x *OneofRequiredWithRequiredField) ClearA() {
	if _, ok := x.O.(*OneofRequiredWithRequiredField_A); ok {
		x.O = nil
	}
}

func (x *OneofRequiredWithRequiredField) ClearB() {
	if _, ok := x.O.(*OneofRequiredWithRequiredField_B); ok {
		x.O = nil
	}
}

const OneofRequiredWithRequiredField_O_not_set_case case_OneofRequiredWithRequiredField_O = 0
const OneofRequiredWithRequiredField_A_case case_OneofRequiredWithRequiredField_O = 1
const OneofRequiredWithRequiredField_B_case case_OneofRequiredWithRequiredField_O = 2

func (x *OneofRequiredWithRequiredField) WhichO() case_OneofRequiredWithRequiredField_O {
	if x == nil {
		return OneofRequiredWithRequiredField_O_not_set_case
	}
	switch x.O.(type) {
	case *OneofRequiredWithRequiredField_A:
		return OneofRequiredWithRequiredField_A_case
	case *OneofRequiredWithRequiredField_B:
		return OneofRequiredWithRequiredField_B_case
	default:
		return OneofRequiredWithRequiredField_O_not_set_case
	}
}

type OneofRequiredWithRequiredField_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof O:
	A *string
	B *string
	// -- end of O
}

func (b0 OneofRequiredWithRequiredField_builder) Build() *OneofRequiredWithRequiredField {
	m0 := &OneofRequiredWithRequiredField{}
	b, x := &b0, m0
	_, _ = b, x
	if b.A != nil {
		x.O = &OneofRequiredWithRequiredField_A{*b.A}
	}
	if b.B != nil {
		x.O = &OneofRequiredWithRequiredField_B{*b.B}
	}
	return m0
}

type case_OneofRequiredWithRequiredField_O protoreflect.FieldNumber

func (x case_OneofRequiredWithRequiredField_O) String() string {
	md := file_buf_validate_conformance_cases_oneofs_proto_msgTypes[4].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isOneofRequiredWithRequiredField_O interface {
	isOneofRequiredWithRequiredField_O()
}

type OneofRequiredWithRequiredField_A struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3,oneof"`
}

type OneofRequiredWithRequiredField_B struct {
	B string `protobuf:"bytes,2,opt,name=b,proto3,oneof"`
}

func (*OneofRequiredWithRequiredField_A) isOneofRequiredWithRequiredField_O() {}

func (*OneofRequiredWithRequiredField_B) isOneofRequiredWithRequiredField_O() {}

var File_buf_validate_conformance_cases_oneofs_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_oneofs_proto_rawDesc = "" +
	"\n" +
	"+buf/validate/conformance/cases/oneofs.proto\x12\x1ebuf.validate.conformance.cases\x1a\x1bbuf/validate/validate.proto\")\n" +
	"\fTestOneofMsg\x12\x19\n" +
	"\x03val\x18\x01 \x01(\bB\a\xbaH\x04j\x02\b\x01R\x03val\"0\n" +
	"\tOneofNone\x12\x0e\n" +
	"\x01x\x18\x01 \x01(\tH\x00R\x01x\x12\x0e\n" +
	"\x01y\x18\x02 \x01(\x05H\x00R\x01yB\x03\n" +
	"\x01o\"\x7f\n" +
	"\x05Oneof\x12\x1a\n" +
	"\x01x\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05:\x03fooH\x00R\x01x\x12\x17\n" +
	"\x01y\x18\x02 \x01(\x05B\a\xbaH\x04\x1a\x02 \x00H\x00R\x01y\x12<\n" +
	"\x01z\x18\x03 \x01(\v2,.buf.validate.conformance.cases.TestOneofMsgH\x00R\x01zB\x03\n" +
	"\x01o\"\xa0\x01\n" +
	"\rOneofRequired\x12\x0e\n" +
	"\x01x\x18\x01 \x01(\tH\x00R\x01x\x12\x0e\n" +
	"\x01y\x18\x02 \x01(\x05H\x00R\x01y\x124\n" +
	"\x15name_with_underscores\x18\x03 \x01(\x05H\x00R\x13nameWithUnderscores\x12-\n" +
	"\x12under_and_1_number\x18\x04 \x01(\x05H\x00R\x0funderAnd1NumberB\n" +
	"\n" +
	"\x01o\x12\x05\xbaH\x02\b\x01\"T\n" +
	"\x1eOneofRequiredWithRequiredField\x12\x16\n" +
	"\x01a\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01H\x00R\x01a\x12\x0e\n" +
	"\x01b\x18\x02 \x01(\tH\x00R\x01bB\n" +
	"\n" +
	"\x01o\x12\x05\xbaH\x02\b\x01B\x95\x02\n" +
	"\"com.buf.validate.conformance.casesB\vOneofsProtoP\x01ZFbuf.build/go/protovalidate/internal/gen/buf/validate/conformance/cases\xa2\x02\x04BVCC\xaa\x02\x1eBuf.Validate.Conformance.Cases\xca\x02\x1eBuf\\Validate\\Conformance\\Cases\xe2\x02*Buf\\Validate\\Conformance\\Cases\\GPBMetadata\xea\x02!Buf::Validate::Conformance::Casesb\x06proto3"

var file_buf_validate_conformance_cases_oneofs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_buf_validate_conformance_cases_oneofs_proto_goTypes = []any{
	(*TestOneofMsg)(nil),                   // 0: buf.validate.conformance.cases.TestOneofMsg
	(*OneofNone)(nil),                      // 1: buf.validate.conformance.cases.OneofNone
	(*Oneof)(nil),                          // 2: buf.validate.conformance.cases.Oneof
	(*OneofRequired)(nil),                  // 3: buf.validate.conformance.cases.OneofRequired
	(*OneofRequiredWithRequiredField)(nil), // 4: buf.validate.conformance.cases.OneofRequiredWithRequiredField
}
var file_buf_validate_conformance_cases_oneofs_proto_depIdxs = []int32{
	0, // 0: buf.validate.conformance.cases.Oneof.z:type_name -> buf.validate.conformance.cases.TestOneofMsg
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_oneofs_proto_init() }
func file_buf_validate_conformance_cases_oneofs_proto_init() {
	if File_buf_validate_conformance_cases_oneofs_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_oneofs_proto_msgTypes[1].OneofWrappers = []any{
		(*OneofNone_X)(nil),
		(*OneofNone_Y)(nil),
	}
	file_buf_validate_conformance_cases_oneofs_proto_msgTypes[2].OneofWrappers = []any{
		(*Oneof_X)(nil),
		(*Oneof_Y)(nil),
		(*Oneof_Z)(nil),
	}
	file_buf_validate_conformance_cases_oneofs_proto_msgTypes[3].OneofWrappers = []any{
		(*OneofRequired_X)(nil),
		(*OneofRequired_Y)(nil),
		(*OneofRequired_NameWithUnderscores)(nil),
		(*OneofRequired_UnderAnd_1Number)(nil),
	}
	file_buf_validate_conformance_cases_oneofs_proto_msgTypes[4].OneofWrappers = []any{
		(*OneofRequiredWithRequiredField_A)(nil),
		(*OneofRequiredWithRequiredField_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_oneofs_proto_rawDesc), len(file_buf_validate_conformance_cases_oneofs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_oneofs_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_oneofs_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_oneofs_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_oneofs_proto = out.File
	file_buf_validate_conformance_cases_oneofs_proto_goTypes = nil
	file_buf_validate_conformance_cases_oneofs_proto_depIdxs = nil
}
