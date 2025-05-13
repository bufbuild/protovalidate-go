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
			src:    &validate.Rule{Id: proto.String("foo"), Message: proto.String("bar")},
			exViol: &validate.Violation{RuleId: proto.String("foo"), Message: proto.String("bar")},
		},
		{
			name:   "invalid string",
			prog:   mockProgram{Val: types.String("bar")},
			src:    &validate.Rule{Id: proto.String("foo")},
			exViol: &validate.Violation{RuleId: proto.String("foo"), Message: proto.String("bar")},
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
			violation, err := expr.eval(&variable{}, &validationConfig{nowFn: timestamppb.Now})
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
				compiledProgram{
					Program: mockProgram{Val: types.True},
					Source:  &validate.Rule{},
				},
				compiledProgram{
					Program: mockProgram{Val: types.String("")},
					Source:  &validate.Rule{},
				},
			},
		},
		{
			name: "runtime error",
			set: programSet{
				compiledProgram{
					Program: mockProgram{Val: types.False},
					Source:  &validate.Rule{},
				},
				compiledProgram{
					Program: mockProgram{Err: errors.New("some error")},
					Source:  &validate.Rule{},
				},
			},
			exErr: true,
		},
		{
			name: "invalid",
			set: programSet{
				compiledProgram{
					Program: mockProgram{Val: types.False},
					Source:  &validate.Rule{Id: proto.String("foo"), Message: proto.String("fizz")},
				},
				compiledProgram{
					Program: mockProgram{Val: types.String("buzz")},
					Source:  &validate.Rule{Id: proto.String("bar")},
				},
			},
			exViols: &validate.Violations{
				Violations: []*validate.Violation{
					{RuleId: proto.String("foo"), Message: proto.String("fizz")},
					{RuleId: proto.String("bar"), Message: proto.String("buzz")},
				},
			},
		},
		{
			name:     "invalid fail fast",
			failFast: true,
			set: programSet{
				compiledProgram{
					Program: mockProgram{Val: types.False},
					Source:  &validate.Rule{Id: proto.String("foo"), Message: proto.String("fizz")},
				},
				compiledProgram{
					Program: mockProgram{Val: types.String("buzz")},
					Source:  &validate.Rule{Id: proto.String("bar")},
				},
			},
			exViols: &validate.Violations{
				Violations: []*validate.Violation{
					{RuleId: proto.String("foo"), Message: proto.String("fizz")},
				},
			},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.set.Eval(protoreflect.ValueOfBool(false), &validationConfig{
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
	mapVal := structMsg.ProtoReflect().
		Get(structMsg.ProtoReflect().Descriptor().Fields().ByName("fields")).
		Map()
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
		name   string
		val    any
		expr   string
		exType *cel.Type
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
			name:   "proto map",
			val:    mapVal,
			expr:   "this['foo']",
			exType: cel.BoolType,
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

	set := programSet{}

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
			res, _, err := prog.Eval(set.bindThis(test.val))
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
