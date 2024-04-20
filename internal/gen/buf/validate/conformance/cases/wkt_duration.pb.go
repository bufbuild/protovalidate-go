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
// source: buf/validate/conformance/cases/wkt_duration.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DurationNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationNone) Reset() {
	*x = DurationNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationNone) ProtoMessage() {}

func (x *DurationNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationNone.ProtoReflect.Descriptor instead.
func (*DurationNone) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{0}
}

func (x *DurationNone) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationRequired struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationRequired) Reset() {
	*x = DurationRequired{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationRequired) ProtoMessage() {}

func (x *DurationRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationRequired.ProtoReflect.Descriptor instead.
func (*DurationRequired) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{1}
}

func (x *DurationRequired) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationConst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationConst) Reset() {
	*x = DurationConst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationConst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationConst) ProtoMessage() {}

func (x *DurationConst) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationConst.ProtoReflect.Descriptor instead.
func (*DurationConst) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{2}
}

func (x *DurationConst) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationIn) Reset() {
	*x = DurationIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationIn) ProtoMessage() {}

func (x *DurationIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationIn.ProtoReflect.Descriptor instead.
func (*DurationIn) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{3}
}

func (x *DurationIn) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationNotIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationNotIn) Reset() {
	*x = DurationNotIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationNotIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationNotIn) ProtoMessage() {}

func (x *DurationNotIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationNotIn.ProtoReflect.Descriptor instead.
func (*DurationNotIn) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{4}
}

func (x *DurationNotIn) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationLT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationLT) Reset() {
	*x = DurationLT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationLT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationLT) ProtoMessage() {}

func (x *DurationLT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationLT.ProtoReflect.Descriptor instead.
func (*DurationLT) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{5}
}

func (x *DurationLT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationLTE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationLTE) Reset() {
	*x = DurationLTE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationLTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationLTE) ProtoMessage() {}

func (x *DurationLTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationLTE.ProtoReflect.Descriptor instead.
func (*DurationLTE) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{6}
}

func (x *DurationLTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationGT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationGT) Reset() {
	*x = DurationGT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationGT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGT) ProtoMessage() {}

func (x *DurationGT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationGT.ProtoReflect.Descriptor instead.
func (*DurationGT) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{7}
}

func (x *DurationGT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationGTE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationGTE) Reset() {
	*x = DurationGTE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationGTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTE) ProtoMessage() {}

func (x *DurationGTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationGTE.ProtoReflect.Descriptor instead.
func (*DurationGTE) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{8}
}

func (x *DurationGTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationGTLT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationGTLT) Reset() {
	*x = DurationGTLT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationGTLT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTLT) ProtoMessage() {}

func (x *DurationGTLT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationGTLT.ProtoReflect.Descriptor instead.
func (*DurationGTLT) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{9}
}

func (x *DurationGTLT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationExLTGT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationExLTGT) Reset() {
	*x = DurationExLTGT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationExLTGT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationExLTGT) ProtoMessage() {}

func (x *DurationExLTGT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationExLTGT.ProtoReflect.Descriptor instead.
func (*DurationExLTGT) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{10}
}

func (x *DurationExLTGT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationGTELTE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationGTELTE) Reset() {
	*x = DurationGTELTE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationGTELTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTELTE) ProtoMessage() {}

func (x *DurationGTELTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationGTELTE.ProtoReflect.Descriptor instead.
func (*DurationGTELTE) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{11}
}

func (x *DurationGTELTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

type DurationExGTELTE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *durationpb.Duration `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DurationExGTELTE) Reset() {
	*x = DurationExGTELTE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationExGTELTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationExGTELTE) ProtoMessage() {}

func (x *DurationExGTELTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationExGTELTE.ProtoReflect.Descriptor instead.
func (*DurationExGTELTE) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{12}
}

func (x *DurationExGTELTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

// Regression for earlier bug where missing Duration field would short circuit
// evaluation in C++.
type DurationFieldWithOtherFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DurationVal *durationpb.Duration `protobuf:"bytes,1,opt,name=duration_val,json=durationVal,proto3" json:"duration_val,omitempty"`
	IntVal      int32                `protobuf:"varint,2,opt,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
}

