package problem0007

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	table := []struct {
		input, output int
	}{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
		{
			1534236469,
			0,
		},
	}

	for _, tbl := range table {
		assert.Equal(t, reverse(tbl.input), tbl.output)
	}

}
