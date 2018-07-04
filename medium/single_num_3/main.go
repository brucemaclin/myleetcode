package main

import "fmt"

func singleNumber(nums []int) []int {
	var result int
	for i, v := range nums {
		if i == 0 {
			result = v
		} else {
			result = result ^ v

		}
	}
	diff := result & (-result)
	var final = make([]int, 2)
	for _, v := range nums {
		if v&diff != 0 {
			final[0] ^= v
		} else {
			final[1] ^= v
		}
	}
	return final
}
func main() {
	fmt.Println(singleNumber([]int{1, 2, 3, 4, 1, 2, 3, 5}))

}
