package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pprev := 1
	prev := 2
	count := 0
	for i := 3; i <= n; i++ {
		count = prev + pprev
		pprev = prev
		prev = count
	}
	return prev
}

func main() {
	fmt.Println(climbStairs(3), climbStairs(10))
}
