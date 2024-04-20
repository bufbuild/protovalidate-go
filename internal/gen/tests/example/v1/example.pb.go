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
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: tests/example/v1/example.proto

package examplev1

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Home  *Coordinates `protobuf:"bytes,4,opt,name=home,proto3" json:"home,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_example_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetHome() *Coordinates {
	if x != nil {
		return x.Home
	}
	return nil
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_example_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type Person2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Person2) Reset() {
	*x = Person2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_example_v1_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person2) ProtoMessage() {}

func (x *Person2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_example_v1_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person2.ProtoReflect.Descriptor instead.
func (*Person2) Descriptor() ([]byte, []int) {
	return file_tests_example_v1_example_proto_rawDescGZIP(), []int{2}
}

func (x *Person2) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person2) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_tests_example_v1_example_proto protoreflect.FileDescriptor

var file_tests_example_v1_example_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0x48, 0x05, 0x32, 0x03, 0x20, 0xe7, 0x07,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x28, 0xba, 0x48, 0x25, 0x72, 0x23, 0x28, 0x80, 0x02, 0x32, 0x1e, 0x5e, 0x5b, 0x5b,
	0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x2b, 0x28, 0x20, 0x5b, 0x5b, 0x3a, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x2b, 0x29, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x04,
	0x68, 0x6f, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x42, 0x17, 0xba, 0x48, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x56, 0x40,
	0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x56, 0xc0, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x29,
	0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xba, 0x48, 0x14,
	0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x66, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x80, 0x66, 0xc0, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x6c, 0x0a, 0x07, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x32, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x32, 0x03, 0x20, 0xe7, 0x07, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x28, 0xba,
	0x48, 0x25, 0x1a, 0x23, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x19, 0x69, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x29, 0x42, 0xd4, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x73, 0x5c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_example_v1_example_proto_rawDescOnce sync.Once
	file_tests_example_v1_example_proto_rawDescData = file_tests_example_v1_example_proto_rawDesc
)

func file_tests_example_v1_example_proto_rawDescGZIP() []byte {
	file_tests_example_v1_example_proto_rawDescOnce.Do(func() {
		file_tests_example_v1_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_example_v1_example_proto_rawDescData)
	})
	return file_tests_example_v1_example_proto_rawDescData
}

var file_tests_example_v1_example_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tests_example_v1_example_proto_goTypes = []interface{}{
	(*Person)(nil),      // 0: tests.example.v1.Person
	(*Coordinates)(nil), // 1: tests.example.v1.Coordinates
	(*Person2)(nil),     // 2: tests.example.v1.Person2
}
var file_tests_example_v1_example_proto_depIdxs = []int32{
	1, // 0: tests.example.v1.Person.home:type_name -> tests.example.v1.Coordinates
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tests_example_v1_example_proto_init() }
func file_tests_example_v1_example_proto_init() {
	if File_tests_example_v1_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_example_v1_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_tests_example_v1_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
		file_tests_example_v1_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person2); i {
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
			RawDescriptor: file_tests_example_v1_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_example_v1_example_proto_goTypes,
		DependencyIndexes: file_tests_example_v1_example_proto_depIdxs,
		MessageInfos:      file_tests_example_v1_example_proto_msgTypes,
	}.Build()
	File_tests_example_v1_example_proto = out.File
	file_tests_example_v1_example_proto_rawDesc = nil
	file_tests_example_v1_example_proto_goTypes = nil
	file_tests_example_v1_example_proto_depIdxs = nil
}
