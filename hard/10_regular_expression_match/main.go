package main

import (
	"fmt"
)

//1, If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
//2, If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];
//3, If p.charAt(j) == '*':
//here are two sub conditions:
//1   if p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2]  //in this case, a* only counts as empty
//2   if p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.':
//dp[i][j] = dp[i-1][j]    //in this case, aaa  a* counts as multiple a
//or dp[i][j] = dp[i][j-1]   // in this case, aa a* counts as single a
//or dp[i][j] = dp[i][j-2]   // in this case,ba  ba* counts as empty
func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	if s == "" && p== "*" {
		return true
	}


	m := len(p)
	n := len(s)
	dp := make([][]bool, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i:=1;i<=m;i++ {
		if p[i-1] == '*' && i>=2 {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '.' ||  p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				if j>=2 &&(p[j-2] == s[i-1]||p[j-2]=='.') {
					dp[i][j] =dp[i][j-1]||dp[i-1][j]||dp[i][j-2]

				} else {
					if j >=2 {
						dp[i][j] = dp[i][j-2]
					}
				}
			}
		}
	}
	return dp[n][m]

}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("ab", ".*c"))
	fmt.Println(isMatch("",".*"))

}
//
//public boolean match(char[] str, char[] pattern) {
//
//int m = str.length, n = pattern.length;
//boolean[][] dp = new boolean[m + 1][n + 1];
//
//dp[0][0] = true;
//for (int i = 1; i <= n; i++)
//if (pattern[i - 1] == '*')
//dp[0][i] = dp[0][i - 2];
//
//for (int i = 1; i <= m; i++)
//for (int j = 1; j <= n; j++)
//if (str[i - 1] == pattern[j - 1] || pattern[j - 1] == '.')
//dp[i][j] = dp[i - 1][j - 1];
//else if (pattern[j - 1] == '*')
//if (pattern[j - 2] == str[i - 1] || pattern[j - 2] == '.') {
//dp[i][j] |= dp[i][j - 1]; // a* counts as single a
//dp[i][j] |= dp[i - 1][j]; // a* counts as multiple a
//dp[i][j] |= dp[i][j - 2]; // a* counts as empty
//} else
//dp[i][j] = dp[i][j - 2];   // a* only counts as empty
//
//return dp[m][n];
//}
