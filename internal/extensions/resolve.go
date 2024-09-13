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

package extensions

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

//nolint:gochecknoglobals // static data, only want single instance
var resolver = newExtensionResolver()

// Resolve resolves extensions without using [proto.GetExtension], in case the
// underlying type of the extension is not the concrete type expected by the
// library. In some cases, particularly when using a dynamic descriptor set, the
// underlying extension value's type will be a dynamicpb.Message. In some cases,
// the extension may not be resolved at all. This function handles reparsing the
// fields as needed to get it into the right concrete message as needed. Resolve
// does not modify the input protobuf message, so it can be used concurrently.
func Resolve[C proto.Message](
	options proto.Message,
	extensionType protoreflect.ExtensionType,
) (typedMessage C) {
	message := resolver.resolve(options, extensionType)
	if message == nil {
		return typedMessage
	} else if typedMessage, ok := message.(C); ok {
		return typedMessage
	}
	typedMessage, _ = typedMessage.ProtoReflect().New().Interface().(C)
	b, _ := proto.Marshal(message)
	_ = proto.Unmarshal(b, typedMessage)
	return typedMessage
}

// extensionResolver implements most of the logic of resolving protovalidate
// extensions.
type extensionResolver struct {
	// registry is a registry that just contains the protovalidate extensions.
	// Consulting the global registry is unnecessary and will incur global mutex
	// locks, so it is best avoided.
	registry *protoregistry.Types

	// legacyExtensionMap is a mapping from current protovalidate extensions to
	// legacy protovalidate extensions, used for backwards compatibility. This
	// map will not be modified, so it is safe to read concurrently.
	legacyExtensionMap map[protoreflect.ExtensionType]protoreflect.ExtensionType
}

// newExtensionResolver creates a new extension resolver. This is only called at
// init and will panic if it fails.
func newExtensionResolver() extensionResolver {
	resolver := extensionResolver{
		registry:           &protoregistry.Types{},
		legacyExtensionMap: make(map[protoreflect.ExtensionType]protoreflect.ExtensionType),
	}
	resolver.register(validate.E_Field)
	resolver.register(validate.E_Message)
	resolver.register(validate.E_Oneof)
	resolver.register(validate.E_Predefined)
	resolver.register(resolver.createLegacyExtension(validate.E_Field))
	resolver.register(resolver.createLegacyExtension(validate.E_Message))
	resolver.register(resolver.createLegacyExtension(validate.E_Oneof))
	return resolver
}

// register registers an extension into the resolver's registry. This is only
// called at init and will panic if it fails.
func (resolver extensionResolver) register(xt protoreflect.ExtensionType) {
	if err := resolver.registry.RegisterExtension(xt); err != nil {
		//nolint:forbidigo // this needs to be a fatal at init
		panic(err)
	}
}

// createLegacyExtension creates a legacy extension, putting it in the legacy
// extension mapping and returning it.
func (resolver extensionResolver) createLegacyExtension(extensionInfo *protoimpl.ExtensionInfo) *protoimpl.ExtensionInfo {
	legacyExtensionInfo := &protoimpl.ExtensionInfo{
		ExtendedType:  extensionInfo.ExtendedType,
		ExtensionType: extensionInfo.ExtensionType,
		Field:         51071,
		Name:          extensionInfo.Name + "_legacy",
		Tag:           strings.Replace(extensionInfo.Tag, newExtensionIndex, previousExtensionIndex, 1),
		Filename:      extensionInfo.Filename,
	}
	resolver.legacyExtensionMap[extensionInfo] = legacyExtensionInfo
	return legacyExtensionInfo
}

// resolve handles the majority of extension resolution logic. This will return
// a proto.Message for the given extension if the message has the tag number of
// the provided extension (or an equivalent legacy extension.) If there was no
// such tag number present in the known or unknown fields, this method will
// return nil. Note that the returned message may be dynamicpb.Message or
// another type, and thus may need to still be reparsed if needed.
func (resolver extensionResolver) resolve(
	options proto.Message,
	extensionType protoreflect.ExtensionType,
) proto.Message {
	msg := resolver.getExtensionOrLegacy(options, extensionType)
	if msg == nil {
		if unknown := options.ProtoReflect().GetUnknown(); len(unknown) > 0 {
			reparsedOptions := options.ProtoReflect().Type().New().Interface()
			if err := (proto.UnmarshalOptions{
				Resolver: resolver.registry,
			}).Unmarshal(unknown, reparsedOptions); err != nil {
				msg = resolver.getExtensionOrLegacy(reparsedOptions, extensionType)
			}
		}
	}
	return msg
}

// getExtensionOrLegacy gets the extension extensionType on message, or if it is
// not found, the corresponding legacy extensionType. Unlike proto.GetExtension,
// this method will not panic if the runtime type of the extension is unexpected
// and returns nil if the extension is not present.
func (resolver extensionResolver) getExtensionOrLegacy(
	message proto.Message,
	extensionType protoreflect.ExtensionType,
) proto.Message {
	reflect := message.ProtoReflect()
	if reflect.Has(extensionType.TypeDescriptor()) {
		extension, _ := reflect.Get(extensionType.TypeDescriptor()).Interface().(protoreflect.Message)
		return extension.Interface()
	}
	// This cast allows us to avoid using a map-of-interfaces for the legacy
	// extensions, which would be undesirable.
	extensionTypeInfo, ok := extensionType.(*protoimpl.ExtensionInfo)
	if !ok {
		return nil
	}
	legacyExtensionType, ok := resolver.legacyExtensionMap[extensionTypeInfo]
	if !ok {
		return nil
	}
	return resolver.getExtensionOrLegacy(message, legacyExtensionType)
}
