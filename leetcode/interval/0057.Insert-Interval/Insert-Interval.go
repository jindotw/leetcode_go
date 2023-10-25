package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			return append(res, intervals[i:]...)
		}
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else {
			newInterval = []int{
				min(newInterval[0], interval[0]),
				max(newInterval[1], interval[1]),
			}
		}
	}
	res = append(res, newInterval)

	return res
}

func main() {
	// [[1,3],[6,9]], newInterval = [2,5]
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
}
