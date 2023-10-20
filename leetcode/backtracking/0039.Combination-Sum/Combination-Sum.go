package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	st := make([]int, 0)

	var dfs func(int, int)
	dfs = func(n, sum int) {
		if sum == target {
			tmp := make([]int, len(st))
			copy(tmp, st)
			results = append(results, tmp)
			return
		} else if sum > target {
			return
		}

		for i := n; i < len(candidates); i++ {
			st = append(st, candidates[i])
			dfs(i, sum+candidates[i])
			st = st[:len(st)-1]
		}
	}
	dfs(0, 0)

	return results
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	results := combinationSum(candidates, target)
	for _, res := range results {
		fmt.Println(res)
	}
}
