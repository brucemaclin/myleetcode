package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var slice = []*TreeNode{root}

	for len(slice) > 0 {
		root = slice[0]
		if root.Right != nil {
			slice = append(slice, root.Right)
		}
		if root.Left != nil {
			slice = append(slice, root.Left)
		}
		slice = slice[1:]
	}
	return root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
