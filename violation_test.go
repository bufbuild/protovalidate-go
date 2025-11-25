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

func TestViolationString(t *testing.T) {
	require.Equal(
		t,
		"one.two: foo",
		(&Violation{
			Proto: &validate.Violation{
				Field: &validate.FieldPath{
					Elements: []*validate.FieldPathElement{
						{
							FieldName: proto.String("one"),
						},
						{
							FieldName: proto.String("two"),
						},
					},
				},
				Message: proto.String("foo"),
				RuleId:  proto.String("bar"),
			},
		}).String(),
	)
	require.Equal(
		t,
		"foo",
		(&Violation{
			Proto: &validate.Violation{
				Message: proto.String("foo"),
				RuleId:  proto.String("bar"),
			},
		}).String(),
	)
	require.Equal(
		t,
		"[bar]",
		(&Violation{
			Proto: &validate.Violation{
				RuleId: proto.String("bar"),
			},
		}).String(),
	)
	require.Equal(
		t,
		"[unknown]",
		(&Violation{
			Proto: &validate.Violation{},
		}).String(),
	)
	require.Equal(
		t,
		"[unknown]",
		(&Violation{}).String(),
	)
}
