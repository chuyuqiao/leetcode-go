package problem0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		inputOne []int
		inputTwo []int
		expected []int
	}{
		{
			[]int{},
			[]int{1, 3, 5, 7},
			[]int{1, 3, 5, 7},
		},
		{

			[]int{1, 3, 5, 7},
			[]int{},
			[]int{1, 3, 5, 7},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{2, 4, 6, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{10, 20, 30},
			[]int{1, 2, 3},
			[]int{1, 2, 3, 10, 20, 30},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 8, 10},
		},
		{
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6, 7, 9},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, l2s(mergeTwoLists(s2l(test.inputOne), s2l(test.inputTwo))))
	}
}

// l2s list to slice
func l2s(l *ListNode) (s []int) {
	if l == nil {
		return
	}
	for l != nil {
		s = append(s, l.Value)
		l = l.Next
	}
	return
}

// s2l slice to list
func s2l(nums []int) (l *ListNode) {
	if len(nums) == 0 {
		return
	}

	l = &ListNode{
		Value: nums[0],
	}
	temp := l
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Value: nums[i],
		}
		temp = temp.Next
	}
	return
}
