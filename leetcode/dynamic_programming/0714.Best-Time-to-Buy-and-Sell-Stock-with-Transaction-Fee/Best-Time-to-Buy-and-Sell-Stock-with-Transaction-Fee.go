package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dump(memo [][]int) {
	for _, rows := range memo {
		for _, col := range rows {
			fmt.Printf(" %02d ", col)
		}
		fmt.Println()
	}
}

func maxProfit(prices []int, fee int) int {
	memo := make([][]int, len(prices))
	for i := range memo {
		memo[i] = make([]int, 2)
	}
	memo[0][0] = -prices[0]
	memo[0][1] = 0

	for i := 1; i < len(prices); i++ {
		fmt.Printf("prices[%d] = %d\n", i, prices[i])
		memo[i][0] = max(memo[i-1][0], -prices[i])
		memo[i][1] = max(memo[i-1][1], prices[i]+memo[i-1][0])
		dump(memo)
		fmt.Println()
	}

	return 0
}

func maxProfit2(prices []int) int {
	cost := prices[0]
	profit := 0

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}

	return profit
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
	fmt.Println(maxProfit2(prices))
}
