package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	var str = A
	var count = 1
	for len(str) < len(B) {
		str += A
		count++
	}
	for i := 0; i < 2; i++ {
		if strings.Contains(str, B) {
			return count
		}
		str += A
		count++

	}
	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))

}
