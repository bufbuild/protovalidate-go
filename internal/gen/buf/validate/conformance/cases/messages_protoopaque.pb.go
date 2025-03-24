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
// source: buf/validate/conformance/cases/messages.proto

//go:build protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	other_package "github.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases/other_package"
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

type TestMsg struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Const  string                 `protobuf:"bytes,1,opt,name=const,proto3"`
	xxx_hidden_Nested *TestMsg               `protobuf:"bytes,2,opt,name=nested,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TestMsg) Reset() {
	*x = TestMsg{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMsg) ProtoMessage() {}

func (x *TestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TestMsg) GetConst() string {
	if x != nil {
		return x.xxx_hidden_Const
	}
	return ""
}

func (x *TestMsg) GetNested() *TestMsg {
	if x != nil {
		return x.xxx_hidden_Nested
	}
	return nil
}

func (x *TestMsg) SetConst(v string) {
	x.xxx_hidden_Const = v
}

func (x *TestMsg) SetNested(v *TestMsg) {
	x.xxx_hidden_Nested = v
}

func (x *TestMsg) HasNested() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Nested != nil
}

func (x *TestMsg) ClearNested() {
	x.xxx_hidden_Nested = nil
}

type TestMsg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Const  string
	Nested *TestMsg
}

func (b0 TestMsg_builder) Build() *TestMsg {
	m0 := &TestMsg{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Const = b.Const
	x.xxx_hidden_Nested = b.Nested
	return m0
}

type MessageNone struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *MessageNone_NoneMsg   `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageNone) Reset() {
	*x = MessageNone{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageNone) ProtoMessage() {}

func (x *MessageNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageNone) GetVal() *MessageNone_NoneMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *MessageNone) SetVal(v *MessageNone_NoneMsg) {
	x.xxx_hidden_Val = v
}

func (x *MessageNone) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *MessageNone) ClearVal() {
	x.xxx_hidden_Val = nil
}

type MessageNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *MessageNone_NoneMsg
}

func (b0 MessageNone_builder) Build() *MessageNone {
	m0 := &MessageNone{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageDisabled struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val uint64                 `protobuf:"varint,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageDisabled) Reset() {
	*x = MessageDisabled{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageDisabled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDisabled) ProtoMessage() {}

func (x *MessageDisabled) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageDisabled) GetVal() uint64 {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return 0
}

func (x *MessageDisabled) SetVal(v uint64) {
	x.xxx_hidden_Val = v
}

type MessageDisabled_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val uint64
}

func (b0 MessageDisabled_builder) Build() *MessageDisabled {
	m0 := &MessageDisabled{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type Message struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *TestMsg               `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Message) GetVal() *TestMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *Message) SetVal(v *TestMsg) {
	x.xxx_hidden_Val = v
}

func (x *Message) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *Message) ClearVal() {
	x.xxx_hidden_Val = nil
}

type Message_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *TestMsg
}

func (b0 Message_builder) Build() *Message {
	m0 := &Message{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageCrossPackage struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *other_package.Embed   `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageCrossPackage) Reset() {
	*x = MessageCrossPackage{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageCrossPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCrossPackage) ProtoMessage() {}

func (x *MessageCrossPackage) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageCrossPackage) GetVal() *other_package.Embed {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *MessageCrossPackage) SetVal(v *other_package.Embed) {
	x.xxx_hidden_Val = v
}

func (x *MessageCrossPackage) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *MessageCrossPackage) ClearVal() {
	x.xxx_hidden_Val = nil
}

type MessageCrossPackage_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *other_package.Embed
}

func (b0 MessageCrossPackage_builder) Build() *MessageCrossPackage {
	m0 := &MessageCrossPackage{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageSkip struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *TestMsg               `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageSkip) Reset() {
	*x = MessageSkip{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSkip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSkip) ProtoMessage() {}

func (x *MessageSkip) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageSkip) GetVal() *TestMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *MessageSkip) SetVal(v *TestMsg) {
	x.xxx_hidden_Val = v
}

