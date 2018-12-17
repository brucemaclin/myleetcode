package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	slice1 := []rune(word1)
	slice2 := []rune(word2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if slice1[i-1] == slice2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return m + n - 2*dp[m][n]
}
func main() {
	fmt.Println(minDistance("sea", "eat"))

}
