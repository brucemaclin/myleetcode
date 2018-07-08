package main

import (
	"fmt"
	"math"
)

func missingNumber(nums []int) int {
	length := int(math.Ceil(float64(len(nums)+1) / float64(32)))
	fmt.Println(length)
	ints := make([]int, length)
	for i := range ints {
		ints[i] = 0xffffffff
	}
	for _, v := range nums {
		index := v / 32
		pos := v % 32

		ints[index] ^= 1 << uint(pos)
	}
	for index, v := range ints {
		if v != 0x0 {
			fmt.Printf("%b\n", v)

			var count int

			for v&0x1 == 0 {
				count++
				v = v >> 1
			}
			return index*32 + count
		}
	}
	return 0
}
func better(nums []int) int {
	var res = len(nums)
	for i, v := range nums {
		res = res ^ i ^ v
	}
	return res
}
func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 8}))
	fmt.Println(better([]int{0, 1, 2, 3, 4, 5, 6, 8}))
}
