package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	size := len(nums)
	memo := make([]int, size)
	for i := range memo {
		memo[i] = 1
	}

	fmt.Println(nums)
	fmt.Println("----------------------------")
	ret := 0
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				memo[i] = max(memo[i], memo[j]+1)
				ret = max(memo[i], ret)
			}
		}
		fmt.Printf("i=%d, num=%d, memo=%v\n", i, nums[i], memo)
	}
	fmt.Println(memo)

	return ret
}

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
