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
// source: tests/example/v1/filter.proto

//go:build protoopaque

package examplev1

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

type InvalidConstraints struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Field int32                  `protobuf:"varint,1,opt,name=field,proto3"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *InvalidConstraints) Reset() {
	*x = InvalidConstraints{}
	mi := &file_tests_example_v1_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvalidConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidConstraints) ProtoMessage() {}

func (x *InvalidConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *InvalidConstraints) GetField() int32 {
	if x != nil {
		return x.xxx_hidden_Field
	}
	return 0
}

func (x *InvalidConstraints) SetField(v int32) {
	x.xxx_hidden_Field = v
}

type InvalidConstraints_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Field int32
}

func (b0 InvalidConstraints_builder) Build() *InvalidConstraints {
	m0 := &InvalidConstraints{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Field = b.Field
	return m0
}

type AllConstraintTypes struct {
	state                    protoimpl.MessageState             `protogen:"opaque.v1"`
	xxx_hidden_Field         int32                              `protobuf:"varint,1,opt,name=field,proto3"`
	xxx_hidden_RequiredOneof isAllConstraintTypes_RequiredOneof `protobuf_oneof:"required_oneof"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *AllConstraintTypes) Reset() {
	*x = AllConstraintTypes{}
	mi := &file_tests_example_v1_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllConstraintTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllConstraintTypes) ProtoMessage() {}

func (x *AllConstraintTypes) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *AllConstraintTypes) GetField() int32 {
	if x != nil {
		return x.xxx_hidden_Field
	}
	return 0
}

func (x *AllConstraintTypes) GetOneofField() string {
	if x != nil {
		if x, ok := x.xxx_hidden_RequiredOneof.(*allConstraintTypes_OneofField); ok {
			return x.OneofField
		}
	}
	return ""
}

func (x *AllConstraintTypes) SetField(v int32) {
	x.xxx_hidden_Field = v
}

func (x *AllConstraintTypes) SetOneofField(v string) {
	x.xxx_hidden_RequiredOneof = &allConstraintTypes_OneofField{v}
}

func (x *AllConstraintTypes) HasRequiredOneof() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_RequiredOneof != nil
}

func (x *AllConstraintTypes) HasOneofField() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_RequiredOneof.(*allConstraintTypes_OneofField)
	return ok
}

func (x *AllConstraintTypes) ClearRequiredOneof() {
	x.xxx_hidden_RequiredOneof = nil
}

func (x *AllConstraintTypes) ClearOneofField() {
	if _, ok := x.xxx_hidden_RequiredOneof.(*allConstraintTypes_OneofField); ok {
		x.xxx_hidden_RequiredOneof = nil
	}
}

const AllConstraintTypes_RequiredOneof_not_set_case case_AllConstraintTypes_RequiredOneof = 0
const AllConstraintTypes_OneofField_case case_AllConstraintTypes_RequiredOneof = 2

func (x *AllConstraintTypes) WhichRequiredOneof() case_AllConstraintTypes_RequiredOneof {
	if x == nil {
		return AllConstraintTypes_RequiredOneof_not_set_case
	}
	switch x.xxx_hidden_RequiredOneof.(type) {
	case *allConstraintTypes_OneofField:
		return AllConstraintTypes_OneofField_case
	default:
		return AllConstraintTypes_RequiredOneof_not_set_case
	}
}

type AllConstraintTypes_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Field int32
	// Fields of oneof xxx_hidden_RequiredOneof:
	OneofField *string
	// -- end of xxx_hidden_RequiredOneof
}

