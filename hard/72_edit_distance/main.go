package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	slice1 := []rune(word1)
	slice2 := []rune(word2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if slice1[i-1] == slice2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("", "a"))

}
