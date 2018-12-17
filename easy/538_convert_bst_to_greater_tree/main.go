package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	curSum := 0
	convert(root, &curSum)
	return root
}
func convert(root *TreeNode, curSum *int) {
	if root == nil {
		return
	}
	convert(root.Right, curSum)
	*curSum += root.Val
	root.Val = *curSum
	convert(root.Left, curSum)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
