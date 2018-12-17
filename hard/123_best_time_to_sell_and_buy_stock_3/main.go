package main

import "fmt"

func maxProfit(prices []int) int {
	hold1 := -1 << 31
	hold2 := -1 << 31
	var sell1 int
	var sell2 int
	for _, v := range prices {
		sell2 = max(sell2, hold2+v)
		hold2 = max(hold2, sell1-v)
		sell1 = max(sell1, hold1+v)
		hold1 = max(hold1, -v)
	}
	return sell2
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))

}
