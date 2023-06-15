package main

import "fmt"

func uniquePaths(m, n int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		grid[i][0] = 1
		for j := 0; j < m; j++ {
			grid[0][j] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[n-1][m-1]
}

func main() {
	y, x := 3, 7
	unique := uniquePaths(y, x)
	fmt.Printf("there are %d unique paths", unique)
}
