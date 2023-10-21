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

func combinationSum2(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	st := make([]int, 0)
	var dfs func(int, int)
	dfs = func(start, sum int) {
		if target == sum {
			tmp := make([]int, len(st))
			copy(tmp, st)
			results = append(results, tmp)
			return
		}
		if start >= len(candidates) || sum > target {
			return
		}

		st = append(st, candidates[start])
		dfs(start, sum+candidates[start])
		st = st[:len(st)-1]

		dfs(start+1, sum)
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
	fmt.Println()
	results = combinationSum2(candidates, target)
	for _, res := range results {
		fmt.Println(res)
	}
}
