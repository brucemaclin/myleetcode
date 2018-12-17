package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || word == "" {
		return true
	}
	m := len(board)
	n := len(board[0])
	visted := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(0, i, j, m, n, visted, word, board) {
				return true
			}
		}
	}
	return false
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func backtracking(curLen int, i, j, m, n int, visited map[string]bool, word string, board [][]byte) bool {
	if curLen == len(word) {
		return true
	}
	key := fmt.Sprintf("%d-%d", i, j)
	if i < 0 || j < 0 || i >= m || j >= n || visited[key] || board[i][j] != []byte(word)[curLen] {
		return false
	}
	visited[key] = true
	for _, v := range directions {
		if backtracking(curLen+1, i+v[0], j+v[1], m, n, visited, word, board) {
			return true
		}
	}
	visited[key] = false
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
