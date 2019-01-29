package main

import (
	"fmt"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var result []string
	backtrack("", root, &result)
	return result

}

func backtrack(path string, node *TreeNode, result *[]string) {
	if path != "" {
		path += "->" + strconv.Itoa(node.Val)
	} else {
		path = strconv.Itoa(node.Val)
	}
	if node.Left == nil && node.Right == nil {
		*result = append(*result, path)
		return
	}

	if node.Left != nil {
		backtrack(path, node.Left, result)
	}
	if node.Right != nil {
		backtrack(path, node.Right, result)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(binaryTreePaths(root))
}
