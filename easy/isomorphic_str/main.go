package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[rune]rune)
	m2 := make(map[rune]rune)
	slice1 := []rune(s)
	slice2 := []rune(t)
	for i := 0; i < len(s); i++ {
		a := slice1[i]
		b := slice2[i]
		get1, ok := m1[a]
		if !ok {
			get2, ok := m2[b]
			if !ok {
				m1[a] = b
				m2[b] = a
			} else {
				if get2 != a {
					return false
				}
			}
		} else {
			if get1 != b {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("ab", "aa"))
}
