package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 1; j <= 9; j++ {
			rowUsed[i][j] = false
			colUsed[i][j] = false
			cubeUsed[i][j] = false
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] != '.' {

				rowUsed[i][board[i][j]-'0'] = true
				colUsed[j][board[i][j]-'0'] = true
				cube := cubeNum(i, j)
				cubeUsed[cube][board[i][j]-'0'] = true
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			backtrack(board, i, j)
		}
	}

}

func backtrack(board [][]byte, row, col int) bool {
	for row < 9 && board[row][col] != '.' {
		if col == 8 {
			row = row + 1
			col = 0
		} else {
			col += 1
		}
	}

	if row == 9 {
		return true
	}
	for i := 1; i <= 9; i++ {
		if rowUsed[row][i] || colUsed[col][i] || cubeUsed[cubeNum(row, col)][i] {
			continue
		}
		rowUsed[row][i] = true
		colUsed[col][i] = true
		cubeUsed[cubeNum(row, col)][i] = true
		board[row][col] = '0' + byte(i)

		if backtrack(board, row, col) {
			return true
		}
		board[row][col] = '.'
		rowUsed[row][i] = false
		colUsed[col][i] = false
		cubeUsed[cubeNum(row, col)][i] = false
	}
	return false
}
func cubeNum(i, j int) int {
	r := i / 3
	c := j / 3
	return r*3 + c
}

var rowUsed [9][10]bool
var colUsed [9][10]bool
var cubeUsed [9][10]bool

func printSudoKu(board [][]byte) {
	for i := 0; i < 9; i++ {
		fmt.Println(string(board[i]))
	}
}
func main() {
	sudo := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(sudo)
	printSudoKu(sudo)
	sudo = [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	solveSudoku(sudo)
	printSudoKu(sudo)
}

//[[],[],[],[],[],[],[]]
