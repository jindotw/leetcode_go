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

func maxProfit(k int, prices []int) int {
	rows := k + 1
	cols := len(prices)
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
	}
	for i := 1; i < rows; i++ {
		maxDiff := memo[i-1][0] - prices[0]
		for j := 1; j < cols; j++ {
			maxDiff = max(maxDiff, memo[i-1][j]-prices[j])
			memo[i][j] = max(memo[i][j-1], maxDiff+prices[j])
		}
	}
	fmt.Println(memo)

	return memo[k][cols-1]
}

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(k, prices))
}
