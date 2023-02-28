package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	grid = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}

// https://leetcode.cn/problems/max-product-of-island/
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				a := product(grid, i, j)
				if a > max {
					max = a
				}
			}
		}
	}
	return max
}

func product(grid [][]int, i, j int) int {
	if !isArea(grid, i, j) {
		return 0
	}

	if grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 2
	return 1 + product(grid, i-1, j) + product(grid, i+1, j) + product(grid, i, j-1) + product(grid, i, j+1)
}

func isArea(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
