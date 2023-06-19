package main

import "fmt"

func dumpGrid(grid [][]int) {
	x := len(grid)
	y := len(grid[0])
	for m := 0; m < x; m++ {
		for n := 0; n < y; n++ {
			fmt.Printf("%d ", grid[m][n])
		}
		fmt.Println()
	}
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	x := len(obstacleGrid)
	y := len(obstacleGrid[0])
	grid := make([][]int, x)
	for m := 0; m < x; m++ {
		grid[m] = make([]int, y)
	}

	for m := 0; m < x && obstacleGrid[m][0] == 0; m++ {
		grid[m][0] = 1
	}
	for n := 0; n < y && obstacleGrid[0][n] == 0; n++ {
		grid[0][n] = 1
	}

	for m := 1; m < x; m++ {
		for n := 1; n < y; n++ {
			if obstacleGrid[m][n] == 1 {
				continue
			}
			grid[m][n] = grid[m-1][n] + grid[m][n-1]
		}
	}

	dumpGrid(grid)
	return grid[x-1][y-1]
}

func main() {
	grid := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	unique := uniquePathsWithObstacles(grid)
	fmt.Printf("There are %d unique paths\n", unique)
}
