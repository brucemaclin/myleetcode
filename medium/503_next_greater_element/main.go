package main

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var stack []int
	var result = make([]int, len(nums))
	var index = 0
	for {
		if index == len(nums) {
			break
		}
		if len(stack) == 0 {
			stack = append(stack, index)
			index++
			continue
		}

		top := stack[len(stack)-1]
		current := nums[index]
		if current > nums[top] {
			result[top] = current
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, index)
			index++
		}

	}

	for {
		if len(stack) == 0 {
			return result
		}
		last := stack[len(stack)-1]
		//fmt.Println("last:", last)
		var setFlag bool
		for i := 0; i < len(nums); i++ {
			fmt.Println(i, last, nums[i], nums[last])
			//	fmt.Println(i, last)
			if nums[i] > nums[last] {
				result[last] = nums[i]
				setFlag = true
				break
			}
		}
		stack = stack[:len(stack)-1]
		if setFlag {
			continue
		}
		result[last] = -1

	}

}

func main() {
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 3, 2, 1}))
	fmt.Println(nextGreaterElements([]int{0, -2, -3}))

}
