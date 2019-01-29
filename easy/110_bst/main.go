package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var result = true
	depth(root, &result)

	return true
}
func depth(root *TreeNode, result *bool) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, result)
	right := depth(root.Right, result)
	if math.Abs(float64(left-right)) > 1 {
		*result = false
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	fmt.Println(isBalanced(root))
	root.Left.Left = &TreeNode{}
	root.Left.Left.Left = &TreeNode{}
	fmt.Println(isBalanced(root))

}
