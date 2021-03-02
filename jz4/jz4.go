package main

import "fmt"

func main() {
	tree := reConstructBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 2, 4, 1, 6, 5, 7})
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	return rebuild(pre, 0, len(pre)-1, vin, 0, len(vin)-1)
}

func rebuild(pre []int, preLeft, preRight int, vin []int, vinLeft, vinRight int) *TreeNode {
	if preLeft > preRight {
		return nil
	}
	root := &TreeNode{Val: pre[preLeft]}
	var rootIndex int
	for i, v := range vin {
		if v == pre[preLeft] {
			rootIndex = i
			break
		}
	}
	root.Left = rebuild(pre, preLeft+1, preLeft+rootIndex-vinLeft, vin, vinLeft, rootIndex-1)
	root.Right = rebuild(pre, preLeft+rootIndex-vinLeft+1, preRight, vin, rootIndex+1, vinRight)
	return root
}
