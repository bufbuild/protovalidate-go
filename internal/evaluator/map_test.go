package evaluator

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

	kv := kvPairs{}
	for _, tc := range tests {
		test := tc
		t.Run(test.expected, func(t *testing.T) {
			t.Parallel()
			actual := kv.formatKey(test.key)
			assert.Equal(t, test.expected, actual)
		})
	}
}
