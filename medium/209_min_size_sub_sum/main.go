package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	var sum int
	var j, i int
	var min = math.MaxInt32
	for j < len(nums) {
		sum += nums[j]
		j++
		for sum >= s {
			if min > j-i {
				min = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
