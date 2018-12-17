package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			if i >= len(v) && string(s[i-len(v):i]) == v {
				dp[i] = dp[i] || dp[i-len(v)]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
