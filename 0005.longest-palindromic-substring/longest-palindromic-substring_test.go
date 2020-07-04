package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	palindrome := longestPalindrome("babad")
	assert.Equal(t, "bab", palindrome)
}
