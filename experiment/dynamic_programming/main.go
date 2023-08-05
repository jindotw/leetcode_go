package main

import (
	"fmt"
	"strings"
)

func callFibonacciSequence() {
	fmt.Println(fibonacciSequence(19))
}

func callKnapsack01() {
	values := []int{15, 20, 30}
	weight := []int{1, 3, 4}
	bagSize := 4
	fmt.Println(knapsack01OneDim(values, weight, bagSize))
}

func callUniquePaths() {
	fmt.Println(uniquePaths(8, 6))
}

func main() {
	callUniquePaths()
}

func dump(memo [][]int) {
	if memo == nil || len(memo) == 0 {
		return
	}
	maxNum := memo[len(memo)-1][len(memo[0])-1]
	maxDigits := 0
	for maxNum > 0 {
		maxNum /= 10
		maxDigits++
	}
	formatStr := fmt.Sprintf("%%%dd", maxDigits)
	for _, r := range memo {
		arr := make([]string, 0)
		for _, c := range r {
			arr = append(arr, fmt.Sprintf(formatStr, c))
		}
		fmt.Println(strings.Join(arr, " "))
	}
}
