package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maximumSwap(num int) int {
	str := strconv.Itoa(num)
	orig := []rune(str)
	tmp := []rune(str)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	max := tmp[len(tmp)-1]
	i := 0
	j := 1
	getBigger := false
	bigger := rune(0)
	biggerIndex := 0
	for {
		if j >= len(orig) {
			if i == len(orig)-1 || i == len(orig)-2 {
				break
			}
			i++
			j = i + 1

		}

		currentVal := orig[i]
		if currentVal == max {
			i++
			j++
			continue
		}
		if orig[j] > currentVal {
			getBigger = true
			for k := j; k < len(orig); k++ {
				if orig[k] >= bigger {
					bigger = orig[k]
					biggerIndex = k
				}
			}
		}
		if getBigger {
			orig[i], orig[biggerIndex] = orig[biggerIndex], orig[i]
			i++
			j = i + 1
			getBigger = false
			biggerIndex = 0
			bigger = 0
		} else {
			j++
		}

	}
	r, _ := strconv.Atoi(string(orig))
	return r
}
func main() {
	max := maximumSwap(2736)
	fmt.Println(max)
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(115))
	fmt.Println(maximumSwap(1993))
}
