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
	"github.com/google/cel-go/interpreter"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type variable struct {
	// Next is the parent activation
	Next interpreter.Activation
	// Name is the variable's name
	Name string
	// Val is the value for this variable
	Val any
}

func bindThis(val any) *variable {
	binding := &variable{}
	binding.Name = "this"

	switch value := val.(type) {
	case protoreflect.Message:
		binding.Val = value.Interface()
	case protoreflect.Map:
		// TODO: expensive to create this copy, but getting this into a ref.Val with
		//  traits.Mapper is not terribly feasible from this type.
		m := make(map[any]any, value.Len())
		value.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
			// TODO Steve - Maybe we need to look up the Kind of key and then use that
			// to set it. bc for example this code works
			m[key.String()] = value.Interface()
			// m[key.Interface()] = value.Interface()
			return true
		})

		binding.Val = m
	default:
		binding.Val = value
	}

	return binding
}

func (v *variable) ResolveName(name string) (any, bool) {
	switch {
	case name == v.Name:
		return v.Val, true
	case v.Next != nil:
		return v.Next.ResolveName(name)
	default:
		return nil, false
	}
}

func (v *variable) Parent() interpreter.Activation { return nil }
