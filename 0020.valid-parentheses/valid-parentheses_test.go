package problem0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
		{
			input:    "({[]})",
			expected: true,
		},
		{
			input:    "(){[({[]})]}",
			expected: true,
		},
		{
			input:    "((([[[{{{",
			expected: false,
		},
		{
			input:    "(())]]",
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, isValid(test.input), test.expected)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("{{{{{[[[[[((((()))))]]]]]}}}}}")
	}
}
