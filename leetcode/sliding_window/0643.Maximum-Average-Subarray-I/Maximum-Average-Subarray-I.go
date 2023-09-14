package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt
	currSum, lft, rgt := 0, 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for rgt < len(nums) {
		currSum += nums[rgt]

		if rgt-lft+1 == k {
			maxSum = max(maxSum, currSum)
			currSum -= nums[lft]
			lft++
		}
		rgt++
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
