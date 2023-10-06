package main

import "fmt"

func search(nums []int, target int) int {
	lft, rgt := 0, len(nums)-1
	for rgt >= lft {
		middle := lft + ((rgt - lft) >> 1)
		num := nums[middle]
		if num == target {
			return middle
		}
		if num > target {
			rgt = middle - 1
		} else {
			lft = middle + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := -2
	pos := search(nums, target)
	fmt.Printf("%d is at position %d\n", target, pos)
}
