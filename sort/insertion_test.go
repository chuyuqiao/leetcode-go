package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertionSort(t *testing.T) {
	table := []struct {
		input  []int
		output []int
	}{
		{
			[]int{3, 4, 5, 6, 1, 2},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tbl := range table {
		insertionNums := insertionSort(tbl.input)
		assert.Equal(t, tbl.output, insertionNums)
	}
}
