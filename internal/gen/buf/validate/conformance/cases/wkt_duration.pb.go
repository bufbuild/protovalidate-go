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
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/wkt_duration.proto

//go:build !protoopaque

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DurationNone struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationNone) Reset() {
	*x = DurationNone{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationNone) ProtoMessage() {}

func (x *DurationNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationNone) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationNone) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationNone) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationNone) ClearVal() {
	x.Val = nil
}

type DurationNone_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationNone_builder) Build() *DurationNone {
	m0 := &DurationNone{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationRequired struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationRequired) Reset() {
	*x = DurationRequired{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationRequired) ProtoMessage() {}

func (x *DurationRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationRequired) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationRequired) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationRequired) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationRequired) ClearVal() {
	x.Val = nil
}

type DurationRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationRequired_builder) Build() *DurationRequired {
	m0 := &DurationRequired{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationConst struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationConst) Reset() {
	*x = DurationConst{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationConst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationConst) ProtoMessage() {}

func (x *DurationConst) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationConst) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationConst) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationConst) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationConst) ClearVal() {
	x.Val = nil
}

type DurationConst_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationConst_builder) Build() *DurationConst {
	m0 := &DurationConst{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationIn struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationIn) Reset() {
	*x = DurationIn{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationIn) ProtoMessage() {}

func (x *DurationIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationIn) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationIn) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationIn) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationIn) ClearVal() {
	x.Val = nil
}

type DurationIn_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationIn_builder) Build() *DurationIn {
	m0 := &DurationIn{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationNotIn struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationNotIn) Reset() {
	*x = DurationNotIn{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationNotIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationNotIn) ProtoMessage() {}

func (x *DurationNotIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationNotIn) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationNotIn) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationNotIn) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationNotIn) ClearVal() {
	x.Val = nil
}

type DurationNotIn_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationNotIn_builder) Build() *DurationNotIn {
	m0 := &DurationNotIn{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationLT struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationLT) Reset() {
	*x = DurationLT{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationLT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationLT) ProtoMessage() {}

func (x *DurationLT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationLT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationLT) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationLT) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationLT) ClearVal() {
	x.Val = nil
}

type DurationLT_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationLT_builder) Build() *DurationLT {
	m0 := &DurationLT{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationLTE struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationLTE) Reset() {
	*x = DurationLTE{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationLTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationLTE) ProtoMessage() {}

func (x *DurationLTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationLTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationLTE) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationLTE) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationLTE) ClearVal() {
	x.Val = nil
}

type DurationLTE_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationLTE_builder) Build() *DurationLTE {
	m0 := &DurationLTE{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationGT struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationGT) Reset() {
	*x = DurationGT{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationGT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGT) ProtoMessage() {}

func (x *DurationGT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationGT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationGT) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationGT) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationGT) ClearVal() {
	x.Val = nil
}

type DurationGT_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationGT_builder) Build() *DurationGT {
	m0 := &DurationGT{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationGTE struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationGTE) Reset() {
	*x = DurationGTE{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationGTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTE) ProtoMessage() {}

func (x *DurationGTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationGTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationGTE) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationGTE) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationGTE) ClearVal() {
	x.Val = nil
}

type DurationGTE_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationGTE_builder) Build() *DurationGTE {
	m0 := &DurationGTE{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationGTLT struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationGTLT) Reset() {
	*x = DurationGTLT{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationGTLT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTLT) ProtoMessage() {}

func (x *DurationGTLT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationGTLT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationGTLT) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationGTLT) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationGTLT) ClearVal() {
	x.Val = nil
}

type DurationGTLT_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationGTLT_builder) Build() *DurationGTLT {
	m0 := &DurationGTLT{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationExLTGT struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationExLTGT) Reset() {
	*x = DurationExLTGT{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationExLTGT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationExLTGT) ProtoMessage() {}

func (x *DurationExLTGT) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationExLTGT) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationExLTGT) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationExLTGT) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationExLTGT) ClearVal() {
	x.Val = nil
}

type DurationExLTGT_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationExLTGT_builder) Build() *DurationExLTGT {
	m0 := &DurationExLTGT{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationGTELTE struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationGTELTE) Reset() {
	*x = DurationGTELTE{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationGTELTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationGTELTE) ProtoMessage() {}

func (x *DurationGTELTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationGTELTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationGTELTE) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationGTELTE) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationGTELTE) ClearVal() {
	x.Val = nil
}

type DurationGTELTE_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationGTELTE_builder) Build() *DurationGTELTE {
	m0 := &DurationGTELTE{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

type DurationExGTELTE struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationExGTELTE) Reset() {
	*x = DurationExGTELTE{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationExGTELTE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationExGTELTE) ProtoMessage() {}

func (x *DurationExGTELTE) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationExGTELTE) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationExGTELTE) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationExGTELTE) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationExGTELTE) ClearVal() {
	x.Val = nil
}

