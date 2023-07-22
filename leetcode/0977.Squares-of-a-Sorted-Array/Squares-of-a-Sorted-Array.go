package main

import "fmt"

func sortedSquares(nums []int) []int {
	k := len(nums)
	ret := make([]int, k)
	lft, rgt := 0, k-1

	for k > 0 {
		k--
		lftNum := nums[lft] * nums[lft]
		rgtNum := nums[rgt] * nums[rgt]
		if lftNum > rgtNum {
			ret[k] = lftNum
			lft++
		} else {
			ret[k] = rgtNum
			rgt--
		}
	}

	return ret
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
