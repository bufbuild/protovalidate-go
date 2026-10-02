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
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

type supplementaryCase struct {
	name string
	// message is an empty message of the type to validate.
	message proto.Message
	// raw replaces text when the contents can't be prototext (unknown fields).
	text     string
	raw      []byte
	failFast bool
	// expect is an error kind the rules require, independent of evaluating CEL.
	expect string
	want   supplementaryResult
}

type supplementaryResult struct {
	// err is the error kind only; the text depends on the CEL engine.
	err        string
	violations []supplementaryViolation
}

type supplementaryViolation struct {
	ruleID  string
	message string
	field   string
	rule    string
	forKey  bool
}

// TestSupplementaryCases runs protovalidate-python's supplementary cases, which
// check what the conformance suite doesn't: rules on fields of the wrong shape
// or type, numeric and calendar extremes, rule order, regex dialect, string
// formats, CEL edge cases, and wrappers. The messages are in
// proto/tests/supplementary/v1/supplementary.proto. Native rules and CEL-only
// mode must agree, and match the expected result.
func TestSupplementaryCases(t *testing.T) {
	t.Parallel()
	validators := map[[2]bool]Validator{}
	for _, disableNative := range []bool{false, true} {
		for _, failFast := range []bool{false, true} {
			var opts []ValidatorOption
			if disableNative {
				opts = append(opts, WithDisableNativeRules())
			}
			if failFast {
				opts = append(opts, WithFailFast())
			}
			validator, err := New(opts...)
			require.NoError(t, err)
			validators[[2]bool{disableNative, failFast}] = validator
		}
	}
	cases := slices.Concat(
		supplementaryShapeCases(),
		supplementaryTypeCases(),
		supplementaryUniqueCases(),
		supplementaryNumericCases(),
		supplementaryTimeCases(),
		supplementaryOrderCases(),
		supplementaryRegexCases(),
		supplementaryFormatCases(),
		supplementaryHeaderCases(),
		supplementaryCelCases(),
		supplementaryIgnoreCases(),
		supplementaryWrapperCases(),
		supplementaryWrapperListCases(),
		supplementaryWrapperMapCases(),
		supplementaryEnumCases(),
		supplementaryAnyCases(),
		supplementaryFieldMaskCases(),
		supplementaryContradictoryCases(),
		supplementaryRecursionCases(),
		supplementaryParityCases(),
		supplementaryFailFastCases(),
		supplementaryNanBoundCases(),
	)
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			msg := testCase.message.ProtoReflect().New().Interface()
			if testCase.raw != nil {
				require.NoError(t, proto.Unmarshal(testCase.raw, msg))
			} else {
				require.NoError(t, prototext.Unmarshal([]byte(testCase.text), msg))
			}
			native := validateSupplementary(validators[[2]bool{false, testCase.failFast}], msg)
			celOnly := validateSupplementary(validators[[2]bool{true, testCase.failFast}], msg)
			assert.Equal(t, celOnly, native, "native rules and CEL-only mode disagree")
			assert.Equal(t, testCase.want, celOnly)
			if testCase.expect != "" {
				assert.Equal(t, testCase.expect, testCase.want.err, "expected error kind not met")
			}
		})
	}
}

func validateSupplementary(validator Validator, msg proto.Message) (result supplementaryResult) {
	// A panic is a result, so one case can't stop the rest.
	defer func() {
		if r := recover(); r != nil {
			result = supplementaryResult{err: fmt.Sprintf("panic: %v", r)}
		}
	}()
	err := validator.Validate(msg)
	var validationErr *ValidationError
	var compilationErr *CompilationError
	var runtimeErr *RuntimeError
	switch {
	case err == nil:
		return supplementaryResult{}
	case errors.As(err, &validationErr):
		violations := make([]supplementaryViolation, len(validationErr.Violations))
		for i, violation := range validationErr.Violations {
			violations[i] = supplementaryViolation{
				ruleID:  violation.Proto.GetRuleId(),
				message: violation.Proto.GetMessage(),
				field:   FieldPathString(violation.Proto.GetField()),
				rule:    FieldPathString(violation.Proto.GetRule()),
				forKey:  violation.Proto.GetForKey(),
			}
		}
		return supplementaryResult{violations: violations}
	case errors.As(err, &compilationErr):
		return supplementaryResult{err: "CompilationError"}
	case errors.As(err, &runtimeErr):
		return supplementaryResult{err: "EvaluationError"}
	default:
		return supplementaryResult{err: fmt.Sprintf("%T", err)}
	}
}