func (x *MessageSkip) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *MessageSkip) ClearVal() {
	x.xxx_hidden_Val = nil
}

type MessageSkip_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *TestMsg
}

func (b0 MessageSkip_builder) Build() *MessageSkip {
	m0 := &MessageSkip{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageRequired struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *TestMsg               `protobuf:"bytes,1,opt,name=val,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageRequired) Reset() {
	*x = MessageRequired{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequired) ProtoMessage() {}

func (x *MessageRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageRequired) GetVal() *TestMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *MessageRequired) SetVal(v *TestMsg) {
	x.xxx_hidden_Val = v
}

func (x *MessageRequired) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *MessageRequired) ClearVal() {
	x.xxx_hidden_Val = nil
}

type MessageRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *TestMsg
}

func (b0 MessageRequired_builder) Build() *MessageRequired {
	m0 := &MessageRequired{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageRequiredButOptional struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Val *TestMsg               `protobuf:"bytes,1,opt,name=val,proto3,oneof"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageRequiredButOptional) Reset() {
	*x = MessageRequiredButOptional{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequiredButOptional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequiredButOptional) ProtoMessage() {}

func (x *MessageRequiredButOptional) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageRequiredButOptional) GetVal() *TestMsg {
	if x != nil {
		return x.xxx_hidden_Val
	}
	return nil
}

func (x *MessageRequiredButOptional) SetVal(v *TestMsg) {
	x.xxx_hidden_Val = v
}

func (x *MessageRequiredButOptional) HasVal() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Val != nil
}

func (x *MessageRequiredButOptional) ClearVal() {
	x.xxx_hidden_Val = nil
}

type MessageRequiredButOptional_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *TestMsg
}

