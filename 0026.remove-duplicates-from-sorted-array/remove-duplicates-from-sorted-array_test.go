package problem0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, removeDuplicates(test.input))
	}
}
