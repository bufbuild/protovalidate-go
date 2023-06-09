// Copyright 2023 Buf Technologies, Inc.
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

package evaluator

import (
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	newExtensionIndex      = "1159"
	previousExtensionIndex = "51071"
)

type StandardConstraintResolver interface {
	ResolveMessageConstraints(desc protoreflect.MessageDescriptor) *validate.MessageConstraints
	ResolveOneofConstraints(desc protoreflect.OneofDescriptor) *validate.OneofConstraints
	ResolveFieldConstraints(desc protoreflect.FieldDescriptor) *validate.FieldConstraints
}

type DefaultResolver struct{}

func (r DefaultResolver) ResolveMessageConstraints(desc protoreflect.MessageDescriptor) *validate.MessageConstraints {
	constraints := resolveExt[protoreflect.MessageDescriptor, *validate.MessageConstraints](desc, validate.E_Message)
	if constraints == nil {
		constraints = resolveDeprecatedIndex[protoreflect.MessageDescriptor, *validate.MessageConstraints](desc, validate.E_Message)
	}
	return constraints
}

func (r DefaultResolver) ResolveOneofConstraints(desc protoreflect.OneofDescriptor) *validate.OneofConstraints {
	constraints := resolveExt[protoreflect.OneofDescriptor, *validate.OneofConstraints](desc, validate.E_Oneof)
	if constraints == nil {
		constraints = resolveDeprecatedIndex[protoreflect.OneofDescriptor, *validate.OneofConstraints](desc, validate.E_Oneof)
	}
	return constraints
}

func (r DefaultResolver) ResolveFieldConstraints(desc protoreflect.FieldDescriptor) *validate.FieldConstraints {
	constraints := resolveExt[protoreflect.FieldDescriptor, *validate.FieldConstraints](desc, validate.E_Field)
	if constraints == nil {
		constraints = resolveDeprecatedIndex[protoreflect.FieldDescriptor, *validate.FieldConstraints](desc, validate.E_Field)
	}
	return constraints
}

// resolveExt does not use proto.GetExtension in the event the underlying type
// of the extension is not the concrete type expected by the library. In some
// circumstances, particularly in dynamic or runtime contexts, the underlying
// extension value's type may be a dynamicpb.Message. In this case, we fall back
// through a proto.[Un]Marshal cycle to get it into the concrete type we expect.
func resolveExt[D protoreflect.Descriptor, C proto.Message](
	desc D,
	extType protoreflect.ExtensionType,
) (constraints C) {
	opts := desc.Options().ProtoReflect()
	fDesc := extType.TypeDescriptor()

	if opts.Has(fDesc) {
		msg := opts.Get(fDesc).Message().Interface()
		if m, ok := msg.(C); ok {
			return m
		}
	}

	unknown := opts.GetUnknown()
	if len(unknown) == 0 {
		return constraints
	}

	opts = opts.Type().New()
	var resolver protoregistry.Types
	if err := resolver.RegisterExtension(extType); err != nil {
		return constraints
	}
	_ = proto.UnmarshalOptions{Resolver: &resolver}.Unmarshal(unknown, opts.Interface())

	if opts.Has(fDesc) {
		msg := opts.Get(fDesc).Message().Interface()
		if m, ok := msg.(C); ok {
			return m
		}
	}
	msg := opts.Get(fDesc).Message().Interface()
	if m, ok := msg.(C); ok {
		return m
	}

	return constraints
}

// resolveDeprecatedIndex is a fallback for the deprecated extension index.
func resolveDeprecatedIndex[D protoreflect.Descriptor, C proto.Message](
	desc D,
	ext *protoimpl.ExtensionInfo,
) C {
	return resolveExt[D, C](desc, &protoimpl.ExtensionInfo{
		ExtendedType:  ext.ExtendedType,
		ExtensionType: ext.ExtensionType,
		Field:         51071,
		Name:          ext.Name,
		Tag:           strings.Replace(ext.Tag, newExtensionIndex, previousExtensionIndex, 1),
		Filename:      ext.Filename,
	})
}
