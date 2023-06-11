package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	loops := len(nums) - k + 1
	currMax := math.MinInt
	var result []int
	for i := 0; i < loops; i++ {
		arr := nums[i : i+k]
		for j := 0; j < len(arr); j++ {
			if arr[j] > currMax {
				currMax = arr[j]
			}
		}
		result = append(result, currMax)
		if arr[len(arr)-1] != currMax {
			currMax = math.MinInt
		}

	}
	return result
}

func main() {
	windows := []int{1, -1}
	k := 1
	arr := maxSlidingWindow(windows, k)
	fmt.Println(arr)
}
