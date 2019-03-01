package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 || n == 0 {
		return 0
	}
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	iMin := 0
	iMax := m
	half := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := half - i
		if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < m && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else {
			var maxOfLeft int
			var minOfRight int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				if nums2[j-1] > nums1[i-1] {
					maxOfLeft = nums2[j-1]
				} else {
					maxOfLeft = nums1[i-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minOfRight = nums2[j]
				} else {
					minOfRight = nums1[i]
				}
			}
			return float64(maxOfLeft+minOfRight) / 2
		}
	}
	return 0
}
func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
