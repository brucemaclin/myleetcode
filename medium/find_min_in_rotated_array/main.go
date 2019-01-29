package main

import "fmt"

func findMin(nums []int) int {
	var start = 0
	var end = len(nums) - 1
	for start < end {
		if nums[start] < nums[end] {
			return nums[start]
		}
		mid := (start + end) >> 1
		if nums[mid] >= nums[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
func main() {
	fmt.Println(findMin([]int{3, 1, 2}))
}
