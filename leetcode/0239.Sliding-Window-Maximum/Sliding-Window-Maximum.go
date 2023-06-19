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

func main() {
	windows := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	arr := maxSlidingWindow(windows, k)
	fmt.Println(arr)
}
