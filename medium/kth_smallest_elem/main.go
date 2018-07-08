package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 2
	left := new(TreeNode)
	left.Val = 1
	right := new(TreeNode)
	right.Val = 3
	root.Left = left
	root.Right = right
	fmt.Println(kthSmallest(root, 2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	if k < 1 {
		return 0
	}
	var index int
	var result int
	getKthSmallest(root, k, &index, &result)
	if index == k {
		return result
	}
	return -1
}

func getKthSmallest(node *TreeNode, k int, index *int, result *int) int {
	if node == nil {
		return -1
	}
	if 0 == getKthSmallest(node.Left, k, index, result) {
		return 0
	}
	*index++
	if *index == k {
		*result = node.Val
		return 0
	}
	return getKthSmallest(node.Right, k, index, result)
}
