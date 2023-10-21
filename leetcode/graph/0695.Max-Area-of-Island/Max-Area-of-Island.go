package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(m, n int) int {
		if m < 0 || n < 0 || m >= rows || n >= cols || grid[m][n] != 1 {
			return 0
		}
		grid[m][n] = 0
		area := 1
		area += dfs(m+1, n)
		area += dfs(m-1, n)
		area += dfs(m, n+1)
		area += dfs(m, n-1)
		return area
	}
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}

	return maxArea
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
