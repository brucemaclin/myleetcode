package main

import (
	"sort"
	"math"
	"strconv"
	"fmt"
)

func findPairs(nums []int, k int) int {
	slice := sort.IntSlice(nums)
	slice.Sort()
	var m = make(map[string]bool)
	for i := range slice {
		for j:=i+1;j<len(slice);j++ {
			diff := int(math.Abs(float64(slice[i]-slice[j])))
			if diff-k == 0 {
				m[strconv.Itoa(slice[i])+"_"+strconv.Itoa(slice[j])] = true
				break
			}
			if diff-k > 0 {
				break
			}
		}
	}
	return len(m)
}

func  main()  {
	fmt.Println(findPairs([]int{1,3,4,5,6},1))
}