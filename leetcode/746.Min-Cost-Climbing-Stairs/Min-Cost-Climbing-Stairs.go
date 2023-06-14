package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dp = make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]
}

func main() {
	arr := []int{0, 2, 2, 1}
	cost := minCostClimbingStairs(arr)
	fmt.Println("The cost is", cost)
}
