package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	if input == "" {
		return nil
	}
	slice := []rune(input)
	var m = make(map[int]int)
	for i := 0; i < len(slice); i++ {
		if slice[i] == '+' || slice[i] == '-' || slice[i] == '*' {
			partOne := slice[:i]
			partTwo := slice[i+1:]
			set1 := diffWaysToCompute(string(partOne))
			set2 := diffWaysToCompute(string(partTwo))
			for _, v := range set1 {
				for _, v2 := range set2 {
					switch slice[i] {
					case '+':
						m[v+v2] ++
					case '-':
						m[v-v2] ++
					case '*':
						if v2 != 0 {
							m[v*v2] ++
						}
					}
				}
			}
		}
	}
	var result []int
	for k ,v:= range m {
		for i:=0;i<v;i++ {
			result = append(result, k)
		}
	}
	if len(m) == 0 {
		i, err := strconv.Atoi(input)
		if err == nil {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
