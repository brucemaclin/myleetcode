package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(j*dp[i-j]), float64(j*(i-j)))))
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(3))
	fmt.Println(integerBreak(10))
}
