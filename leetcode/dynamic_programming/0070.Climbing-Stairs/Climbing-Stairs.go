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

func climbStairs2(n int) int {
	memo := make([]int, n+1)
	memo[n] = 1
	memo[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		memo[i] = memo[i+1] + memo[i+2]
	}
	fmt.Println(memo)
	return memo[0]
}

func main() {
	n := 5
	fmt.Printf("There are %d ways to climb a %d staircases\n", climbStairs(n), n)
	fmt.Printf("There are %d ways to climb a %d staircases\n", climbStairs2(n), n)
}
