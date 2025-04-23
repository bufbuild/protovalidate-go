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
	"time"

	pb "github.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/sourcecontextpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				&pb.HasMsgExprs{X: 2, Y: 43},
				false,
			},
			{
				&pb.HasMsgExprs{X: 9, Y: 8},
				true,
			},
		}

		for _, test := range tests {
			err := val.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestValidator_ValidateGlobal(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				&pb.HasMsgExprs{X: 2, Y: 43},
				false,
			},
			{
				&pb.HasMsgExprs{X: 9, Y: 8},
				true,
			},
		}

		for _, test := range tests {
			err := Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestGlobalValidator(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				&pb.HasMsgExprs{X: 2, Y: 43},
				false,
			},
			{
				&pb.HasMsgExprs{X: 9, Y: 8},
				true,
			},
		}

		for _, test := range tests {
			err := GlobalValidator.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		}
	})
}

func TestRecursive(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)

	selfRec := &pb.SelfRecursive{X: 123, Turtle: &pb.SelfRecursive{X: 456}}
	err = val.Validate(selfRec)
	require.NoError(t, err)

	loopRec := &pb.LoopRecursiveA{B: &pb.LoopRecursiveB{}}
	err = val.Validate(loopRec)
	require.NoError(t, err)
}

func TestValidator_ValidateOneof(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	oneofMessage := &pb.MsgHasOneof{O: &pb.MsgHasOneof_X{X: "foo"}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Y{Y: 42}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Msg{Msg: &pb.HasMsgExprs{X: 4, Y: 50}}}
	err = val.Validate(oneofMessage)
	require.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{}
	err = val.Validate(oneofMessage)
	assert.Error(t, err)
}

func TestValidator_ValidateRepeatedFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	repeatMessage := &pb.MsgHasRepeated{
		X: []float32{1, 2, 3},
		Y: []string{"foo", "bar"},
		Z: []*pb.HasMsgExprs{
			{
				X: 4,
				Y: 55,
			}, {
				X: 4,
				Y: 60,
			},
		},
	}
	err = val.Validate(repeatMessage)
	require.NoError(t, err)
}

func TestValidator_ValidateMapFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	mapMessage := &pb.MsgHasMap{
		Int32Map:   map[int32]int32{-1: 1, 2: 2},
		StringMap:  map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		MessageMap: map[int64]*pb.LoopRecursiveA{0: nil},
	}
	err = val.Validate(mapMessage)
	require.Error(t, err)
}

func TestValidator_Validate_TransitiveFieldConstraints(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.TransitiveFieldConstraint{
		Mask: &fieldmaskpb.FieldMask{Paths: []string{"foo", "bar"}},
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_MultipleStepsTransitiveFieldConstraints(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.MultipleStepsTransitiveFieldConstraints{
		Api: &apipb.Api{
			SourceContext: &sourcecontextpb.SourceContext{
				FileName: "path/file",
			},
		},
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_FieldOfTypeAny(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	simple := &pb.Simple{S: "foo"}
	anyFromSimple, err := anypb.New(simple)
	require.NoError(t, err)
	msg := &pb.FieldOfTypeAny{
		Any: anyFromSimple,
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_CelMapOnARepeated(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.CelMapOnARepeated{Values: []*pb.CelMapOnARepeated_Value{
		{Name: "foo"},
		{Name: "bar"},
		{Name: "baz"},
	}}
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.Values = append(msg.Values, &pb.CelMapOnARepeated_Value{Name: "foo"})
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
}

func TestValidator_Validate_RepeatedItemCel(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.RepeatedItemCel{Paths: []string{"foo"}}
	err = val.Validate(msg)
	require.NoError(t, err)
	msg.Paths = append(msg.Paths, " bar")
	err = val.Validate(msg)
	valErr := &ValidationError{}
	require.ErrorAs(t, err, &valErr)
	assert.Equal(t, "paths.no_space", valErr.Violations[0].Proto.GetConstraintId())
	pathsFd := msg.ProtoReflect().Descriptor().Fields().ByName("paths")
	err = val.Validate(msg, WithFilter(FilterFunc(func(m protoreflect.Message, d protoreflect.Descriptor) bool {
		return !(m.Interface() == msg && d == pathsFd)
	})))
	require.NoError(t, err)
}

func TestValidator_Validate_Filter(t *testing.T) {
	t.Parallel()

	t.Run("FilterField", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.Person{}
		err = val.Validate(msg)
		valErr := &ValidationError{}
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 3)
		idFd := msg.ProtoReflect().Descriptor().Fields().ByName("id")
		err = val.Validate(msg, WithFilter(FilterFunc(func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
			return d == idFd
		})))
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 1)
	})

	t.Run("FilterInvalid", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.InvalidConstraints{}
		err = val.Validate(msg)
		require.Error(t, err)
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, _ protoreflect.Descriptor) bool {
				return false
			},
		)))
		require.NoError(t, err)
	})

	t.Run("FilterNested", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.NestedConstraints{
			Field:         &pb.AllConstraintTypes{},
			RepeatedField: []*pb.AllConstraintTypes{{}},
			MapField:      map[string]*pb.AllConstraintTypes{"test": {}},
		}
		descs := []string{}
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				descs = append(descs, string(d.FullName()))
				return false
			},
		)))
		require.Equal(t, []string{
			"tests.example.v1.NestedConstraints",
			"tests.example.v1.NestedConstraints.required_oneof",
			"tests.example.v1.NestedConstraints.field",
			"tests.example.v1.NestedConstraints.field2",
			"tests.example.v1.NestedConstraints.repeated_field",
			"tests.example.v1.NestedConstraints.map_field",
		}, descs)
		require.NoError(t, err)
		descs = []string{}
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				descs = append(descs, string(d.FullName()))
				return true
			},
		)))
		require.Equal(t, []string{
			"tests.example.v1.NestedConstraints",
			"tests.example.v1.NestedConstraints.required_oneof",
			"tests.example.v1.NestedConstraints.field",
			"tests.example.v1.AllConstraintTypes",
			"tests.example.v1.AllConstraintTypes.required_oneof",
			"tests.example.v1.AllConstraintTypes.field",
			"tests.example.v1.NestedConstraints.field2",
			"tests.example.v1.NestedConstraints.repeated_field",
			"tests.example.v1.AllConstraintTypes",
			"tests.example.v1.AllConstraintTypes.required_oneof",
			"tests.example.v1.AllConstraintTypes.field",
			"tests.example.v1.NestedConstraints.map_field",
			"tests.example.v1.AllConstraintTypes",
			"tests.example.v1.AllConstraintTypes.required_oneof",
			"tests.example.v1.AllConstraintTypes.field",
		}, descs)
		require.Error(t, err)
	})

	t.Run("FilterIncludeCompilationError", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.MixedValidInvalidConstraints{
			StringFieldBoolConstraint: "foo",
			ValidStringConstraint:     "bar",
		}
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				return d == msg.ProtoReflect().Descriptor().Fields().Get(0)
			},
		)))
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})

	t.Run("FilterExcludeCompilationError", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.MixedValidInvalidConstraints{
			ValidStringConstraint:     "bar",
			StringFieldBoolConstraint: "foo",
		}
		err = val.Validate(msg, WithFilter(FilterFunc(
			func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
				return d == msg.ProtoReflect().Descriptor().Fields().Get(1)
			},
		)))
		require.Error(t, err)
		compErr := &CompilationError{}
		require.NotErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.ErrorAs(t, err, &valErr)
		require.Len(t, valErr.Violations, 1)
	})
}

