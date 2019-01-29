package main

import "fmt"

func findShortestSubArray(nums []int) int {
	type info struct {
		start int
		end   int
		count int
	}
	m := make(map[int]info)
	for i, v := range nums {
		tmp, ok := m[v]
		if !ok {
			tmp.start = i
			tmp.end = i
			tmp.count = 1
		} else {
			tmp.end = i
			tmp.count++
		}
		m[v] = tmp

	}
	var max int
	var min int
	for _, v := range m {
		if max < v.count {
			max = v.count
			min = v.end - v.start + 1
		} else if max == v.count {
			cur := v.end - v.start + 1
			if min > cur {
				min = cur
			}
		}
	}
	return min
}
func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
