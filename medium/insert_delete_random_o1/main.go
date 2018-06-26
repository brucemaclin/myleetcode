package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set      map[int]int64
	nodeCount int64
	nodes     []int
}

const maxCap = 1 << 32
const minInt32 = -(1 << 31)

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{set:make(map[int]int64,nodes:make([]int,1),}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	index,ok := this.set[val]
	if ok {
		return false
	}
	this.set[val] = this.nodeCount
	this.nodes[this.nodeCount] = val
	this.nodes = append(this.nodes,val)
	this.nodeCount++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index,ok := this.set[val]
	if !ok {
		return false
	}

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.nodeCount == 0 {
		return 0
	}
	index := rand.Int63n(this.nodeCount)
	return this.nodes[index]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
