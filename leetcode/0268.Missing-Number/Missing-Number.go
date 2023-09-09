package main

import "fmt"

func missingNumber(nums []int) int {
	xor := 0

	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}
	// fmt.Printf("%08b\n", xor)
	for _, num := range nums {
		xor ^= num
	}
	// fmt.Printf("%08b\n", xor)

	return xor
}

func main() {
	nums := []int{0, 1, 2, 3}
	fmt.Println(missingNumber(nums))
}
