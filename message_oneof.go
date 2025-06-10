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

import (
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// oneof performs validation on a oneof union.
type oneofEvaluator struct {
	Fields   []string
	Required bool
}

func (o oneofEvaluator) formatFields() string {
	quoted := make([]string, len(o.Fields))
	for idx, val := range o.Fields {
		quoted[idx] = fmt.Sprintf("'%s'", val)
	}
	return fmt.Sprintf("[%s]", strings.Join(quoted, ", "))

}

func (o oneofEvaluator) Evaluate(_ protoreflect.Message, val protoreflect.Value, cfg *validationConfig) error {
	return o.EvaluateMessage(val.Message(), cfg)
}

func (o oneofEvaluator) EvaluateMessage(msg protoreflect.Message, cfg *validationConfig) error {
	if !cfg.filter.ShouldValidate(msg, msg.Descriptor()) {
		return nil
	}
	err := &ValidationError{}
	if len(o.Fields) > 0 {
		count := 0
		for _, v := range o.Fields {
			fd := msg.Descriptor().Fields().ByName(protoreflect.Name(v))
			if fd != nil && msg.Has(fd) {
				count++
			}
		}

		if o.Required && count != 1 {
			err.Violations = append(err.Violations, &Violation{
				Proto: &validate.Violation{
					RuleId:  proto.String("message.oneof"),
					Message: proto.String(fmt.Sprintf("one of %s must be set", o.formatFields())),
				},
			})
			return err
		}
		if count > 1 {
			err.Violations = append(err.Violations, &Violation{
				Proto: &validate.Violation{
					RuleId:  proto.String("message.oneof"),
					Message: proto.String(fmt.Sprintf("only one of %s can be set", o.formatFields())),
				},
			})
			return err
		}
	}
	return nil
}

func (o oneofEvaluator) Tautology() bool {
	return false
}

var _ messageEvaluator = oneofEvaluator{}
