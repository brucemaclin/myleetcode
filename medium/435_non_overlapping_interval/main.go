package main

import (
	"fmt"
	"sort"
)

//Definition for an interval.
type Interval struct {
	Start int
	End   int
}
type intervalSlice []Interval

func (m intervalSlice) Less(i, j int) bool {
	return m[i].End < m[j].End
}
func (m intervalSlice) Len() int {
	return len(m)
}
func (m intervalSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]

}
func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	slice := intervalSlice(intervals)
	sort.Sort(slice)
	count := 1
	end := intervals[0].End
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= end {
			end = intervals[i].End
			count++
		}
	}
	return len(intervals) - count
}

func main() {
	interval1 := Interval{Start: 1, End: 2}
	interval2 := Interval{Start: 1, End: 3}
	interval3 := Interval{Start: 2, End: 3}
	var intervals []Interval
	intervals = append(intervals, interval1, interval2, interval3)
	fmt.Println(eraseOverlapIntervals(intervals))

}
