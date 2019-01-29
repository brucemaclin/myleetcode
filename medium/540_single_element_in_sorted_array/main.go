package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l := 0
	h := len(nums) - 1
	for l < h {
		mid := l + (h-l)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			h = mid
		}
	}
	return nums[l]
}
func main() {
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
