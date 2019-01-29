package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := startWithRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
	return ret
}

func startWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if sum == root.Val {
		ret++
	}
	ret += startWithRoot(root.Left, sum-root.Val) + startWithRoot(root.Right, sum-root.Val)
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
