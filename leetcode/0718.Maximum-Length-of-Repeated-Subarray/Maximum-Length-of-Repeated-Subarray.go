package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	memo := make([][]int, m+1)
	for i, _ := range memo {
		memo[i] = make([]int, n+1)
	}
	max := 0
	for i, num1 := range nums1 {
		for j, num2 := range nums2 {
			if num1 == num2 {
				memo[i+1][j+1] = memo[i][j] + 1
				if max < memo[i+1][j+1] {
					max = memo[i+1][j+1]
				}
			}
		}
	}

	return max
}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}
