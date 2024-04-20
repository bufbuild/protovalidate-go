// Copyright 2023 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/wkt_nested.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WktLevelOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Two *WktLevelOne_WktLevelTwo `protobuf:"bytes,1,opt,name=two,proto3" json:"two,omitempty"`
}

func (x *WktLevelOne) Reset() {
	*x = WktLevelOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WktLevelOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WktLevelOne) ProtoMessage() {}

func (x *WktLevelOne) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WktLevelOne.ProtoReflect.Descriptor instead.
func (*WktLevelOne) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_nested_proto_rawDescGZIP(), []int{0}
}

func (x *WktLevelOne) GetTwo() *WktLevelOne_WktLevelTwo {
	if x != nil {
		return x.Two
	}
	return nil
}

type WktLevelOne_WktLevelTwo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Three *WktLevelOne_WktLevelTwo_WktLevelThree `protobuf:"bytes,1,opt,name=three,proto3" json:"three,omitempty"`
}

func (x *WktLevelOne_WktLevelTwo) Reset() {
	*x = WktLevelOne_WktLevelTwo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WktLevelOne_WktLevelTwo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WktLevelOne_WktLevelTwo) ProtoMessage() {}

func (x *WktLevelOne_WktLevelTwo) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WktLevelOne_WktLevelTwo.ProtoReflect.Descriptor instead.
func (*WktLevelOne_WktLevelTwo) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_nested_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WktLevelOne_WktLevelTwo) GetThree() *WktLevelOne_WktLevelTwo_WktLevelThree {
	if x != nil {
		return x.Three
	}
	return nil
}

type WktLevelOne_WktLevelTwo_WktLevelThree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *WktLevelOne_WktLevelTwo_WktLevelThree) Reset() {
	*x = WktLevelOne_WktLevelTwo_WktLevelThree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WktLevelOne_WktLevelTwo_WktLevelThree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WktLevelOne_WktLevelTwo_WktLevelThree) ProtoMessage() {}

func (x *WktLevelOne_WktLevelTwo_WktLevelThree) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WktLevelOne_WktLevelTwo_WktLevelThree.ProtoReflect.Descriptor instead.
func (*WktLevelOne_WktLevelTwo_WktLevelThree) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_nested_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *WktLevelOne_WktLevelTwo_WktLevelThree) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_buf_validate_conformance_cases_wkt_nested_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_wkt_nested_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x77, 0x6b, 0x74, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x02, 0x0a, 0x0b, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x12, 0x51,
	0x0a, 0x03, 0x74, 0x77, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x57, 0x6b, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x2e, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x54, 0x77, 0x6f, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x74, 0x77,
	0x6f, 0x1a, 0xa1, 0x01, 0x0a, 0x0b, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x77,
	0x6f, 0x12, 0x63, 0x0a, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x6e, 0x65, 0x2e, 0x57, 0x6b,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x77, 0x6f, 0x2e, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x1a, 0x2d, 0x0a, 0x0d, 0x57, 0x6b, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0xa2, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42, 0x0e, 0x57, 0x6b,
	0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_wkt_nested_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_wkt_nested_proto_rawDescData = file_buf_validate_conformance_cases_wkt_nested_proto_rawDesc
)

func file_buf_validate_conformance_cases_wkt_nested_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_wkt_nested_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_wkt_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_wkt_nested_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_wkt_nested_proto_rawDescData
}

var file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_validate_conformance_cases_wkt_nested_proto_goTypes = []interface{}{
	(*WktLevelOne)(nil),                           // 0: buf.validate.conformance.cases.WktLevelOne
	(*WktLevelOne_WktLevelTwo)(nil),               // 1: buf.validate.conformance.cases.WktLevelOne.WktLevelTwo
	(*WktLevelOne_WktLevelTwo_WktLevelThree)(nil), // 2: buf.validate.conformance.cases.WktLevelOne.WktLevelTwo.WktLevelThree
}
var file_buf_validate_conformance_cases_wkt_nested_proto_depIdxs = []int32{
	1, // 0: buf.validate.conformance.cases.WktLevelOne.two:type_name -> buf.validate.conformance.cases.WktLevelOne.WktLevelTwo
	2, // 1: buf.validate.conformance.cases.WktLevelOne.WktLevelTwo.three:type_name -> buf.validate.conformance.cases.WktLevelOne.WktLevelTwo.WktLevelThree
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_wkt_nested_proto_init() }
func file_buf_validate_conformance_cases_wkt_nested_proto_init() {
	if File_buf_validate_conformance_cases_wkt_nested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WktLevelOne); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WktLevelOne_WktLevelTwo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WktLevelOne_WktLevelTwo_WktLevelThree); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_wkt_nested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_wkt_nested_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_wkt_nested_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_wkt_nested_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_wkt_nested_proto = out.File
	file_buf_validate_conformance_cases_wkt_nested_proto_rawDesc = nil
	file_buf_validate_conformance_cases_wkt_nested_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_nested_proto_depIdxs = nil
}
