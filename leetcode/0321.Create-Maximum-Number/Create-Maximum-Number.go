package main

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	max := make([]int, 0)
	len1, len2 := len(nums1), len(nums2)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	isLeftBigger := func(left []int, right []int) bool {
		for len(left) > 0 && len(right) > 0 && left[0] == right[0] {
			left = left[1:]
			right = right[1:]
		}
		if len(left) == 0 {
			return false
		}
		if (len(right)) == 0 {
			return true
		}

		return left[0] > right[0]
	}
	findMaxNumber := func(arr []int, retain int) []int {
		op := len(arr) - retain
		st := make([]int, 0)
		for _, val := range arr {
			last := len(st) - 1
			for op > 0 && last >= 0 && val > st[last] {
				st = st[:last]
				last--
				op--
			}
			st = append(st, val)
		}
		if len(st) > retain {
			return st[0:retain]
		}
		return st
	}
	merge := func(arr1 []int, arr2 []int) []int {
		ret := make([]int, 0)
		i, j := 0, 0
		len1 := len(arr1)
		len2 := len(arr2)
		for len(ret) <= k && i < len1 && j < len2 {
			if isLeftBigger(arr1[i:], arr2[j:]) {
				ret = append(ret, arr1[i])
				i++
			} else {
				ret = append(ret, arr2[j])
				j++
			}
		}
		diff := k - len(ret)
		if diff > 0 {
			if i < len1 {
				last := min(i+diff, len1)
				ret = append(ret, arr1[i:last]...)
			} else {
				last := min(j+diff, len2)
				ret = append(ret, arr2[j:last]...)
			}
		}
		return ret
	}

	for i := 0; i <= k; i++ {
		if i <= len1 && (k-i) <= len2 {
			arr1 := findMaxNumber(nums1, i)
			arr2 := findMaxNumber(nums2, k-i)
			maxNum := merge(arr1, arr2)
			if isLeftBigger(maxNum, max) {
				max = maxNum
			}
		}
	}
	return max
}

func main() {
	nums1 := []int{3, 4, 6, 5}
	nums2 := []int{9, 1, 2, 5, 8, 3}
	//nums1 := []int{2, 8}
	//nums2 := []int{6, 0, 4}
	//nums1 := []int{8, 5, 9, 5, 1, 6, 9}
	//nums2 := []int{2, 6, 4, 3, 8, 4, 1, 0, 7, 2, 9, 2, 8}
	k := 3
	fmt.Println(maxNumber(nums1, nums2, k))
}
