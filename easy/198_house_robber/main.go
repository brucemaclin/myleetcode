package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	pre1 := 0
	pre2 := 0
	for _, v := range nums {
		cur := int(math.Max(float64(pre2+v), float64(pre1)))
		pre2 = pre1
		pre1 = cur
	}
	if pre1 > pre2 {
		return pre1
	}
	return pre2
}
func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
