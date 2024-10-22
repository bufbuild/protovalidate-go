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

package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// getFieldValue returns the field value at a given path, using the provided
// registry to resolve extensions.
func getFieldValue(
	registry protoregistry.ExtensionTypeResolver,
	message proto.Message,
	path string,
) (field protoreflect.Value, descriptor protoreflect.FieldDescriptor, err error) {
	var name, subscript string
	var atEnd, isExt bool
	reflectMessage := message.ProtoReflect()
	for !atEnd {
		name, subscript, path, atEnd, isExt = parsePathElement(path)
		if name == "" {
			return protoreflect.Value{}, nil, errors.New("empty field name")
		}
		var descriptor protoreflect.FieldDescriptor
		if isExt {
			extension, err := registry.FindExtensionByName(protoreflect.FullName(name))
			if err != nil {
				return protoreflect.Value{}, nil, fmt.Errorf("resolving extension: %w", err)
			}
			descriptor = extension.TypeDescriptor()
		} else {
			descriptor = reflectMessage.Descriptor().Fields().ByTextName(name)
		}
		if descriptor == nil {
			return protoreflect.Value{}, nil, fmt.Errorf("field %s not found", name)
		}
		field = reflectMessage.Get(descriptor)
		if subscript != "" {
			descriptor, field, err = traverseSubscript(descriptor, subscript, field, name)
			if err != nil {
				return protoreflect.Value{}, nil, err
			}
		} else if descriptor.IsList() || descriptor.IsMap() {
			if atEnd {
				break
			}
			return protoreflect.Value{}, nil, fmt.Errorf("missing subscript on field %s", name)
		}
		if descriptor.Message() != nil {
			reflectMessage = field.Message()
		}
	}
	return field, descriptor, nil
}

func traverseSubscript(
	descriptor protoreflect.FieldDescriptor,
	subscript string,
	field protoreflect.Value,
	name string,
) (protoreflect.FieldDescriptor, protoreflect.Value, error) {
	switch {
	case descriptor.IsList():
		i, err := strconv.Atoi(subscript)
		if err != nil {
			return nil, protoreflect.Value{}, fmt.Errorf("invalid list index: %s", subscript)
		}
		if !field.IsValid() || i >= field.List().Len() {
			return nil, protoreflect.Value{}, fmt.Errorf("index %d out of bounds of field %s", i, name)
		}
		field = field.List().Get(i)
	case descriptor.IsMap():
		key, err := parseMapKey(descriptor, subscript)
		if err != nil {
			return nil, protoreflect.Value{}, err
		}
		field = field.Map().Get(key)
		if !field.IsValid() {
			return nil, protoreflect.Value{}, fmt.Errorf("key %s not present on field %s", subscript, name)
		}
		descriptor = descriptor.MapValue()
	default:
		return nil, protoreflect.Value{}, fmt.Errorf("unexpected subscript on field %s", name)
	}
	return descriptor, field, nil
}

func parseMapKey(mapDescriptor protoreflect.FieldDescriptor, subscript string) (protoreflect.MapKey, error) {
	switch mapDescriptor.MapKey().Kind() {
	case protoreflect.BoolKind:
		if boolValue, err := strconv.ParseBool(subscript); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfBool(boolValue)), nil
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if intValue, err := strconv.ParseInt(subscript, 10, 32); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfInt32(int32(intValue))), nil
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if intValue, err := strconv.ParseInt(subscript, 10, 64); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfInt64(intValue)), nil
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if intValue, err := strconv.ParseUint(subscript, 10, 32); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfUint32(uint32(intValue))), nil
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if intValue, err := strconv.ParseUint(subscript, 10, 64); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfUint64(intValue)), nil
		}
	case protoreflect.StringKind:
		if stringValue, err := strconv.Unquote(subscript); err == nil {
			return protoreflect.MapKey(protoreflect.ValueOfString(stringValue)), nil
		}
	case protoreflect.EnumKind, protoreflect.FloatKind, protoreflect.DoubleKind,
		protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
		fallthrough
	default:
		// This should not occur, but it might if the rules are relaxed in the
		// future.
		return protoreflect.MapKey{}, fmt.Errorf("unsupported map key type: %s", mapDescriptor.MapKey().Kind())
	}
	return protoreflect.MapKey{}, fmt.Errorf("invalid map key: %s", subscript)
}

// parsePathElement parses a single
func parsePathElement(path string) (name, subscript, rest string, atEnd bool, isExt bool) {
	// Scan extension name.
	if len(path) > 0 && path[0] == '[' {
		if i := strings.IndexByte(path, ']'); i >= 0 {
			isExt = true
			name, path = path[1:i], path[i+1:]
		}
	}
	// Scan field name.
	if !isExt {
		if i := strings.IndexAny(path, ".["); i >= 0 {
			name, path = path[:i], path[i:]
		} else {
			name, path = path, ""
		}
	}
	// No subscript: At end of path.
	if len(path) == 0 {
		return name, "", path, true, isExt
	}
	// No subscript: At end of path element.
	if path[0] == '.' {
		return name, "", path[1:], false, isExt
	}
	// Malformed subscript
	if len(path) == 1 || path[1] == '.' {
		name, path = name+path[:1], path[1:]
		return name, "", path, true, isExt
	}
	switch path[1] {
	case ']':
		// Empty subscript
		name, path = name+path[:2], path[2:]
	case '`', '"', '\'':
		// String subscript: must scan string.
		var err error
		subscript, err = strconv.QuotedPrefix(path[1:])
		if err == nil {
			path = path[len(subscript)+2:]
		}
	default:
		// Other subscript; can skip to next ]
		if i := strings.IndexByte(path, ']'); i >= 0 {
			subscript, path = path[1:i], path[i+1:]
		} else {
			// Unterminated subscript
			return name + path, "", "", true, isExt
		}
	}
	// No subscript: At end of path.
	if len(path) == 0 {
		return name, subscript, path, true, isExt
	}
	// No subscript: At end of path element.
	if path[0] == '.' {
		return name, subscript, path[1:], false, isExt
	}
	// Malformed element
	return name, subscript, path, false, isExt
}
