package main

import "fmt"

func longestOnes(nums []int, k int) int {
	lft, maxStreaks, zeroes := 0, 0, 0

	for rgt := 0; rgt < len(nums); rgt++ {
		if nums[rgt] == 0 {
			zeroes++
		}

		for zeroes > k {
			if nums[lft] == 0 {
				zeroes--
			}
			lft++
		}

		currStreaks := rgt - lft + 1
		if currStreaks > maxStreaks {
			maxStreaks = currStreaks
		}
	}
	return maxStreaks
}

func main() {
	nums := []int{0, 0, 0, 1}
	k := 4
	fmt.Println(longestOnes(nums, k))
}
