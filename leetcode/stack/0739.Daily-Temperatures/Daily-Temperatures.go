package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	result := make([]int, size)
	st := make([]int, 0)

	for idx, temp := range temperatures {
		result[idx] = 0
		last := len(st) - 1
		for last >= 0 && temp > temperatures[st[last]] {
			result[st[last]] = idx - st[last]
			st = st[:last]
			last--
		}

		st = append(st, idx)
	}

	return result
}

func main() {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temps))
}
