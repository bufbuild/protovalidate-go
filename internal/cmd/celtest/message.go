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

package main

import (
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// message performs validation on a protoreflect.Message.
type message struct {
	// Err stores if there was a compilation error constructing this evaluator.
	// It is cached here so that it can be stored in the registry's lookup table.
	Err error

	program   *compiledProgram
	fieldEval field

	// evaluator

	// evaluators are the individual evaluators that are applied to a message.
	// evaluators messageEvaluators

	// // nestedEvaluators are the evaluators that are applied to nested fields and
	// // oneofs.
	// nestedEvaluators messageEvaluators
}

func (m *message) Eval(msg protoreflect.Message) (ref.Val, error) {
	return m.fieldEval.EvaluateMessage(msg)
}
