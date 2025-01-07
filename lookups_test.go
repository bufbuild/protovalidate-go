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

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestExpectedWrapperConstraints(t *testing.T) {
	t.Parallel()

	tests := map[protoreflect.FullName]*string{
		"google.protobuf.BoolValue":   proto.String("buf.validate.FieldConstraints.bool"),
		"google.protobuf.BytesValue":  proto.String("buf.validate.FieldConstraints.bytes"),
		"google.protobuf.DoubleValue": proto.String("buf.validate.FieldConstraints.double"),
		"google.protobuf.FloatValue":  proto.String("buf.validate.FieldConstraints.float"),
		"google.protobuf.Int32Value":  proto.String("buf.validate.FieldConstraints.int32"),
		"google.protobuf.Int64Value":  proto.String("buf.validate.FieldConstraints.int64"),
		"google.protobuf.StringValue": proto.String("buf.validate.FieldConstraints.string"),
		"google.protobuf.UInt32Value": proto.String("buf.validate.FieldConstraints.uint32"),
		"google.protobuf.UInt64Value": proto.String("buf.validate.FieldConstraints.uint64"),
		"foo.bar":                     nil,
	}

	for name, cons := range tests {
		fqn, constraint := name, cons
		t.Run(string(fqn), func(t *testing.T) {
			t.Parallel()
			desc, ok := expectedWrapperConstraints(fqn)
			if constraint != nil {
				assert.Equal(t, *constraint, string(desc.FullName()))
				assert.True(t, ok)
			} else {
				assert.False(t, ok)
			}
		})
	}
}
