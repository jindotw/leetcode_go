package main

import "fmt"

func pivotIndex(nums []int) int {
	lftSum, rgtSum := 0, 0
	for _, num := range nums {
		rgtSum += num
	}

	for i, num := range nums {
		rgtSum -= num
		if lftSum == rgtSum {
			return i
		}
		lftSum += num
	}

	return -1
}

func main() {
	nums := []int{-1, 1, -1}
	fmt.Println(pivotIndex(nums))
}
