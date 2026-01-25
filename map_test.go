// Copyright 2023-2025 Buf Technologies, Inc.
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		key      any
		expected string
	}{
		{
			key:      int32(32),
			expected: "32",
		},
		{
			key:      int64(64),
			expected: "64",
		},
		{
			key:      uint32(32),
			expected: "32",
		},
		{
			key:      uint32(64),
			expected: "64",
		},
		{
			key:      true,
			expected: "true",
		},
		{
			key:      false,
			expected: "false",
		},
		{
			key:      `"foobar"`,
			expected: `"\"foobar\""`,
		},
	}

	kv := &kvPairs{}
	for _, tc := range tests {
		test := tc
		t.Run(test.expected, func(t *testing.T) {
			t.Parallel()
			actual := kv.formatKey(test.key)
			assert.Equal(t, test.expected, actual)
		})
	}
}
