package main

import "fmt"

func maxOperations(nums []int, k int) int {
	freq := make(map[int]int)

	op := 0
	for _, val := range nums {
		if count, ok := freq[val]; ok {
			op++
			if count <= 1 {
				delete(freq, val)
			} else {
				freq[val]--
			}
		} else {
			diff := k - val
			freq[diff]++
		}
	}

	return op
}

func main() {
	nums := []int{1, 1, 1, 2, 1, 2, 2, 2, 2}
	k := 3
	fmt.Println(maxOperations(nums, k))
}
