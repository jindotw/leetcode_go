package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	matrix := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int, int)
	dfs = func(start int, depth int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			matrix = append(matrix, tmp)
			return
		}

		for i := start; i <= n-(k-len(path))+1; i++ {
			for x := 0; x < depth; x++ {
				fmt.Print("  ")
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, i)
			fmt.Printf("i=%d, cond<=%d, path_before=%v, path_after=%v\n", i, n-(k-len(path))+1, path, tmp)
			path = append(path, i)
			dfs(i+1, depth+1)
			path = path[:len(path)-1]
			for x := 0; x < depth; x++ {
				fmt.Print("  ")
			}
			fmt.Printf("Backtrack - next i=%d, cond<=%d, path=%v\n", i+1, n-(k-len(path))+1, path)
		}
	}

	dfs(1, 0)

	return matrix
}

func main() {
	n := 4
	k := 3
	matrix := combine(n, k)
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%02d ", val)
		}
		fmt.Println()
	}
}
