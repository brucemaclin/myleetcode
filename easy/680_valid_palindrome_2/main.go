package main

import "fmt"

func validPalindrome(s string) bool {
	slice := []rune(s)
	i := 0
	j := len(slice) - 1
	for i < j {
		if slice[i] != slice[j] {
			return isPalindrome(slice, i+1, j) || isPalindrome(slice, i, j-1)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(slice []rune, i, j int) bool {

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
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
}
