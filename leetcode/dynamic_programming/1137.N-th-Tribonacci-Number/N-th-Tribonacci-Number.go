package main

import "fmt"

func tribonacci(n int) int {
	dp := []int{0, 1, 1}
	if n <= 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp[2], dp[1], dp[0] = dp[2]+dp[1]+dp[0], dp[2], dp[1]
	}

	return dp[2]
}

func main() {
	tn := tribonacci(25)
	fmt.Println("tn=", tn)
}
