// Copyright 2023-2026 Buf Technologies, Inc.
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
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals
var (
	boolConstSite = makeRuleSite(
		fieldRulesDesc.Fields().ByName("bool"),
		(*validate.BoolRules)(nil).ProtoReflect().Descriptor().Fields().ByName("const"),
		"bool.const", "",
	)
)

// tryBuildNativeBoolRules attempts to build a native Go evaluator for
// bool rules. Returns nil if the rules can't be handled natively.
func tryBuildNativeBoolRules(base base, rules *validate.BoolRules) evaluator {
	if rules == nil {
		return nil
	}
	if len(rules.ProtoReflect().GetUnknown()) > 0 {
		return nil
	}
	if !rules.HasConst() {
		return nil
	}
	constVal := rules.GetConst()
	rules.ProtoReflect().Clear(boolConstSite.desc)
	return nativeBoolEval{
		base:     base,
		constVal: constVal,
	}
}

var _ evaluator = nativeBoolEval{}

// nativeBoolEval is a native Go evaluator for bool const rules.
type nativeBoolEval struct {
	base
	constVal bool
}

func (n nativeBoolEval) Evaluate(_ protoreflect.Message, val protoreflect.Value, _ *validationConfig) error {
	if val.Bool() != n.constVal {
		return &ValidationError{Violations: []*Violation{n.newViolation(boolConstSite,
			"bool.const", fmt.Sprintf("must equal %t", n.constVal),
			val, protoreflect.ValueOfBool(n.constVal)),
		}}
	}
	return nil
}

func (n nativeBoolEval) Tautology() bool {
	return false
}
