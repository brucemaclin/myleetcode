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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return depth(root)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left > right {
		return right + 1
	}
	return left + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func main() {
	root := newNode(3)
	root.Left = newNode(9)
	root.Right = newNode(20)
	root.Right.Left = newNode(15)
	root.Right.Right = newNode(7)
	fmt.Println(minDepth(root))

}
