package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	maxDiameter(root, &max)
	return max
}

func maxDiameter(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := maxDiameter(root.Left, max)
	right := maxDiameter(root.Right, max)
	if *max < left+right {
		*max = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1

}
func main() {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	fmt.Println(diameterOfBinaryTree(root))
	root.Left.Left = &TreeNode{}
	root.Left.Left.Left = &TreeNode{}
	fmt.Println(diameterOfBinaryTree(root))

}