func (x *DurationFieldWithOtherFields) Reset() {
	*x = DurationFieldWithOtherFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationFieldWithOtherFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationFieldWithOtherFields) ProtoMessage() {}

func (x *DurationFieldWithOtherFields) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationFieldWithOtherFields.ProtoReflect.Descriptor instead.
func (*DurationFieldWithOtherFields) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP(), []int{13}
}

func (x *DurationFieldWithOtherFields) GetDurationVal() *durationpb.Duration {
	if x != nil {
		return x.DurationVal
	}
	return nil
}

func (x *DurationFieldWithOtherFields) GetIntVal() int32 {
	if x != nil {
		return x.IntVal
	}
	return 0
}

var File_buf_validate_conformance_cases_wkt_duration_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x77, 0x6b, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3b, 0x0a, 0x0c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x47, 0x0a,
	0x10, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x33, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x48, 0x0a, 0x0d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xaa, 0x01, 0x04, 0x12, 0x02, 0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x4a, 0x0a, 0x0a, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x12, 0x3c,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0xaa, 0x01, 0x09, 0x3a, 0x02,
	0x08, 0x01, 0x3a, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x46, 0x0a, 0x0d,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x12, 0x35, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xba, 0x48, 0x05, 0xaa, 0x01, 0x02, 0x42, 0x00, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0a, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x54, 0x12, 0x35, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xba, 0x48, 0x05, 0xaa,
	0x01, 0x02, 0x1a, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x46, 0x0a, 0x0b, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x54, 0x45, 0x12, 0x37, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0xaa, 0x01, 0x04, 0x22, 0x02, 0x08, 0x01, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x46, 0x0a, 0x0a, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x54, 0x12,
	0x38, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xaa, 0x01, 0x05, 0x2a,
	0x03, 0x10, 0xe8, 0x07, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x48, 0x0a, 0x0b, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x54, 0x45, 0x12, 0x39, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0c, 0xba, 0x48, 0x09, 0xaa, 0x01, 0x06, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x49, 0x0a, 0x0c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x54, 0x4c, 0x54, 0x12, 0x39, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0xba, 0x48, 0x09,
	0xaa, 0x01, 0x06, 0x1a, 0x02, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x4b,
	0x0a, 0x0e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x4c, 0x54, 0x47, 0x54,
	0x12, 0x39, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0xba, 0x48, 0x09, 0xaa, 0x01, 0x06,
	0x1a, 0x00, 0x2a, 0x02, 0x08, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x4e, 0x0a, 0x0e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x54, 0x45, 0x4c, 0x54, 0x45, 0x12, 0x3c, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0xaa, 0x01, 0x09, 0x22, 0x03, 0x08,
	0x90, 0x1c, 0x32, 0x02, 0x08, 0x3c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x50, 0x0a, 0x10, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x47, 0x54, 0x45, 0x4c, 0x54, 0x45, 0x12,
	0x3c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0xaa, 0x01, 0x09, 0x22,
	0x02, 0x08, 0x3c, 0x32, 0x03, 0x08, 0x90, 0x1c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x8a, 0x01,
	0x0a, 0x1c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x48,
	0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xaa, 0x01, 0x04, 0x22, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02,
	0x20, 0x10, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x42, 0xa4, 0x02, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x42, 0x10, 0x57, 0x6b, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02,
	0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca,
	0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73,
	0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21,
	0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_wkt_duration_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_wkt_duration_proto_rawDescData = file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc
)

func file_buf_validate_conformance_cases_wkt_duration_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_wkt_duration_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_wkt_duration_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_wkt_duration_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_wkt_duration_proto_rawDescData
}

