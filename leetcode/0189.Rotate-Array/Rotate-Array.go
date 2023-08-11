package main

import "fmt"

func bruteForce(nums []int, k int) {
	steps := k % len(nums)
	last := len(nums) - 1
	for steps > 0 {
		tmp := nums[last]
		for i := last - 1; i >= 0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = tmp
		steps--
	}
}

func rotateRight(nums []int, k int) {
	size := len(nums)
	last := size - 1
	k %= size
	reverse := func(lft, rgt int) {
		for lft < rgt {
			nums[lft], nums[rgt] = nums[rgt], nums[lft]
			lft, rgt = lft+1, rgt-1
		}
	}
	reverse(0, last)
	reverse(0, k-1)
	reverse(k, last)
}

func rotateLeft(nums []int, k int) {
	steps := k % len(nums)

	lft := nums[:steps]
	rgt := nums[steps:]

	nums = nums[:0]
	nums = append(nums, append(rgt, lft...)...)
}

func rotate(nums []int, k int) {
	bruteForce(nums, k)
}

func main() {
	// org: 1 2 3 4
	// rev: 4 3 2 1
	// lft: 2 3 4 1 (0 ~ k, k+1 ~ last)
	// rgt: 4 1 2 3 (0 ~ k-1, k ~ last)
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 6
	rotate(nums, k)
	fmt.Println(nums)
}
