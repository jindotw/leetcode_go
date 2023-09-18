package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for prev, v := range prices[1:] {
		dayProfit := v - prices[prev]
		if dayProfit > 0 {
			profit += dayProfit
		}
	}

	return profit
}

func maxProfit2(prices []int) int {
	memo := make([][]int, len(prices))
	for i := range memo {
		memo[i] = make([]int, 2)
	}
	memo[0][0] = -prices[0]
	memo[0][1] = 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		memo[i][0] = max(memo[i-1][0], memo[i-1][1]-prices[i])
		memo[i][1] = max(memo[i-1][1], prices[i]+memo[i-1][0])
		fmt.Println(memo)
	}

	return 0
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println("Profit=", profit)
	maxProfit2(prices)
}
