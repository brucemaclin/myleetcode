package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	slice1 := []rune(s)
	slice2 := []rune(p)
	i:=0
	j:=0
	startIndex := -1
	match := 0
	for i<len(slice1){
		if j <len(slice2) && (slice1[i] == slice2[j]||slice2[j] == '?') {
			i++
			j++
		} else if j <len(slice2) && slice2[j] == '*' {
				startIndex = j
				match = i
				j++
		}  else if startIndex != -1 {
					j = startIndex+1
					match++
					i = match
		} else {
			return false
		}
	}
	for j < len(slice2) && slice2[j] == '*' {
		j++
	}
	return j == len(slice2)
}
//"mississippi"
//"m??*ss*?i*pi"
func main() {
	fmt.Println(isMatch("aa","a"))
	fmt.Println(isMatch("aa","*"))
	fmt.Println(isMatch("cb","?a"))
	fmt.Println(isMatch("adceb","*a*b"))
	fmt.Println(isMatch("acdcb","a*c?b"))
	fmt.Println(isMatch("aab","c*a*b"))
	fmt.Println(isMatch("mississippi","m??*ss*?i*pi"))
}
//int s = 0, p = 0, match = 0, starIdx = -1;
//while (s < str.length()){
//// advancing both pointers
//if (p < pattern.length()  && (pattern.charAt(p) == '?' || str.charAt(s) == pattern.charAt(p))){
//s++;
//p++;
//}
//// * found, only advancing pattern pointer
//else if (p < pattern.length() && pattern.charAt(p) == '*'){
//starIdx = p;
//match = s;
//p++;
//}
//// last pattern pointer was *, advancing string pointer
//else if (starIdx != -1){
//p = starIdx + 1;
//match++;
//s = match;
//}
////current pattern pointer is not star, last patter pointer was not *
////characters do not match
//else return false;
//}
//
////check for remaining characters in pattern
//while (p < pattern.length() && pattern.charAt(p) == '*')
//p++;
//
//return p == pattern.length();