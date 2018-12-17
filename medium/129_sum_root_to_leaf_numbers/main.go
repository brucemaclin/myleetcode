package main

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getSum(root, 0)
}

func getSum(n *TreeNode, total int) int {
	if n == nil {
		return 0
	}
	total = total*10 + n.val
	if n.left != nil && n.right != nil {
		return total
	}

	return getSum(n.left, total) + getSum(n.right, total)
}
