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

package expression

import (
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// ProtoKindToCELType maps a protoreflect.Kind to a compatible cel.Type.
func ProtoKindToCELType(kind protoreflect.Kind) *cel.Type {
	switch kind {
	case
		protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return cel.DoubleType
	case
		protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.EnumKind:
		return cel.IntType
	case
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Fixed64Kind:
		return cel.UintType
	case protoreflect.BoolKind:
		return cel.BoolType
	case protoreflect.StringKind:
		return cel.StringType
	case protoreflect.BytesKind:
		return cel.BytesType
	case
		protoreflect.MessageKind,
		protoreflect.GroupKind:
		return cel.DynType
	default:
		return cel.DynType
	}
}

// ProtoFieldToCELType resolves the CEL value type for the provided
// FieldDescriptor. If generic is true, the specific subtypes of map and
// repeated fields will be replaced with cel.DynType. If forItems is true, the
// type for the repeated list items is returned instead of the list type itself.
func ProtoFieldToCELType(fieldDesc protoreflect.FieldDescriptor, generic, forItems bool) *cel.Type {
	if !forItems {
		switch {
		case fieldDesc.IsMap():
			if generic {
				return cel.MapType(cel.DynType, cel.DynType)
			}
			keyType := ProtoFieldToCELType(fieldDesc.MapKey(), false, true)
			valType := ProtoFieldToCELType(fieldDesc.MapValue(), false, true)
			return cel.MapType(keyType, valType)
		case fieldDesc.IsList():
			if generic {
				return cel.ListType(cel.DynType)
			}
			itemType := ProtoFieldToCELType(fieldDesc, false, true)
			return cel.ListType(itemType)
		}
	}

	if fieldDesc.Kind() == protoreflect.MessageKind {
		switch fqn := fieldDesc.Message().FullName(); fqn {
		case "google.protobuf.Any":
			return cel.AnyType
		case "google.protobuf.Duration":
			return cel.DurationType
		case "google.protobuf.Timestamp":
			return cel.TimestampType
		default:
			return cel.ObjectType(string(fqn))
		}
	}
	return ProtoKindToCELType(fieldDesc.Kind())
}

// RequiredCELEnvOptions returns the options required to have expressions which
// rely on the provided descriptor.
func RequiredCELEnvOptions(fieldDesc protoreflect.FieldDescriptor) []cel.EnvOption {
	if fieldDesc.IsMap() {
		return append(
			RequiredCELEnvOptions(fieldDesc.MapKey()),
			RequiredCELEnvOptions(fieldDesc.MapValue())...,
		)
	}
	if fieldDesc.Kind() == protoreflect.MessageKind {
		return []cel.EnvOption{
			cel.Types(dynamicpb.NewMessage(fieldDesc.Message())),
		}
	}
	return nil
}
