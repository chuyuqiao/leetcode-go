package probleam0003

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TC struct {
	input    string
	expected int
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	tcs := []TC{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.expected, lengthOfLongestSubstring(tc.input))
	}
}
