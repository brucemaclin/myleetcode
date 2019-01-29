package main

import "fmt"

func solve(board [][]byte) {

	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0, m, n)
		if n > 1 {
			dfs(board, i, n-1, m, n)
		}
	}
	for j := 0; j < n; j++ {
		dfs(board, 0, j, m, n)
		if m > 1 {
			dfs(board, m-1, j, m, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(board [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'T'
	for _, v := range directions {
		dfs(board, i+v[0], j+v[1], m, n)
	}
}
func main() {
	var data = [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(data)
	d := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	fmt.Println(data)
	solve(d)
	fmt.Println(d)

}
