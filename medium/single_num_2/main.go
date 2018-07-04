package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		if i+2 <= len(nums) {
			if nums[i] == nums[i+2] {
				i += 3
			} else {
				return nums[i]
			}
		} else {
			return nums[i]
		}
	}
	return 0
}
func better(nums []int) int {
	var ones int
	var twos int
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones

	}
	return ones
}
func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 99, 0, 1, 0, 1, 1}))
	fmt.Println(singleNumber(createNums()))
	fmt.Println(better(createNums()))
}
func createNums() []int {
	var nums []int
	for j := 0; j < 3; j++ {
		for i := 0; i < 100; i++ {
			nums = append(nums, i)
		}
	}
	nums = append(nums, 100)
	return nums
}
