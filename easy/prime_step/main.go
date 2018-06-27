package main

import (
	"math/big"
	"fmt"
)

func step(g, m, n int) []int {
	// your code
	if n < m || n-m <g {

		return nil
	}
	var result []int
	for i:=m;i<=n;i++ {
		if i+g > n {
			return nil
		}
		ints := big.NewInt(int64(i))
		isPrime :=  ints.ProbablyPrime(20)
		fmt.Println(isPrime, i)
		if isPrime {
			bigger := big.NewInt(int64(i+g))
			if bigger.ProbablyPrime(20) {
				result = append(result,i,i+g)
				return result
			}
		}
	}
	return nil
}

func main() {

	result := step(1,2,5)
	fmt.Println(result)
}