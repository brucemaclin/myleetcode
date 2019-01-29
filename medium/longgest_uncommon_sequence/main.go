package  main

import "fmt"

func findLUSlength(strs []string) int {
	var m = make(map[string]int)
	var maxStr string
	for _, v := range strs {
		if len(v) >= len(maxStr) {
			maxStr = v
		}
		c,ok := m[v]
		if ok {
			c++

		}
		m[v] = c
	}
	c := m[maxStr]
	if c > 1 {
		return -1
	}
	return len(maxStr)
}

func main() {
	fmt.Println( findLUSlength([]string{"abc","cba","ece"}))
}