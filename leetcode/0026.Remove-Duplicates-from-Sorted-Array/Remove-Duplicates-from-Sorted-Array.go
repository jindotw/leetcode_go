package main

import "fmt"

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	i := 0
	for j := 1; j < size; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)

	return i + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// 0, 1, 1, 1, 0, 2, 2, 3, 3, 4
	//

	remaining := removeDuplicates(nums)
	fmt.Printf("Remaining elements: %d\n", remaining)
}
