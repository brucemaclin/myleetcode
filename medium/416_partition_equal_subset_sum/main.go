package main

import "fmt"

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, v := range nums {
		for i := sum; i >= 0; i-- {
			if i >= v {
				dp[i] = dp[i] || dp[i-v]
			}
		}
	}
	return dp[sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
