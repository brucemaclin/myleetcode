package main

import (
	"fmt"
	"sort"
	"strconv"
)

func combinationSum(candidates []int, target int) [][]int {
	var result = make([][]int, 0)

	sort.Ints(candidates)
	var path []int
	existed := make(map[string]bool)
	backtrack(candidates, target, path, existed, &result)
	return result
}

func backtrack(candidate []int, target int, path []int, existed map[string]bool, result *[][]int) {
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
	for _, v := range candidate {
		if v > target {
			break
		}
		path = append(path, v)
		newTarget := target - v
		backtrack(candidate, newTarget, path, existed, result)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
