package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {

	r1 := getRob(nums, 0, len(nums)-2)
	r2 := getRob(nums, 1, len(nums)-1)
	if r1 > r2 {
		return r1
	}
	return r2
}

func getRob(nums []int, start, last int) int {
	pre1 := 0
	pre2 := 0
	for i := start; i <= last; i++ {
		cur := int(math.Max(float64(pre2+nums[i]), float64(pre1)))
		pre2 = pre1
		pre1 = cur
	}
	if pre1 > pre2 {
		return pre1
	}
	return pre2
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
