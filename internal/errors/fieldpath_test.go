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

package errors

import (
	"testing"

	pb "github.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestGetFieldValue(t *testing.T) {
	t.Parallel()

	type (
		Simple   = pb.FieldPathSimple
		Nested   = pb.FieldPathNested
		Repeated = pb.FieldPathRepeated
		Maps     = pb.FieldPathMaps
	)

	withExtension := &Simple{}
	proto.SetExtension(withExtension, pb.E_Ext, []uint64{10})

	testGetFieldValueMatch(t, uint64(64), &Simple{Val: 64}, "val")
	testGetFieldValueMatch(t, uint64(0), (*Simple)(nil), "val")
	testGetFieldValueMatch(t, float64(1.0), &Simple{Nested: &Nested{Val: 1.0}}, "nested.val")
	testGetFieldValueMatch(t, float64(0.0), (*Simple)(nil), "nested.val")
	testGetFieldValueMatch(t, float64(0.0), &Simple{Nested: nil}, "nested.val")
	testGetFieldValueMatch(t, int32(2), &Repeated{Val: []int32{1, 2, 3}}, "val[1]")
	testGetFieldValueMatch(t, uint64(32), &Repeated{Msg: []*Simple{{Val: 64}, {Val: 32}}}, "msg[1].val")
	testGetFieldValueMatch(t, int32(4), &Maps{Int32Int32Map: map[int32]int32{1: 2, 2: 4, 4: 8}}, "int32_int32_map[2]")
	testGetFieldValueMatch(t, uint64(32), &Maps{Int32Map: map[int32]*Simple{1: {Val: 32}}}, "int32_map[1].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Int64Map: map[int64]*Simple{2: {Val: 64}}}, "int64_map[2].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Sint32Map: map[int32]*Simple{2: {Val: 64}}}, "sint32_map[2].val")
	testGetFieldValueMatch(t, uint64(32), &Maps{Sint64Map: map[int64]*Simple{1: {Val: 32}}}, "sint64_map[1].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Sfixed32Map: map[int32]*Simple{2: {Val: 64}}}, "sfixed32_map[2].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Sfixed64Map: map[int64]*Simple{2: {Val: 64}}}, "sfixed64_map[2].val")
	testGetFieldValueMatch(t, uint64(32), &Maps{Uint32Map: map[uint32]*Simple{1: {Val: 32}}}, "uint32_map[1].val")
	testGetFieldValueMatch(t, uint64(32), &Maps{Uint64Map: map[uint64]*Simple{1: {Val: 32}}}, "uint64_map[1].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Fixed32Map: map[uint32]*Simple{2: {Val: 64}}}, "fixed32_map[2].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{Fixed64Map: map[uint64]*Simple{2: {Val: 64}}}, "fixed64_map[2].val")
	testGetFieldValueMatch(t, uint64(64), &Maps{StringMap: map[string]*Simple{"a": {Val: 64}}}, `string_map["a"].val`)
	testGetFieldValueMatch(t, uint64(64), &Maps{StringMap: map[string]*Simple{`".[]][`: {Val: 64}}}, `string_map["\".[]]["].val`)
	testGetFieldValueMatch(t, uint64(1), &Maps{BoolMap: map[bool]*Simple{true: {Val: 1}}}, "bool_map[true].val")
	testGetFieldValueMatch(t, uint64(0), &Maps{Int32Map: map[int32]*Simple{1: nil}}, "int32_map[1].val")
	testGetFieldValueMatch(t, uint64(10), withExtension, "[tests.example.v1.ext][0]")
	testGetFieldValueMatch(t, uint64(10), &Repeated{Msg: []*Simple{withExtension}}, "msg[0].[tests.example.v1.ext][0]")

	testGetFieldValueError(t, "field nofield not found", &Simple{Val: 1}, "nofield")
	testGetFieldValueError(t, "field nofield not found", &Simple{Val: 1}, "val.nofield")
	testGetFieldValueError(t, "unexpected subscript on field val", &Simple{Val: 1}, "val[1]")
	testGetFieldValueError(t, "empty field name", &Simple{Val: 1}, "nested.")
	testGetFieldValueError(t, "empty field name", &Simple{Val: 1}, "nested..b")
	testGetFieldValueError(t, "field ] not found", &Simple{Val: 1}, "]")
	testGetFieldValueError(t, "field a] not found", &Simple{Val: 1}, "a]")
	testGetFieldValueError(t, "field a[] not found", &Simple{Val: 1}, "a[]")
	testGetFieldValueError(t, "field a not found", &Simple{Val: 1}, "a[1]")
	testGetFieldValueError(t, "invalid list index: #", &Repeated{}, "val[#]")
	testGetFieldValueError(t, "index 1 out of bounds of field val", &Repeated{}, "val[1]")
	testGetFieldValueError(t, "index 1 out of bounds of field val", &Repeated{Val: []int32{1}}, "val[1]")
	testGetFieldValueError(t, "missing subscript on field msg", &Repeated{}, "msg.val")
	testGetFieldValueError(t, "key 2 not present on field int32_map", &Maps{Int32Map: map[int32]*Simple{1: {Val: 1}}}, "int32_map[2].val")
	testGetFieldValueError(t, `missing subscript on field string_map`, &Maps{}, `string_map["not a string]`)
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "int32_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "int64_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "sint32_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "sint64_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "sfixed32_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "sfixed64_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "uint32_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "uint64_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "fixed32_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "fixed64_map[#]")
	testGetFieldValueError(t, `invalid map key: #`, &Maps{}, "bool_map[#]")
}

func BenchmarkGetFieldValue(b *testing.B) {
	simpleMessage := &pb.FieldPathSimple{Val: 64}
	nestedMessage := &pb.FieldPathSimple{
		Nested: &pb.FieldPathNested{Val: 1},
	}
	repeatedMessage := &pb.FieldPathRepeated{
		Msg: []*pb.FieldPathSimple{{Val: 1}},
	}
	mapMessage := &pb.FieldPathMaps{
		StringMap: map[string]*pb.FieldPathSimple{"abc": {Val: 1}},
	}

	b.Run("Simple", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, _, err := getFieldValue(protoregistry.GlobalTypes, simpleMessage, "val")
				require.NoError(b, err)
			}
		})
	})

	b.Run("Nested", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, _, err := getFieldValue(protoregistry.GlobalTypes, nestedMessage, "nested.val")
				require.NoError(b, err)
			}
		})
	})

	b.Run("Repeated", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, _, err := getFieldValue(protoregistry.GlobalTypes, repeatedMessage, "msg[0].val")
				require.NoError(b, err)
			}
		})
	})

	b.Run("Map", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, _, err := getFieldValue(protoregistry.GlobalTypes, mapMessage, `string_map["abc"].val`)
				require.NoError(b, err)
			}
		})
	})

	b.Run("Error", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_, _, err := getFieldValue(protoregistry.GlobalTypes, mapMessage, `string_map["z"].val`)
				require.Error(b, err)
			}
		})
	})
}

func testGetFieldValueMatch(t *testing.T, expected any, message proto.Message, path string) {
	t.Helper()

	val, _, err := getFieldValue(protoregistry.GlobalTypes, message, path)
	require.NoError(t, err)
	assert.Equal(t, expected, val.Interface())
}

func testGetFieldValueError(t *testing.T, errString string, message proto.Message, path string) {
	t.Helper()

	_, _, err := getFieldValue(protoregistry.GlobalTypes, message, path)
	assert.EqualError(t, err, errString)
}
