package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	j := int(math.Sqrt(float64(c)))
	i := 0
	for i <= j {
		tmp := i*i + j*j
		if tmp == c {
			return true
		}
		if tmp > c {
			j--
		} else {
			i++
		}

	}
	return false
}
func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(9))
}
