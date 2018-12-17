package main

import "fmt"

func permute(nums []int) [][]int {
	visited := make(map[int]bool)
	var path []int
	var result [][]int
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
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5, 4, 2, 6}))
}
