package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var nums []int
	inOrder(root, &nums)
	var i = 0
	var j = len(nums) - 1
	for {
		if i >= j {
			break
		}
		add := nums[i] + nums[j]
		if add == k {
			return true
		}
		if add > k {
			j--
		} else {
			i++
		}
	}
	return false
}

func inOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
