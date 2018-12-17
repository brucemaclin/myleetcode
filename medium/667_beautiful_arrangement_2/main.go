package main

import "fmt"

func constructArray(n int, k int) []int {
	ret := make([]int, n)
	ret[0] = 1
	i := 1
	j := k
	for {
		if i > k {
			break
		}
		if i%2 == 0 {
			ret[i] = ret[i-1] - j
		} else {
			ret[i] = ret[i-1] + j

		}

		i++
		j--
	}
	for i := k + 1; i < n; i++ {
		ret[i] = i + 1
	}
	return ret
}
func main() {
	//fmt.Println(constructArray(3, 1))
	//mt.Println(constructArray(3, 2))
	fmt.Println(constructArray(5, 4))
}
