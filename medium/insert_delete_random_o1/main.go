package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	set      map[int]*list.Element
	nodes    []*list.Element
	count    int
	nodeList *list.List
	r        *rand.Rand
}

type node struct {
	index int
	val   int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {

	return RandomizedSet{
		set:      make(map[int]*list.Element),
		count:    0,
		nodeList: new(list.List),
		r:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}

}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	} else {
		tmp := new(node)
		tmp.index = this.count
		tmp.val = val
		element := new(list.Element)
		element.Value = tmp

		this.nodeList.PushFront(element)
		if len(this.nodes) >= (this.count + 1) {
			this.nodes[this.count] = element
		} else {
			this.nodes = append(this.nodes, element)
		}
		this.set[val] = element
		this.count++
		return true
	}
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	element, ok := this.set[val]
	if !ok {
		return false
	}
	tmp := element.Value.(*node)
	index := tmp.index
	lastNodeElement := this.nodes[this.count-1]
	lastNode := lastNodeElement.Value.(*node)
	lastNode.index = index
	lastNodeElement.Value = lastNode
	this.nodes[index] = lastNodeElement
	this.count--
	this.nodeList.Remove(element)
	delete(this.set, val)

	return true

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.count == 0 {
		return -1
	}

	index := this.r.Int31n(int32(this.count))
	fmt.Println("index:", index)
	element := this.nodes[index]
	return element.Value.(*node).val

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)

	param_2 := obj.Insert(2)
	param_3 := obj.GetRandom()
	fmt.Println(param_1, param_2, param_3)
	obj.Remove(2)
	param_3 = obj.GetRandom()
	fmt.Println(param_3)
}
