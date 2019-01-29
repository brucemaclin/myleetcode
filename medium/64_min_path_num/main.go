package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j >= 1 {
					dp[j] = dp[j-1]
				}

			} else {
				if j >= 1 {
					dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j])))
				}

			}
			dp[j] += grid[i][j]
			//fmt.Println(i, j, dp[j])
		}
	}
	return dp[n-1]
}
func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2}, {1, 1}}))
}
