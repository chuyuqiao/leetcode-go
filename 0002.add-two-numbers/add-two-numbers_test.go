package probleam0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
		},
	}
	l2 := &ListNode{
		Val: 1,
	}

	sum := addTwoNumbers(l1, l2)

	assert.Equal(t, 0, sum.Val)
	assert.Equal(t, 0, sum.Next.Val)
	assert.Equal(t, 1, sum.Next.Next.Val)
}
