package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{
			"",
			"",
		},
		{
			"aaaaa",
			"aaaaa",
		},
		{
			"abccbf",
			"bccb",
		},
		{
			"abcdef",
			"a",
		},
	}

	for _, tbl := range table {
		assert.Equal(t, longestPalindrome(tbl.input), tbl.output)
	}
}
