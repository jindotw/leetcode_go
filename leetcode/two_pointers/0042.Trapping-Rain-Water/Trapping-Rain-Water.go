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

func trap2(height []int) int {
	lft, rgt := 0, len(height)-1
	lftMax, rgtMax := height[lft], height[rgt]
	trapped := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for lft < rgt {
		if lftMax <= rgtMax {
			lft++
			lftMax = max(lftMax, height[lft])
			trapped += lftMax - height[lft]
		} else {
			rgt--
			rgtMax = max(rgtMax, height[rgt])
			trapped += rgtMax - height[rgt]
		}
	}
	return trapped
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap2(height))
}
