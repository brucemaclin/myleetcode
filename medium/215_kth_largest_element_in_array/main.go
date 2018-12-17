package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	l := 0
	h := len(nums) - 1
	k = len(nums) - k
	for {
		pos := partition(nums, l, h)
		fmt.Println(pos)
		if pos == k {
			return nums[pos]
		}
		if pos > k {
			h = pos - 1
		} else {
			l = pos + 1
		}
	}

}
func partition(nums []int, l, h int) int {
	i := l
	j := h + 1
	for {
		for {
			i++
			if nums[i] < nums[l] && i < h {
				continue
			} else {
				break
			}
		}

		for {
			j--
			if nums[j] > nums[l] && j > l {
				continue
			} else {
				break
			}
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, l, j)
	return j
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 3, 3, 3, 3}, 1))
}
