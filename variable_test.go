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
	"github.com/google/cel-go/interpreter"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ interpreter.Activation = (*variables)(nil)

func TestPredefinedRuleVariable(t *testing.T) {
	t.Parallel()
	var (
		this  = 123
		rules = &validate.Int32Rules{In: []int32{1}}
		rule  = rules.GetIn()
	)
	v := variables{
		This:  newOptional[any](this),
		Rules: rules,
		Rule:  rule,
	}
	gotThis, ok := v.ResolveName("this")
	assert.True(t, ok)
	assert.Equal(t, this, gotThis)
	gotRules, ok := v.ResolveName("rules")
	assert.True(t, ok)
	assert.Empty(t, cmp.Diff(rules, gotRules, protocmp.Transform()))
	gotRule, ok := v.ResolveName("rule")
	assert.True(t, ok)
	assert.Equal(t, rule, gotRule)
	gotNow, ok := v.ResolveName("now")
	assert.False(t, ok)
	assert.Nil(t, gotNow)
	v.This = optional[any]{}
	gotThis, ok = v.ResolveName("this")
	assert.False(t, ok)
	assert.Nil(t, gotThis)
	now := timestamppb.Now()
	v.NowFn = func() *timestamppb.Timestamp {
		return now
	}
	gotNow, ok = v.ResolveName("now")
	assert.True(t, ok)
	assert.Equal(t, now, gotNow)
}
