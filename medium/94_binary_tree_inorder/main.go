package main

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nodes []*TreeNode
	cur := root
	var result []int
	for cur != nil || len(nodes) > 0 {
		for cur != nil {
			nodes = append(nodes, cur)
			cur = cur.Left
		}
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		result = append(result, node.Val)
		cur = node.Right
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
