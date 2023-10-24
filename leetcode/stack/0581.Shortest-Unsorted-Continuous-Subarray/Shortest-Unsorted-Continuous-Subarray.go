package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	var st []int
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	lftBound, rgtBound := len(nums)-1, 0
	for pos, val := range nums {
		for last := len(st) - 1; last >= 0; last = len(st) - 1 {
			cmp := nums[st[last]]
			if val >= cmp {
				break
			}
			lftBound = min(lftBound, st[last])
			fmt.Printf("val=%d, comparing with %d, lftBound=%d\n", val, nums[st[last]], lftBound)
			st = st[:last]
		}
		st = append(st, pos)
	}

	st = st[:0]
	for pos := len(nums) - 1; pos >= 0; pos-- {
		val := nums[pos]
		for last := len(st) - 1; last >= 0; last = len(st) - 1 {
			cmp := nums[st[last]]
			if val <= cmp {
				break
			}
			rgtBound = max(rgtBound, st[last])
			st = st[:last]
		}
		st = append(st, pos)
	}
	elements := rgtBound - lftBound
	if elements > 0 {
		return elements + 1
	}
	return 0
}

func findUnsortedSubarray2(nums []int) int {
	si := -1
	s := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < s {
			si = i
		} else {
			s = nums[i]
		}
	}
	ei := 0
	e := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > e {
			ei = i
		} else {
			e = nums[i]
		}
	}

	return si - ei + 1

}

func main() {
	nums := []int{1, 4, 6, 8, 3, 2, 15}
	elem := findUnsortedSubarray(nums)
	fmt.Println(elem)
	fmt.Println(findUnsortedSubarray2(nums))
}
