package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
