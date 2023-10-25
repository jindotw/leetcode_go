package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := len(res) - 1
		if interval[0] > res[last][1] {
			res = append(res, interval)
		} else {
			res[last][1] = max(interval[1], res[last][1])
		}
	}

	return res
}

func main() {
	// [[1,3],[2,6],[8,10],[15,18]]
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {9, 18}}
	intervals := [][]int{{1, 4}, {0, 0}}
	fmt.Println(merge(intervals))
}
