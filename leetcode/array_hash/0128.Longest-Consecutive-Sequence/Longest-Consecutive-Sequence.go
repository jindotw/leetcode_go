package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	cache := make(map[int]struct{})
	for _, val := range nums {
		cache[val] = struct{}{}
	}

	longest := 0
	for v := range cache {
		if _, hasPrev := cache[v-1]; hasPrev {
			continue
		}

		length := 1
		for {
			if _, ok := cache[v+length]; !ok {
				break
			}
			delete(cache, v+length)
			length++
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
