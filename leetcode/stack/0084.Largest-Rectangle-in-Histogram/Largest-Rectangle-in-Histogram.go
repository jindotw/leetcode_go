package main

import "fmt"

type pair struct {
	idx int
	hgt int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	st := make([]pair, 0)
	for i, height := range heights {
		start := i
		last := len(st) - 1
		for last >= 0 && height < st[last].hgt {
			bar := st[last]
			maxArea = max(maxArea, (i-bar.idx)*bar.hgt)
			st = st[:last]
			last--
			start = bar.idx
		}

		st = append(st, pair{
			idx: start,
			hgt: height,
		})
	}
	for _, bar := range st {
		maxArea = max(maxArea, (len(heights)-bar.idx)*bar.hgt)
	}

	return maxArea
}

func main() {
	heights := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(largestRectangleArea(heights))
}
