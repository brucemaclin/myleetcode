package main

import (
	"fmt"
	"math"
)

func dpNumSquares(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[0] = 0
	squares := generateSquares(n)
	for i := 2; i <= n; i++ {
		min := math.MaxInt64
		for _, v := range squares {
			if v > i {
				break
			}
			min = int(math.Min(float64(min), float64(dp[i-v]+1)))
		}
		dp[i] = min
	}
	return dp[n]
}
func numSquares(n int) int {
	squares := generateSquares(n)
	m := make(map[int]bool)
	var level int
	list := []int{n}
	for {

		length := len(list)
		if length == 0 {
			break
		}
		level++
		for i := 0; i < length; i++ {
			cur := list[i]
			for _, v := range squares {
				next := cur - v
				if next == 0 {
					return level
				}
				if next < 0 {
					break
				}
				if m[next] {
					continue
				}
				m[next] = true
				list = append(list, next)
			}
		}
		list = list[length:]
	}
	return n
}
func generateSquares(n int) []int {
	var result []int
	first := 1
	diff := 3
	for {
		if first <= n {
			result = append(result, first)
		} else {
			break
		}
		first += diff
		diff += 2
	}
	return result

}
func main() {
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(12))

	fmt.Println(dpNumSquares(13))
	fmt.Println(dpNumSquares(12))
	fmt.Println(numSquares(2))
	fmt.Println(dpNumSquares(2))
}