type DurationExGTELTE_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationExGTELTE_builder) Build() *DurationExGTELTE {
	m0 := &DurationExGTELTE{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

// Regression for earlier bug where missing Duration field would short circuit
// evaluation in C++.
type DurationFieldWithOtherFields struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	DurationVal   *durationpb.Duration   `protobuf:"bytes,1,opt,name=duration_val,json=durationVal,proto3" json:"duration_val,omitempty"`
	IntVal        int32                  `protobuf:"varint,2,opt,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationFieldWithOtherFields) Reset() {
	*x = DurationFieldWithOtherFields{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationFieldWithOtherFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationFieldWithOtherFields) ProtoMessage() {}

func (x *DurationFieldWithOtherFields) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
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

func (x *DurationFieldWithOtherFields) SetDurationVal(v *durationpb.Duration) {
	x.DurationVal = v
}

func (x *DurationFieldWithOtherFields) SetIntVal(v int32) {
	x.IntVal = v
}

func (x *DurationFieldWithOtherFields) HasDurationVal() bool {
	if x == nil {
		return false
	}
	return x.DurationVal != nil
}

func (x *DurationFieldWithOtherFields) ClearDurationVal() {
	x.DurationVal = nil
}

type DurationFieldWithOtherFields_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DurationVal *durationpb.Duration
	IntVal      int32
}

func (b0 DurationFieldWithOtherFields_builder) Build() *DurationFieldWithOtherFields {
	m0 := &DurationFieldWithOtherFields{}
	b, x := &b0, m0
	_, _ = b, x
	x.DurationVal = b.DurationVal
	x.IntVal = b.IntVal
	return m0
}

type DurationExample struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Val           *durationpb.Duration   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DurationExample) Reset() {
	*x = DurationExample{}
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DurationExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationExample) ProtoMessage() {}

func (x *DurationExample) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DurationExample) GetVal() *durationpb.Duration {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *DurationExample) SetVal(v *durationpb.Duration) {
	x.Val = v
}

func (x *DurationExample) HasVal() bool {
	if x == nil {
		return false
	}
	return x.Val != nil
}

func (x *DurationExample) ClearVal() {
	x.Val = nil
}

type DurationExample_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Val *durationpb.Duration
}

func (b0 DurationExample_builder) Build() *DurationExample {
	m0 := &DurationExample{}
	b, x := &b0, m0
	_, _ = b, x
	x.Val = b.Val
	return m0
}

var File_buf_validate_conformance_cases_wkt_duration_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc = string([]byte{
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
	0x20, 0x10, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x22, 0x4a, 0x0a, 0x0f, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x37, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xaa, 0x01, 0x04, 0x4a, 0x02, 0x08,
	0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0xa4, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42, 0x10, 0x57,
	0x6b, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75,
	0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42,
	0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a,
	0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_buf_validate_conformance_cases_wkt_duration_proto_goTypes = []any{
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
	(*DurationExample)(nil),              // 14: buf.validate.conformance.cases.DurationExample
	(*durationpb.Duration)(nil),          // 15: google.protobuf.Duration
}
var file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs = []int32{
	15, // 0: buf.validate.conformance.cases.DurationNone.val:type_name -> google.protobuf.Duration
	15, // 1: buf.validate.conformance.cases.DurationRequired.val:type_name -> google.protobuf.Duration
	15, // 2: buf.validate.conformance.cases.DurationConst.val:type_name -> google.protobuf.Duration
	15, // 3: buf.validate.conformance.cases.DurationIn.val:type_name -> google.protobuf.Duration
	15, // 4: buf.validate.conformance.cases.DurationNotIn.val:type_name -> google.protobuf.Duration
	15, // 5: buf.validate.conformance.cases.DurationLT.val:type_name -> google.protobuf.Duration
	15, // 6: buf.validate.conformance.cases.DurationLTE.val:type_name -> google.protobuf.Duration
	15, // 7: buf.validate.conformance.cases.DurationGT.val:type_name -> google.protobuf.Duration
	15, // 8: buf.validate.conformance.cases.DurationGTE.val:type_name -> google.protobuf.Duration
	15, // 9: buf.validate.conformance.cases.DurationGTLT.val:type_name -> google.protobuf.Duration
	15, // 10: buf.validate.conformance.cases.DurationExLTGT.val:type_name -> google.protobuf.Duration
	15, // 11: buf.validate.conformance.cases.DurationGTELTE.val:type_name -> google.protobuf.Duration
	15, // 12: buf.validate.conformance.cases.DurationExGTELTE.val:type_name -> google.protobuf.Duration
	15, // 13: buf.validate.conformance.cases.DurationFieldWithOtherFields.duration_val:type_name -> google.protobuf.Duration
	15, // 14: buf.validate.conformance.cases.DurationExample.val:type_name -> google.protobuf.Duration
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_wkt_duration_proto_init() }
func file_buf_validate_conformance_cases_wkt_duration_proto_init() {
	if File_buf_validate_conformance_cases_wkt_duration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc), len(file_buf_validate_conformance_cases_wkt_duration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_wkt_duration_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_wkt_duration_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_wkt_duration_proto = out.File
	file_buf_validate_conformance_cases_wkt_duration_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_duration_proto_depIdxs = nil
}
