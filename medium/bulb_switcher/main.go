package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {

	var count int
	for i := 1; i*i <= n; i++ {
		count++
	}
	return count
}
func better(n int) int {
	return int(math.Sqrt(float64(n)))
}
func main() {
	fmt.Println(bulbSwitch(6), better(6))
	fmt.Println(bulbSwitch(1), better(1))
	fmt.Println(bulbSwitch(3), better(3))
}
