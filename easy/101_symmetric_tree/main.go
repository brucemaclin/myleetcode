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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}
func checkSymmetric(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return checkSymmetric(s.Left, t.Right) && checkSymmetric(s.Right, t.Left)
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
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(2)
	root.Left.Left = newNode(3)
	root.Left.Right = newNode(4)
	root.Right.Left = newNode(4)
	root.Right.Right = newNode(3)
	fmt.Println(isSymmetric(root))

}
