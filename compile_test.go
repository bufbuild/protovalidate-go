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

	"buf.build/go/protovalidate/gen/buf/validate"
	"github.com/google/cel-go/cel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestCompile(t *testing.T) {
	t.Parallel()

	baseEnv, err := cel.NewEnv()
	baseEnv.Compile("true")
	require.NoError(t, err)

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		var exprs expressions
		set, err := compile(exprs, baseEnv)
		assert.Nil(t, set)
		require.NoError(t, err)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		exprs := expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Id: proto.String("foo"), Expression: proto.String("this == 123")}.Build(),
				validate.Rule_builder{Id: proto.String("bar"), Expression: proto.String("'a string'")}.Build(),
			},
		}
		set, err := compile(exprs, baseEnv, cel.Variable("this", cel.IntType))
		assert.Len(t, set, len(exprs.Rules))
		require.NoError(t, err)
	})

	t.Run("env extension err", func(t *testing.T) {
		t.Parallel()
		exprs := expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Id: proto.String("foo"), Expression: proto.String("0 != 0")}.Build(),
			},
		}
		set, err := compile(exprs, baseEnv, cel.Types(true))
		assert.Nil(t, set)
		var compErr *CompilationError
		require.ErrorAs(t, err, &compErr)
	})

	t.Run("bad syntax", func(t *testing.T) {
		t.Parallel()
		exprs := expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Id: proto.String("foo"), Expression: proto.String("!@#$%^&")}.Build(),
			},
		}
		set, err := compile(exprs, baseEnv)
		assert.Nil(t, set)
		var compErr *CompilationError
		require.ErrorAs(t, err, &compErr)
	})

	t.Run("invalid output type", func(t *testing.T) {
		t.Parallel()
		exprs := expressions{
			Rules: []*validate.Rule{
				validate.Rule_builder{Id: proto.String("foo"), Expression: proto.String("1.23")}.Build(),
			},
		}
		set, err := compile(exprs, baseEnv)
		assert.Nil(t, set)
		var compErr *CompilationError
		require.ErrorAs(t, err, &compErr)
	})
}
