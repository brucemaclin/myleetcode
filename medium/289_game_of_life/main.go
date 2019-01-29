package main

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := liveCount(board, m, n, i, j)
			if board[i][j] == 1 && (live >= 2) && live <= 3 {
				board[i][j] = 3
			}
			if board[i][j] == 0 && live == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}

	}

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func liveCount(board [][]int, m, n, i, j int) int {
	live := 0
	for k := max(i-1, 0); k <= min(i+1, m-1); k++ {
		for l := max(j-1, 0); l <= min(j+1, n-1); l++ {
			live += board[k][l] & 1
		}
	}
	live -= board[i][j] & 1
	return live
}
