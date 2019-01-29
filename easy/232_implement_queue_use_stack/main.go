package main

import "container/list"

type MyQueue struct {
	inStack  *list.List
	outStack *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{inStack: new(list.List), outStack: new(list.List)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	e := new(list.Element)
	e.Value = x
	this.inStack.PushBack(e)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.outStack.Len() != 0 {
		e := this.outStack.Back()
		this.outStack.Remove(e)
		get := e.Value.(int)
		return get

	}
	e := this.inStack.Back()
	for e.Prev() != nil {
		prev := e.Prev()
		this.outStack.PushBack(e)
		this.inStack.Remove(e)
		e = prev
	}
	return e.Value.(int)

}

/** Get the front element. */
func (this *MyQueue) Peek() int {

}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
