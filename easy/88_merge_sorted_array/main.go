package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m+n {
		return
	}
	mergeIndex := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[mergeIndex] = nums2[j]
			mergeIndex--
			j--
		} else if j < 0 {
			nums1[mergeIndex] = nums1[i]
			mergeIndex--
			i--
		} else if nums1[i] > nums2[j] {
			nums1[mergeIndex] = nums1[i]
			mergeIndex--
			i--
		} else {
			nums1[mergeIndex] = nums2[j]
			mergeIndex--
			j--
		}
	}
	return
}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
