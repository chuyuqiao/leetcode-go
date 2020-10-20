package _020_10_20

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	levelOrder := [][]int{}
	if root == nil {
		return levelOrder
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	count := 0
	for len(queue) > 0 {
		count++
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if count%2 == 0 {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-1-i] = level[len(level)-1-i], level[i]
			}
		}
		levelOrder = append(levelOrder, level)
	}
	return levelOrder
}
