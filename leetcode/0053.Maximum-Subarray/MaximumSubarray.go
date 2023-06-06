package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]

	for _, val := range nums[1:] {
		if currSum > 0 {
			currSum += val
		} else {
			currSum = val
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := maxSubArray(nums)
	fmt.Println("maxSum=", maxSum)
}
