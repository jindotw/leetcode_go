package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	lft, sum := 0, 0
	minLen := math.MaxInt
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for rgt := range nums {
		sum += nums[rgt]
		for sum >= target {
			minLen = min(minLen, rgt-lft+1)
			sum -= nums[lft]
			lft++
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func main() {
	target := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(target, nums))
}
