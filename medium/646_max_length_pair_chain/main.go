package main

import "fmt"

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 || len(pairs[0]) == 0 {
		return 0
	}
	dp := make([]int, len(pairs))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	var ret int
	for i := 0; i < len(dp); i++ {
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret

}
func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))

}
