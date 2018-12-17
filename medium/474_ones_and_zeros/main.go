package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for _, v := range strs {
		slice := []rune(v)
		ones := 0
		zeros := 0
		for _, r := range slice {
			if r == '0' {
				zeros++
			} else {
				ones++
			}
		}
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i][j] < dp[i-zeros][j-ones]+1 {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
