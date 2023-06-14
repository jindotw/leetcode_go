package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	var queue []int
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
		for true {
			size := len(queue)
			if size == 0 {
				break
			}
			size--
			cmp := nums[queue[size]]
			if val >= cmp {
				break
			}
			lftBound = min(lftBound, queue[size])
			queue = queue[:size]
		}
		queue = append(queue, pos)
	}

	queue = queue[:0]
	for pos := len(nums) - 1; pos >= 0; pos-- {
		for true {
			size := len(queue)
			if size == 0 {
				break
			}
			size--
			cmp := nums[queue[size]]
			val := nums[pos]
			if val <= cmp {
				break
			}
			rgtBound = max(rgtBound, queue[size])
			queue = queue[:size]
		}
		queue = append(queue, pos)
	}
	elements := rgtBound - lftBound
	if elements > 0 {
		return elements + 1
	}
	return 0
}

func main() {
	elem := findUnsortedSubarray([]int{1, 2, 3, 4})
	fmt.Println(elem)
}
