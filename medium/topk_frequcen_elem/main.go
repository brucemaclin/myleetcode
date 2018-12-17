package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type frequceInfo struct {
	val   int
	count int
}
type infoHeap []*frequceInfo

func (info infoHeap) Len() int {
	return len(info)
}
func (info infoHeap) Swap(i, j int) {
	info[i], info[j] = info[j], info[i]
}
func (info infoHeap) Less(i, j int) bool {
	return info[i].count < info[j].count
}
func (pq *infoHeap) Push(x interface{}) {
	item := x.(*frequceInfo)
	*pq = append(*pq, item)
}

func (pq *infoHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)
	for _, v := range nums {
		c := m[v]
		c++
		m[v] = c
	}

	var infos = &infoHeap{}
	heap.Init(infos)
	var minItem *frequceInfo
	for val, count := range m {

		var tmp frequceInfo
		tmp.val = val
		tmp.count = count
		if minItem == nil && infos.Len() >= k {
			minItem = heap.Pop(infos).(*frequceInfo)

		}
		if minItem != nil {
			if tmp.count > minItem.count {
				heap.Push(infos, &tmp)

			} else {
				heap.Push(infos, minItem)

			}
			minItem = nil
		} else {
			heap.Push(infos, &tmp)
		}
	}
	var result = make([]int, k)
	var index int
	for infos.Len() > 0 {
		items := heap.Pop(infos).(*frequceInfo)
		result[k-index-1] = items.val
		index++
	}
	return result
}
func better(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	var counts []int
	for k := range m {
		counts = append(counts, m[k])
	}
	sort.Ints(counts)
	min := counts[len(counts)-k]
	var result []int
	for num, count := range m {
		if count >= min {
			result = append(result, num)
		}
	}
	return result
}
func main() {
	var tests = []int{1, 1, 1, 2, 2, 3, 4, 5}
	fmt.Println(topKFrequent(tests, 2))
	fmt.Println(better(tests, 2))
}
