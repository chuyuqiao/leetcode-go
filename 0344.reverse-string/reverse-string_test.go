package problem0344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	bytes := []byte("abcdefg")
	reverseString(bytes)
	assert.Equal(t, "gfedcba", bytes)
}
