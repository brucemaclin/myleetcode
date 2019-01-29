package main

func secondMin(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil && root.Right == nil {
		return -1
	}
	left := root.Left.Val
	right := root.Right.Val
	if left == root.Val {
		left = secondMin(root.Left)
	}
	if right == root.Val {
		right = secondMin(root.Right)
	}
	if left != -1 && right != -1 {
		return min(left, right)
	}
	if left == -1 && right == -1 {
		return -1
	}
	if left == -1 {
		return right
	}
	return left
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
