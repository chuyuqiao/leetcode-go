package _008_string_to_integer_atoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myatoi(t *testing.T) {
	table := []struct {
		input  string
		output int
	}{
		{
			"42",
			42,
		},
		{
			"   -42",
			-42,
		},
		{
			"4193 with words",
			4193,
		},
		{
			"words and 987",
			0,
		},
		{
			"-91283472332",
			-2147483648,
		},
	}

	for _, tbl := range table {
		assert.Equal(t, myAtoi(tbl.input), tbl.output)
	}
}
