package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] += dp[i-1][j] + dp[i][j-1]

		}
	}
	return dp[m][n]
}
func better(m int, n int) int {
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}
	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			dp[i] += dp[i-1]

		}
	}
	return dp[m-1]
}
func main() {
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(better(7, 3))
}
