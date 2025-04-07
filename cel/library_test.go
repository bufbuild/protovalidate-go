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

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCELLib(t *testing.T) {
	t.Parallel()

	env, err := cel.NewEnv(cel.Lib(NewLibrary()))
	require.NoError(t, err)

	t.Run("ext", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			expr string
			ex   bool
		}{
			// {"0.0.isInf()", false},
			// {"0.0.isNan()", false},
			// {"(1.0/0.0).isInf()", true},
			// {"(1.0/0.0).isInf(0)", true},
			// {"(1.0/0.0).isInf(1)", true},
			// {"(1.0/0.0).isInf(-1)", false},
			// {"(-1.0/0.0).isInf()", true},
			// {"(-1.0/0.0).isInf(0)", true},
			// {"(-1.0/0.0).isInf(1)", false},
			// {"(-1.0/0.0).isInf(-1)", true},
			// {"(0.0/0.0).isNan()", true},
			// {"(0.0/0.0).isInf()", false},
			// {"(1.0/0.0).isNan()", false},
			// {
			// 	"[].unique()",
			// 	true,
			// },
			// {
			// 	"[true].unique()",
			// 	true,
			// },
			// {
			// 	"[true, false].unique()",
			// 	true,
			// },
			// {
			// 	"[true, true].unique()",
			// 	false,
			// },
			// {
			// 	"[1, 2, 3].unique()",
			// 	true,
			// },
			// {
			// 	"[1, 2, 1].unique()",
			// 	false,
			// },
			// {
			// 	"[1u, 2u, 3u].unique()",
			// 	true,
			// },
			// {
			// 	"[1u, 2u, 2u].unique()",
			// 	false,
			// },
			// {
			// 	"[1.0, 2.0, 3.0].unique()",
			// 	true,
			// },
			// {
			// 	"[3.0,2.0,3.0].unique()",
			// 	false,
			// },
			// {
			// 	"['abc', 'def'].unique()",
			// 	true,
			// },
			// {
			// 	"['abc', 'abc'].unique()",
			// 	false,
			// },
			// {
			// 	"[b'abc', b'123'].unique()",
			// 	true,
			// },
			// {
			// 	"[b'123', b'123'].unique()",
			// 	false,
			// },
			// {
			// 	"'1.2.3.0/24'.isIpPrefix()",
			// 	true,
			// },
			// {
			// 	"'1.2.3.4/24'.isIpPrefix()",
			// 	true,
			// },
			// {
			// 	"'1.2.3.0/24'.isIpPrefix(true)",
			// 	true,
			// },
			// {
			// 	"'1.2.3.4/24'.isIpPrefix(true)",
			// 	false,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:4000/118'.isIpPrefix()",
			// 	true,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118'.isIpPrefix()",
			// 	true,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118'.isIpPrefix(true)",
			// 	false,
			// },
			// {
			// 	"'1.2.3.4'.isIpPrefix()",
			// 	false,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:430b'.isIpPrefix()",
			// 	false,
			// },
			// {
			// 	"'1.2.3.0/24'.isIpPrefix(4)",
			// 	true,
			// },
			// {
			// 	"'1.2.3.4/24'.isIpPrefix(4)",
			// 	true,
			// },
			// {
			// 	"'1.2.3.0/24'.isIpPrefix(4,true)",
			// 	true,
			// },
			// {
			// 	"'1.2.3.4/24'.isIpPrefix(4,true)",
			// 	false,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:4000/118'.isIpPrefix(4)",
			// 	false,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:4000/118'.isIpPrefix(6)",
			// 	true,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118'.isIpPrefix(6)",
			// 	true,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:4000/118'.isIpPrefix(6,true)",
			// 	true,
			// },
			// {
			// 	"'fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118'.isIpPrefix(6,true)",
			// 	false,
			// },
			// {
			// 	"'1.2.3.0/24'.isIpPrefix(6)",
			// 	false,
			// },
			// {
			// 	"'foo@example.com'.isEmail()",
			// 	true,
			// },
			// {
			// 	"'<foo@example.com>'.isEmail()",
			// 	false,
			// },
			// {
			// 	"'  foo@example.com'.isEmail()",
			// 	false,
			// },
			// {
			// 	"'foo@example.com    '.isEmail()",
			// 	false,
			// },
			{
				"'::1q'.isIp()",
				false,
			},
			{
				"'1234567890abcdef'.isIp()",
				false,
			},
		}

		for _, tc := range tests {
			test := tc
			t.Run(test.expr, func(t *testing.T) {
				t.Parallel()
				prog := buildTestProgram(t, env, test.expr)
				val, _, err := prog.Eval(interpreter.EmptyActivation())
				require.NoError(t, err)
				isUnique, ok := val.Value().(bool)
				require.True(t, ok)
				assert.Equal(t, test.ex, isUnique)
			})
		}
	})
}

func buildTestProgram(t *testing.T, env *cel.Env, expr string) cel.Program {
	t.Helper()
	ast, issues := env.Compile(expr)
	require.NoError(t, issues.Err())
	prog, err := env.Program(ast)
	require.NoError(t, err)
	return prog
}
