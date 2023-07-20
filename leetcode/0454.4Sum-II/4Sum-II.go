package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumCache := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			sumCache[a+b]++
		}
	}
	tuples := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			tuples += sumCache[0-(c+d)]
		}
	}
	return tuples
}

func main() {
	n1 := []int{1, 2}
	n2 := []int{-2, -1}
	n3 := []int{-1, 2}
	n4 := []int{0, 2}

	count := fourSumCount(n1, n2, n3, n4)
	fmt.Printf("%d occurrences\n", count)
}
