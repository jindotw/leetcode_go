package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, len(nums1)-1; k >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	/*
		nums1 := []int{2,4,8,3,5,0}
		nums1 := []int{2,4,8,3,0,0}

	*/

	m, n := 3, 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
