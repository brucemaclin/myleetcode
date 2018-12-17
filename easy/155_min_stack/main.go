package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: math.MaxInt64}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.stack = append(this.stack, this.min)
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	a := this.stack[len(this.stack)-1]
	if a == this.min {
		this.min = this.stack[len(this.stack)-2]
		this.stack = this.stack[:len(this.stack)-2]
	} else {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())

}
