package main

import (
	"fmt"
	"strings"
)

func findLongestWord(s string, d []string) string {
	var longgest string
	for _, v := range d {
		if len(longgest) > len(v) || len(longgest) == len(v) && strings.Compare(longgest, v) > 0 {
			continue
		}
		if isValid(s, v) {
			longgest = v
		}
	}

	return longgest
}
func isValid(s, d string) bool {
	slice1 := []byte(s)
	slice2 := []byte(d)

	var i, j int
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] == slice2[j] {
			j++

		}
		i++
	}

	return j == len(d)
}
func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
