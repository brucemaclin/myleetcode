package main

import "fmt"

func findCircleNum(M [][]int) int {
	m := len(M)
	visited := make([]bool, m)
	circle := 0
	for i := 0; i < m; i++ {
		if !visited[i] {
			dfs(M, i, visited)
			circle++
		}
	}
	return circle
}

func dfs(M [][]int, i int, visited []bool) {
	visited[i] = true
	for k := 0; k < len(M[i]); k++ {
		if M[i][k] == 1 && !visited[k] {
			dfs(M, k, visited)
		}
	}
}
func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	fmt.Println(findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
}
