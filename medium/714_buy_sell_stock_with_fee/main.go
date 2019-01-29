package main

import (
	"fmt"
)

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := make([]int, len(prices))
	hold := make([]int, len(prices))
	skip := make([]int, len(prices))
	sell := make([]int, len(prices))
	buy[0] -= prices[0]
	hold[0] -= prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i] = max(sell[i-1], skip[i-1]) - prices[i]
		hold[i] = max(hold[i-1], buy[i-1])
		skip[i] = max(skip[i-1], sell[i-1])
		sell[i] = max(hold[i-1], buy[i-1]) + prices[i] - fee
	}
	return max(buy[len(prices)-1], max(hold[len(prices)-1], max(sell[len(prices)-1], skip[len(prices)-1])))

}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
