package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if pos, ok := cache[diff]; ok {
			return []int{pos, i}
		}
		cache[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(nums, target)
	fmt.Println(arr)
}
