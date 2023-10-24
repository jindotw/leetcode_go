package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	st := make([]int, 0)

	for _, num := range nums2 {
		cache[num] = -1
		last := len(st) - 1
		for last >= 0 && num > st[last] {
			cache[st[last]] = num
			st = st[:last]
			last--
		}
		st = append(st, num)
	}

	result := make([]int, len(nums1))
	for idx, num := range nums1 {
		result[idx] = cache[num]
	}
	return result
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
