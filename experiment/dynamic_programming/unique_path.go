package main

func uniquePaths(m, n int) int {
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		memo[i][0] = 1
	}
	for j, _ := range memo[0] {
		memo[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i-1][j] + memo[i][j-1]
		}
	}
	dump(memo)
	return memo[m-1][n-1]
}
