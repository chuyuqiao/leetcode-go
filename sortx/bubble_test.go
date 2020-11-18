package sortx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubbleSort(t *testing.T) {
	table := []struct {
		input  []int
		output []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{3},
			[]int{3},
		},
		{
			[]int{3, 4, 5, 6, 1, 2},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tbl := range table {
		assert.Equal(t, tbl.output, bubbleSort(tbl.input))
	}
}
