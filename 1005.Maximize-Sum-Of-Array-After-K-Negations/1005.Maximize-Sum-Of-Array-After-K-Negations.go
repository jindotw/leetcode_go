package main

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})
	fmt.Println(nums)

	for i := len(nums) - 1; k > 0 && i >= 0; i-- {
		if nums[i] < 0 {
			nums[i] *= -1
			k--
		}
	}

	if k%2 == 1 {
		nums[0] *= -1
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	nums := []int{-2, 5, 0, 2, -2}
	k := 3
	sum := largestSumAfterKNegations(nums, k)
	fmt.Println("max sum:", sum)
}
