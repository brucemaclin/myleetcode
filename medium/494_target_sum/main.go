package main

import "fmt"

//P all positive N negative
//sum(P)- sum(N) = S ==> sum(P)-sum(N) +sum(P)+sum(N) = S +sum(P)+sum(N) = S+sum = 2*sum(P)
func findTargetSumWays(nums []int, S int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum < S || (sum+S)&1 == 1 {
		return 0
	}

	sum = (sum + S) / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, v := range nums {
		for i := sum; i >= v; i-- {
			if i >= v {
				dp[i] = dp[i] + dp[i-v]
			}
		}
	}
	return dp[sum]

}
func dfs(nums []int, start int, target int) int {
	if start == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}

	return dfs(nums, start+1, target+nums[start]) +
		dfs(nums, start+1, target-nums[start])

}
func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(dfs([]int{1, 1, 1, 1, 1}, 0, 3))
}
