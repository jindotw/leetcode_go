package main

import "fmt"

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(nums) == 1 {
		return nums[0]
	}
	loot0 := nums[0]
	loot1 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// 1 3 1 3 100
		// 1 3 3 6 103
		loot0, loot1 = loot1, max(nums[i]+loot0, loot1)
	}

	return loot1
}

func main() {
	houses := []int{1, 3, 1, 3, 100}
	money := rob(houses)
	fmt.Println("money robbed is", money)
}
