package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i < v {
				break
			}
			if dp[i] == 0 {
				dp[i] = amount + 1
			}
			if i-v != 0 && dp[i-v] == 0 {
				dp[i-v] = amount + 1
			}
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}

		}
	}
	if dp[amount] > amount || dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 1))
}
