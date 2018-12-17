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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var result []float64
	slice := []*TreeNode{root}
	for {
		length := len(slice)
		if length == 0 {
			break
		}
		var tmp float64
		for i := 0; i < length; i++ {
			tmp += float64(slice[i].Val)
			if slice[i].Left != nil {
				slice = append(slice, slice[i].Left)
			}
			if slice[i].Right != nil {
				slice = append(slice, slice[i].Right)
			}
		}
		average := tmp / float64(length)
		result = append(result, average)
		slice = slice[length:]
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

func main() {
	root := newNode(3)
	root.Left = newNode(9)
	root.Right = newNode(20)
	root.Right.Left = newNode(15)
	root.Right.Right = newNode(7)
	fmt.Println(averageOfLevels(root))

}
