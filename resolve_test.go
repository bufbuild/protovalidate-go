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

package protovalidate

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestResolve(t *testing.T) {
	t.Parallel()

	expectedRules := &validate.FieldRules{
		Cel: []*validate.Rule{
			{Message: proto.String("test")},
		},
	}
	expectedRulesBytes, err := proto.Marshal(expectedRules)
	require.NoError(t, err)

	tests := []struct {
		name    string
		builder func() proto.Message
	}{
		{
			name: "Normal",
			builder: func() proto.Message {
				options := &descriptorpb.FieldOptions{}
				proto.SetExtension(options, validate.E_Field, expectedRules)
				return options
			},
		},
		{
			name: "Dynamic",
			builder: func() proto.Message {
				var unknownBytes []byte
				unknownBytes = protowire.AppendTag(
					unknownBytes,
					validate.E_Field.TypeDescriptor().Number(),
					protowire.BytesType,
				)
				unknownBytes = protowire.AppendBytes(
					unknownBytes,
					expectedRulesBytes,
				)
				options := &descriptorpb.FieldOptions{}
				options.ProtoReflect().SetUnknown(protoreflect.RawFields(unknownBytes))
				return options
			},
		},
		{
			name: "Unknown",
			builder: func() proto.Message {
				var unknownBytes []byte
				unknownBytes = protowire.AppendTag(
					unknownBytes,
					validate.E_Field.TypeDescriptor().Number(),
					protowire.BytesType,
				)
				unknownBytes = protowire.AppendBytes(
					unknownBytes,
					expectedRulesBytes,
				)
				options := &descriptorpb.FieldOptions{}
				options.ProtoReflect().SetUnknown(protoreflect.RawFields(unknownBytes))
				return options
			},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			pb := test.builder()
			extension, err := resolve[*validate.FieldRules](pb, validate.E_Field)
			require.NoError(t, err)
			require.NotNil(t, extension)
			require.Equal(t, "test", extension.GetCel()[0].GetMessage())
		})
	}
}

func TestResolveNone(t *testing.T) {
	t.Parallel()
	extension, err := resolve[*validate.FieldRules](
		&descriptorpb.FieldOptions{},
		validate.E_Field,
	)
	require.NoError(t, err)
	require.Nil(t, extension)
}
