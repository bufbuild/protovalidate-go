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

package cel

import (
	"bytes"
	"math"
	"strings"
	"sync"

	"buf.build/go/protovalidate/internal/rules"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/overloads"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"github.com/google/cel-go/ext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

// NewLibrary creates a new CEL library that specifies all of the functions and
// settings required by protovalidate beyond the standard definitions of the CEL
// Specification:
//
//	https://github.com/google/cel-spec/blob/master/doc/langdef.md#list-of-standard-definitions
//
// Using this function, you can create a CEL environment that is identical to
// the one used to evaluate protovalidate CEL expressions.
func NewLibrary() cel.Library {
	return &library{
		uniqueScalarPool: sync.Pool{New: func() any {
			return map[ref.Val]struct{}{}
		}},
		uniqueBytesPool: sync.Pool{New: func() any {
			return map[string]struct{}{}
		}},
	}
}

// library is the collection of functions and settings required by protovalidate
// beyond the standard definitions of the CEL Specification:
//
//	https://github.com/google/cel-spec/blob/master/doc/langdef.md#list-of-standard-definitions
//
// All implementations of protovalidate MUST implement these functions and
// should avoid exposing additional functions as they will not be portable.
type library struct {
	uniqueScalarPool sync.Pool
	uniqueBytesPool  sync.Pool
}

func (l *library) CompileOptions() []cel.EnvOption { //nolint:funlen,gocyclo
	return []cel.EnvOption{
		cel.TypeDescs(protoregistry.GlobalFiles),
		cel.DefaultUTCTimeZone(true),
		cel.CrossTypeNumericComparisons(true),
		cel.EagerlyValidateDeclarations(true),
		// TODO: reduce this to just the functionality we want to support
		ext.Strings(ext.StringsValidateFormatCalls(true)),
		cel.Variable("now", cel.TimestampType),
		cel.Function("unique",
			l.uniqueMemberOverload(cel.BoolType, l.uniqueScalar),
			l.uniqueMemberOverload(cel.IntType, l.uniqueScalar),
			l.uniqueMemberOverload(cel.UintType, l.uniqueScalar),
			l.uniqueMemberOverload(cel.DoubleType, l.uniqueScalar),
			l.uniqueMemberOverload(cel.StringType, l.uniqueScalar),
			l.uniqueMemberOverload(cel.BytesType, l.uniqueBytes),
		),
		cel.Function("getField",
			cel.Overload(
				"get_field_any_string",
				[]*cel.Type{cel.DynType, cel.StringType},
				cel.DynType,
				cel.FunctionBinding(func(values ...ref.Val) ref.Val {
					message, ok := values[0].(traits.Indexer)
					if !ok {
						return types.UnsupportedRefValConversionErr(values[0])
					}
					fieldName, ok := values[1].Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(values[1])
					}
					return message.Get(types.String(fieldName))
				}),
			),
		),
		cel.Function("isNan",
			cel.MemberOverload(
				"double_is_nan_bool",
				[]*cel.Type{cel.DoubleType},
				cel.BoolType,
				cel.UnaryBinding(func(value ref.Val) ref.Val {
					num, ok := value.Value().(float64)
					if !ok {
						return types.UnsupportedRefValConversionErr(value)
					}
					return types.Bool(math.IsNaN(num))
				}),
			),
		),
		cel.Function("isInf",
			cel.MemberOverload(
				"double_is_inf_bool",
				[]*cel.Type{cel.DoubleType},
				cel.BoolType,
				cel.UnaryBinding(func(value ref.Val) ref.Val {
					num, ok := value.Value().(float64)
					if !ok {
						return types.UnsupportedRefValConversionErr(value)
					}
					return types.Bool(math.IsInf(num, 0))
				}),
			),
			cel.MemberOverload(
				"double_int_is_inf_bool",
				[]*cel.Type{cel.DoubleType, cel.IntType},
				cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					num, ok := lhs.Value().(float64)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					sign, ok := rhs.Value().(int64)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					return types.Bool(math.IsInf(num, int(sign)))
				}),
			),
		),
		cel.Function("isHostname",
			cel.MemberOverload(
				"string_is_hostname_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					host, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsHostname(host))
				}),
			),
		),
		cel.Function("isEmail",
			cel.MemberOverload(
				"string_is_email_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					addr, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsEmail(addr))
				}),
			),
		),
		cel.Function("isIp",
			cel.MemberOverload(
				"string_is_ip_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					addr, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIP(addr, 0))
				}),
			),
			cel.MemberOverload(
				"string_int_is_ip_bool",
				[]*cel.Type{cel.StringType, cel.IntType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					addr, aok := args[0].Value().(string)
					vers, vok := args[1].Value().(int64)
					if !aok || !vok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIP(addr, vers))
				})),
		),
		cel.Function("isIpPrefix",
			cel.MemberOverload(
				"string_is_ip_prefix_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					prefix, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIPPrefix(prefix, 0, false))
				})),
			cel.MemberOverload(
				"string_int_is_ip_prefix_bool",
				[]*cel.Type{cel.StringType, cel.IntType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					prefix, pok := args[0].Value().(string)
					vers, vok := args[1].Value().(int64)
					if !pok || !vok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIPPrefix(prefix, vers, false))
				})),
			cel.MemberOverload(
				"string_bool_is_ip_prefix_bool",
				[]*cel.Type{cel.StringType, cel.BoolType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					prefix, pok := args[0].Value().(string)
					strict, sok := args[1].Value().(bool)
					if !pok || !sok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIPPrefix(prefix, 0, strict))
				})),
			cel.MemberOverload(
				"string_int_bool_is_ip_prefix_bool",
				[]*cel.Type{cel.StringType, cel.IntType, cel.BoolType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					prefix, pok := args[0].Value().(string)
					vers, vok := args[1].Value().(int64)
					strict, sok := args[2].Value().(bool)
					if !pok || !vok || !sok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsIPPrefix(prefix, vers, strict))
				})),
		),
		cel.Function("isUri",
			cel.MemberOverload(
				"string_is_uri_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					s, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsURI(s))
				}),
			),
		),
		cel.Function("isUriRef",
			cel.MemberOverload(
				"string_is_uri_ref_bool",
				[]*cel.Type{cel.StringType},
				cel.BoolType,
				cel.FunctionBinding(func(args ...ref.Val) ref.Val {
					s, ok := args[0].Value().(string)
					if !ok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsURIRef(s))
				}),
			),
		),
		cel.Function(overloads.Contains,
			cel.MemberOverload(
				overloads.ContainsString, []*cel.Type{cel.StringType, cel.StringType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					substr, ok := rhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(strings.Contains(value, substr))
				}),
			),
			cel.MemberOverload("contains_bytes", []*cel.Type{cel.BytesType, cel.BytesType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					substr, ok := rhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(bytes.Contains(value, substr))
				}),
			),
		),
		cel.Function(overloads.EndsWith,
			cel.MemberOverload(
				overloads.EndsWithString, []*cel.Type{cel.StringType, cel.StringType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					suffix, ok := rhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(strings.HasSuffix(value, suffix))
				}),
			),
			cel.MemberOverload("ends_with_bytes", []*cel.Type{cel.BytesType, cel.BytesType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					suffix, ok := rhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(bytes.HasSuffix(value, suffix))
				}),
			),
		),
		cel.Function(overloads.StartsWith,
			cel.MemberOverload(
				overloads.StartsWithString, []*cel.Type{cel.StringType, cel.StringType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					prefix, ok := rhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().(string)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(strings.HasPrefix(value, prefix))
				}),
			),
			cel.MemberOverload("starts_with_bytes", []*cel.Type{cel.BytesType, cel.BytesType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					prefix, ok := rhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(rhs)
					}
					value, ok := lhs.Value().([]byte)
					if !ok {
						return types.UnsupportedRefValConversionErr(lhs)
					}
					return types.Bool(bytes.HasPrefix(value, prefix))
				}),
			),
		),
		cel.Function("isHostAndPort",
			cel.MemberOverload("string_bool_is_host_and_port_bool",
				[]*cel.Type{cel.StringType, cel.BoolType}, cel.BoolType,
				cel.BinaryBinding(func(lhs ref.Val, rhs ref.Val) ref.Val {
					val, vok := lhs.Value().(string)
					portReq, pok := rhs.Value().(bool)
					if !vok || !pok {
						return types.Bool(false)
					}
					return types.Bool(rules.IsHostAndPort(val, portReq))
				}),
			),
		),
	}
}

