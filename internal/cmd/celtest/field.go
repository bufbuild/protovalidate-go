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

type field struct {
	// Value is the evaluator to apply to the field's value
	Value value

	Err error
}

type value struct {
	// Descriptor is the FieldDescriptor targeted by this evaluator
	Descriptor protoreflect.FieldDescriptor

	program *compiledProgram
}

func (v value) Evaluate(_ protoreflect.Message, val protoreflect.Value) (ref.Val, error) {
	binding := bindThis(val.Interface())

	return v.program.eval(binding)
}

func (f field) EvaluateMessage(msg protoreflect.Message) (ref.Val, error) {
	if f.Err != nil {
		return nil, f.Err
	}

	val := msg.Get(f.Value.Descriptor)
	return f.Value.Evaluate(msg, val)
}
