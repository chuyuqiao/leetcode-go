package _020_10_18

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	//hair := &ListNode{Next: head}
	//pre := hair
	//for head != nil {
	//	tail := pre
	//	for i := 0; i < k; i++ {
	//		tail = tail.Next
	//		if tail == nil {
	//			return hair.Next
	//		}
	//	}
	//	nex := tail.Next
	//	// 翻转
	//	head, tail = myReverse(head, tail)
	//	pre.Next = head
	//	tail.Next = nex
	//	// 前进
	//	pre = tail
	//	head = tail.Next
	//}
	//return hair.Next

	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		// 翻转前
		nex := tail.Next
		// 翻转
		head, tail = myReverse(head, tail)
		// 翻转后
		pre.Next = head
		tail.Next = nex

		// 下一次
		pre = tail
		head = tail.Next
	}
	return hair
}

func myReverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	//pre := tail.Next
	//p := head
	//for pre != tail {
	//	nex := p.Next
	//	p.Next = pre
	//	pre = p
	//	p = nex
	//}
	//return tail, head

	pre := tail.Next
	p := head
	for pre != tail {
		nex := p.Next
		p.Next = pre
		pre = p
		p = nex
	}

	return tail, head
}
