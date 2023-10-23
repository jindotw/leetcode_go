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

// qs2 finds the kth largest element in an unsorted array.
func qs2(nums []int, k int) int {
	// Convert kth largest to kth smallest index for easier comparison during partition.
	k = len(nums) - k

	// quickSelect is a recursive function that partitions the array
	// and selects the desired kth element.
	var quickSelect func(int, int) int
	quickSelect = func(lft, rgt int) int {
		// Start with the rightmost element as the pivot.
		ptr, pivot := lft, nums[rgt]

		// Reorder the array such that elements less than or equal to the pivot
		// are on the left side and elements greater than the pivot are on the right side.
		for i := lft; i < rgt; i++ {
			if nums[i] <= pivot {
				nums[i], nums[ptr] = nums[ptr], nums[i]
				ptr++
			}
		}

		// Place the pivot in its sorted position.
		nums[ptr], nums[rgt] = nums[rgt], nums[ptr]

		if k > ptr {
			// If kth largest is on the right side of the pivot, search to the right.
			return quickSelect(ptr+1, rgt)
		} else if k < ptr {
			// If kth largest is on the left side of the pivot, search to the left.
			return quickSelect(lft, ptr-1)
		}

		// If kth smallest is the pivot itself, return it.
		return nums[ptr]
	}

	// Begin the quickSelect process starting with the whole array.
	defer func() {
		fmt.Println(nums)
	}()
	return quickSelect(0, len(nums)-1)
}

func findKthLargest(nums []int, k int) int {
	// return qs(nums, len(nums)-k)
	fmt.Println(qs(nums, len(nums)-k))
	return qs2(nums, k)
}

func main() {
	// 1, 2, 3, 4, 5, 6
	nums := []int{3, 2, 1, 5, 6, 4, 100}
	k := 1
	fmt.Println(findKthLargest(nums, k))
}
