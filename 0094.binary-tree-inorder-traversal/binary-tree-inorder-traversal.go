package problem0094

import "leetcode-go/pkg/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, inorderTraversalRecursive(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversalRecursive(root.Right)...)
	}
	return res
}

func inorderTraversalIteration(root *TreeNode) []int {
	var res []int
	s := stack.New()
	curr := root
	for curr != nil || s.Len() > 0 {
		for curr != nil {
			s.Push(curr)
			curr = curr.Left
		}
		pop := s.Pop().(*TreeNode)
		res = append(res, pop.Val)
		curr = pop.Right
	}
	return res
}
