package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min, mid := math.MaxInt32, math.MaxInt32

	for _, val := range nums {
		if val < min {
			min = val
		} else if val > min && val < mid {
			mid = val
		} else if val > mid {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums))
}
