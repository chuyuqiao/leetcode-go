package _020_10_18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	table := []struct {
		input  string
		output bool
	}{
		{
			"[]",
			true,
		},
		{
			"(([]))",
			true,
		},
		{
			"([]",
			false,
		},
		{
			"{}[]",
			true,
		},
		{
			"[]",
			true,
		},
	}
	for _, tbl := range table {
		assert.Equal(t, tbl.output, isValid(tbl.input))
	}
}
