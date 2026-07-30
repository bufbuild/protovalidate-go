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

package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsUri(t *testing.T) {
	t.Parallel()
	require.True(t, IsURI("A://"))
}

func TestIsHostname(t *testing.T) {
	t.Parallel()
	require.True(t, IsHostname("foo.example.com"))
	require.True(t, IsHostname("A.ISI.EDU"))
	require.False(t, IsHostname("İ"))

	// The 253-character limit excludes the optional trailing dot.
	name253 := strings.Repeat("a", 63) + "." + strings.Repeat("a", 63) + "." +
		strings.Repeat("a", 63) + "." + strings.Repeat("a", 61)
	name254 := strings.Repeat("a", 63) + "." + strings.Repeat("a", 63) + "." +
		strings.Repeat("a", 63) + "." + strings.Repeat("a", 62)
	require.True(t, IsHostname(name253))
	require.True(t, IsHostname(name253+"."))
	require.False(t, IsHostname(name254))
	require.False(t, IsHostname(name254+"."))
}

func TestIsHostAndPort(t *testing.T) {
	t.Parallel()
	require.False(t, IsHostAndPort("example.com:080", false))
	require.False(t, IsHostAndPort("example.com:00", false))
}
