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
	"github.com/google/cel-go/cel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestASTSet_Merge(t *testing.T) {
	t.Parallel()

	var set astSet
	other := astSet{
		{AST: &cel.Ast{}},
		{AST: &cel.Ast{}},
	}
	merged := set.Merge(other)
	assert.Equal(t, other, merged)

	another := astSet{
		{AST: &cel.Ast{}},
		{AST: &cel.Ast{}},
		{AST: &cel.Ast{}},
	}
	merged = other.Merge(another)
	assert.Equal(t, other, merged[0:2])
	assert.Equal(t, another, merged[2:])
}

func TestASTSet_ToProgramSet(t *testing.T) {
	t.Parallel()

	env, err := cel.NewEnv(cel.Lib(pvcel.NewLibrary()))
	require.NoError(t, err)

	asts, err := compileASTs(
		expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Expression: proto.String("foo")}.Build(),
			},
		},
		env,
		cel.Variable("foo", cel.BoolType),
	)
	require.NoError(t, err)
	assert.Len(t, asts, 1)
	set, err := asts.ToProgramSet()
	require.NoError(t, err)
	assert.Len(t, set, 1)
	assert.Equal(t, asts[0].Source, set[0].Source)

	empty := astSet{}
	set, err = empty.ToProgramSet()
	assert.Empty(t, set)
	require.NoError(t, err)
}

func TestASTSet_ReduceResiduals(t *testing.T) {
	t.Parallel()

	env, err := cel.NewEnv(cel.Lib(pvcel.NewLibrary()))
	require.NoError(t, err)

	asts, err := compileASTs(
		expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Expression: proto.String("foo")}.Build(),
			},
		},
		env,
		cel.Variable("foo", cel.BoolType),
	)
	require.NoError(t, err)
	assert.Len(t, asts, 1)
	set, err := asts.ReduceResiduals(
		(&validate.StringRules{}).ProtoReflect(),
		cel.Globals(map[string]any{"foo": true}),
	)
	require.NoError(t, err)
	assert.Empty(t, set)
}
