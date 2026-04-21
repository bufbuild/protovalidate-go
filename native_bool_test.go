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

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestNativeBoolConst_True(t *testing.T) {
	t.Parallel()
	eval := buildNativeBool(t, validate.BoolRules_builder{Const: proto.Bool(true)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBool(true), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfBool(false), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	require.Len(t, valErr.Violations, 1)
	assert.Equal(t, "bool.const", valErr.Violations[0].Proto.GetRuleId())
	assert.Equal(t, "must equal true", valErr.Violations[0].Proto.GetMessage())
}

func TestNativeBoolConst_False(t *testing.T) {
	t.Parallel()
	eval := buildNativeBool(t, validate.BoolRules_builder{Const: proto.Bool(false)}.Build())
	require.NotNil(t, eval)

	require.NoError(t, eval.Evaluate(nil, protoreflect.ValueOfBool(false), &validationConfig{}))

	err := eval.Evaluate(nil, protoreflect.ValueOfBool(true), &validationConfig{})
	require.Error(t, err)
	var valErr *ValidationError
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "must equal false", valErr.Violations[0].Proto.GetMessage())
}

func TestTryBuildNativeBoolRules_ReturnsNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rules *validate.BoolRules
	}{
		{"nil_rules", nil},
		{"no_const", validate.BoolRules_builder{}.Build()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Nil(t, tryBuildNativeBoolRules(base{}, tt.rules))
		})
	}
}

func TestNativeBoolTautology(t *testing.T) {
	t.Parallel()
	eval := buildNativeBool(t, validate.BoolRules_builder{Const: proto.Bool(true)}.Build())
	require.NotNil(t, eval)
	assert.False(t, eval.Tautology())
}

func buildNativeBool(t testing.TB, rules *validate.BoolRules) evaluator {
	t.Helper()
	fdesc := newFieldDescriptor(t, descriptorpb.FieldDescriptorProto_TYPE_BOOL,
		descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum())
	b := base{
		Descriptor:       fdesc,
		FieldPathElement: fieldPathElement(fdesc),
	}
	return tryBuildNativeBoolRules(b, rules)
}
