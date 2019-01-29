package main

import "fmt"

func dailyTemperature(data []int) []int {
	if len(data) == 0 {
		return nil
	}
	var stack []int
	var result = make([]int, len(data))
	var index = 0
	for {
		if index == len(data) {
			break
		}
		if len(stack) == 0 {
			stack = append(stack, index)
			index++
			continue
		}

		top := stack[len(stack)-1]
		current := data[index]
		if current > data[top] {
			result[top] = index - top
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, index)
			index++
		}

	}
	return result
}
func main() {
	fmt.Println(dailyTemperature([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
