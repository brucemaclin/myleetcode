package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	visited := make(map[*TreeNode]int)
	return robSub(root, visited)
}
func robSub(root *TreeNode, visited map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := visited[root]; ok {
		return v
	}
	val1 := root.Val
	if root.Left != nil {
		val1 += robSub(root.Left.Left, visited) + robSub(root.Left.Right, visited)
	}
	if root.Right != nil {
		val1 += robSub(root.Right.Left, visited) + robSub(root.Right.Right, visited)
	}
	val2 := robSub(root.Left, visited) + robSub(root.Right, visited)
	if val1 > val2 {
		visited[root] = val1
		return val1
	}
	visited[root] = val2
	return val2
}
func better(root *TreeNode) int {
	a, b := betterRobSub(root)
	return max(a, b)
}
func betterRobSub(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	la, lb := betterRobSub(root.Left)
	ra, rb := betterRobSub(root.Right)
	val1 := root.Val + lb + rb
	val2 := max(la, lb) + max(ra, rb)
	return val1, val2
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	root := newNode(3)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Right = newNode(3)
	root.Right.Right = newNode(1)
	fmt.Println(rob(root))

}
