package main

import (
	"fmt"
)

func partition(nums []int) int {
	lft, rgt := 0, len(nums)-1

	for lft < rgt {
		if nums[lft] >= nums[lft+1] {
			nums[lft], nums[lft+1] = nums[lft+1], nums[lft]
			lft++
		} else if nums[rgt] > nums[lft] {
			rgt--
		} else {
			nums[lft+1], nums[rgt] = nums[rgt], nums[lft+1]
		}
	}

	return lft
}

func qs(nums []int, k int) int {
	pivot := partition(nums)
	if pivot == k {
		return nums[pivot]
	} else if pivot > k {
		return qs(nums[:pivot], k)
	} else {
		return qs(nums[pivot+1:], k-(pivot+1))
	}
}

func findKthLargest(nums []int, k int) int {
	return qs(nums, len(nums)-k)
}

func main() {
	// 1, 2, 3, 4, 5, 6
	nums := []int{100}
	k := 1
	fmt.Println(findKthLargest(nums, k))
}
