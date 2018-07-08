package main

import "fmt"

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}
	magic := make([]int, n+1)
	magic[0] = 1
	magic[1] = 2
	magic[2] = 2
	head := 2
	tail := 3
	num := 1
	var count = 1
	for tail < n {
		for i := 0; i < magic[head]; i++ {
			magic[tail] = num
			if num == 1 && tail < n {
				count++
			}
			tail++
		}
		num ^= 3
		head++
	}
	return count

}
func main() {
	fmt.Println(magicalString(3))
	fmt.Println(magicalString(6))
}