func (b0 AllConstraintTypes_builder) Build() *AllConstraintTypes {
	m0 := &AllConstraintTypes{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Field = b.Field
	if b.OneofField != nil {
		x.xxx_hidden_RequiredOneof = &allConstraintTypes_OneofField{*b.OneofField}
	}
	return m0
}

type case_AllConstraintTypes_RequiredOneof protoreflect.FieldNumber

func (x case_AllConstraintTypes_RequiredOneof) String() string {
	md := file_tests_example_v1_filter_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isAllConstraintTypes_RequiredOneof interface {
	isAllConstraintTypes_RequiredOneof()
}

type allConstraintTypes_OneofField struct {
	OneofField string `protobuf:"bytes,2,opt,name=oneof_field,json=oneofField,proto3,oneof"`
}

func (*allConstraintTypes_OneofField) isAllConstraintTypes_RequiredOneof() {}

type NestedConstraints struct {
	state                    protoimpl.MessageState            `protogen:"opaque.v1"`
	xxx_hidden_Field         *AllConstraintTypes               `protobuf:"bytes,1,opt,name=field,proto3"`
	xxx_hidden_Field2        string                            `protobuf:"bytes,2,opt,name=field2,proto3"`
	xxx_hidden_RepeatedField *[]*AllConstraintTypes            `protobuf:"bytes,3,rep,name=repeated_field,json=repeatedField,proto3"`
	xxx_hidden_MapField      map[string]*AllConstraintTypes    `protobuf:"bytes,4,rep,name=map_field,json=mapField,proto3" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_RequiredOneof isNestedConstraints_RequiredOneof `protobuf_oneof:"required_oneof"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *NestedConstraints) Reset() {
	*x = NestedConstraints{}
	mi := &file_tests_example_v1_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NestedConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedConstraints) ProtoMessage() {}

func (x *NestedConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *NestedConstraints) GetField() *AllConstraintTypes {
	if x != nil {
		return x.xxx_hidden_Field
	}
	return nil
}

func (x *NestedConstraints) GetField2() string {
	if x != nil {
		return x.xxx_hidden_Field2
	}
	return ""
}

func (x *NestedConstraints) GetRepeatedField() []*AllConstraintTypes {
	if x != nil {
		if x.xxx_hidden_RepeatedField != nil {
			return *x.xxx_hidden_RepeatedField
		}
	}
	return nil
}

func (x *NestedConstraints) GetMapField() map[string]*AllConstraintTypes {
	if x != nil {
		return x.xxx_hidden_MapField
	}
	return nil
}

func (x *NestedConstraints) GetOneofField() string {
	if x != nil {
		if x, ok := x.xxx_hidden_RequiredOneof.(*nestedConstraints_OneofField); ok {
			return x.OneofField
		}
	}
	return ""
}

func (x *NestedConstraints) SetField(v *AllConstraintTypes) {
	x.xxx_hidden_Field = v
}

func (x *NestedConstraints) SetField2(v string) {
	x.xxx_hidden_Field2 = v
}

func (x *NestedConstraints) SetRepeatedField(v []*AllConstraintTypes) {
	x.xxx_hidden_RepeatedField = &v
}

func (x *NestedConstraints) SetMapField(v map[string]*AllConstraintTypes) {
	x.xxx_hidden_MapField = v
}

func (x *NestedConstraints) SetOneofField(v string) {
	x.xxx_hidden_RequiredOneof = &nestedConstraints_OneofField{v}
}

func (x *NestedConstraints) HasField() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Field != nil
}

func (x *NestedConstraints) HasRequiredOneof() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_RequiredOneof != nil
}

func (x *NestedConstraints) HasOneofField() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_RequiredOneof.(*nestedConstraints_OneofField)
	return ok
}

func (x *NestedConstraints) ClearField() {
	x.xxx_hidden_Field = nil
}

func (x *NestedConstraints) ClearRequiredOneof() {
	x.xxx_hidden_RequiredOneof = nil
}

func (x *NestedConstraints) ClearOneofField() {
	if _, ok := x.xxx_hidden_RequiredOneof.(*nestedConstraints_OneofField); ok {
		x.xxx_hidden_RequiredOneof = nil
	}
}

const NestedConstraints_RequiredOneof_not_set_case case_NestedConstraints_RequiredOneof = 0
const NestedConstraints_OneofField_case case_NestedConstraints_RequiredOneof = 5

func (x *NestedConstraints) WhichRequiredOneof() case_NestedConstraints_RequiredOneof {
	if x == nil {
		return NestedConstraints_RequiredOneof_not_set_case
	}
	switch x.xxx_hidden_RequiredOneof.(type) {
	case *nestedConstraints_OneofField:
		return NestedConstraints_OneofField_case
	default:
		return NestedConstraints_RequiredOneof_not_set_case
	}
}

type NestedConstraints_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Field         *AllConstraintTypes
	Field2        string
	RepeatedField []*AllConstraintTypes
	MapField      map[string]*AllConstraintTypes
	// Fields of oneof xxx_hidden_RequiredOneof:
	OneofField *string
	// -- end of xxx_hidden_RequiredOneof
}

