package main

import (
	"fmt"
	"sort"
	"strconv"
)

func combinationSum4(nums []int, target int) int {
	if target == 0 {
		return 0
	}
	sort.Ints(nums)
	var count int
	has := make(map[string]bool)
	dfs(nums, target, nil, &count, has)
	return count

}

func dpCombinationSum4(nums []int, target int) int {
	if target == 0 {
		return 0
	}
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i < v {
				break
			}
			dp[i] += dp[i-v]
		}
	}
	return dp[target]
}

func dfs(nums []int, target int, path []int, count *int, has map[string]bool) {
	if target == 0 {
		key := ""
		for _, v := range path {
			key += strconv.Itoa(v) + "_"
		}
		if !has[key] {
			*count++
		}
		return
	}
	for _, v := range nums {
		if v > target {
			break
		}
		path = append(path, v)
		dfs(nums, target-v, path, count, has)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(dpCombinationSum4([]int{1, 2, 3}, 4))
}
