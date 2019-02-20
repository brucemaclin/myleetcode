package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var result []string
	start := nums[0]
	end := start
	for i := 1; i < len(nums); i++ {

		if nums[i] == end+1 {
			end++
		} else {
			if end != start {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
				start = nums[i]
				end = start
			} else {
				result = append(result, strconv.Itoa(start))
				start = nums[i]
				end = start
			}
		}
	}
	if start == end {
		result = append(result, strconv.Itoa(start))
	}
	if end != start {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return result
}
