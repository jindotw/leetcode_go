package main

func fibonacciSequence(n int) int {
	if n <= 1 {
		return n
	}
	memo := [2]int{0, 1}

	for i := 2; i <= n; i++ {
		memo[0], memo[1] = memo[1], memo[0]+memo[1]
	}

	return memo[1]
}
