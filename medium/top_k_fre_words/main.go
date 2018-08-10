package main

import (
	"fmt"
	"sort"
	"strings"
)

type pair struct {
	word  string
	count int
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	if p[i].count > p[j].count {
		return true
	} else if p[i].count == p[j].count {
		return strings.Compare(p[i].word, p[j].word) < 0
	}
	return false
}
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	var result []string
	var p pairs
	for word, count := range m {
		tmp := pair{}
		tmp.word = word
		tmp.count = count
		p = append(p, tmp)
	}
	sort.Sort(p)
	i := k
	for {
		if i == 0 {
			break
		}
		result = append(result, p[k-i].word)
		i--
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
