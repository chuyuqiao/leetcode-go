package probleam0003

import (
	"github.com/stretchr/testify/assert"
	"leetcode-go/pkg/test"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tcs := []test.TC{
		{
			Input:    "abcabcbb",
			Expected: 3,
		},
		{
			Input:    "bbbbb",
			Expected: 1,
		},
		{
			Input:    "pwwkew",
			Expected: 3,
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.Expected, lengthOfLongestSubstring(tc.Input))
	}
}
