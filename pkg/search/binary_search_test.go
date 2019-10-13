package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type bsTest struct {
	inNums   []int
	val      int
	expected int
}

var tests = []bsTest{
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      2,
		expected: 1,
	},
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      8,
		expected: 4,
	},
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      6,
		expected: 3,
	},
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      12,
		expected: 6,
	},
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      0,
		expected: -1,
	},
	{
		inNums:   []int{1, 2, 3, 6, 8, 9, 12},
		val:      13,
		expected: -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, BinarySearch(test.inNums, test.val), test.expected)
	}
}
