package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	q := MonotonicQueue{elem: []int{}}
	for i := 0; i < k; i++ {
		q.Enqueue(nums[i])
	}

	val, _ := q.Front()
	result = append(result, val)
	for i := k; i < len(nums); i++ {
		q.Dequeue(nums[i-k])
		q.Enqueue(nums[i])
		val, _ = q.Front()
		result = append(result, val)
	}

	return result
}

func maxSlidingWindow2(nums []int, k int) []int {
	result := make([]int, 0)
	queue := make([]int, 0)

	lft := 0
	for rgt, num := range nums {
		last := len(queue) - 1
		for last >= 0 && num > queue[last] {
			queue = queue[:last]
			last--
		}
		queue = append(queue, num)
		if rgt+1 >= k {
			result = append(result, queue[0])
			last := len(queue) - 1
			if last >= 0 && queue[0] == nums[lft] {
				queue = queue[1:]
				last--
			}
			lft++
		}
	}

	return result
}

func main() {
	windows := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	k := 4
	fmt.Println(maxSlidingWindow(windows, k))
	fmt.Println(maxSlidingWindow2(windows, k))
}
