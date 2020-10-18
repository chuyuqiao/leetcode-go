package _020_10_18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	table := []struct {
		input  []int
		output [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{3, 0, -2, -1, 1, 2},
			[][]int{
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, tbl := range table {
		assert.Equal(t, tbl.output, threeSum(tbl.input))
	}
}
