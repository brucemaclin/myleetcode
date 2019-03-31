package main

import "fmt"

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		//start := string(s[i])
		for j := i + len(max); j < len(s); j++ {
			checkStr := s[i : j+1]
			if len(checkStr) > len(max) {
				if isPalindrome(checkStr) {
					max = checkStr
				}
			}
		}
	}
	return max
}

func isPalindrome(s string) bool {
	slice := []rune(s)
	i := 0
	j := len(slice) - 1
	for i < j {
		if slice[i] != slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("bb"))
}