func TestValidator_ValidateCompilationError(t *testing.T) {
	t.Parallel()

	t.Run("CompilationErrorNoViolations", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.MismatchConstraints{}
		err = val.Validate(msg)
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})

	t.Run("CompilationErrorWithViolations", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.MixedValidInvalidConstraints{
			StringFieldBoolConstraint: "foo",
			ValidStringConstraint:     "bar",
		}
		err = val.Validate(msg)
		require.Error(t, err)
		compErr := &CompilationError{}
		require.ErrorAs(t, err, &compErr)
		valErr := &ValidationError{}
		require.NotErrorAs(t, err, &valErr)
	})
}

func TestValidator_WithNowFunc_Issue211(t *testing.T) {
	t.Parallel()

	nowFn := func() *timestamppb.Timestamp {
		return timestamppb.New(time.Now().Add(time.Hour))
	}

	msg := &pb.Issue211{
		Value: timestamppb.New(time.Now().Add(time.Minute)),
	}
	val, err := New()
	require.NoError(t, err)
	err = val.Validate(msg)
	require.NoError(t, err)
	err = val.Validate(msg, WithNowFunc(nowFn))
	require.Error(t, err)

	val, err = New(WithNowFunc(nowFn))
	require.NoError(t, err)
	err = val.Validate(msg)
	require.Error(t, err)
	err = val.Validate(msg, WithNowFunc(timestamppb.Now))
	require.NoError(t, err)
}

func TestValidator_Validate_Issue141(t *testing.T) {
	t.Parallel()

	t.Run("FieldWithIssue", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.FieldWithIssue{}
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})

	t.Run("OneTwo", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.OneTwo{
			Field1: &pb.F1{
				Field: &pb.FieldWithIssue{},
			},
		}
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})

	t.Run("TwoOne", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)
		msg := &pb.TwoOne{
			Field1: &pb.F1{
				Field: &pb.FieldWithIssue{},
			},
		}
		err = val.Validate(msg)
		var valErr *ValidationError
		require.ErrorAs(t, err, &valErr)
	})
}

func TestValidator_Validate_Issue148(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.Issue148{Test: proto.Int32(1)}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_Issue187(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := pb.Issue187_builder{
		FalseField: proto.Bool(false),
		TrueField:  proto.Bool(true),
	}.Build()
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_RailroadGin(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.RailroadGin{
		Name: 69,
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}

func TestValidator_Validate_WrongType_Timestamp(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	msg := &pb.RulesWrongTypeExample{
		F: &pb.WrongType{
			Name: "test",
		},
	}
	err = val.Validate(msg)
	require.NoError(t, err)
}
