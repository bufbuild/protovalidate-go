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

import "google.golang.org/protobuf/reflect/protoreflect"

// wrappedValueEval adapts a native evaluator built against a wrapper WKT's
// inner scalar field (e.g. google.protobuf.Int32Value.value) so it can run
// when the outer value is the wrapper message itself. At evaluation time it
// extracts the inner scalar via Message().Get(innerField) and delegates.
//
// processWrapperRules calls buildValue with the inner "value" field
// descriptor, but appends the resulting evaluators onto the outer value
// whose Descriptor still points at the wrapper message field. Without this
// adapter, native evaluators that call val.Int()/Float()/etc. would see the
// wrapper message and panic.
type wrappedValueEval struct {
	innerField protoreflect.FieldDescriptor
	inner      evaluator
}

func (w wrappedValueEval) Evaluate(msg protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	return w.inner.Evaluate(msg, val.Message().Get(w.innerField), cfg)
}

func (w wrappedValueEval) Tautology() bool {
	return w.inner.Tautology()
}

var _ evaluator = wrappedValueEval{}
