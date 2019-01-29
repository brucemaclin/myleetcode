package main

import "fmt"

func arrayNesting(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		var cnt int
		for j := i; nums[j] != -1; {
			cnt++
			t := nums[j]
			nums[j] = -1
			j = t
		}
		if max < cnt {
			max = cnt
		}
	}
	return max
}
func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))

}
