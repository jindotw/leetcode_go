package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	k := 3
	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, len(prices))
	}
	for i := 1; i < len(memo); i++ {
		maxDiff := memo[i-1][0] - prices[0]
		for j := 1; j < len(memo[i]); j++ {
			memo[i][j] = max(memo[i][j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, memo[i-1][j]-prices[j])
		}
	}
	for _, row := range memo {
		fmt.Println(row)
	}

	return memo[k][len(prices)-1]
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	// prices := []int{1, 2, 4}
	fmt.Println(maxProfit(prices))
}
