package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var prevBuy int
	var prevSell int
	var buy = int(math.MinInt32)
	var sell int
	for _, v := range prices {
		prevBuy = buy
		if prevSell-v > prevBuy {
			buy = prevSell - v
		}
		prevSell = sell
		if prevBuy+v > prevSell {
			sell = prevBuy + v
		}
	}
	return sell
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
