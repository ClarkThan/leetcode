package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	fmt.Println(islandPerimeter(grid))
}

// https://leetcode.cn/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	cnt := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// cnt += (count(grid, i-1, j) + count(grid, i+1, j) + count(grid, i, j-1) + count(grid, i, j+1))
				if i-1 < 0 || grid[i-1][j] == 0 {
					cnt++
				}
				if i+1 >= m || grid[i+1][j] == 0 {
					cnt++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					cnt++
				}
				if j+1 >= n || grid[i][j+1] == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}

func count(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 1
	}

	return 0
}
