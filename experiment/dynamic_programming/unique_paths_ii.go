package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	for row := 0; row < m && obstacleGrid[row][0] == 0; row++ {
		memo[row][0] = 1
	}
	for col := 0; col < n && obstacleGrid[0][col] == 0; col++ {
		memo[0][col] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				memo[i][j] = memo[i-1][j] + memo[i][j-1]
			}
		}
	}
	return memo[m-1][n-1]
}
