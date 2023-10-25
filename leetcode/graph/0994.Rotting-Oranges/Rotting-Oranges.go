package main

import (
	"fmt"
)

type coordinate [2]int

func orangesRotting(grid [][]int) int {
	freshOranges, days := 0, 0
	rows, cols := len(grid), len(grid[0])
	queue := make([]coordinate, 0)
	bfs := func() {
		directions := []coordinate{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		visited := make(map[coordinate]struct{})

		for len(queue) > 0 && freshOranges > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				v := queue[0]
				queue = queue[1:]
				if _, present := visited[v]; present {
					continue
				}
				visited[v] = struct{}{}
				for _, dir := range directions {
					r, c := v[0]+dir[0], v[1]+dir[1]
					if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != 1 {
						continue
					}

					grid[r][c] = 2
					freshOranges--
					queue = append(queue, coordinate{r, c})
				}
			}
			days++
		}
	}

	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				freshOranges++
			} else if val == 2 {
				queue = append(queue, coordinate{i, j})
			}
		}
	}
	bfs()

	if freshOranges > 0 {
		return -1
	}
	return days
}

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	// grid := [][]int{{0, 2}}
	// grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Println(orangesRotting(grid))
}
