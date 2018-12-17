package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	set := "aeiouAEIOU"
	var result = make([]rune, len(s))
	i := 0
	j := len(s) - 1
	tmp := []rune(s)
	for i <= j {
		if !strings.Contains(set, string(tmp[i])) {
			result[i] = tmp[i]
			i++
		} else if !strings.Contains(set, string(tmp[j])) {
			result[j] = tmp[j]
			j--
		} else {
			result[i], result[j] = tmp[j], tmp[i]
			i++
			j--
		}
	}
	return string(result)
}
func main() {
	fmt.Println(reverseVowels("leetcode"))
}
