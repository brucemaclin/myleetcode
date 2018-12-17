package main

import "fmt"

func sortColors(nums []int) {
	var zero = -1
	ones := 0
	twos := len(nums) - 1
	for ones < twos {
		if nums[ones] == 0 {
			zero++
			swap(nums, zero, ones)
			ones++
		} else if nums[ones] == 2 {
			swap(nums, ones, twos)
			twos--
		} else {
			ones++
		}

	}
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)

}
