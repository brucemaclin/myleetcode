package main

import (
	"github.com/derekparker/trie"
)

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}
	t := trie.New()

	for _, v := range words {
		t.Add(v, v)
		t.Add(reverse(v), v)
	}
	var result []string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, "", i, j, t, &result)
		}
	}
	return result
}

func dfs(board [][]byte, path string, i, j int, t *trie.Trie, result *[]string) {
	c := board[i][j]
	if c == '#' {
		return
	}
	path += string(c)
	if t.HasKeysWithPrefix(path) {
		if node, exist := t.Find(path); exist {
			*result = append(*result, node.Meta().(string))
		}
		board[i][j] = '#'
		if i > 0 {
			dfs(board, path, i-1, j, t, result)
		}
		if j > 0 {
			dfs(board, path, i, j-1, t, result)
		}
		if i < len(board)-1 {
			dfs(board, path, i+1, j, t, result)
		}
		if j < len(board[0])-1 {
			dfs(board, path, i, j+1, t, result)
		}
		board[i][j] = c
	}

}

func reverse(str string) string {
	slice := []rune(str)
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return string(slice)
}

func main() {

}
