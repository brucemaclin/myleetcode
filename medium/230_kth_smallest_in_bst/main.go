package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func kth(root *TreeNode, k int) int {
	var count int
	var result int
	inOrder(root, k, &count, &result)
	return result
}
func inOrder(node *TreeNode, k int, count *int, result *int) {
	if node == nil {
		return
	}
	inOrder(node.Left, k, count, result)
	*count++
	if *count == k {
		*result = node.Val
		return
	}
	inOrder(node.Right, k, count, result)
}
