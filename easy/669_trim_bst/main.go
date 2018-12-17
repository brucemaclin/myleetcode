package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == L {
		root.Left = nil
	}
	if root.Val == R {
		root.Right = nil
	}
	if root.Val > L {
		root.Left = trimBST(root.Left, L, R)
	}
	if root.Val < R {
		root.Right = trimBST(root.Right, L, R)
	}
	if root.Val < L {
		root = trimBST(root.Right, L, R)
	}
	if root.Val > R {
		root = trimBST(root.Left, L, R)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
