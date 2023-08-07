package main

import "fmt"

func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, v := range stones {
		sum += v
	}
	fmt.Println(stones, sum)
	memo := make([]int, sum+1)
	for _, stone := range stones {
		for j := sum; j >= stone; j-- {
			memo[j] = max(memo[j], memo[j-stone]+stone)
		}
	}
	fmt.Println(memo)
	return 0
}
