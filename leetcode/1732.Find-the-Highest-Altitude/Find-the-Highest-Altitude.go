package main

import "fmt"

func largestAltitude(gain []int) int {
	highest, curr := 0, 0
	for _, val := range gain {
		curr += val
		if curr > highest {
			highest = curr
		}
	}
	return highest
}

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}
