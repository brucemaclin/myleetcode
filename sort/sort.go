package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(child+first, child+first+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}
func mySort(data Interface) {
	heapSort(data, 0, data.Len())
	return
}
func heapSort(data Interface, a, b int) {
	lo := 0
	first := a
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first+i, first)
		siftDown(data, lo, i, first)
	}
}
func main() {
	data := []int{1, 2, 7, 5, 3, 8, 9}
	slice := sort.IntSlice(data)
	mySort(slice)
	fmt.Println(data)
}
