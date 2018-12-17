package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var longgest int
	for k, v := range m {
		plus := m[k+1]
		if longgest < plus+v {
			longgest = plus + v
		}
	}
	return longgest
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))

}
