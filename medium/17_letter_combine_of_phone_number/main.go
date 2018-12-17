package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	slice := []rune(digits)
	var result []string
	doCombine("", &result, slice)
	return result
}
func doCombine(prefix string, result *[]string, slice []rune) {
	fmt.Println(len(prefix), len(slice))
	if len(prefix) == len(slice) {
		*result = append(*result, prefix)
		return
	}

	curRune := slice[len(prefix)]
	runes := numMap[string(curRune)]
	fmt.Println(prefix, result, string(curRune), runes)
	for _, v := range runes {
		prefix += string(v)
		doCombine(prefix, result, slice)
		prefix = string([]rune(prefix)[:len(prefix)-1])
	}
}

var numMap = map[string][]rune{
	"2": {'a', 'b', 'c'},
	"3": {'d', 'e', 'f'},
	"4": {'g', 'h', 'i'},
	"5": {'j', 'k', 'l'},
	"6": {'m', 'n', 'o'},
	"7": {'p', 'q', 'r', 's'},
	"8": {'t', 'u', 'v'},
	"9": {'w', 'x', 'y', 'z'},
}

func main() {
	fmt.Println(letterCombinations("23"))
}
