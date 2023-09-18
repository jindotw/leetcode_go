package main

import "fmt"

func change(amount int, coins []int) int {
	memo := make([]int, amount+1)
	memo[0] = 1

	for i, coin := range coins {
		fmt.Printf("========== coins[%d] = %d ==========\n", i, coin)
		for j := coin; j <= amount; j++ {
			memo[j] += memo[j-coin]
			fmt.Printf("j=%d, memo=%v\n", j, memo)
		}
		fmt.Println()
	}

	return memo[amount]
}

func main() {
	coins := []int{2, 3}
	amount := 2
	fmt.Println(change(amount, coins))
}
