package main

import "fmt"

func longestSubarray(nums []int) int {
	lft, maxStreaks, zeros := 0, 0, 0

	for rgt := 0; rgt < len(nums); rgt++ {
		if nums[rgt] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[lft] == 0 {
				zeros--
			}
			lft++
		}
		currStreaks := rgt - lft + 1
		if currStreaks > maxStreaks {
			maxStreaks = currStreaks
		}
	}
	return maxStreaks - 1
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
