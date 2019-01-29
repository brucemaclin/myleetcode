package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	slice := []rune(s)
	type info struct {
		substr    string
		lastIndex int
	}

	var m = make(map[string]info)
	var max int
	for i := 0; i < len(slice); i++ {
		for j := i; j >= 0; j-- {
			if j == i {
				if max < 1 {
					max = 1
				}
				var tmp info
				tmp.lastIndex = j
				tmp.substr = string(slice[i])
				m[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = tmp
			} else {
				prev := m[strconv.Itoa(j)+"_"+strconv.Itoa(i-1)]
				fmt.Println("prev:", prev)
				if strings.Contains(prev.substr, string(slice[i])) {
					m[strconv.Itoa(j)+"_"+strconv.Itoa(i)] = prev
					if max < len(prev.substr) {
						max = len(prev.substr)
					}
				} else {
					if prev.lastIndex == i-1 {
						var tmp info
						tmp.substr = prev.substr + string(slice[i])
						tmp.lastIndex = i
						m[strconv.Itoa(j)+"_"+strconv.Itoa(i)] = tmp

						if max < len(prev.substr)+1 {
							max = len(prev.substr) + 1
						}
					} else {
						m[strconv.Itoa(j)+"_"+strconv.Itoa(i)] = prev
					}
				}
			}
		}
	}
	return max
}
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcfsac"))
}
