package problem0021

type ListNode struct {
	Value int
	Next  *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	/**
	l1 1->2->3
	l2 2->4->5
	*/
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Value < l2.Value {
		head, node = l1, l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return head
}
