package main

import (
	"fmt"
	"strings"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var str string
	for _, v := range wordList {
		str += " " + v
	}
	if !strings.Contains(str, endWord) {
		return 0
	}
	list := []string{beginWord}
	m := make(map[string]bool)
	m[beginWord] = true
	var level int
	for {
		length := len(list)

		if length == 0 {
			break
		}
		level++
		for i := 0; i < length; i++ {
			cur := list[i]
			if cur == endWord {
				return level
			}
			for _, v := range wordList {
				if m[v] {
					continue
				}
				if diffCount(v, cur) {
					list = append(list, v)

					m[v] = true
				}

			}

		}
		list = list[length:]
	}
	return 0
}
func diffCount(s, t string) bool {
	sSlice := []rune(s)
	tSlice := []rune(t)
	var count int
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] != tSlice[i] {
			count++
		}
	}
	return count == 1
}
func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cot"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "cog"}))
}
