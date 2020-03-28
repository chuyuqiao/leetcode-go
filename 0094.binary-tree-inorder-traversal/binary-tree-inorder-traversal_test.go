package problem0094

import "testing"

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	t.Log(inorderTraversalRecursive(root))
	t.Log(inorderTraversalIteration(root))
}
