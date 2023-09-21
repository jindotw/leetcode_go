package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows < 1 {
		return false
	}
	cols := len(matrix[0])
	lft, rgt := 0, rows*cols

	for lft < rgt {
		mid := lft + ((rgt - lft) >> 1)
		r := mid / cols
		c := mid % cols
		num := matrix[r][c]
		if num == target {
			return true
		} else if num < target {
			lft = mid + 1
		} else {
			rgt = mid
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 5
	fmt.Println(searchMatrix(matrix, target))
}
