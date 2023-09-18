package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
	}

	for i := 0; i < len(coins); i++ {
		fmt.Printf("====== coins[%d] = %d ======\n", i, coins[i])
		for j := coins[i]; j <= amount; j++ {
			memo[j] = min(memo[j], memo[j-coins[i]]+1)
			fmt.Printf("amount=%d, memo=%v\n", j, memo)
		}
		fmt.Println("*****************************************")
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]
}

func main() {
	coins := []int{1, 5, 10}
	amount := 14
	fmt.Println(coinChange(coins, amount))
}
