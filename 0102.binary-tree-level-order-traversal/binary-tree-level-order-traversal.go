package problem0102

import "leetcode-go/pkg/tree"

type TreeNode = tree.Node

func levelOrder(root *TreeNode) [][]int {

	var ret [][]int

	var bfs func(*TreeNode, int)

	bfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[level] = append(ret[level], root.Val)

		bfs(root.Left, level+1)
		bfs(root.Right, level+1)
	}

	bfs(root, 0)

	return ret
}
