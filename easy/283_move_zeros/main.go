package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	l := 0
	h := len(nums) - 1
	index := -1
	for l <= h {
		if nums[l] == 0 {
			if index == -1 {
				index = l
			}
			l++
		} else {
			if index != -1 {
				tmp := index
				if l > index {

					nums[l], nums[index] = nums[index], nums[l]
					for j := index + 1; j <= h; j++ {
						if nums[j] == 0 {
							index = j
							break
						}
					}

					if index == tmp {
						index = -1
					}
				}

			}

			l++
		}
	}

}
func better(nums []int) {
	insertPos := 0
	for _, v := range nums {
		if v != 0 {
			nums[insertPos] = v
			insertPos++
		}
	}
	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}

}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{2, 1}
	moveZeroes(nums)
	fmt.Print(nums)
}
