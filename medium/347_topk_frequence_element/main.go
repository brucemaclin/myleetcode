package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	countM := make(map[int]int)
	for _, v := range nums {
		count := countM[v]
		count++
		countM[v] = count
	}
	array := make([][]int, len(nums)+1)
	var result []int
	for k, v := range countM {
		get := array[v]

		get = append(get, k)
		array[v] = get

	}
	for i := len(array) - 1; i >= 0 && len(result) < k; i-- {
		if get := array[i]; get != nil {
			result = append(result, get...)
		}
	}
	return result

}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
