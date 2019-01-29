package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	sSlice := []rune(s)
	tSlice := []rune(t)
	i := 0
	j := 0
	for i < len(sSlice) && j < len(tSlice) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(sSlice)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
