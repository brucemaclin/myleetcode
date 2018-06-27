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
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	}
func convertBST(root *TreeNode) *TreeNode {
	curSum := 0
	convert(root,&curSum)
	return root
}
func convert(root *TreeNode, curSum *int) {
	if root == nil {
		return
	}
	convert(root.Right,curSum)
	*curSum += root.Val
	root.Val = *curSum
	convert(root.Left,curSum)
	return
}
func main() {
	root := new(TreeNode)
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 13
	root = convertBST(root)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
}