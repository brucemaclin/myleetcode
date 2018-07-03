package main

import "fmt"

func bulbSwitch(n int) int {

	var count int
	for i := 1; i*i <= n; i++ {
		count++
	}
	return count
}
func main() {
	fmt.Println(bulbSwitch(6))
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(3))
}
