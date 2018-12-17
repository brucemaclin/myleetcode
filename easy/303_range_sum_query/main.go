package main

import "fmt"

type NumArray struct {
	subSum []int
}

func Constructor(nums []int) NumArray {
	tmp := NumArray{}
	tmp.subSum = make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		tmp.subSum[i] = tmp.subSum[i-1] + nums[i-1]
	}
	return tmp
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.subSum[j+1] - this.subSum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
