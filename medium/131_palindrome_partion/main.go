package main

import "fmt"

func partition(s string) [][]string {
	var result [][]string
	var path []string
	doPartition([]rune(s), path, &result)
	return result
}

func doPartition(s []rune, path []string, result *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(s); i++ {
		tmp := string(s[:i+1])
		if isPalind([]rune(tmp), 0, i) {
			path = append(path, tmp)
			doPartition(s[i+1:], path, result)
			path = path[:len(path)-1]
		}
	}
}

func isPalind(s []rune, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	result := partition("aab")
	fmt.Println(result)
}
