package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	var up = 1
	var down = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down

}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))

}
