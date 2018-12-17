package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	min := prices[0]
	var max int

	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			profit := prices[i] - min
			if profit > max {
				max = profit
			}
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
