package main

import "fmt"

var dp = make(map[int]int)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	if val, ok := dp[n]; ok {
		return val
	}

	val := climbStairs(n-2) + climbStairs(n-1)
	dp[n] = val

	return val
}

func main() {
	n := 2
	ways := climbStairs(n)
	fmt.Printf("There are %d ways to climb a %d staircases\n", ways, n)
}