func (l *library) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{
		cel.EvalOptions(
			cel.OptOptimize,
		),
	}
}

func (l *library) uniqueMemberOverload(itemType *cel.Type, overload func(lister traits.Lister) ref.Val) cel.FunctionOpt {
	return cel.MemberOverload(
		itemType.String()+"_unique_bool",
		[]*cel.Type{cel.ListType(itemType)},
		cel.BoolType,
		cel.UnaryBinding(func(value ref.Val) ref.Val {
			list, ok := value.(traits.Lister)
			if !ok {
				return types.UnsupportedRefValConversionErr(value)
			}
			return overload(list)
		}),
	)
}

func (l *library) uniqueScalar(list traits.Lister) ref.Val {
	size, ok := list.Size().Value().(int64)
	if !ok {
		return types.UnsupportedRefValConversionErr(list.Size().Value())
	}
	if size <= 1 {
		return types.Bool(true)
	}
	exist := l.uniqueScalarPool.Get().(map[ref.Val]struct{}) //nolint:errcheck // guaranteed to match
	defer func() {
		clear(exist)
		l.uniqueScalarPool.Put(exist)
	}()
	for i := range size {
		val := list.Get(types.Int(i))
		if _, ok := exist[val]; ok {
			return types.Bool(false)
		}
		exist[val] = struct{}{}
	}
	return types.Bool(true)
}

