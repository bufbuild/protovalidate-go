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
// source: buf/validate/conformance/cases/yet_another_package/embed2.proto

//go:build !protoopaque

package yet_another_package

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

type Embed_Enumerated int32

const (
	Embed_ENUMERATED_UNSPECIFIED Embed_Enumerated = 0
	Embed_ENUMERATED_VALUE       Embed_Enumerated = 1
)

// Enum value maps for Embed_Enumerated.
var (
	Embed_Enumerated_name = map[int32]string{
		0: "ENUMERATED_UNSPECIFIED",
		1: "ENUMERATED_VALUE",
	}
	Embed_Enumerated_value = map[string]int32{
		"ENUMERATED_UNSPECIFIED": 0,
		"ENUMERATED_VALUE":       1,
	}
)

func (x Embed_Enumerated) Enum() *Embed_Enumerated {
	p := new(Embed_Enumerated)
	*p = x
	return p
}

func (x Embed_Enumerated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Embed_Enumerated) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes[0].Descriptor()
}

func (Embed_Enumerated) Type() protoreflect.EnumType {
	return &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes[0]
}

func (x Embed_Enumerated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Validate message embedding across packages.
type Embed struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           int64                  `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Embed) Reset() {
	*x = Embed{}
	mi := &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed) ProtoMessage() {}

func (x *Embed) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Embed) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

func (x *Embed) SetVal(v int64) {
	x.Val = v
}

type Embed_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val int64
}

func (b0 Embed_builder) Build() *Embed {
	m0 := &Embed{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_yet_another_package_embed2_proto protoreflect.FileDescriptor

const file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc = "" +
	"\n" +
	"?buf/validate/conformance/cases/yet_another_package/embed2.proto\x122buf.validate.conformance.cases.yet_another_package\x1a\x1bbuf/validate/validate.proto\"b\n" +
	"\x05Embed\x12\x19\n" +
	"\x03val\x18\x01 \x01(\x03B\a\xbaH\x04\"\x02 \x00R\x03val\">\n" +
	"\n" +
	"Enumerated\x12\x1a\n" +
	"\x16ENUMERATED_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10ENUMERATED_VALUE\x10\x01B\x91\x03\n" +
	"6com.buf.validate.conformance.cases.yet_another_packageB\vEmbed2ProtoP\x01Zdgithub.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases/yet_another_package\xa2\x02\x05BVCCY\xaa\x020Buf.Validate.Conformance.Cases.YetAnotherPackage\xca\x020Buf\\Validate\\Conformance\\Cases\\YetAnotherPackage\xe2\x02<Buf\\Validate\\Conformance\\Cases\\YetAnotherPackage\\GPBMetadata\xea\x024Buf::Validate::Conformance::Cases::YetAnotherPackageb\x06proto3"

var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes = []any{
	(Embed_Enumerated)(0), // 0: buf.validate.conformance.cases.yet_another_package.Embed.Enumerated
	(*Embed)(nil),         // 1: buf.validate.conformance.cases.yet_another_package.Embed
}
var file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_yet_another_package_embed2_proto_init() }
func file_buf_validate_conformance_cases_yet_another_package_embed2_proto_init() {
	if File_buf_validate_conformance_cases_yet_another_package_embed2_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc), len(file_buf_validate_conformance_cases_yet_another_package_embed2_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs,
		EnumInfos:         file_buf_validate_conformance_cases_yet_another_package_embed2_proto_enumTypes,
		MessageInfos:      file_buf_validate_conformance_cases_yet_another_package_embed2_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_yet_another_package_embed2_proto = out.File
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_goTypes = nil
	file_buf_validate_conformance_cases_yet_another_package_embed2_proto_depIdxs = nil
}
