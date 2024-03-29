package main

import "fmt"

const shiftAmount = 16

func encode(i, j int) int {
	return (i << shiftAmount) | j
}

func decode(key int) (int, int) {
	return key >> shiftAmount, key & (1<<shiftAmount - 1)
}

func numIslands(grid [][]byte) int {
	islands := 0
	set := make(map[int]struct{})
	rows, cols := len(grid), len(grid[0])

	bfs := func(m, n int) {
		directions := [][]int{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		}
		queue := make([]int, 0)
		pos := encode(m, n)
		set[pos] = struct{}{}
		queue = append(queue, pos)

		for len(queue) > 0 {
			pos = queue[0]
			m, n = decode(pos)
			queue = queue[1:]
			for _, dir := range directions {
				r, c := m+dir[0], n+dir[1]
				pos = encode(r, c)
				if _, ok := set[pos]; r >= 0 && c >= 0 && r < rows && c < cols && grid[r][c] == '1' && !ok {
					set[pos] = struct{}{}
					queue = append(queue, pos)
				}
			}
		}
	}

	for m, row := range grid {
		for n, val := range row {
			pos := encode(m, n)
			if _, ok := set[pos]; val == '1' && !ok {
				bfs(m, n)
				islands++
			}
		}
	}

	return islands
}

func numIslands2(grid [][]byte) int {
	islands := 0
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(m, n int) {
		if m < 0 || n < 0 || m >= rows || n >= cols || grid[m][n] == '0' {
			return
		}
		// mark grid[m][n] as visited
		grid[m][n] = '0'
		dfs(m+1, n)
		dfs(m-1, n)
		dfs(m, n+1)
		dfs(m, n-1)
	}

	for m, row := range grid {
		for n, val := range row {
			if val == '1' {
				islands++
				dfs(m, n)
			}
		}
	}

	return islands
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '1'},
	}
	fmt.Println(numIslands(grid))
	fmt.Println(numIslands2(grid))
}
