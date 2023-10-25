package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removed := 0
	prevEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] >= prevEnd {
			prevEnd = interval[1]
		} else {
			removed++
			if interval[1] < prevEnd {
				prevEnd = interval[1]
			}
		}
	}

	return removed
}

func main() {
	intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
