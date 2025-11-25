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
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestValidationErrorString(t *testing.T) {
	require.Equal(
		t,
		`validation error: one.two: foo`,
		(&ValidationError{
			Violations: []*Violation{
				{
					Proto: validate.Violation_builder{
						Field: validate.FieldPath_builder{
							Elements: []*validate.FieldPathElement{
								validate.FieldPathElement_builder{
									FieldName: proto.String("one"),
								}.Build(),
								validate.FieldPathElement_builder{
									FieldName: proto.String("two"),
								}.Build(),
							},
						}.Build(),
						Message: proto.String("foo"),
						RuleId:  proto.String("bar"),
					}.Build(),
				},
			},
		}).Error(),
	)
	require.Equal(
		t,
		`validation errors:
 - one.two: foo
 - one.three: [bar]`,
		(&ValidationError{
			Violations: []*Violation{
				{
					Proto: validate.Violation_builder{
						Field: validate.FieldPath_builder{
							Elements: []*validate.FieldPathElement{
								validate.FieldPathElement_builder{
									FieldName: proto.String("one"),
								}.Build(),
								validate.FieldPathElement_builder{
									FieldName: proto.String("two"),
								}.Build(),
							},
						}.Build(),
						Message: proto.String("foo"),
						RuleId:  proto.String("bar"),
					}.Build(),
				},
				{
					Proto: validate.Violation_builder{
						Field: validate.FieldPath_builder{
							Elements: []*validate.FieldPathElement{
								validate.FieldPathElement_builder{
									FieldName: proto.String("one"),
								}.Build(),
								validate.FieldPathElement_builder{
									FieldName: proto.String("three"),
								}.Build(),
							},
						}.Build(),
						RuleId: proto.String("bar"),
					}.Build(),
				},
			},
		}).Error(),
	)
}
