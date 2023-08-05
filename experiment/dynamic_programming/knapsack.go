package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack01(values, weight []int, bagSize int) int {
	items := len(values)
	memo := make([][]int, items)
	for i, _ := range memo {
		memo[i] = make([]int, bagSize+1)
		memo[i][0] = 0
	}
	for j := weight[0]; j <= bagSize; j++ {
		memo[0][j] = values[0]
	}

	for i := 1; i < items; i++ {
		for j := 1; j <= bagSize; j++ {
			if weight[i] > j {
				memo[i][j] = memo[i-1][j]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i-1][j-weight[i]]+values[i])
			}
		}
	}
	dump(memo)
	return memo[items-1][bagSize]
}

func knapsack01OneDim(values, weight []int, bagSize int) int {
	memo := make([]int, bagSize+1)
	for i, _ := range memo {
		memo[i] = 0
	}
	for i := 0; i < len(weight); i++ {
		for j := bagSize; j >= weight[i]; j-- {
			memo[j] = max(memo[j], memo[j-weight[i]]+values[i])
		}
	}
	return memo[bagSize]
}
