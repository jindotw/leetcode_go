package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	avg := sum / 2
	memo := make([]int, avg+1)
	memo[0] = 0
	for _, val := range nums {
		for j := avg; j >= val; j-- {
			memo[j] = max(memo[j], memo[j-val]+val)
		}
	}

	fmt.Println(memo)

	return avg == memo[avg]
}

func main() {
	input := []int{80, 70, 30, 120}
	yesOrNo := canPartition(input)
	fmt.Println(input, "can be partitioned?", yesOrNo)
}
