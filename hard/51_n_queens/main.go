package main

import "fmt"

var rowUsed []bool
var colUsed []bool
var diagonal45 []bool
var diagonal135 []bool

func solveNQueens(n int) [][]string {
	rowUsed = make([]bool, n+1)
	colUsed = make([]bool, n+1)
	diagonal45 = make([]bool, 2*n)
	diagonal135 = make([]bool, 2*n)
	var path []string
	var result [][]string
	path = make([]string, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i] += "."
		}
	}
	backtrack(0, n, path, &result)
	return result
}

func backtrack(row, n int, path []string, result *[][]string) {
	if row == n {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}

	i := row
	for j := 0; j < n; j++ {
		if rowUsed[i] || colUsed[j] || diagonal45[i+j] || diagonal135[n-1-(i-j)] {
			continue
		}
		rowUsed[i] = true
		colUsed[j] = true
		diagonal45[i+j] = true
		diagonal135[n-1-(i-j)] = true
		path[row] = getTempString(j, true, path[row])
		backtrack(row+1, n, path, result)
		rowUsed[i] = false
		colUsed[j] = false
		diagonal45[i+j] = false
		diagonal135[n-1-(i-j)] = false
		path[row] = getTempString(j, false, path[row])
	}

}

func getTempString(j int, flag bool, str string) string {
	slices := []rune(str)
	if flag {
		slices[j] = 'Q'
	} else {
		slices[j] = '.'
	}
	return string(slices)
}

func main() {
	fmt.Println(solveNQueens(4))
}
