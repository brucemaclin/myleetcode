package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var total int = 1
	var zeroCount int
	for _, v := range nums {
		if v != 0 {
			total *= v
		} else {

			zeroCount++
		}
	}
	for i, v := range nums {
		if v == 0 {
			if zeroCount > 1 {
				nums[i] = 0
			} else {
				nums[i] = total
			}
		} else {
			if zeroCount > 0 {
				nums[i] = 0
			} else {
				nums[i] = total / v
			}
		}
	}
	return nums
}
func main() {
	var tests = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(tests))
}
