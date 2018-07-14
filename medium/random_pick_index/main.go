package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Solution struct {
	saved []int
}

func Constructor(nums []int) Solution {
	var tmp Solution

	tmp.saved = nums
	return tmp
}

var randSrc *rand.Rand

func (this *Solution) Pick(target int) int {
	//rand.Seed(time.Now().Unix())
	var count int
	var res = -1
	for i := 0; i < len(this.saved); i++ {
		if this.saved[i] == target {
			count++
			x := rand.Intn(count)
			if x == 0 {
				res = i
			}
		}
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func main() {
	randSrc = rand.New(rand.NewSource(1))
	fmt.Println(sort.SearchInts([]int{1, 2, 3, 3, 3, 3, 3, 3}, 3))
	obj := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
}
