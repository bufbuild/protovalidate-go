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
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// timestampPB is a specialized evaluator for applying some validate.TimestampRules (only the
// `valid` rule currently) to a google.protobuf.Timestamp message. This is handled outside
// CEL which handles google.protobuf.Timestamp as an abstract type, thus not allowing access
// to the message fields.
type timestampPB struct {
	SecondsDescriptor protoreflect.FieldDescriptor
	NanosDescriptor   protoreflect.FieldDescriptor
	Valid             bool
}

func (t timestampPB) Evaluate(val protoreflect.Value, failFast bool) error {
	seconds := val.Message().Get(t.SecondsDescriptor).Int()
	nanos := val.Message().Get(t.NanosDescriptor).Int()

	timestamp := &timestamppb.Timestamp{Seconds: seconds, Nanos: int32(nanos)}

	err := &errors.ValidationError{Violations: []*validate.Violation{}}
	if t.Valid {
		terr := timestamp.CheckValid()
		if terr != nil {
			err.Violations = append(err.Violations, &validate.Violation{
				ConstraintId: "timestamp.valid",
				Message:      terr.Error(),
			})
			if failFast {
				return err
			}
		}
	}

	if len(err.Violations) > 0 {
		return err
	}
	return nil
}

func (t timestampPB) Tautology() bool {
	return !t.Valid
}

var _ evaluator = timestampPB{}
