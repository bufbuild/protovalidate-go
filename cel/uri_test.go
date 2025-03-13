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

package cel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUri(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		str   string
		valid bool
	}{
		{
			"bad_hash",
			"https://example.com##",
			false,
		},
		{
			"valid/authority_path-abempty_with_query_and_fragment",
			"foo://example.com/0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ%20!$&'()*+,;=:@%20",
			true,
		},
		{
			"invalid/fragment_bad_pct-encoded/b",
			"https://example.com#%",
			false,
		},
		{
			"valid/query-extra",
			"https://example.com?/?",
			true,
		},
		{
			"valid/userinfo_sub-delims",
			"https://!$&'()*+,;=@example.com",
			true,
		},
		{
			"valid/port_65536",
			"https://example.com:65536",
			true,
		},
		// {
		// 	"invalid/host_reg-name_pct-encoded_invalid_utf8",
		// 	"https://foo%c3x%96",
		// 	false,
		// },
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			uri := NewURI(test.str)
			if test.valid {
				assert.True(t, uri.uri())
			} else {
				assert.False(t, uri.uri())
			}
		})
	}
}
