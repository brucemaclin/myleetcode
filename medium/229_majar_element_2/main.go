package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0
	for _, v := range nums {
		if candidate1 == v {
			count1++
		} else if candidate2 == v {
			count2++
		} else if count1 == 0 {
			candidate1 = v
			count1 = 1
		} else if count2 == 0 {
			candidate2 = v
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	n := len(nums) / 3
	count1 = 0
	count2 = 0
	for _, v := range nums {
		if v == candidate1 {
			count1++
		} else if v == candidate2 {
			count2++
		}
	}
	var result []int
	if count1 > n {
		result = append(result, candidate1)
	}
	if count2 > n {
		result = append(result, candidate2)
	}
	return result
}
func partition(nums []int, l, h int) int {
	i := l + 1
	j := h
	v := nums[l]
	for {
		for {
			if i != h && nums[i] < v {
				i++
			} else {
				break
			}
		}
		for {
			if j != l && nums[j] > v {
				j--
			} else {
				break
			}
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