// uniqueBytes is an overload implementation of the unique function that
// compares bytes type CEL values. This function is used instead of uniqueScalar
// as the bytes ([]uint8) type is not hashable in Go; we cheat this by converting
// the value to a string.
func (l *library) uniqueBytes(list traits.Lister) ref.Val {
	size, ok := list.Size().Value().(int64)
	if !ok {
		return types.UnsupportedRefValConversionErr(list.Size().Value())
	}
	if size <= 1 {
		return types.Bool(true)
	}
	exist := l.uniqueBytesPool.Get().(map[string]struct{}) //nolint:errcheck // guaranteed to match
	defer func() {
		clear(exist)
		l.uniqueBytesPool.Put(exist)
	}()
	for i := range size {
		val := list.Get(types.Int(i)).Value()
		b, ok := val.([]byte)
		if !ok {
			return types.NewErr("expected bytes, got %v", val)
		}
		str := string(b)
		if _, ok := exist[str]; ok {
			return types.Bool(false)
		}
		exist[str] = struct{}{}
	}
	return types.Bool(true)
}

// RequiredEnvOptions returns the options required to have expressions which
// rely on the provided descriptor.
func RequiredEnvOptions(fieldDesc protoreflect.FieldDescriptor) []cel.EnvOption {
	if fieldDesc.IsMap() {
		return append(
			RequiredEnvOptions(fieldDesc.MapKey()),
			RequiredEnvOptions(fieldDesc.MapValue())...,
		)
	}
	if fieldDesc.Kind() == protoreflect.MessageKind ||
		fieldDesc.Kind() == protoreflect.GroupKind {
		return []cel.EnvOption{
			cel.Types(dynamicpb.NewMessage(fieldDesc.Message())),
		}
	}
	return nil
}
