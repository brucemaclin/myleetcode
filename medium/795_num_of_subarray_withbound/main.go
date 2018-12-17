package main

func numSubarrayBoundedMax(A []int, L int, R int) int {
	result := 0
	left := -1
	right := -1
	for i := 0; i < len(A); i++ {
		if A[i] > R {
			left = i
		}
		if A[i] >= L {
			right = i
		}
		result += right - left
	}
	return result
}
