package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	if n <= 0 || n > 55 || k <= 0 {
		return nil
	}
	visited := make(map[int]bool)
	var result [][]int
	var path []int
	backtrack(1, k, n, path, visited, &result)
	return result
}

func backtrack(start, k, n int, path []int, visited map[int]bool, result *[][]int) {
	if n == 0 && len(path) == k {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}
	for i := start; i <= 9; i++ {
		if visited[i] {
			continue
		}
		if i > n {
			break
		}
		newTarget := n - i
		visited[i] = true
		path = append(path, i)
		backtrack(i+1, k, newTarget, path, visited, result)
		visited[i] = false
		path = path[:len(path)-1]
	}
}
func main() {
	fmt.Println(combinationSum3(3, 7))
}
