package main

import "fmt"

func subsequence(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	countM := make(map[int]int)
	for k := range m {
		forward(k, countM, m)
	}
	var max int
	for _, v := range countM {
		if v > max {
			max = v
		}
	}
	return max
}

func forward(i int, countM map[int]int, m map[int]bool) int {
	if _, ok := m[i]; !ok {
		return 0
	}
	if get := countM[i]; get >= 1 {
		return get
	}
	count := forward(i+1, countM, m) + 1
	countM[i] = count
	return count
}
func main() {
	fmt.Println(subsequence([]int{100, 4, 200, 1, 3, 2}))

}
