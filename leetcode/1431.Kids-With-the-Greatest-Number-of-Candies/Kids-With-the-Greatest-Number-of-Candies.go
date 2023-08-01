package main

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := -1
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	results := make([]bool, 0)
	for _, candy := range candies {
		results = append(results, candy+extraCandies >= max)
	}

	return results
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
