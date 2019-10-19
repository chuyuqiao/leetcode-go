package problem0206

// 单链表数据结构
type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode

	for head != nil {
		temp := head.Next

		head.Next = pre
		pre = head

		head = temp
	}
	return pre
}
