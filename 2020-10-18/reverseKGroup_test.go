package _020_10_18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	table := []struct {
		input  []int
		k      int
		output []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{2, 1, 4, 3, 5},
		},
	}

	for _, tbl := range table {
		assert.Equal(t, tbl.output, node2slice(reverseKGroup(slice2node(tbl.input), tbl.k)))
	}
}

func node2slice(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func slice2node(input []int) *ListNode {
	head := &ListNode{Val: input[0]}
	p := head
	for i := 1; i < len(input); i++ {
		p.Next = &ListNode{Val: input[i]}
		p = p.Next
	}
	fmt.Println(head)
	return head
}
