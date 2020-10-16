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
		assert.Equal(t, tbl.output, myAtoi(tbl.input))
	}
}

func Test_myatoi2(t *testing.T) {
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
		assert.Equal(t, tbl.output, myAtoi2(tbl.input))
	}
}
