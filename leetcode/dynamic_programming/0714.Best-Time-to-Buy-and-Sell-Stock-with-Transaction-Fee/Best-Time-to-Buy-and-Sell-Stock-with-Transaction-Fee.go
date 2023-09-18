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
	/*
	  memo is a two-dimensional array used for memoization. It keeps track of the maximum profit you can have at each day
	  and for two possible states:
	    - whether you hold a stock (memo[i][0]) or
	    - whether you don't hold a stock (memo[i][1])

	  In simple terms, memo[i][0] represents the maximum profit at the end of day i if you hold a stock on that day,
	  and memo[i][1] represents the maximum profit at the end of day i if you don't hold a stock. The dynamic
	  programming approach keeps track of these two states to find the overall maximum profit considering the
	  transaction fee.
	*/
	memo := make([][]int, len(prices))
	for i := range memo {
		memo[i] = make([]int, 2)
	}
	/*
		memo[0][0] is initialized to -prices[0], representing that on the first day, you can choose to buy a stock (hence
		the negative value, as it represents an expense).
	*/
	memo[0][0] = -prices[0]

	/*
		The loop starts from the second day (i = 1) and iterates through the entire list of stock prices.

		memo[i][0] is updated with the maximum between:
		  - memo[i-1][0]: Representing the maximum profit if you continue holding a stock from the previous day.
		  - memo[i-1][1] - Representing the profit gained from selling the stock on the previous day to buy a stock
			at the price on the current day (prices[i]).
		memo[i][1] is updated with the maximum between:
		  - memo[i-1][1]: Representing the maximum profit if you continue not holding a stock (i.e., you didn't sell on the
			previous day).
		  - memo[i-1][0] + prices[i] -  Representing the profit after selling the stock on the current day and
			accounting for the transaction fee.
	*/
	for i := 1; i < len(prices); i++ {
		memo[i][0] = max(memo[i-1][0], memo[i-1][1]-prices[i])
		memo[i][1] = max(memo[i-1][1], memo[i-1][0]+prices[i]-fee)
	}

	dump(memo)

	return memo[len(prices)-1][1]
}

func maxProfit2(prices []int, fee int) int {
	hold, free := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		hold, free = max(hold, free-prices[i]), max(free, hold+prices[i]-fee)
	}

	return free

}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	// prices := []int{1, 3, 7, 5, 10, 3}
	// fee := 3
	fmt.Println(maxProfit(prices, fee))
	fmt.Println(maxProfit2(prices, fee))
}
