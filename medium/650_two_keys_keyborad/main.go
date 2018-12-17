package main

import (
	"fmt"
	"math"
)

func minSteps(n int) int {
	if n <= 1 {
		return 0
	}
	var s int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			s += i
			n = n / i
		}
	}
	return s
}
func dp(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j <= int(math.Sqrt(float64(n))); j++ {
			if i%j == 0 {
				dp[i] = dp[j] + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(minSteps(3))
}
