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
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
)

// A ValidationError is returned if one or more constraint violations were
// detected.
type ValidationError validate.Violations

func (err *ValidationError) Error() string {
	bldr := &strings.Builder{}
	bldr.WriteString("validation error:")
	for _, violation := range err.Violations {
		bldr.WriteString("\n - ")
		if fieldPath := violation.GetFieldPath(); fieldPath != "" {
			bldr.WriteString(fieldPath)
			bldr.WriteString(": ")
		}
		_, _ = fmt.Fprintf(bldr, "%s [%s]",
			violation.GetMessage(),
			violation.GetConstraintId())
	}
	return bldr.String()
}

// ToProto converts this error into its proto.Message form.
func (err *ValidationError) ToProto() *validate.Violations {
	return (*validate.Violations)(err)
}

// PrefixFieldPaths prepends to the provided prefix to the error's internal
// field paths.
func PrefixFieldPaths(err *ValidationError, format string, args ...any) {
	prefix := fmt.Sprintf(format, args...)
	for _, violation := range err.Violations {
		switch {
		case violation.GetFieldPath() == "": // no existing field path
			violation.FieldPath = proto.String(prefix)
		case violation.GetFieldPath()[0] == '[': // field is a map/list
			violation.FieldPath = proto.String(prefix + violation.GetFieldPath())
		default: // any other field
			violation.FieldPath = proto.String(fmt.Sprintf("%s.%s", prefix, violation.GetFieldPath()))
		}
	}
}
