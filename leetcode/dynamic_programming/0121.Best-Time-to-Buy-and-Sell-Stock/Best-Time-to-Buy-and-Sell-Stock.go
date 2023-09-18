package main

import "fmt"

func maxProfit(prices []int) int {
	cost := -prices[0]
	profit := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, price := range prices[1:] {
		profit = max(profit, price+cost)
		cost = max(cost, -price)
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
