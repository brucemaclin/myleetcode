package main

import (
	"fmt"
	"sort"
	"strconv"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result = make([][]int, 0)

	sort.Ints(candidates)
	var path []int
	existed := make(map[string]bool)
	visited := make(map[int]bool)
	backtrack(candidates, target, path, existed, visited, &result)
	return result
}

func backtrack(candidate []int, target int, path []int, existed map[string]bool, visited map[int]bool, result *[][]int) {
	if target == 0 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		sort.Ints(newPath)
		var key string
		for _, v := range newPath {
			key += strconv.Itoa(v)
		}
		if !existed[key] {
			*result = append(*result, newPath)
			existed[key] = true
		}
		return
	}
	for i, v := range candidate {
		if v > target {
			break
		}
		if i != 0 && candidate[i-1] == candidate[i] && !visited[i-1] {
			continue
		}
		if visited[i] {
			continue
		}
		visited[i] = true
		path = append(path, v)
		newTarget := target - v
		backtrack(candidate, newTarget, path, existed, visited, result)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
