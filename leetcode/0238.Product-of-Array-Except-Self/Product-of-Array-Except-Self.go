package main

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)
	lftPrd, rgtPrd := make([]int, size), make([]int, size)
	curr := 1

	for idx, num := range nums {
		lftPrd[idx] = curr
		curr *= num
	}
	curr = 1
	for idx := size - 1; idx >= 0; idx-- {
		rgtPrd[idx] = curr
		curr *= nums[idx]
	}

	result := make([]int, size)
	for idx := 0; idx < size; idx++ {
		result[idx] = lftPrd[idx] * rgtPrd[idx]
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	curr := 1

	for idx, val := range nums {
		result[idx] = curr
		curr *= val
	}
	curr = 1
	for idx := size - 1; idx >= 0; idx-- {
		result[idx] *= curr
		curr *= nums[idx]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelf2(nums))
}
