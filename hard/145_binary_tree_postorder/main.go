package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var nodes = []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	for i := 0; i < len(result); i++ {
		j := len(result) - i - 1
		if i > j {
			break
		}
		result[i], result[j] = result[j], result[i]
	}
	return result
}
