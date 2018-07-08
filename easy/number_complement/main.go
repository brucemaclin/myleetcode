package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	s := fmt.Sprintf("%b", num)
	var final string
	for _, v := range s {
		if v == '0' {
			final += "1"
		} else {
			final += "0"
		}
	}
	result, _ := strconv.ParseInt(final, 2, 32)
	return int(result)
}

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))

}
