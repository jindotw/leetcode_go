package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res, lft, rgt := nums[0], 0, len(nums)-1
	for lft <= rgt {
		if nums[lft] <= nums[rgt] {
			res = min(res, nums[lft])
			break
		}

		mid := lft + ((rgt - lft) >> 1)
		res = min(res, nums[mid])
		if nums[mid] >= nums[lft] {
			lft = mid + 1
		} else {
			rgt = mid - 1
		}
	}
	return res
}

func main() {
	// nums := []int{3, 4, 5, 1, 2}
	//fmt.Println(findMin([]int{1, 2, 3, 4, 5}))
	//fmt.Println(findMin([]int{2, 3, 4, 5, 1}))
	//fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
	//fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
}
