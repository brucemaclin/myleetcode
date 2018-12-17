package main

import "fmt"

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	var result [][]int
	n := len(matrix[0])
	var boolP = make(map[string]bool)
	var boolA = make(map[string]bool)
	for i := 0; i < m; i++ {
		dfs(matrix, i, 0, m, n, boolP)
		dfs(matrix, i, n-1, m, n, boolA)
	}
	for i := 0; i < n; i++ {
		dfs(matrix, 0, i, m, n, boolP)
		dfs(matrix, m-1, i, m, n, boolA)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			key := fmt.Sprintf("%d-%d", i, j)
			if boolP[key] && boolA[key] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(matrix [][]int, i, j, m, n int, reach map[string]bool) {
	key := fmt.Sprintf("%d-%d", i, j)
	if reach[key] {
		return
	}
	reach[key] = true
	for _, v := range directions {
		nextI := i + v[0]
		nextJ := j + v[1]
		if nextI < 0 || nextJ < 0 || nextI >= m || nextJ >= n || matrix[nextI][nextJ] < matrix[i][j] {
			continue
		}
		dfs(matrix, nextI, nextJ, m, n, reach)
	}
}

func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}
