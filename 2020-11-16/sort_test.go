package _020_11_16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sort(t *testing.T) {
	table := []struct {
		input  []int
		output []int
	}{
		{
			[]int{5, 4, 1, 2, 3},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tbl := range table {
		assert.Equal(t, tbl.output, bubbleSort(tbl.input))
		assert.Equal(t, tbl.output, insertionSort(tbl.input))
		assert.Equal(t, tbl.output, quickSort(tbl.input))
	}

}
