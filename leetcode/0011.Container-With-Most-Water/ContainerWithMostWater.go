package main

import (
	"fmt"
)

func maxArea(height []int) int {
	lft, rgt := 0, len(height)-1
	max := 0

	for lft < rgt {
		currArea := rgt - lft
		if height[lft] < height[rgt] {
			currArea *= height[lft]
			lft++
		} else {
			currArea *= height[rgt]
			rgt--
		}
		if currArea > max {
			max = currArea
		}
	}

	return max
}

func main() {
	height := []int{1, 3, 2, 5, 25, 24, 5}
	area := maxArea(height)
	fmt.Println("Area =", area)
}
