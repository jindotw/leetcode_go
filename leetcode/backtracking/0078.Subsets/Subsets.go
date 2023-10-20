package main

import "fmt"

func subsets(nums []int) [][]int {
	var results [][]int
	st := make([]int, 0)
	var dfs func(int)
	dfs = func(n int) {
		tmp := make([]int, len(st))
		copy(tmp, st)
		results = append(results, tmp)

		for i := n; i < len(nums); i++ {
			st = append(st, nums[i])
			dfs(i + 1)
			st = st[:len(st)-1]
		}
	}
	dfs(0)

	return results
}

func main() {
	nums := []int{1, 2, 3}
	results := subsets(nums)
	for _, result := range results {
		fmt.Println(result)
	}
}
