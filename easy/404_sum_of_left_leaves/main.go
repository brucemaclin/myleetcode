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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeftLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeftLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
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
	fmt.Println(sumOfLeftLeaves(root))

}
