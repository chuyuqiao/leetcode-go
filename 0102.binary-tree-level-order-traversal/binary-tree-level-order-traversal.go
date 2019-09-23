package problem0102

import "leetcode-go/pkg/tree"

type TreeNode = tree.TreeNode

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[level] = append(ret[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return ret
}
