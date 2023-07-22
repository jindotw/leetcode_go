package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func removeElement2(nums []int, target int) int {
	lft, rgt := 0, len(nums)-1

	for lft <= rgt {
		for lft <= rgt && nums[lft] != target {
			lft++
		}
		for lft <= rgt && nums[rgt] == target {
			rgt--
		}
		if lft < rgt {
			nums[lft] = nums[rgt]
			lft++
			rgt--
		}
	}

	return lft
}

func main() {
	nums := []int{1}
	val := 1
	remaining := removeElement2(nums, val)
	fmt.Printf("%d remains\n", remaining)
}
