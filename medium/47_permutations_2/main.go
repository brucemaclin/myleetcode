package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	visited := make(map[int]bool)
	var path []int
	var result [][]int
	sort.Ints(nums)
	backtrack(nums, visited, path, &result)
	return result
}
func backtrack(nums []int, visited map[int]bool, path []int, result *[][]int) {
	if len(path) == len(nums) {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}
	for i, v := range nums {
		if i > 0 && !visited[i-1] && v == nums[i-1] {
			continue
		}
		if visited[i] {
			continue
		}
		path = append(path, v)
		visited[i] = true
		backtrack(nums, visited, path, result)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 1}))

}
