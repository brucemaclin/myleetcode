package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var stack []rune
	runes := []rune(s)
	for _, v := range runes {
		if v == '{' || v == '(' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if v == '}' && pop != '{' {
				return false
			}
			if v == ']' && pop != '[' {
				return false
			}
			if v == ')' && pop != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