func (b0 NestedConstraints_builder) Build() *NestedConstraints {
	m0 := &NestedConstraints{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Field = b.Field
	x.xxx_hidden_Field2 = b.Field2
	x.xxx_hidden_RepeatedField = &b.RepeatedField
	x.xxx_hidden_MapField = b.MapField
	if b.OneofField != nil {
		x.xxx_hidden_RequiredOneof = &nestedConstraints_OneofField{*b.OneofField}
	}
	return m0
}

type case_NestedConstraints_RequiredOneof protoreflect.FieldNumber

func (x case_NestedConstraints_RequiredOneof) String() string {
	md := file_tests_example_v1_filter_proto_msgTypes[2].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isNestedConstraints_RequiredOneof interface {
	isNestedConstraints_RequiredOneof()
}

type nestedConstraints_OneofField struct {
	OneofField string `protobuf:"bytes,5,opt,name=oneof_field,json=oneofField,proto3,oneof"`
}

func (*nestedConstraints_OneofField) isNestedConstraints_RequiredOneof() {}

var File_tests_example_v1_filter_proto protoreflect.FileDescriptor

var file_tests_example_v1_filter_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1,
	0x01, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x5e, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x48, 0xba, 0x48, 0x45, 0xba, 0x01, 0x42, 0x0a, 0x10, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x1a, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4b, 0xba, 0x48, 0x48, 0x1a, 0x46, 0x0a, 0x12, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x12, 0x22, 0x74, 0x68, 0x69, 0x73, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x1a, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x22, 0xf3, 0x01, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x43, 0xba, 0x48, 0x40, 0xba, 0x01, 0x3d,
	0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x12, 0x22, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73,
	0x20, 0x66, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x46, 0xba, 0x48, 0x43, 0x1a, 0x41, 0x0a, 0x12,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x12, 0x24, 0x74, 0x68, 0x69, 0x73, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x61, 0x6c, 0x77, 0x61,
	0x79, 0x73, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x42,
	0x17, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0xbe, 0x04, 0x0a, 0x11, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x86,
	0x01, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x42, 0x4a, 0xba, 0x48, 0x47, 0xba, 0x01, 0x44, 0x0a, 0x17, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x20, 0x61, 0x6c, 0x77,
	0x61, 0x79, 0x73, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x64, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4c, 0xba, 0x48, 0x49, 0xba, 0x01, 0x46, 0x0a,
	0x19, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x32, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x20, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x05,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x4b, 0x0a,
	0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4e, 0x0a, 0x09, 0x6d, 0x61,
	0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x61, 0x0a,
	0x0d, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x17, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42, 0xd3, 0x01, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73,
	0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x65,
	0x73, 0x74, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_tests_example_v1_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tests_example_v1_filter_proto_goTypes = []any{
	(*InvalidConstraints)(nil), // 0: tests.example.v1.InvalidConstraints
	(*AllConstraintTypes)(nil), // 1: tests.example.v1.AllConstraintTypes
	(*NestedConstraints)(nil),  // 2: tests.example.v1.NestedConstraints
	nil,                        // 3: tests.example.v1.NestedConstraints.MapFieldEntry
}
var file_tests_example_v1_filter_proto_depIdxs = []int32{
	1, // 0: tests.example.v1.NestedConstraints.field:type_name -> tests.example.v1.AllConstraintTypes
	1, // 1: tests.example.v1.NestedConstraints.repeated_field:type_name -> tests.example.v1.AllConstraintTypes
	3, // 2: tests.example.v1.NestedConstraints.map_field:type_name -> tests.example.v1.NestedConstraints.MapFieldEntry
	1, // 3: tests.example.v1.NestedConstraints.MapFieldEntry.value:type_name -> tests.example.v1.AllConstraintTypes
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tests_example_v1_filter_proto_init() }
func file_tests_example_v1_filter_proto_init() {
	if File_tests_example_v1_filter_proto != nil {
		return
	}
	file_tests_example_v1_filter_proto_msgTypes[1].OneofWrappers = []any{
		(*allConstraintTypes_OneofField)(nil),
	}
	file_tests_example_v1_filter_proto_msgTypes[2].OneofWrappers = []any{
		(*nestedConstraints_OneofField)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tests_example_v1_filter_proto_rawDesc), len(file_tests_example_v1_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_example_v1_filter_proto_goTypes,
		DependencyIndexes: file_tests_example_v1_filter_proto_depIdxs,
		MessageInfos:      file_tests_example_v1_filter_proto_msgTypes,
	}.Build()
	File_tests_example_v1_filter_proto = out.File
	file_tests_example_v1_filter_proto_goTypes = nil
	file_tests_example_v1_filter_proto_depIdxs = nil
}
