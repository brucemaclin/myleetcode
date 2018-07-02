package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	sort.Ints(nums)
	var m = make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := m[itoa(nums[i])]
		if !ok {
			var needAdd []string
			for k := range m {
				k = k + "_" + itoa(nums[i])
				needAdd = append(needAdd, k)
			}
			for _, v := range needAdd {
				m[v] = true
			}
			m[itoa(nums[i])] = true
		}
	}
	for k := range m {

		slice := strings.Split(k, "_")
		var tmp []int
		for _, v := range slice {
			tmp = append(tmp, atoi(v))
		}
		result = append(result, tmp)
	}
	return result
}
func itoa(i int) string {
	return strconv.Itoa(i)
}
func atoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
