package main

import "fmt"

func numIslands(grid [][]byte) int {
	var numsOfIsLands int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] != 0 {
				dfs(grid, i, j, len(grid), len(grid[i]))
				numsOfIsLands++

			}
		}
	}
	return numsOfIsLands
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(grid [][]byte, i, j int, m, n int) {

	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == byte(0) {
		return
	}
	grid[i][j] = 0
	for _, v := range directions {
		dfs(grid, i+v[0], j+v[1], m, n)
	}
}

func main() {
	fmt.Println(numIslands([][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}))
	fmt.Println(numIslands([][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}))
}
