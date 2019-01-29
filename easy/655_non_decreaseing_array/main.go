package main

import "fmt"

func checkPossibility(nums []int) bool {
	var cnt int
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		cnt++
		if i-2 >= 0 && nums[i] < nums[i-2] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return cnt <= 1
}
func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
}
