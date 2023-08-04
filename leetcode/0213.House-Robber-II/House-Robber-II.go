package main

import "fmt"

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	loot := func(houses []int) int {
		loot0 := houses[0]
		loot1 := max(houses[0], houses[1])
		for i := 2; i < len(houses); i++ {
			loot0, loot1 = loot1, max(loot1, houses[i]+loot0)
		}
		return loot1
	}
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		return max(loot(nums[0:len(nums)-1]), loot(nums[1:]))
	}
}

func main() {
	nums := []int{1}
	fmt.Println(rob(nums))
}
