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

package constraints

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go/celext"
	"github.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func getFieldDesc(t *testing.T, msg proto.Message, fld protoreflect.Name) protoreflect.FieldDescriptor {
	t.Helper()
	desc := msg.ProtoReflect().Descriptor().Fields().ByName(fld)
	require.NotNil(t, desc)
	return desc
}

func TestCache_BuildStandardConstraints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		desc     protoreflect.FieldDescriptor
		cons     *validate.FieldConstraints
		forItems bool
		exCt     int
		exErr    bool
	}{
		{
			name: "no constraints",
			desc: getFieldDesc(t, &cases.FloatNone{}, "val"),
			cons: &validate.FieldConstraints{},
			exCt: 0,
		},
		{
			name: "nil constraints",
			desc: getFieldDesc(t, &cases.FloatNone{}, "val"),
			cons: nil,
			exCt: 0,
		},
		{
			name: "list constraints",
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			cons: &validate.FieldConstraints{Type: &validate.FieldConstraints_Repeated{Repeated: &validate.RepeatedRules{
				MinItems: proto.Uint64(3),
			}}},
			exCt: 1,
		},
		{
			name: "list item constraints",
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			cons: &validate.FieldConstraints{Type: &validate.FieldConstraints_Int64{Int64: &validate.Int64Rules{
				NotIn: []int64{123},
				Const: proto.Int64(456),
			}}},
			forItems: true,
			exCt:     2,
		},
		{
			name: "map constraints",
			desc: getFieldDesc(t, &cases.MapNone{}, "val"),
			cons: &validate.FieldConstraints{Type: &validate.FieldConstraints_Map{Map: &validate.MapRules{
				MinPairs: proto.Uint64(2),
			}}},
			exCt: 1,
		},
		{
			name: "mismatch constraints",
			desc: getFieldDesc(t, &cases.AnyNone{}, "val"),
			cons: &validate.FieldConstraints{Type: &validate.FieldConstraints_Float{Float: &validate.FloatRules{
				Const: proto.Float32(1.23),
			}}},
			exErr: true,
		},
	}

	env, err := celext.DefaultEnv(false)
	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			require.NoError(t, err)
			c := NewCache()

			set, err := c.Build(env, test.desc, test.cons, protoregistry.GlobalTypes, test.forItems)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Len(t, set, test.exCt)
			}
		})
	}
}

func TestCache_LoadOrCompileStandardConstraint(t *testing.T) {
	t.Parallel()

	env, err := celext.DefaultEnv(false)
	require.NoError(t, err)

	msg := &cases.FloatIn{}
	desc := getFieldDesc(t, msg, "val")
	require.NotNil(t, desc)

	cache := NewCache()
	_, ok := cache.cache[desc]
	assert.False(t, ok)

	asts, err := cache.loadOrCompileStandardConstraint(env, desc)
	require.NoError(t, err)
	assert.NotNil(t, asts)

	cached, ok := cache.cache[desc]
	assert.True(t, ok)
	assert.Equal(t, cached, asts)

	asts, err = cache.loadOrCompileStandardConstraint(env, desc)
	require.NoError(t, err)
	assert.Equal(t, cached, asts)
}

func TestCache_GetExpectedConstraintDescriptor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     protoreflect.FieldDescriptor
		forItems bool
		ex       protoreflect.FieldDescriptor
	}{
		{
			desc: getFieldDesc(t, &cases.MapNone{}, "val"),
			ex:   mapFieldConstraintsDesc,
		},
		{
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			ex:   repeatedFieldConstraintsDesc,
		},
		{
			desc:     getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			forItems: true,
			ex:       expectedStandardConstraints[protoreflect.Int64Kind],
		},
		{
			desc: getFieldDesc(t, &cases.AnyNone{}, "val"),
			ex:   expectedWKTConstraints["google.protobuf.Any"],
		},
		{
			desc: getFieldDesc(t, &cases.TimestampNone{}, "val"),
			ex:   expectedWKTConstraints["google.protobuf.Timestamp"],
		},
		{
			desc: getFieldDesc(t, &cases.DurationNone{}, "val"),
			ex:   expectedWKTConstraints["google.protobuf.Duration"],
		},
		{
			desc: getFieldDesc(t, &cases.StringNone{}, "val"),
			ex:   expectedStandardConstraints[protoreflect.StringKind],
		},
		{
			desc: getFieldDesc(t, &cases.MessageNone{}, "val"),
			ex:   nil,
		},
	}

	c := NewCache()
	for _, tc := range tests {
		test := tc
		t.Run(string(test.desc.FullName()), func(t *testing.T) {
			t.Parallel()
			out, ok := c.getExpectedConstraintDescriptor(test.desc, test.forItems)
			if test.ex != nil {
				assert.True(t, ok)
				assert.Equal(t, test.ex.FullName(), out.FullName())
			} else {
				assert.False(t, ok)
			}
		})
	}
}
