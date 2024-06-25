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

package evaluator

import (
	"fmt"
	"strconv"

	"github.com/bufbuild/protovalidate-go/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// kvPairs performs validation on a map field's KV Pairs.
type kvPairs struct {
	// KeyConstraints are checked on the map keys
	KeyConstraints value
	// ValueConstraints are checked on the map values
	ValueConstraints value
}

func (m kvPairs) Evaluate(val protoreflect.Value, failFast bool) (err error) {
	var ok bool
	val.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		evalErr := m.evalPairs(key, value, failFast)
		if evalErr != nil {
			errors.PrefixErrorPaths(evalErr, "[%s]", m.formatKey(key.Interface()))
		}
		ok, err = errors.Merge(err, evalErr, failFast)
		return ok
	})
	return err
}

func (m kvPairs) evalPairs(key protoreflect.MapKey, value protoreflect.Value, failFast bool) (err error) {
	evalErr := m.KeyConstraints.Evaluate(key.Value(), failFast)
	errors.MarkForKey(evalErr)
	ok, err := errors.Merge(err, evalErr, failFast)
	if !ok {
		return err
	}

	evalErr = m.ValueConstraints.Evaluate(value, failFast)
	_, err = errors.Merge(err, evalErr, failFast)
	return err
}

func (m kvPairs) Tautology() bool {
	return m.KeyConstraints.Tautology() &&
		m.ValueConstraints.Tautology()
}

func (m kvPairs) formatKey(key any) string {
	switch k := key.(type) {
	case string:
		return strconv.Quote(k)
	default:
		return fmt.Sprintf("%v", key)
	}
}

var _ evaluator = kvPairs{}
