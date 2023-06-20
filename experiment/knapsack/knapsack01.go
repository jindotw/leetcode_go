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

func main() {
	capacity := 7
	weight := []int{4, 3, 1}
	value := []int{15, 20, 30}
	solution1D(capacity, weight, value)
}
