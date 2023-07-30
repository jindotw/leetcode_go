package main

import "fmt"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	size := len(height)
	st := make([]int, 0)
	st = append(st, 0)
	sum := 0

	for i := 1; i < size; i++ {
		if len(st) > 0 {
			if height[i] < height[st[len(st)-1]] {
				st = append(st, i)
			} else if height[i] == height[st[len(st)-1]] {
				st[len(st)-1] = i
			} else {
				rgt := i
				for len(st) > 0 && height[i] > height[st[len(st)-1]] {
					mid := st[len(st)-1]
					st = st[:len(st)-1]
					if len(st) > 0 {
						lft := st[len(st)-1]
						h := min(height[lft], height[rgt]) - height[mid]
						w := rgt - lft - 1
						sum += h * w
					}
				}
				st = append(st, i)
			}
		}
	}

	return sum
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
