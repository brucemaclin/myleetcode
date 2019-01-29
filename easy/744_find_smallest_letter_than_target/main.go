package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	if len(letters) <= 0 {
		return 0
	}
	if target > letters[len(letters)-1] {
		return letters[0]
	}
	l := 0
	h := len(letters) - 1
	for l <= h {
		mid := l + (h-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	if l >= len(letters) {
		return letters[0]
	}
	return letters[l]
}
func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
}
