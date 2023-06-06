package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	numElements := len(nums)
	if numElements <= 1 {
		return numElements
	}
	currDiff := 0
	preDiff := 0
	numElements--
	result := 1
	for i := 0; i < numElements; i++ {
		currDiff = nums[i+1] - nums[i]
		if (preDiff <= 0 && currDiff > 0) || (preDiff >= 0 && currDiff < 0) {
			preDiff = currDiff
			result++
		}
	}
	return result
}

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	max := wiggleMaxLength(nums)
	fmt.Println("The longest wiggle sequence is", max)
}
