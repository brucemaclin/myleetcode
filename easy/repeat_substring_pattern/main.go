package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	if s == "" || len(s) == 1 {
		return false
	}
	start := s[0]
	fmt.Println(start)
	var index = 1
	var substr string
	for {
	loop:
		fmt.Println(index)
		if index > len(s) {
			fmt.Println(index, len(s))
			return false
		}
		getIndex := strings.IndexByte(string(s[index:]), start)
		if getIndex == -1 {
			return false
		}
		index += getIndex
		substr = string(s[0:index])
		if len(s)%len(substr) != 0 {
			index += 1
		} else {
			fmt.Println("break with:", substr)
			for i := index; i < len(s); i += len(substr) {
				if substr != string(s[i:i+len(substr)]) {
					index += 1
					goto loop
				}
			}
			return true
		}
	}
	return true

}

func main() {
	fmt.Println(repeatedSubstringPattern("abaababaab"))

}
