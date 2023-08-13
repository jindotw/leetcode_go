package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	ans := make([]int, len(digits)+1)
	ans[0] = 1
	return ans
}

func main() {
	nums := []int{9, 9}
	fmt.Println(plusOne(nums))
}
