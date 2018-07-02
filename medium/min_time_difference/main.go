package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	var ints []int
	for _, v := range timePoints {
		var tmp [2]int
		fmt.Sscanf(v, "%02d:%02d", &tmp[0], &tmp[1])

		ints = append(ints, 60*tmp[0]+tmp[1])
	}
	sort.Ints(ints)
	var min int = 24 * 3600
	for i, v := range ints {

		var diff int
		if i == len(ints)-1 {
			diff = v - ints[0]
		} else {
			diff = v - ints[i+1]
		}

		if diff < 0 {

			diff = -diff

		}
		if diff > 12*60 {
			diff = 24*60 - diff
		}

		if min > diff {
			min = diff
		}
	}
	return min
}
func better(timePoints []string) int {
	a := make(map[int]bool)
	for _, v := range timePoints {
		if _, ok := a[timeToInt(v)]; ok {
			return 0
		}
		a[timeToInt(v)] = true
	}
	var b []int
	for k := range a {
		b = append(b, k)
	}
	sort.Ints(b)

	var min = 1 << 32
	for i := 1; i < len(b); i++ {
		if b[i]-b[i-1] < min {
			min = b[i] - b[i-1]
		}
	}
	if (b[0] - b[len(b)-1] + 24*60) < min {
		min = b[0] - b[len(b)-1] + 24*60
	}
	return min

}
func timeToInt(s string) int {
	a, _ := strconv.Atoi(s[:2])
	b, _ := strconv.Atoi(s[3:])
	return 60*a + b
}
func main() {
	timePoints := []string{"23:59", "00:00"}
	fmt.Println(findMinDifference(timePoints))
	fmt.Println(better(timePoints))
}
