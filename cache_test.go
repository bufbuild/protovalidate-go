// Copyright 2023-2025 Buf Technologies, Inc.
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

package protovalidate

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pvcel "buf.build/go/protovalidate/cel"
	"buf.build/go/protovalidate/internal/gen/buf/validate/conformance/cases"
	"github.com/google/cel-go/cel"
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

func TestCache_BuildStandardRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		desc     protoreflect.FieldDescriptor
		cons     *validate.FieldRules
		forItems bool
		exCt     int
		exErr    bool
	}{
		{
			name: "no rules",
			desc: getFieldDesc(t, &cases.FloatNone{}, "val"),
			cons: &validate.FieldRules{},
			exCt: 0,
		},
		{
			name: "nil rules",
			desc: getFieldDesc(t, &cases.FloatNone{}, "val"),
			cons: nil,
			exCt: 0,
		},
		{
			name: "list rules",
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			cons: validate.FieldRules_builder{
				Repeated: validate.RepeatedRules_builder{
					MinItems: proto.Uint64(3),
				}.Build(),
			}.Build(),
			exCt: 1,
		},
		{
			name: "list item rules",
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			cons: validate.FieldRules_builder{
				Int64: validate.Int64Rules_builder{
					NotIn: []int64{123},
					Const: proto.Int64(456),
				}.Build(),
			}.Build(),
			forItems: true,
			exCt:     2,
		},
		{
			name: "map rules",
			desc: getFieldDesc(t, &cases.MapNone{}, "val"),
			cons: validate.FieldRules_builder{
				Map: validate.MapRules_builder{
					MinPairs: proto.Uint64(2),
				}.Build(),
			}.Build(),
			exCt: 1,
		},
		{
			name: "mismatch rules",
			desc: getFieldDesc(t, &cases.AnyNone{}, "val"),
			cons: validate.FieldRules_builder{
				Float: validate.FloatRules_builder{
					Const: proto.Float32(1.23),
				}.Build(),
			}.Build(),
			exErr: true,
		},
	}

	env, err := cel.NewEnv(cel.Lib(pvcel.NewLibrary()))
	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			require.NoError(t, err)
			c := newCache()

			set, err := c.Build(env, test.desc, test.cons, protoregistry.GlobalTypes, false, test.forItems, nil)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Len(t, set, test.exCt)
			}
		})
	}
}

func TestCache_LoadOrCompileStandardRule(t *testing.T) {
	t.Parallel()

	env, err := cel.NewEnv(cel.Lib(pvcel.NewLibrary()))
	require.NoError(t, err)

	rules := &validate.FieldRules{}
	oneOfDesc := rules.ProtoReflect().Descriptor().Oneofs().ByName("type").Fields().ByName("float")
	msg := &cases.FloatIn{}
	desc := getFieldDesc(t, msg, "val")
	require.NotNil(t, desc)

	cache := newCache()
	_, ok := cache.cache[desc]
	assert.False(t, ok)

	asts, err := cache.loadOrCompileStandardRule(env, oneOfDesc, desc)
	require.NoError(t, err)
	assert.Nil(t, asts)

	cached, ok := cache.cache[desc]
	assert.True(t, ok)
	assert.Equal(t, cached, asts)

	asts, err = cache.loadOrCompileStandardRule(env, oneOfDesc, desc)
	require.NoError(t, err)
	assert.Equal(t, cached, asts)
}

func TestCache_GetExpectedRuleDescriptor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     protoreflect.FieldDescriptor
		forItems bool
		ex       protoreflect.FieldDescriptor
	}{
		{
			desc: getFieldDesc(t, &cases.MapNone{}, "val"),
			ex:   mapFieldRulesDesc,
		},
		{
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			ex:   repeatedFieldRulesDesc,
		},
		{
			desc:     getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			forItems: true,
			ex:       expectedStandardRules[protoreflect.Int64Kind],
		},
		{
			desc: getFieldDesc(t, &cases.AnyNone{}, "val"),
			ex:   expectedWKTRules["google.protobuf.Any"],
		},
		{
			desc: getFieldDesc(t, &cases.TimestampNone{}, "val"),
			ex:   expectedWKTRules["google.protobuf.Timestamp"],
		},
		{
			desc: getFieldDesc(t, &cases.DurationNone{}, "val"),
			ex:   expectedWKTRules["google.protobuf.Duration"],
		},
		{
			desc: getFieldDesc(t, &cases.FieldMaskNone{}, "val"),
			ex:   expectedWKTRules["google.protobuf.FieldMask"],
		},
		{
			desc: getFieldDesc(t, &cases.StringNone{}, "val"),
			ex:   expectedStandardRules[protoreflect.StringKind],
		},
		{
			desc: getFieldDesc(t, &cases.MessageNone{}, "val"),
			ex:   nil,
		},
	}

	c := newCache()
	for _, tc := range tests {
		test := tc
		t.Run(string(test.desc.FullName()), func(t *testing.T) {
			t.Parallel()
			out, ok := c.getExpectedRuleDescriptor(test.desc, test.forItems)
			if test.ex != nil {
				assert.True(t, ok)
				assert.Equal(t, test.ex.FullName(), out.FullName())
			} else {
				assert.False(t, ok)
			}
		})
	}
}
