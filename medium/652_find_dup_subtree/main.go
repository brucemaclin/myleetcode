package main

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]*TreeNode)
	result := make(map[string]*TreeNode)
	preOrder(root, m, result)
	var ret []*TreeNode
	for _, v := range result {
		ret = append(ret, v)
	}
	return ret
}

func preOrder(node *TreeNode, m map[string]*TreeNode, result map[string]*TreeNode) string {
	if node == nil {
		return "$"
	}
	s := strconv.Itoa(node.Val) + "," + preOrder(node.Left, m, result) + "," + preOrder(node.Right, m, result)
	if n, ok := m[s]; ok {
		result[s] = n
	} else {
		m[s] = node
	}
	return s

}
