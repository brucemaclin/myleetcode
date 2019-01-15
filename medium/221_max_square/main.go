package main

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)

	if m == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	var max int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
