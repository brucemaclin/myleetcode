package main

import "fmt"

func mySqrt(x int) int {
	l := 0
	h := x
	for l <= h {
		mid := l + (h-l)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		}
		if sqrt > mid {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return h
}
func main() {
	fmt.Println(mySqrt(8))
}
