// Copyright 2023-2026 Buf Technologies, Inc.
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
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestExpectedWrapperRules(t *testing.T) {
	t.Parallel()

	tests := map[protoreflect.FullName]*string{
		"google.protobuf.BoolValue":   new("buf.validate.FieldRules.bool"),
		"google.protobuf.BytesValue":  new("buf.validate.FieldRules.bytes"),
		"google.protobuf.DoubleValue": new("buf.validate.FieldRules.double"),
		"google.protobuf.FloatValue":  new("buf.validate.FieldRules.float"),
		"google.protobuf.Int32Value":  new("buf.validate.FieldRules.int32"),
		"google.protobuf.Int64Value":  new("buf.validate.FieldRules.int64"),
		"google.protobuf.StringValue": new("buf.validate.FieldRules.string"),
		"google.protobuf.UInt32Value": new("buf.validate.FieldRules.uint32"),
		"google.protobuf.UInt64Value": new("buf.validate.FieldRules.uint64"),
		"foo.bar":                     nil,
	}

	for name, cons := range tests {
		fqn, rule := name, cons
		t.Run(string(fqn), func(t *testing.T) {
			t.Parallel()
			desc, ok := expectedWrapperRules(fqn)
			if rule != nil {
				assert.Equal(t, *rule, string(desc.FullName()))
				assert.True(t, ok)
			} else {
				assert.False(t, ok)
			}
		})
	}
}
