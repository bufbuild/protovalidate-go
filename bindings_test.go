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

var _ interpreter.Activation = (*bindings)(nil)

func TestBindings(t *testing.T) {
	t.Parallel()
	var (
		this  = 123
		rules = validate.Int32Rules_builder{In: []int32{1}}.Build()
		rule  = rules.GetIn()
	)
	bind := bindings{
		This:  newOptional[any](this),
		Rules: rules,
		Rule:  rule,
	}
	gotThis, ok := bind.ResolveName("this")
	assert.True(t, ok)
	assert.Equal(t, this, gotThis)
	gotRules, ok := bind.ResolveName("rules")
	assert.True(t, ok)
	assert.Empty(t, cmp.Diff(rules, gotRules, protocmp.Transform()))
	gotRule, ok := bind.ResolveName("rule")
	assert.True(t, ok)
	assert.Equal(t, rule, gotRule)
	gotNow, ok := bind.ResolveName("now")
	assert.False(t, ok)
	assert.Nil(t, gotNow)
	bind.This = optional[any]{}
	gotThis, ok = bind.ResolveName("this")
	assert.False(t, ok)
	assert.Nil(t, gotThis)
	now := timestamppb.Now()
	bind.NowFn = func() *timestamppb.Timestamp {
		return now
	}
	gotNow, ok = bind.ResolveName("now")
	assert.True(t, ok)
	assert.Equal(t, now, gotNow)
}
