package _006_zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Convert(t *testing.T) {
	s := convert("PAYPALISHIRING", 3)
	assert.Equal(t, "PAHNAPLSIIGYIR", s)
}
