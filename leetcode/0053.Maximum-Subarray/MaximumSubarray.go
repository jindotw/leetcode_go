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

func maxSubArray2(nums []int) int {
	maxHere, maxSoFar := nums[0], nums[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, num := range nums[1:] {
		maxHere = max(num, maxHere+num)
		maxSoFar = max(maxHere, maxSoFar)
	}

	return maxSoFar
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{1, 2, 3, -5, 4, 5, 6}
	maxSum := maxSubArray(nums)
	maxSum2 := maxSubArray2(nums)
	fmt.Println("maxSum=", maxSum)
	fmt.Println("maxSum2=", maxSum2)
}
