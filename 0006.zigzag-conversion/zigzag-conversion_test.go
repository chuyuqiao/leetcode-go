package problem0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	s := convert("LEETCODEISHIRING", 3)
	assert.Equal(t, "LCIRETOESIIGEDHN", s)
}
