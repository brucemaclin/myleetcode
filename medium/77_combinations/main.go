package main

import "fmt"

func combine(n int, k int) [][]int {
	if k == 0 || n == 0 {
		return nil
	}
	var result [][]int
	if k == 1 {
		for i := 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	var path []int

	backtrack(path, 1, n, k, &result)
	return result
}
func backtrack(path []int, start, n, k int, result *[][]int) {
	if len(path) == k {
		newPath := make([]int, k)
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		backtrack(path, i+1, n, k, result)
		path = path[:len(path)-1]
	}

}
func main() {
	fmt.Println(combine(4, 2))
}
