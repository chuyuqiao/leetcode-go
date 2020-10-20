package _020_10_20

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//if l1 == nil{
	//	return l2
	//}
	//if l2 == nil{
	//	return l1
	//}
	//if l1.Val < l2.Val{
	//	l1.Next = mergeTwoLists(l1.Next,l2)
	//	return l1
	//}else {
	//	l2.Next = mergeTwoLists(l2.Next,l1)
	//	return l2
	//}
	hair := &ListNode{}
	pre := hair
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return hair.Next
}
