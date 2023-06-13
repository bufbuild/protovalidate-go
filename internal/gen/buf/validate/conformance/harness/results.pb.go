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
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/validate/conformance/harness/results.proto

package harness

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResultOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuiteFilter string `protobuf:"bytes,1,opt,name=suite_filter,json=suiteFilter,proto3" json:"suite_filter,omitempty"`
	CaseFilter  string `protobuf:"bytes,2,opt,name=case_filter,json=caseFilter,proto3" json:"case_filter,omitempty"`
	Verbose     bool   `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Strict      bool   `protobuf:"varint,4,opt,name=strict,proto3" json:"strict,omitempty"`
	// If the violation message must be an exact match.
	StrictMessage bool `protobuf:"varint,5,opt,name=strict_message,json=strictMessage,proto3" json:"strict_message,omitempty"`
	// If the distinciton between runtime and compile time errors must be exact.
	StrictError bool `protobuf:"varint,6,opt,name=strict_error,json=strictError,proto3" json:"strict_error,omitempty"`
}

func (x *ResultOptions) Reset() {
	*x = ResultOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultOptions) ProtoMessage() {}

func (x *ResultOptions) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultOptions.ProtoReflect.Descriptor instead.
func (*ResultOptions) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_results_proto_rawDescGZIP(), []int{0}
}

func (x *ResultOptions) GetSuiteFilter() string {
	if x != nil {
		return x.SuiteFilter
	}
	return ""
}

func (x *ResultOptions) GetCaseFilter() string {
	if x != nil {
		return x.CaseFilter
	}
	return ""
}

func (x *ResultOptions) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

func (x *ResultOptions) GetStrict() bool {
	if x != nil {
		return x.Strict
	}
	return false
}

func (x *ResultOptions) GetStrictMessage() bool {
	if x != nil {
		return x.StrictMessage
	}
	return false
}

func (x *ResultOptions) GetStrictError() bool {
	if x != nil {
		return x.StrictError
	}
	return false
}

type ResultSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successes int32           `protobuf:"varint,1,opt,name=successes,proto3" json:"successes,omitempty"`
	Failures  int32           `protobuf:"varint,2,opt,name=failures,proto3" json:"failures,omitempty"`
	Suites    []*SuiteResults `protobuf:"bytes,3,rep,name=suites,proto3" json:"suites,omitempty"`
	Options   *ResultOptions  `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ResultSet) Reset() {
	*x = ResultSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSet) ProtoMessage() {}

func (x *ResultSet) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSet.ProtoReflect.Descriptor instead.
func (*ResultSet) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_results_proto_rawDescGZIP(), []int{1}
}

func (x *ResultSet) GetSuccesses() int32 {
	if x != nil {
		return x.Successes
	}
	return 0
}

func (x *ResultSet) GetFailures() int32 {
	if x != nil {
		return x.Failures
	}
	return 0
}

func (x *ResultSet) GetSuites() []*SuiteResults {
	if x != nil {
		return x.Suites
	}
	return nil
}

func (x *ResultSet) GetOptions() *ResultOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type SuiteResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Successes int32                           `protobuf:"varint,2,opt,name=successes,proto3" json:"successes,omitempty"`
	Failures  int32                           `protobuf:"varint,3,opt,name=failures,proto3" json:"failures,omitempty"`
	Cases     []*CaseResult                   `protobuf:"bytes,4,rep,name=cases,proto3" json:"cases,omitempty"`
	Fdset     *descriptorpb.FileDescriptorSet `protobuf:"bytes,5,opt,name=fdset,proto3" json:"fdset,omitempty"`
}

func (x *SuiteResults) Reset() {
	*x = SuiteResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuiteResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuiteResults) ProtoMessage() {}

func (x *SuiteResults) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuiteResults.ProtoReflect.Descriptor instead.
func (*SuiteResults) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_results_proto_rawDescGZIP(), []int{2}
}

func (x *SuiteResults) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuiteResults) GetSuccesses() int32 {
	if x != nil {
		return x.Successes
	}
	return 0
}

func (x *SuiteResults) GetFailures() int32 {
	if x != nil {
		return x.Failures
	}
	return 0
}

func (x *SuiteResults) GetCases() []*CaseResult {
	if x != nil {
		return x.Cases
	}
	return nil
}

func (x *SuiteResults) GetFdset() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.Fdset
	}
	return nil
}

type CaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Success bool        `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Wanted  *TestResult `protobuf:"bytes,3,opt,name=wanted,proto3" json:"wanted,omitempty"`
	Got     *TestResult `protobuf:"bytes,4,opt,name=got,proto3" json:"got,omitempty"`
	Input   *anypb.Any  `protobuf:"bytes,5,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *CaseResult) Reset() {
	*x = CaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaseResult) ProtoMessage() {}

func (x *CaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_results_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaseResult.ProtoReflect.Descriptor instead.
func (*CaseResult) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_results_proto_rawDescGZIP(), []int{3}
}

func (x *CaseResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CaseResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CaseResult) GetWanted() *TestResult {
	if x != nil {
		return x.Wanted
	}
	return nil
}

func (x *CaseResult) GetGot() *TestResult {
	if x != nil {
		return x.Got
	}
	return nil
}

func (x *CaseResult) GetInput() *anypb.Any {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_buf_validate_conformance_harness_results_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_harness_results_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x1a, 0x2e, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x68, 0x61, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcf, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0xd8, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x75, 0x69,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x06, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x73, 0x12, 0x49, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61,
	0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xda, 0x01, 0x0a,
	0x0c, 0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x38, 0x0a, 0x05, 0x66, 0x64, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x74, 0x52, 0x05, 0x66, 0x64, 0x73, 0x65, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x0a, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x03,
	0x67, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x03, 0x67, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0xac, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73,
	0x73, 0x42, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x68, 0x61,
	0x72, 0x6e, 0x65, 0x73, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x48, 0xaa, 0x02, 0x20, 0x42,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0xca,
	0x02, 0x20, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x48, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0xe2, 0x02, 0x2c, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x48, 0x61,
	0x72, 0x6e, 0x65, 0x73, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x23, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a,
	0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_harness_results_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_harness_results_proto_rawDescData = file_buf_validate_conformance_harness_results_proto_rawDesc
)

func file_buf_validate_conformance_harness_results_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_harness_results_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_harness_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_harness_results_proto_rawDescData)
	})
	return file_buf_validate_conformance_harness_results_proto_rawDescData
}

var file_buf_validate_conformance_harness_results_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_validate_conformance_harness_results_proto_goTypes = []interface{}{
	(*ResultOptions)(nil),                  // 0: buf.validate.conformance.harness.ResultOptions
	(*ResultSet)(nil),                      // 1: buf.validate.conformance.harness.ResultSet
	(*SuiteResults)(nil),                   // 2: buf.validate.conformance.harness.SuiteResults
	(*CaseResult)(nil),                     // 3: buf.validate.conformance.harness.CaseResult
	(*descriptorpb.FileDescriptorSet)(nil), // 4: google.protobuf.FileDescriptorSet
	(*TestResult)(nil),                     // 5: buf.validate.conformance.harness.TestResult
	(*anypb.Any)(nil),                      // 6: google.protobuf.Any
}
var file_buf_validate_conformance_harness_results_proto_depIdxs = []int32{
	2, // 0: buf.validate.conformance.harness.ResultSet.suites:type_name -> buf.validate.conformance.harness.SuiteResults
	0, // 1: buf.validate.conformance.harness.ResultSet.options:type_name -> buf.validate.conformance.harness.ResultOptions
	3, // 2: buf.validate.conformance.harness.SuiteResults.cases:type_name -> buf.validate.conformance.harness.CaseResult
	4, // 3: buf.validate.conformance.harness.SuiteResults.fdset:type_name -> google.protobuf.FileDescriptorSet
	5, // 4: buf.validate.conformance.harness.CaseResult.wanted:type_name -> buf.validate.conformance.harness.TestResult
	5, // 5: buf.validate.conformance.harness.CaseResult.got:type_name -> buf.validate.conformance.harness.TestResult
	6, // 6: buf.validate.conformance.harness.CaseResult.input:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_harness_results_proto_init() }
func file_buf_validate_conformance_harness_results_proto_init() {
	if File_buf_validate_conformance_harness_results_proto != nil {
		return
	}
	file_buf_validate_conformance_harness_harness_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_harness_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultOptions); i {
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
		file_buf_validate_conformance_harness_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSet); i {
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
		file_buf_validate_conformance_harness_results_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuiteResults); i {
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
		file_buf_validate_conformance_harness_results_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaseResult); i {
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
			RawDescriptor: file_buf_validate_conformance_harness_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_harness_results_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_harness_results_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_harness_results_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_harness_results_proto = out.File
	file_buf_validate_conformance_harness_results_proto_rawDesc = nil
	file_buf_validate_conformance_harness_results_proto_goTypes = nil
	file_buf_validate_conformance_harness_results_proto_depIdxs = nil
}
