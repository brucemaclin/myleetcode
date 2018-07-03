package main

import "fmt"

func grayCode(n int) []int {
	result := make([]int, 0)
	result = append(result, 0)
	for i := 0; i < n; i++ {
		size := len(result)
		for j := size - 1; j >= 0; j-- {
			result = append(result, result[j]|1<<uint(i))
		}
	}
	return result
}
func main() {
	printBinaryBit(grayCode(0), 1)
	printBinaryBit(grayCode(2), 2)
	printBinaryBit(grayCode(3), 3)
}

func printBinaryBit(data []int, n int) {
	params := "%0" + fmt.Sprintf(`%d`, n) + "b,"
	for _, v := range data {
		fmt.Printf(params, v)
	}
	fmt.Println()
}
