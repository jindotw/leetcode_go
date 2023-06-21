package main

import (
	"fmt"
	"math"
)

func solution(inputArray []int) int {
	max := math.MinInt
	size := len(inputArray)
	for i, val := range inputArray {
		next := i + 1
		if next == size {
			break
		}
		product := val * inputArray[next]
		if product > max {
			max = product
		}
	}

	return max
}

func main() {
	arr := []int{-23, 4, -3, 8, -12}
	product := solution(arr)
	fmt.Println(product)
}
