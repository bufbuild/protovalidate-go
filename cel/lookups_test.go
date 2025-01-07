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

package cel

import (
	"testing"

	"github.com/bufbuild/protovalidate-go/internal/gen/buf/validate/conformance/cases"
	"github.com/google/cel-go/cel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestCache_GetCELType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     protoreflect.FieldDescriptor
		generic  bool
		forItems bool
		ex       *cel.Type
	}{
		{
			desc: getFieldDesc(t, &cases.MapNone{}, "val"),
			ex:   cel.MapType(cel.UintType, cel.BoolType),
		},
		{
			desc:    getFieldDesc(t, &cases.MapNone{}, "val"),
			generic: true,
			ex:      cel.MapType(cel.DynType, cel.DynType),
		},
		{
			desc: getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			ex:   cel.ListType(cel.IntType),
		},
		{
			desc:    getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			generic: true,
			ex:      cel.ListType(cel.DynType),
		},
		{
			desc:     getFieldDesc(t, &cases.RepeatedNone{}, "val"),
			forItems: true,
			ex:       cel.IntType,
		},
		{
			desc: getFieldDesc(t, &cases.AnyNone{}, "val"),
			ex:   cel.AnyType,
		},
		{
			desc: getFieldDesc(t, &cases.DurationNone{}, "val"),
			ex:   cel.DurationType,
		},
		{
			desc: getFieldDesc(t, &cases.TimestampNone{}, "val"),
			ex:   cel.TimestampType,
		},
		{
			desc: getFieldDesc(t, &cases.MessageNone{}, "val"),
			ex:   cel.ObjectType(string(((&cases.MessageNone{}).GetVal()).ProtoReflect().Descriptor().FullName())),
		},
		{
			desc: getFieldDesc(t, &cases.Int32None{}, "val"),
			ex:   cel.IntType,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(string(test.desc.FullName()), func(t *testing.T) {
			t.Parallel()
			typ := ProtoFieldToCELType(test.desc, test.generic, test.forItems)
			assert.Equal(t, test.ex.String(), typ.String())
		})
	}
}

func TestProtoKindToCELType(t *testing.T) {
	t.Parallel()

	tests := map[protoreflect.Kind]*cel.Type{
		protoreflect.FloatKind:    cel.DoubleType,
		protoreflect.DoubleKind:   cel.DoubleType,
		protoreflect.Int32Kind:    cel.IntType,
		protoreflect.Int64Kind:    cel.IntType,
		protoreflect.Uint32Kind:   cel.UintType,
		protoreflect.Uint64Kind:   cel.UintType,
		protoreflect.Sint32Kind:   cel.IntType,
		protoreflect.Sint64Kind:   cel.IntType,
		protoreflect.Fixed32Kind:  cel.UintType,
		protoreflect.Fixed64Kind:  cel.UintType,
		protoreflect.Sfixed32Kind: cel.IntType,
		protoreflect.Sfixed64Kind: cel.IntType,
		protoreflect.BoolKind:     cel.BoolType,
		protoreflect.StringKind:   cel.StringType,
		protoreflect.BytesKind:    cel.BytesType,
		protoreflect.EnumKind:     cel.IntType,
		protoreflect.MessageKind:  cel.DynType,
		protoreflect.GroupKind:    cel.DynType,
		protoreflect.Kind(0):      cel.DynType,
	}

	for k, ty := range tests {
		kind, typ := k, ty
		t.Run(kind.String(), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, typ, protoKindToCELType(kind))
		})
	}
}

func getFieldDesc(t *testing.T, msg proto.Message, fld protoreflect.Name) protoreflect.FieldDescriptor {
	t.Helper()
	desc := msg.ProtoReflect().Descriptor().Fields().ByName(fld)
	require.NotNil(t, desc)
	return desc
}
