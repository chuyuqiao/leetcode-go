package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(printListFromTailToHead(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}))
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	var arr []int
	if head == nil {
		return arr
	}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	size := len(arr)
	for i := 0; i < size/2; i++ {
		arr[i], arr[size-1-i] = arr[size-1-i], arr[i]
	}
	return arr
}
