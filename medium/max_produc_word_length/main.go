package main

import "fmt"

func maxProduct(words []string) int {
	var values []int = make([]int, len(words))
	for i, w := range words {
		for _, r := range w {
			values[i] |= 1 << uint32(r-'a')
		}
	}
	var maxProduct int
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if (values[i]&values[j] == 0) && (len(words[i])*len(words[j]) > maxProduct) {
				maxProduct = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProduct

}
func main() {
	fmt.Println(maxProduct([]string{"a", "abc", "ab", "cd"}))
}
