package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sum := 0
	for _, weight := range stones {
		sum += weight
	}
	half := sum / 2

	memo := make([]int, half+1)
	for _, stone := range stones {
		for j := half; j >= stone; j-- {
			memo[j] = max(memo[j], memo[j-stone]+stone)
		}
	}

	return sum - memo[half]*2
}

func main() {
	input := []int{2, 7, 4, 1, 8, 1}
	left := lastStoneWeightII(input)
	fmt.Println("The smallest stone left is", left)
}
