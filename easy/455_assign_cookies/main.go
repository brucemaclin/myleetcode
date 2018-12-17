package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	gLen := len(g)
	sLen := len(s)
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	j := 0
	for i < gLen && j < sLen {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
