package  main

import "fmt"

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return false
	}
	if len(word) == 1 {
		return true
	}
	firstRune := []rune( word)[0]
	var firstUpperFlag bool
	if checkRuneUpper(firstRune) {
		firstUpperFlag = true
	}
	var allUpperFlag bool
	var secondUppserFlag bool
	secondRune := []rune(word)[1]
	if checkRuneUpper(secondRune) {
		secondUppserFlag = true
	}
	if firstUpperFlag && secondUppserFlag {
		allUpperFlag = true
	}
	if !firstUpperFlag && secondUppserFlag {
		return false
	}
	slice := []rune(word)
	for i:=2;i<len(slice);i++ {
		if allUpperFlag {
			if !checkRuneUpper(slice[i]) {
				return false
			}
		} else {
			if checkRuneUpper(slice[i]) {
				return false
			}
		}
	}
	return true

}
func perfectCheck(word string) bool {
	slice := []rune(word)
	firstRUne := slice[0]
	var cnt int
	for i := 0; i < len(slice); i++ {
		if 'Z' - slice[i] >= 0 {
			cnt++
		}
	}
	return cnt ==0 || cnt == len(slice)||(cnt==1 && ('Z'-firstRUne) >= 0)
}
func checkRuneUpper(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		 return true
	}
	return false
}
func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Google"))
	fmt.Println(detectCapitalUse("flagG"))

	fmt.Println(perfectCheck("Google"))
	fmt.Println(perfectCheck("flagG"))
}