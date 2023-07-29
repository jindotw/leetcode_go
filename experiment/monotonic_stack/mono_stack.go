package main

import "fmt"

func increasingStack(nums []int) {
	st := make([]int, 0, len(nums))
	for _, num := range nums {
		for true {
			last := len(st) - 1
			if last < 0 || st[last] < num {
				break
			}
			st = st[:last]
		}
		st = append(st, num)
	}

	fmt.Println(st)
}

func decreasingStack(nums []int) {
	st := make([]int, 0, len(nums))
	for _, num := range nums {
		for true {
			last := len(st) - 1
			if last < 0 || st[last] > num {
				break
			}
			st = st[:last]
		}
		st = append(st, num)
	}

	fmt.Println(st)
}

func nextGreaterElement(nums []int) {
	size := len(nums)
	result := make([]int, len(nums))
	for i := 0; i < size; i++ {
		result[i] = -1
		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] {
				result[i] = nums[j]
				break
			}
		}
	}

	fmt.Println(result)
}

func nextGreaterElement2(nums []int) {
	size := len(nums)
	result := make([]int, size)
	st := make([]int, 0, size)

	for idx, num := range nums {
		result[idx] = -1
		last := len(st) - 1
		for last >= 0 && nums[st[last]] < num {
			result[st[last]] = num
			st = st[:last]
			last--
		}
		st = append(st, idx)
	}

	fmt.Println(result)
}

func nextSmallerElement(nums []int) {
	size := len(nums)
	result := make([]int, size)
	st := make([]int, 0, size)
	for idx, num := range nums {
		result[idx] = -1
		last := len(st) - 1
		if last >= 0 && num < nums[st[last]] {
			result[st[last]] = num
			st = st[:last]
			last--
		}
		st = append(st, idx)
	}

	fmt.Println(result)
}

func main() {
	nums := []int{4, 5, 2, 25, 7, 18}
	// increasingStack(nums)
	// decreasingStack(nums)
	nextSmallerElement(nums)
}
