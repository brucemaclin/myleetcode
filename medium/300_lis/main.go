package main

import (
	"fmt"
	"math"
)

func nlognLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tails := make([]int, len(nums))

	var count int
	for i := 0; i < len(nums); i++ {
		index := binarySearch(tails, count, nums[i])
		tails[index] = nums[i]
		if index == count {
			count++
		}
	}
	return count
}
func binarySearch(tails []int, count, target int) int {
	l := 0
	h := count
	for l < h {
		m := l + (h-l)/2
		if tails[m] == target {
			return m
		}
		if tails[m] > target {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = int(math.Max(float64(max), float64(dp[j]+1)))
			}
		}
		dp[i] = max
	}
	var ret int
	for i := 0; i < len(dp); i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(nlognLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(nlognLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