func (b0 MessageRequiredButOptional_builder) Build() *MessageRequiredButOptional {
	m0 := &MessageRequiredButOptional{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Val = b.Val
	return m0
}

type MessageRequiredOneof struct {
	state          protoimpl.MessageState     `protogen:"opaque.v1"`
	xxx_hidden_One isMessageRequiredOneof_One `protobuf_oneof:"one"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MessageRequiredOneof) Reset() {
	*x = MessageRequiredOneof{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequiredOneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequiredOneof) ProtoMessage() {}

func (x *MessageRequiredOneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MessageRequiredOneof) GetVal() *TestMsg {
	if x != nil {
		if x, ok := x.xxx_hidden_One.(*messageRequiredOneof_Val); ok {
			return x.Val
		}
	}
	return nil
}

func (x *MessageRequiredOneof) SetVal(v *TestMsg) {
	if v == nil {
		x.xxx_hidden_One = nil
		return
	}
	x.xxx_hidden_One = &messageRequiredOneof_Val{v}
}

func (x *MessageRequiredOneof) HasOne() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_One != nil
}

func (x *MessageRequiredOneof) HasVal() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_One.(*messageRequiredOneof_Val)
	return ok
}

func (x *MessageRequiredOneof) ClearOne() {
	x.xxx_hidden_One = nil
}

func (x *MessageRequiredOneof) ClearVal() {
	if _, ok := x.xxx_hidden_One.(*messageRequiredOneof_Val); ok {
		x.xxx_hidden_One = nil
	}
}

const MessageRequiredOneof_One_not_set_case case_MessageRequiredOneof_One = 0
const MessageRequiredOneof_Val_case case_MessageRequiredOneof_One = 1

func (x *MessageRequiredOneof) WhichOne() case_MessageRequiredOneof_One {
	if x == nil {
		return MessageRequiredOneof_One_not_set_case
	}
	switch x.xxx_hidden_One.(type) {
	case *messageRequiredOneof_Val:
		return MessageRequiredOneof_Val_case
	default:
		return MessageRequiredOneof_One_not_set_case
	}
}

type MessageRequiredOneof_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof xxx_hidden_One:
	Val *TestMsg
	// -- end of xxx_hidden_One
}

func (b0 MessageRequiredOneof_builder) Build() *MessageRequiredOneof {
	m0 := &MessageRequiredOneof{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Val != nil {
		x.xxx_hidden_One = &messageRequiredOneof_Val{b.Val}
	}
	return m0
}

type case_MessageRequiredOneof_One protoreflect.FieldNumber

func (x case_MessageRequiredOneof_One) String() string {
	md := file_buf_validate_conformance_cases_messages_proto_msgTypes[8].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isMessageRequiredOneof_One interface {
	isMessageRequiredOneof_One()
}

type messageRequiredOneof_Val struct {
	Val *TestMsg `protobuf:"bytes,1,opt,name=val,proto3,oneof"`
}

func (*messageRequiredOneof_Val) isMessageRequiredOneof_One() {}

type MessageWith3DInside struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageWith3DInside) Reset() {
	*x = MessageWith3DInside{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageWith3DInside) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWith3DInside) ProtoMessage() {}

func (x *MessageWith3DInside) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type MessageWith3DInside_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 MessageWith3DInside_builder) Build() *MessageWith3DInside {
	m0 := &MessageWith3DInside{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type MessageNone_NoneMsg struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageNone_NoneMsg) Reset() {
	*x = MessageNone_NoneMsg{}
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageNone_NoneMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageNone_NoneMsg) ProtoMessage() {}

func (x *MessageNone_NoneMsg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_messages_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type MessageNone_NoneMsg_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 MessageNone_NoneMsg_builder) Build() *MessageNone_NoneMsg {
	m0 := &MessageNone_NoneMsg{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_buf_validate_conformance_cases_messages_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_messages_proto_rawDesc = "" +
	"\n" +
	"-buf/validate/conformance/cases/messages.proto\x12\x1ebuf.validate.conformance.cases\x1a8buf/validate/conformance/cases/other_package/embed.proto\x1a\x1bbuf/validate/validate.proto\"l\n" +
	"\aTestMsg\x12 \n" +
	"\x05const\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\n" +
	"\x03fooR\x05const\x12?\n" +
	"\x06nested\x18\x02 \x01(\v2'.buf.validate.conformance.cases.TestMsgR\x06nested\"_\n" +
	"\vMessageNone\x12E\n" +
	"\x03val\x18\x01 \x01(\v23.buf.validate.conformance.cases.MessageNone.NoneMsgR\x03val\x1a\t\n" +
	"\aNoneMsg\"3\n" +
	"\x0fMessageDisabled\x12\x19\n" +
	"\x03val\x18\x01 \x01(\x04B\a\xbaH\x042\x02 {R\x03val:\x05\xbaH\x02\b\x01\"D\n" +
	"\aMessage\x129\n" +
	"\x03val\x18\x01 \x01(\v2'.buf.validate.conformance.cases.TestMsgR\x03val\"\\\n" +
	"\x13MessageCrossPackage\x12E\n" +
	"\x03val\x18\x01 \x01(\v23.buf.validate.conformance.cases.other_package.EmbedR\x03val\"P\n" +
	"\vMessageSkip\x12A\n" +
	"\x03val\x18\x01 \x01(\v2'.buf.validate.conformance.cases.TestMsgB\x06\xbaH\x03\xd8\x01\x03R\x03val\"T\n" +
	"\x0fMessageRequired\x12A\n" +
	"\x03val\x18\x01 \x01(\v2'.buf.validate.conformance.cases.TestMsgB\x06\xbaH\x03\xc8\x01\x01R\x03val\"l\n" +
	"\x1aMessageRequiredButOptional\x12F\n" +
	"\x03val\x18\x01 \x01(\v2'.buf.validate.conformance.cases.TestMsgB\x06\xbaH\x03\xc8\x01\x01H\x00R\x03val\x88\x01\x01B\x06\n" +
	"\x04_val\"i\n" +
	"\x14MessageRequiredOneof\x12C\n" +
	"\x03val\x18\x01 \x01(\v2'.buf.validate.conformance.cases.TestMsgB\x06\xbaH\x03\xc8\x01\x01H\x00R\x03valB\f\n" +
	"\x03one\x12\x05\xbaH\x02\b\x01\"\x15\n" +
	"\x13MessageWith3dInsideB\xa1\x02\n" +
	"\"com.buf.validate.conformance.casesB\rMessagesProtoP\x01ZPgithub.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases\xa2\x02\x04BVCC\xaa\x02\x1eBuf.Validate.Conformance.Cases\xca\x02\x1eBuf\\Validate\\Conformance\\Cases\xe2\x02*Buf\\Validate\\Conformance\\Cases\\GPBMetadata\xea\x02!Buf::Validate::Conformance::Casesb\x06proto3"

var file_buf_validate_conformance_cases_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_buf_validate_conformance_cases_messages_proto_goTypes = []any{
	(*TestMsg)(nil),                    // 0: buf.validate.conformance.cases.TestMsg
	(*MessageNone)(nil),                // 1: buf.validate.conformance.cases.MessageNone
	(*MessageDisabled)(nil),            // 2: buf.validate.conformance.cases.MessageDisabled
	(*Message)(nil),                    // 3: buf.validate.conformance.cases.Message
	(*MessageCrossPackage)(nil),        // 4: buf.validate.conformance.cases.MessageCrossPackage
	(*MessageSkip)(nil),                // 5: buf.validate.conformance.cases.MessageSkip
	(*MessageRequired)(nil),            // 6: buf.validate.conformance.cases.MessageRequired
	(*MessageRequiredButOptional)(nil), // 7: buf.validate.conformance.cases.MessageRequiredButOptional
	(*MessageRequiredOneof)(nil),       // 8: buf.validate.conformance.cases.MessageRequiredOneof
	(*MessageWith3DInside)(nil),        // 9: buf.validate.conformance.cases.MessageWith3dInside
	(*MessageNone_NoneMsg)(nil),        // 10: buf.validate.conformance.cases.MessageNone.NoneMsg
	(*other_package.Embed)(nil),        // 11: buf.validate.conformance.cases.other_package.Embed
}
var file_buf_validate_conformance_cases_messages_proto_depIdxs = []int32{
	0,  // 0: buf.validate.conformance.cases.TestMsg.nested:type_name -> buf.validate.conformance.cases.TestMsg
	10, // 1: buf.validate.conformance.cases.MessageNone.val:type_name -> buf.validate.conformance.cases.MessageNone.NoneMsg
	0,  // 2: buf.validate.conformance.cases.Message.val:type_name -> buf.validate.conformance.cases.TestMsg
	11, // 3: buf.validate.conformance.cases.MessageCrossPackage.val:type_name -> buf.validate.conformance.cases.other_package.Embed
	0,  // 4: buf.validate.conformance.cases.MessageSkip.val:type_name -> buf.validate.conformance.cases.TestMsg
	0,  // 5: buf.validate.conformance.cases.MessageRequired.val:type_name -> buf.validate.conformance.cases.TestMsg
	0,  // 6: buf.validate.conformance.cases.MessageRequiredButOptional.val:type_name -> buf.validate.conformance.cases.TestMsg
	0,  // 7: buf.validate.conformance.cases.MessageRequiredOneof.val:type_name -> buf.validate.conformance.cases.TestMsg
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_messages_proto_init() }
func file_buf_validate_conformance_cases_messages_proto_init() {
	if File_buf_validate_conformance_cases_messages_proto != nil {
		return
	}
	file_buf_validate_conformance_cases_messages_proto_msgTypes[7].OneofWrappers = []any{}
	file_buf_validate_conformance_cases_messages_proto_msgTypes[8].OneofWrappers = []any{
		(*messageRequiredOneof_Val)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_messages_proto_rawDesc), len(file_buf_validate_conformance_cases_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_messages_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_messages_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_messages_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_messages_proto = out.File
	file_buf_validate_conformance_cases_messages_proto_goTypes = nil
	file_buf_validate_conformance_cases_messages_proto_depIdxs = nil
}
