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

func main() {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println("Profit=", profit)
}
