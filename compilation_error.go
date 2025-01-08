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

import "strings"

// A CompilationError is returned if a CEL expression cannot be compiled &
// type-checked or if invalid standard constraints are applied.
type CompilationError struct {
	cause error
}

func (err *CompilationError) Error() string {
	if err == nil {
		return ""
	}
	var sb strings.Builder
	_, _ = sb.WriteString("compilation error: ")
	if err.cause != nil {
		_, _ = sb.WriteString(": ")
		_, _ = sb.WriteString(err.cause.Error())
	}
	return sb.String()
}

func (err *CompilationError) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.cause
}
