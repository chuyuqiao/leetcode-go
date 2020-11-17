package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	table := []struct {
		innput int
		output int
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
	}

	for _, tbl := range table {
		assert.Equal(t, tbl.output, climbStairs(tbl.innput))
		assert.Equal(t, tbl.output, climbStairs2(tbl.innput))
	}
}
