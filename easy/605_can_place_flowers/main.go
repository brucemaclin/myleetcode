package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	var cnt int
	for index, v := range flowerbed {
		if v == 1 {
			continue
		}
		var pre int
		var next int
		if index > 0 {
			pre = flowerbed[index-1]
		} else {
			pre = 0
		}
		if index < len(flowerbed)-1 {
			next = flowerbed[index+1]
		} else {
			next = 0
		}
		if pre == 0 && next == 0 {
			cnt++
			flowerbed[index] = 1
		}
	}
	return cnt >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
