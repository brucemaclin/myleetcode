package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	var path []string
	var count int
	dfs(s, path, &count)

	return count
}

func dpNumDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	slice := []rune(s)
	if slice[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= len(slice); i++ {
		a, _ := strconv.Atoi(string(slice[i-1 : i]))
		if a != 0 {
			dp[i] += dp[i-1]
		}
		if slice[i-2] == '0' {
			continue
		}
		b, _ := strconv.Atoi(string(slice[i-2 : i]))
		if b <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
func dfs(s string, path []string, result *int) {
	if len(s) == 0 {
		*result++
		return
	}
	slice := []rune(s)

	for i := 0; i < len(slice); i++ {
		if slice[0] == '0' {
			break
		}
		tmp := string(slice[:i+1])

		a, _ := strconv.Atoi(tmp)
		if a > 26 {
			break
		}
		path = append(path, string('A'+a-1))
		dfs(string(slice[i+1:]), path, result)
		path = path[:len(path)-1]

	}
}

func main() {
	fmt.Println(numDecodings("2210"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("1168963884134125126536966946586868418993819971673459188478711546479621135253876271366851168524933185"))
	fmt.Println(dpNumDecodings("2210"))
	fmt.Println(dpNumDecodings("12"))
	fmt.Println(dpNumDecodings("0"))
	fmt.Println(dpNumDecodings("1168963884134125126536966946586868418993819971673459188478711546479621135253876271366851168524933185"))
}
