package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}
type info struct {
	start int
	index int
}
type infoSlice []info

func (m infoSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m infoSlice) Len() int {
	return len(m)
}

func (m infoSlice) Less(i, j int) bool {
	return m[i].start < m[j].start
}
func findRightInterval(intervals []Interval) []int {

	var infos []info
	for index, v := range intervals {
		tmp := info{}
		tmp.start = v.Start
		tmp.index = index
		infos = append(infos, tmp)
	}
	sort.Sort(infoSlice(infos))
	var result []int
	for _, v := range intervals {
		l := 0
		h := len(intervals) - 1
		for l < h {
			m := l + (h-l)/2
			if infos[m].start >= v.End {
				h = m
			} else {
				l = m + 1
			}
		}
		if infos[l].start >= v.End {
			result = append(result, infos[l].index)
		} else {
			result = append(result, -1)
		}

	}
	return result

}

func main() {
	fmt.Println(findRightInterval([]Interval{{3, 4}, {2, 3}, {1, 2}}))
}
