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

package protovalidate

import "google.golang.org/protobuf/reflect/protoreflect"

// The Filter interface determines which constraints should be validated.
type Filter interface {
	// ShouldValidate returns whether or not a given message, field or oneof
	// should be validated. Note that this only determines whether constraints
	// on the message, field, or oneof itself will be evaluated: nested
	// constraints will still be evaluated unless ShouldValidate returns false
	// for those, too.
	// For a message, the message argument provides the message itself. For a
	// field or oneof, the message argument provides the containing message.
	ShouldValidate(message protoreflect.Message, descriptor protoreflect.Descriptor) bool
}

// FilterFunc is a function type that implements the Filter interface, as a
// convenience for simple filters. A FilterFunc should follow the same semantics
// as the ShouldValidate method of Filter.
type FilterFunc func(protoreflect.Message, protoreflect.Descriptor) bool

func (f FilterFunc) ShouldValidate(message protoreflect.Message, descriptor protoreflect.Descriptor) bool {
	return f(message, descriptor)
}

type nopFilter struct{}

func (nopFilter) ShouldValidate(_ protoreflect.Message, _ protoreflect.Descriptor) bool {
	return true
}

var _ Filter = nopFilter{}
