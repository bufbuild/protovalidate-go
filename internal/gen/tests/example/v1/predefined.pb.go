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
// source: tests/example/v1/predefined.proto

//go:build !protoopaque

package examplev1

import (
	validate "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// https://github.com/bufbuild/protovalidate-go/issues/148
type Issue148 struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Test          *int32                 `protobuf:"varint,1,opt,name=test" json:"test,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Issue148) Reset() {
	*x = Issue148{}
	mi := &file_tests_example_v1_predefined_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Issue148) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Issue148) ProtoMessage() {}

func (x *Issue148) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_predefined_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Issue148) GetTest() int32 {
	if x != nil && x.Test != nil {
		return *x.Test
	}
	return 0
}

func (x *Issue148) SetTest(v int32) {
	x.Test = &v
}

func (x *Issue148) HasTest() bool {
	if x == nil {
		return false
	}
	return x.Test != nil
}

func (x *Issue148) ClearTest() {
	x.Test = nil
}

type Issue148_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Test *int32
}

func (b0 Issue148_builder) Build() *Issue148 {
	m0 := &Issue148{}
	b, x := &b0, m0
	_, _ = b, x
	x.Test = b.Test
	return m0
}

// https://github.com/bufbuild/protovalidate-go/issues/187
type Issue187 struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	FalseField    *bool                  `protobuf:"varint,1,opt,name=false_field,json=falseField" json:"false_field,omitempty"`
	TrueField     *bool                  `protobuf:"varint,2,opt,name=true_field,json=trueField" json:"true_field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Issue187) Reset() {
	*x = Issue187{}
	mi := &file_tests_example_v1_predefined_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Issue187) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Issue187) ProtoMessage() {}

func (x *Issue187) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_predefined_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Issue187) GetFalseField() bool {
	if x != nil && x.FalseField != nil {
		return *x.FalseField
	}
	return false
}

func (x *Issue187) GetTrueField() bool {
	if x != nil && x.TrueField != nil {
		return *x.TrueField
	}
	return false
}

func (x *Issue187) SetFalseField(v bool) {
	x.FalseField = &v
}

func (x *Issue187) SetTrueField(v bool) {
	x.TrueField = &v
}

func (x *Issue187) HasFalseField() bool {
	if x == nil {
		return false
	}
	return x.FalseField != nil
}

func (x *Issue187) HasTrueField() bool {
	if x == nil {
		return false
	}
	return x.TrueField != nil
}

func (x *Issue187) ClearFalseField() {
	x.FalseField = nil
}

func (x *Issue187) ClearTrueField() {
	x.TrueField = nil
}

type Issue187_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	FalseField *bool
	TrueField  *bool
}

func (b0 Issue187_builder) Build() *Issue187 {
	m0 := &Issue187{}
	b, x := &b0, m0
	_, _ = b, x
	x.FalseField = b.FalseField
	x.TrueField = b.TrueField
	return m0
}

var file_tests_example_v1_predefined_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*validate.Int32Rules)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         1800,
		Name:          "tests.example.v1.abs_not_in",
		Tag:           "varint,1800,rep,name=abs_not_in",
		Filename:      "tests/example/v1/predefined.proto",
	},
	{
		ExtendedType:  (*validate.BoolRules)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1800,
		Name:          "tests.example.v1.this_equals_rule",
		Tag:           "varint,1800,opt,name=this_equals_rule",
		Filename:      "tests/example/v1/predefined.proto",
	},
}

// Extension fields to validate.Int32Rules.
var (
	// repeated int32 abs_not_in = 1800;
	E_AbsNotIn = &file_tests_example_v1_predefined_proto_extTypes[0]
)

// Extension fields to validate.BoolRules.
var (
	// optional bool this_equals_rule = 1800;
	E_ThisEqualsRule = &file_tests_example_v1_predefined_proto_extTypes[1]
)

var File_tests_example_v1_predefined_proto protoreflect.FileDescriptor

const file_tests_example_v1_predefined_proto_rawDesc = "" +
	"\n" +
	"!tests/example/v1/predefined.proto\x12\x10tests.example.v1\x1a\x1bbuf/validate/validate.proto\"4\n" +
	"\bIssue148\x12(\n" +
	"\x04test\x18\x01 \x01(\x05B\x14\xbaH\x11\x1a\x0f\xc0p\x01\xc0p\xfe\xff\xff\xff\xff\xff\xff\xff\xff\x01R\x04test\"^\n" +
	"\bIssue187\x12)\n" +
	"\vfalse_field\x18\x01 \x01(\bB\b\xbaH\x05j\x03\xc0p\x00R\n" +
	"falseField\x12'\n" +
	"\n" +
	"true_field\x18\x02 \x01(\bB\b\xbaH\x05j\x03\xc0p\x01R\ttrueField:\xa6\x01\n" +
	"\n" +
	"abs_not_in\x12\x18.buf.validate.Int32Rules\x18\x88\x0e \x03(\x05Bm\xc2Hj\n" +
	"h\n" +
	"\x10int32.abs_not_in\x12+value must not be in absolute value of list\x1a'this in rule || this in rule.map(n, -n)R\babsNotIn:\xb1\x01\n" +
	"\x10this_equals_rule\x12\x17.buf.validate.BoolRules\x18\x88\x0e \x01(\bBm\xc2Hj\n" +
	"h\n" +
	"\x15bool.this_equals_rule\x1aOthis == rule ? '' : 'this = %s, rule = %s'.format([string(this), string(rule)])R\x0ethisEqualsRuleB\xd7\x01\n" +
	"\x14com.tests.example.v1B\x0fPredefinedProtoP\x01ZLgithub.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1;examplev1\xa2\x02\x03TEX\xaa\x02\x10Tests.Example.V1\xca\x02\x10Tests\\Example\\V1\xe2\x02\x1cTests\\Example\\V1\\GPBMetadata\xea\x02\x12Tests::Example::V1"

var file_tests_example_v1_predefined_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_example_v1_predefined_proto_goTypes = []any{
	(*Issue148)(nil),            // 0: tests.example.v1.Issue148
	(*Issue187)(nil),            // 1: tests.example.v1.Issue187
	(*validate.Int32Rules)(nil), // 2: buf.validate.Int32Rules
	(*validate.BoolRules)(nil),  // 3: buf.validate.BoolRules
}
var file_tests_example_v1_predefined_proto_depIdxs = []int32{
	2, // 0: tests.example.v1.abs_not_in:extendee -> buf.validate.Int32Rules
	3, // 1: tests.example.v1.this_equals_rule:extendee -> buf.validate.BoolRules
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_example_v1_predefined_proto_init() }
func file_tests_example_v1_predefined_proto_init() {
	if File_tests_example_v1_predefined_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tests_example_v1_predefined_proto_rawDesc), len(file_tests_example_v1_predefined_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_tests_example_v1_predefined_proto_goTypes,
		DependencyIndexes: file_tests_example_v1_predefined_proto_depIdxs,
		MessageInfos:      file_tests_example_v1_predefined_proto_msgTypes,
		ExtensionInfos:    file_tests_example_v1_predefined_proto_extTypes,
	}.Build()
	File_tests_example_v1_predefined_proto = out.File
	file_tests_example_v1_predefined_proto_goTypes = nil
	file_tests_example_v1_predefined_proto_depIdxs = nil
}
