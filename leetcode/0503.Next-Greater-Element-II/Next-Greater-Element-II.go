package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	doubled := make([]int, size*2)
	for idx, num := range nums {
		doubled[idx] = num
		doubled[idx+size] = num
	}

	st := make([]int, 0)
	result := make([]int, size*2)
	for idx, num := range doubled {
		result[idx] = -1
		last := len(st) - 1
		for last >= 0 && num > doubled[st[last]] {
			result[st[last]] = num
			st = st[:last]
			last--
		}
		st = append(st, idx)
	}

	return result[:size]
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}
