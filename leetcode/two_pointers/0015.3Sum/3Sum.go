package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	last := len(nums) - 1
	for idx, a := range nums {
		if a > 0 {
			return ans
		}
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}
		lft := idx + 1
		rgt := last
		for lft < rgt {
			b := nums[lft]
			c := nums[rgt]
			sum := a + b + c
			if sum == 0 {
				ans = append(ans, []int{a, b, c})
				for rgt > lft && nums[rgt] == nums[rgt-1] {
					rgt--
				}
				for lft < rgt && nums[lft] == nums[lft+1] {
					lft++
				}
				rgt--
				lft++
			} else if sum > 0 {
				rgt--
			} else {
				lft++
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-2, 0, 0, 2, 2}
	ans := threeSum(nums)
	fmt.Println(ans)
}
