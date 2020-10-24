package _020_10_24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	table := []struct {
		input1 []int
		input2 []int
		output float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			[]int{},
			[]int{1},
			1.0,
		},
	}

	for _, tbl := range table {
		assert.Equal(t, tbl.output, findMedianSortedArrays(tbl.input1, tbl.input2))
	}
}
