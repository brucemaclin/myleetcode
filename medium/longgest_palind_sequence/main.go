package main

import (
	"fmt"
	"math"
)

func longestPalindromeSubseq(s string) int {
	slice := []rune(s)
	var dp = make([][]int, len(slice))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(slice))
	}
	for i := len(slice) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[0][len(slice)-1]
}
func main() {
	fmt.Println(longestPalindromeSubseq("abbc"))
	fmt.Println(longestPalindromeSubseq("bbbab"))

}
