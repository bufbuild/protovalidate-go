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

package resolver

import (
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	newExtensionIndex      = "1159"  // protovalidate versions >= v0.2.0
	previousExtensionIndex = "51071" // protovalidate versions < v0.2.0
)

// DefaultResolver resolves protovalidate constraints options from descriptors.
type DefaultResolver struct{}

// ResolveMessageConstraints returns the MessageConstraints option set for the
// MessageDescriptor.
func (r DefaultResolver) ResolveMessageConstraints(desc protoreflect.MessageDescriptor) *validate.MessageConstraints {
	return resolveConstraints[validate.MessageConstraints](desc, validate.E_Message)
}

// ResolveOneofConstraints returns the OneofConstraints option set for the
// OneofDescriptor.
func (r DefaultResolver) ResolveOneofConstraints(desc protoreflect.OneofDescriptor) *validate.OneofConstraints {
	return resolveConstraints[validate.OneofConstraints](desc, validate.E_Oneof)
}

// ResolveFieldConstraints returns the FieldConstraints option set for the
// FieldDescriptor.
func (r DefaultResolver) ResolveFieldConstraints(desc protoreflect.FieldDescriptor) *validate.FieldConstraints {
	return resolveConstraints[validate.FieldConstraints](desc, validate.E_Field)
}

func resolveConstraints[C any, CP interface {
	*C
	proto.Message
}](
	desc protoreflect.Descriptor,
	extType *protoimpl.ExtensionInfo,
) (constraints CP) {
	constraints = resolveExt[CP](desc.Options(), extType)
	if constraints == nil {
		constraints = resolveDeprecatedIndex[CP](desc.Options(), extType)
	}
	return constraints
}

// resolveExt does not use proto.GetExtension in the event the underlying type
// of the extension is not the concrete type expected by the library. In some
// circumstances, particularly in dynamic or runtime contexts, the underlying
// extension value's type may be a dynamicpb.Message. In this case, we fall back
// through a proto.[Un]Marshal cycle to get it into the concrete type we expect.
func resolveExt[C proto.Message](
	options proto.Message,
	extType protoreflect.ExtensionType,
) (constraints C) {
	num := extType.TypeDescriptor().Number()
	var msg proto.Message

	proto.RangeExtensions(options, func(typ protoreflect.ExtensionType, i interface{}) bool {
		if num != typ.TypeDescriptor().Number() {
			return true
		}
		msg, _ = i.(proto.Message)
		return false
	})

	if msg == nil {
		return constraints
	} else if m, ok := msg.(C); ok {
		return m
	}

	constraints, _ = constraints.ProtoReflect().New().Interface().(C)
	b, _ := proto.Marshal(msg)
	_ = proto.Unmarshal(b, constraints)
	return constraints
}

// resolveDeprecatedIndex is a fallback for the deprecated extension index.
func resolveDeprecatedIndex[C proto.Message](
	options proto.Message,
	ext *protoimpl.ExtensionInfo,
) C {
	extInfo := &protoimpl.ExtensionInfo{
		ExtendedType:  ext.ExtendedType,
		ExtensionType: ext.ExtensionType,
		Field:         51071,
		Name:          ext.Name,
		Tag:           strings.Replace(ext.Tag, newExtensionIndex, previousExtensionIndex, 1),
		Filename:      ext.Filename,
	}

	// detect and handle if there are unknown options
	if unknown := options.ProtoReflect().GetUnknown(); len(unknown) > 0 {
		opts := options.ProtoReflect().Type().New()
		resolver := &protoregistry.Types{}
		if err := resolver.RegisterExtension(extInfo); err == nil {
			if err = (&proto.UnmarshalOptions{Resolver: resolver}).Unmarshal(unknown, opts.Interface()); err == nil {
				options = opts.Interface()
			}
		}
	}

	return resolveExt[C](options, extInfo)
}
