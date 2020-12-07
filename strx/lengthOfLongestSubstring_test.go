package strx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	table := []struct {
		input  string
		output int
	}{
		{
			"",
			0,
		},
		{
			"aaaaa",
			1,
		},
		{
			"abbcd",
			3,
		},
	}

	for _, tbl := range table {
		assert.Equal(t, lengthOfLongestSubstring(tbl.input), tbl.output)
	}
}
