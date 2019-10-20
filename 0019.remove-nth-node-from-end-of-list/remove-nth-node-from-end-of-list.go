package problem0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd {
		// 删除头节点
		return head.Next
	}

	// 删除中间节点
	d.Next = d.Next.Next

	return head
}

// 获取倒数第N个节点的父节点
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head
	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}

		n--
		head = head.Next
	}
	// n==0 说明链的长度为n
	headIsNthFromEnd = n == 0
	return
}
