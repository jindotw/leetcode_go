package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution(capacity int, weight []int, value []int) {
	memo := make([][]int, len(weight))
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}

	for j := capacity; j >= weight[0]; j-- {
		memo[0][j] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 1; j <= capacity; j++ {
			if j < weight[i] {
				memo[i][j] = memo[i-1][j]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(memo)
}

func solution1D(capacity int, weight []int, value []int) {
	memo := make([]int, capacity+1)

	for i, wt := range weight {
		for j := capacity; j >= weight[i]; j-- {
			memo[j] = max(memo[j], memo[j-wt]+value[i])
		}
	}
	fmt.Println(memo)
}

func solutionRecurInner(memo [][]int, weight []int, value []int, n int, capacity int) int {
	if n == 0 || capacity == 0 {
		return 0
	}
	if memo[n][capacity] != 0 {
		fmt.Printf("memo hit on %d %d\n", n, capacity)
		return memo[n][capacity]
	}
	if weight[n] > capacity {
		memo[n][capacity] = solutionRecurInner(memo, weight, value, n-1, capacity)
		return memo[n][capacity]
	}

	result := max(
		solutionRecurInner(memo, weight, value, n-1, capacity-weight[n])+value[n],
		solutionRecurInner(memo, weight, value, n-1, capacity),
	)
	memo[n][capacity] = result

	return result
}

func solutionRecur(weight []int, value []int, n int, capacity int) int {
	memo := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		memo[i] = make([]int, capacity+1)
	}
	for i := weight[1]; i <= capacity; i++ {
		memo[1][i] = value[1]
	}
	val := solutionRecurInner(memo, weight, value, n, capacity)
	fmt.Println(memo)

	return val
}

func main() {
	capacity := 8
	weight := []int{0, 2, 3, 1, 5, 6, 4}
	value := []int{0, 60, 20, 30, 25, 35, 40}
	// solution1D(capacity, weight, value)
	val := solutionRecur(weight, value, len(weight)-1, capacity)
	fmt.Println(val)
}
