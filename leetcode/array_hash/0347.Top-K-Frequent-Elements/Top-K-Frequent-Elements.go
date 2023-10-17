package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	counters := make([][]int, len(nums)+1)
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}
	for num, freq := range frequencies {
		counters[freq] = append(counters[freq], num)
	}
	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		for _, num := range counters[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
