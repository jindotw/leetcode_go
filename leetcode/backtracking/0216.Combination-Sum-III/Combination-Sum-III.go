package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	matrix := make([][]int, 0)
	path := make([]int, 0)
	sum := 0

	var dfs func(int)
	dfs = func(start int) {
		if len(path) == k {
			if sum > n {
				return
			}
			if sum == n {
				tmp := make([]int, len(path))
				copy(tmp, path)
				matrix = append(matrix, tmp)
			}
			return
		}
		// i <= 9 - (k - path.size()) + 1
		for i := start; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			sum += i
			dfs(i + 1)
			sum -= i
			path = path[:len(path)-1]
		}
	}

	dfs(1)

	return matrix
}

func main() {
	k := 3
	n := 9
	fmt.Println(combinationSum3(k, n))
}
