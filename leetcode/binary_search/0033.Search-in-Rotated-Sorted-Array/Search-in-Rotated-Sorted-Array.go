package main

import "fmt"

func search(nums []int, target int) int {
	lft, rgt := 0, len(nums)-1

	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[lft] <= nums[mid] {
			if target > nums[mid] || target < nums[lft] {
				lft = mid + 1
			} else {
				rgt = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[rgt] {
				rgt = mid - 1
			} else {
				lft = mid + 1
			}

		}
	}
	return -1
}

func main() {
	target := 6
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, target))
	fmt.Println(search([]int{7, 0, 1, 2, 4, 5, 6}, target))
	fmt.Println(search([]int{6, 7, 0, 1, 2, 4, 5}, target))
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, target))

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, target))
	fmt.Println(search([]int{2, 4, 5, 6, 7, 0, 1}, target))
	fmt.Println(search([]int{1, 2, 4, 5, 6, 7, 0}, target))
}
