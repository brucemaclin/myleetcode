package main

import "fmt"

func frequencySort(s string) string {
	slice := []rune(s)
	m := make(map[rune]int)
	for _, v := range slice {
		m[v]++
	}
	var result = make([][]rune, len(slice)+1)
	for k, v := range m {
		result[v] = append(result[v], k)
	}
	var finalStr string
	for i := len(result) - 1; i >= 0; i-- {
		get := result[i]
		for _, r := range get {
			for j := 0; j < i; j++ {
				finalStr += string(r)
			}
		}
	}
	return finalStr
}
func main() {
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("tree"))
}
