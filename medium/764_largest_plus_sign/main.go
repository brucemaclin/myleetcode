package main

import (
	"fmt"
)

type numOf1Info struct {
	left  int
	right int
	up    int
	down  int
}

func orderOfLargestPlusSign(N int, mines [][]int) int {
	if N == 1 {
		for _, v := range mines {
			if v[0] == 0 && v[1] == 0 {
				return 0
			}
		}
		return 1
	}

	grid := make([][]int, N)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, N)
		for j := 0; j < N; j++ {
			grid[i][j] = N
		}
	}
	for _, v := range mines {
		grid[v[0]][v[1]] = 0
	}

	var max int

	for i := 0; i < N; i++ {
		var l, r, u, d int
		for j, k := 0, N-1; j < N; {
			if grid[i][j] == 0 {
				l = 0
			} else {
				grid[i][j] = getMin(grid[i][j], l+1)
				l = l + 1
			}
			if grid[i][k] == 0 {
				r = 0
			} else {
				grid[i][k] = getMin(grid[i][k], r+1)
				r = r + 1
			}
			if grid[j][i] == 0 {
				u = 0
			} else {
				grid[j][i] = getMin(grid[j][i], u+1)
				u = u + 1
			}
			if grid[k][i] == 0 {
				d = 0
			} else {
				grid[k][i] = getMin(grid[k][i], d+1)
				d = d + 1
			}
			j++
			k--
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {

			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
	}
	return max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	//fmt.Println(orderOfLargestPlusSign(2, nil))
	//fmt.Println(orderOfLargestPlusSign(1, [][]int{{0, 0}}))
}
