package _020_10_20

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
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
		nex := tail.Next

		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex

		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
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
