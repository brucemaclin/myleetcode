package main

import (
	"fmt"
	"sort"
	"strings"
)

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	type indexSort struct {
		index    int
		target   string
		position int
		src      string
	}
	type sortSlice []indexSort
	var tmpSlice sortSlice
	for i := 0; i < len(indexes); i++ {
		var tmp indexSort
		tmp.position = i
		tmp.index = indexes[i]
		tmp.target = targets[i]
		tmp.src = sources[i]
		tmpSlice = append(tmpSlice, tmp)
	}
	sort.Slice(tmpSlice, func(i, j int) bool {
		return tmpSlice[i].index > tmpSlice[j].index
	})
	for _, src := range tmpSlice {
		index := src.index
		getIndex := strings.Index(string(S[index:]), src.src)
		fmt.Println(index, src.src, src.target, S)
		if getIndex == 0 {
			S = string(S[0:index]) + src.target + string(S[index+len(src.src):])
		}

	}
	return S

}

func main() {
	fmt.Println(findReplaceString("wreorttvosuidhrxvmvo", []int{14, 12, 10, 5, 0, 18}, []string{"rxv", "dh", "ui", "ttv", "wreor", "vo"}, []string{"frs", "c", "ql", "qpir", "gwbeve", "n"}))
}
