package problem0217

import (
	"github.com/stretchr/testify/assert"
	"leetcode-go/pkg/test"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tcs := []test.TCSB{
		{
			Input:    []int{1, 2, 3, 1},
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3, 4},
			Expected: false,
		},
		{
			Input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			Expected: true,
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.Expected, containsDuplicate(tc.Input), tc.Input)
	}
}
