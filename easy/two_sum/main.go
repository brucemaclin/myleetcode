package  main

import (
	"sort"
	"fmt"
)
type info struct {
	val int
	oriPos int
}
type infoSlice []info
func (info infoSlice) Len() int {
	return len(info)
}
func (info infoSlice) Less(i,j int) bool{
	return info[i].val < info[j].val
}
func (info infoSlice) Swap(i,j int) {
	info[i],info[j] = info[j],info[i]
}
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var infos []info
	for i, v := range nums {
		var tmp info
		tmp.val = v
		tmp.oriPos = i
		infos = append(infos, tmp)
	}
	slice := infoSlice(infos)
	sort.Sort(slice)

	lo := 0
	hi := lo+1
	for lo < hi  && hi < len(nums){
		if slice[lo].val + slice[hi].val == target {
			return []int{slice[lo].oriPos, slice[hi].oriPos}
		}
		if slice[lo].val + slice[hi].val < target {
			if hi+1 == slice.Len() {
				lo++
				hi = lo+1
				continue
			}
			hi++
			continue
		}
		if slice[lo].val + slice[hi].val > target {
			lo++
			hi = lo+1
		}
	}
	return nil

}
func better(nums []int, target int) []int {
	var m = make(map[int]int)
	var result = make([]int,2)
	for i,v := range nums {
		if _, ok := m[target-v];ok {
			result[1] = i
			result[0] = m[target-v]
			break
		}
		m[v] = i
	}
	return result
}
func main() {
	ints := []int{2,5,5,11}
	fmt.Println(twoSum(ints, 10))
	fmt.Println(better(ints, 10))
}