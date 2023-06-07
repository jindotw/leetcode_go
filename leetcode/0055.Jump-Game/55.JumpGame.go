package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	covered := 0
	for i := 0; i <= covered; i++ {
		steps := i + nums[i]
		if steps > covered {
			covered = steps
		}
		if covered >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	jumpable := canJump([]int{1, 0, 1})
	fmt.Println("Can Jump:", jumpable)
}
