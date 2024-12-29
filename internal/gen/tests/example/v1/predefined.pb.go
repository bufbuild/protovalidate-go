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
// source: tests/example/v1/predefined.proto

//go:build !protoopaque

package examplev1

import (
	validate "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var file_tests_example_v1_predefined_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*validate.Int32Rules)(nil),
		ExtensionType: ([]int32)(nil),
		Field:         1800,
		Name:          "tests.example.v1.abs_not_in",
		Tag:           "varint,1800,rep,name=abs_not_in",
		Filename:      "tests/example/v1/predefined.proto",
	},
}

// Extension fields to validate.Int32Rules.
var (
	// repeated int32 abs_not_in = 1800;
	E_AbsNotIn = &file_tests_example_v1_predefined_proto_extTypes[0]
)

var File_tests_example_v1_predefined_proto protoreflect.FileDescriptor

var file_tests_example_v1_predefined_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x34, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x31, 0x34, 0x38, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0xba, 0x48,
	0x11, 0x1a, 0x0f, 0xc0, 0x70, 0x01, 0xc0, 0x70, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa6, 0x01, 0x0a, 0x0a, 0x61, 0x62, 0x73,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x88, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x42, 0x6d, 0xc2, 0x48, 0x6a, 0x0a, 0x68, 0x0a,
	0x10, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x2e, 0x61, 0x62, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x69,
	0x6e, 0x12, 0x2b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x27,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x72, 0x75, 0x6c, 0x65, 0x20, 0x7c, 0x7c, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x28, 0x6e, 0x2c, 0x20, 0x2d, 0x6e, 0x29, 0x52, 0x08, 0x61, 0x62, 0x73, 0x4e, 0x6f, 0x74, 0x49,
	0x6e, 0x42, 0xd7, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x50, 0x72, 0x65, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x45,
	0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x5c, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x73, 0x74, 0x73, 0x5c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x3a,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31,
}

var file_tests_example_v1_predefined_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_example_v1_predefined_proto_goTypes = []any{
	(*Issue148)(nil),            // 0: tests.example.v1.Issue148
	(*validate.Int32Rules)(nil), // 1: buf.validate.Int32Rules
}
var file_tests_example_v1_predefined_proto_depIdxs = []int32{
	1, // 0: tests.example.v1.abs_not_in:extendee -> buf.validate.Int32Rules
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
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
			RawDescriptor: file_tests_example_v1_predefined_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_tests_example_v1_predefined_proto_goTypes,
		DependencyIndexes: file_tests_example_v1_predefined_proto_depIdxs,
		MessageInfos:      file_tests_example_v1_predefined_proto_msgTypes,
		ExtensionInfos:    file_tests_example_v1_predefined_proto_extTypes,
	}.Build()
	File_tests_example_v1_predefined_proto = out.File
	file_tests_example_v1_predefined_proto_rawDesc = nil
	file_tests_example_v1_predefined_proto_goTypes = nil
	file_tests_example_v1_predefined_proto_depIdxs = nil
}
