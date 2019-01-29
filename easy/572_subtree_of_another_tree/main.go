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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtreeWithRoot(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeWithRoot(root *TreeNode, t *TreeNode) bool {
	if root == nil && t == nil {
		return true
	}
	if root == nil || t == nil {
		return false
	}
	if root.Val != t.Val {
		return false
	}
	return isSubtreeWithRoot(root.Left, t.Left) && isSubtreeWithRoot(root.Right, t.Right)
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
	root1 := newNode(3)
	root1.Left = newNode(4)
	root1.Right = newNode(5)
	root1.Left.Left = newNode(1)
	root1.Left.Right = newNode(2)
	root2 := newNode(4)
	root2.Left = newNode(1)
	root2.Right = newNode(2)
	fmt.Println(isSubtree(root1, root2))
}
