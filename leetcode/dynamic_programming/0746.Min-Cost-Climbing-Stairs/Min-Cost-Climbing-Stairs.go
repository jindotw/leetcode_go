package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	var dp = make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]
}

func minCostClimbingStairs2(cost []int) int {
	cost = append(cost, 0)
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func main() {
	// arr := []int{0, 2, 2, 1}
	arr := []int{10, 15, 20}
	fmt.Printf("The cost is %d\n", minCostClimbingStairs(arr))
	fmt.Printf("The cost is %d\n", minCostClimbingStairs2(arr))
}
