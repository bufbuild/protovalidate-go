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
	"context"
	"errors"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCompiled(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		prog   cel.Program
		src    *validate.Rule
		exViol *validate.Violation
		exErr  bool
	}{
		{
			name: "valid bool",
			prog: mockProgram{Val: types.True},
		},
		{
			name: "valid string",
			prog: mockProgram{Val: types.String("")},
		},
		{
			name:   "invalid bool",
			prog:   mockProgram{Val: types.False},
			src:    validate.Rule_builder{Id: proto.String("foo"), Message: proto.String("bar")}.Build(),
			exViol: validate.Violation_builder{RuleId: proto.String("foo"), Message: proto.String("bar")}.Build(),
		},
		{
			name:   "invalid string",
			prog:   mockProgram{Val: types.String("bar")},
			src:    validate.Rule_builder{Id: proto.String("foo")}.Build(),
			exViol: validate.Violation_builder{RuleId: proto.String("foo"), Message: proto.String("bar")}.Build(),
		},
		{
			name:  "eval error",
			prog:  mockProgram{Err: errors.New("some error")},
			exErr: true,
		},
		{
			name:  "invalid type",
			prog:  mockProgram{Val: types.Double(1.23)},
			exErr: true,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			expr := compiledProgram{
				Program: test.prog,
				Source:  test.src,
			}
			violation, err := expr.eval(&bindings{}, &validationConfig{nowFn: timestamppb.Now})
			if test.exErr {
				require.Error(t, err)
			} else {
				if test.exViol == nil {
					assert.Nil(t, violation)
				} else {
					assert.True(t, proto.Equal(test.exViol, violation.Proto))
				}
			}
		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()

	env, err := cel.NewEnv()
	require.NoError(t, err)

	tests := []struct {
		name     string
		set      programSet
		failFast bool
		exViols  *validate.Violations
		exErr    bool
	}{
		{
			name: "empty",
		},
		{
			name: "success",
			set: programSet{
				programs: []compiledProgram{
					{
						Program: mockProgram{Val: types.True},
						Source:  &validate.Rule{},
					},
					{
						Program: mockProgram{Val: types.String("")},
						Source:  &validate.Rule{},
					},
				},
				env: env,
			},
		},
		{
			name: "runtime error",
			set: programSet{
				programs: []compiledProgram{
					{
						Program: mockProgram{Val: types.False},
						Source:  &validate.Rule{},
					},
					{
						Program: mockProgram{Err: errors.New("some error")},
						Source:  &validate.Rule{},
					},
				},
				env: env,
			},
			exErr: true,
		},
		{
			name: "invalid",
			set: programSet{
				programs: []compiledProgram{
					{
						Program: mockProgram{Val: types.False},
						Source:  validate.Rule_builder{Id: proto.String("foo"), Message: proto.String("fizz")}.Build(),
					},
					{
						Program: mockProgram{Val: types.String("buzz")},
						Source:  validate.Rule_builder{Id: proto.String("bar")}.Build(),
					},
				},
				env: env,
			},
			exViols: validate.Violations_builder{
				Violations: []*validate.Violation{
					validate.Violation_builder{RuleId: proto.String("foo"), Message: proto.String("fizz")}.Build(),
					validate.Violation_builder{RuleId: proto.String("bar"), Message: proto.String("buzz")}.Build(),
				},
			}.Build(),
		},
		{
			name:     "invalid fail fast",
			failFast: true,
			set: programSet{
				programs: []compiledProgram{
					{
						Program: mockProgram{Val: types.False},
						Source:  validate.Rule_builder{Id: proto.String("foo"), Message: proto.String("fizz")}.Build(),
					},
					{
						Program: mockProgram{Val: types.String("buzz")},
						Source:  validate.Rule_builder{Id: proto.String("bar")}.Build(),
					},
				},
				env: env,
			},
			exViols: validate.Violations_builder{
				Violations: []*validate.Violation{
					validate.Violation_builder{RuleId: proto.String("foo"), Message: proto.String("fizz")}.Build(),
				},
			}.Build(),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.set.Eval(protoreflect.ValueOfBool(false), nil, &validationConfig{
				failFast: test.failFast,
			})
			switch {
			case test.exViols != nil:
				var viols *ValidationError
				require.ErrorAs(t, err, &viols)
				require.True(t, proto.Equal(test.exViols, viols.ToProto()))
			case test.exErr:
				require.Error(t, err)
			default:
				require.NoError(t, err)
			}
		})
	}
}

func TestSet_BindThis(t *testing.T) {
	t.Parallel()
	structMsg := &structpb.Struct{Fields: map[string]*structpb.Value{
		"foo": {Kind: &structpb.Value_BoolValue{BoolValue: true}},
	}}
	mapFieldDesc := structMsg.ProtoReflect().Descriptor().Fields().ByName("fields")
	mapVal := structMsg.ProtoReflect().Get(mapFieldDesc).Map()
	listMsg := &structpb.ListValue{
		Values: []*structpb.Value{
			{Kind: &structpb.Value_BoolValue{BoolValue: true}},
			{Kind: &structpb.Value_BoolValue{BoolValue: false}},
		},
	}
	listVal := listMsg.ProtoReflect().
		Get(listMsg.ProtoReflect().Descriptor().Fields().ByName("values")).
		List()

	tests := []struct {
		name      string
		val       any
		fieldDesc protoreflect.FieldDescriptor
		expr      string
		exType    *cel.Type
	}{
		{
			name:   "reflect message",
			val:    timestamppb.Now().ProtoReflect(),
			expr:   "this",
			exType: cel.TimestampType,
		},
		{
			name:   "proto message",
			val:    timestamppb.Now(),
			expr:   "this",
			exType: cel.TimestampType,
		},
		{
			name:      "proto map",
			val:       mapVal,
			fieldDesc: mapFieldDesc,
			expr:      "this['foo']",
			exType:    cel.BoolType,
		},
		{
			name:   "proto list",
			val:    listVal,
			expr:   "this[1]",
			exType: cel.BoolType,
		},
		{
			name:   "scalar",
			val:    "foo",
			expr:   "this",
			exType: cel.StringType,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			env, err := cel.NewEnv(cel.Variable("this", cel.DynType))
			require.NoError(t, err)
			ast, issues := env.Compile(test.expr)
			require.NoError(t, issues.Err())
			prog, err := env.Program(ast)
			require.NoError(t, err)
			res, _, err := prog.Eval(&bindings{This: newOptional(thisToCel(test.val, test.fieldDesc, env.CELTypeAdapter()))})
			require.NoError(t, err)
			assert.Equal(t, test.exType.String(), res.Type().TypeName())
		})
	}
}

type mockProgram struct {
	Val ref.Val
	Err error
}

func (m mockProgram) Eval(_ any) (ref.Val, *cel.EvalDetails, error) {
	return m.Val, nil, m.Err
}

func (m mockProgram) ContextEval(_ context.Context, _ any) (ref.Val, *cel.EvalDetails, error) {
	return m.Val, nil, m.Err
}
