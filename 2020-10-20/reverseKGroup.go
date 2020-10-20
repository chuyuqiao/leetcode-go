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

		head, tail = myReserve(head, tail)

		pre.Next = head
		tail.Next = nex

		pre = tail
		head = nex

	}

	return hair.Next
}

func myReserve(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	for pre != tail {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return tail, head
}
