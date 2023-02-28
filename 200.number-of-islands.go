package main

import (
	"fmt"
)

func main() {
	grid := [][]byte{
		{'1', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '1'},
		{'0', '1', '1', '1', '0', '1'},
		{'1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '1', '0', '1'},
	}
	// grid = [][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }

	fmt.Println(numIslands(grid))
}

// https://leetcode.cn/problems/number-of-islands/?favorite=2cktkvj
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	num := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				visit(grid, r, c)
				num++
			}
		}
	}

	return num
}

func visit(grid [][]byte, r, c int) {
	m, n := len(grid), len(grid[0])
	if grid[r][c] == '1' {
		grid[r][c] = '2'
		if r > 0 {
			visit(grid, r-1, c)
		}
		if r < m-1 {
			visit(grid, r+1, c)
		}
		if c > 0 {
			visit(grid, r, c-1)
		}
		if c < n-1 {
			visit(grid, r, c+1)
		}
	}
}

func numIslandsV0(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	num := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if area(grid, r, c) > 0 {
				num++
			}
		}
	}

	return num
}

func area(grid [][]byte, r, c int) int {
	if !inArea(grid, r, c) {
		return 0
	}

	if grid[r][c] == '1' {
		grid[r][c] = '2'
		return 1 + area(grid, r+1, c) + area(grid, r-1, c) + area(grid, r, c+1) + area(grid, r, c-1)
	}

	return 0
}

func inArea(grid [][]byte, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