var file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_buf_validate_conformance_cases_wkt_duration_proto_goTypes = []interface{}{
	(*DurationNone)(nil),                 // 0: buf.validate.conformance.cases.DurationNone
	(*DurationRequired)(nil),             // 1: buf.validate.conformance.cases.DurationRequired
	(*DurationConst)(nil),                // 2: buf.validate.conformance.cases.DurationConst
	(*DurationIn)(nil),                   // 3: buf.validate.conformance.cases.DurationIn
	(*DurationNotIn)(nil),                // 4: buf.validate.conformance.cases.DurationNotIn
	(*DurationLT)(nil),                   // 5: buf.validate.conformance.cases.DurationLT
	(*DurationLTE)(nil),                  // 6: buf.validate.conformance.cases.DurationLTE
	(*DurationGT)(nil),                   // 7: buf.validate.conformance.cases.DurationGT
	(*DurationGTE)(nil),                  // 8: buf.validate.conformance.cases.DurationGTE
	(*DurationGTLT)(nil),                 // 9: buf.validate.conformance.cases.DurationGTLT
	(*DurationExLTGT)(nil),               // 10: buf.validate.conformance.cases.DurationExLTGT
	(*DurationGTELTE)(nil),               // 11: buf.validate.conformance.cases.DurationGTELTE
	(*DurationExGTELTE)(nil),             // 12: buf.validate.conformance.cases.DurationExGTELTE
	(*DurationFieldWithOtherFields)(nil), // 13: buf.validate.conformance.cases.DurationFieldWithOtherFields
	(*durationpb.Duration)(nil),          // 14: google.protobuf.Duration
}
var file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs = []int32{
	14, // 0: buf.validate.conformance.cases.DurationNone.val:type_name -> google.protobuf.Duration
	14, // 1: buf.validate.conformance.cases.DurationRequired.val:type_name -> google.protobuf.Duration
	14, // 2: buf.validate.conformance.cases.DurationConst.val:type_name -> google.protobuf.Duration
	14, // 3: buf.validate.conformance.cases.DurationIn.val:type_name -> google.protobuf.Duration
	14, // 4: buf.validate.conformance.cases.DurationNotIn.val:type_name -> google.protobuf.Duration
	14, // 5: buf.validate.conformance.cases.DurationLT.val:type_name -> google.protobuf.Duration
	14, // 6: buf.validate.conformance.cases.DurationLTE.val:type_name -> google.protobuf.Duration
	14, // 7: buf.validate.conformance.cases.DurationGT.val:type_name -> google.protobuf.Duration
	14, // 8: buf.validate.conformance.cases.DurationGTE.val:type_name -> google.protobuf.Duration
	14, // 9: buf.validate.conformance.cases.DurationGTLT.val:type_name -> google.protobuf.Duration
	14, // 10: buf.validate.conformance.cases.DurationExLTGT.val:type_name -> google.protobuf.Duration
	14, // 11: buf.validate.conformance.cases.DurationGTELTE.val:type_name -> google.protobuf.Duration
	14, // 12: buf.validate.conformance.cases.DurationExGTELTE.val:type_name -> google.protobuf.Duration
	14, // 13: buf.validate.conformance.cases.DurationFieldWithOtherFields.duration_val:type_name -> google.protobuf.Duration
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_wkt_duration_proto_init() }
func file_buf_validate_conformance_cases_wkt_duration_proto_init() {
	if File_buf_validate_conformance_cases_wkt_duration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationNone); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationRequired); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationConst); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationIn); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationNotIn); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationLT); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationLTE); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationGT); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationGTE); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationGTLT); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationExLTGT); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationGTELTE); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationExGTELTE); i {
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
		file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationFieldWithOtherFields); i {
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
			RawDescriptor: file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_wkt_duration_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_wkt_duration_proto = out.File
	file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc = nil
	file_buf_validate_conformance_cases_wkt_duration_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs = nil
}
