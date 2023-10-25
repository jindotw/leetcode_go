package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt
	}
	prices[src] = 0

	for i := 0; i < k+1; i++ {
		tmpPrice := make([]int, len(prices))
		copy(tmpPrice, prices)
		for _, flight := range flights {
			from, to, weight := flight[0], flight[1], flight[2]
			if prices[from] == math.MaxInt {
				continue
			}
			if prices[from]+weight < tmpPrice[to] {
				tmpPrice[to] = prices[from] + weight
			}
		}
		prices = tmpPrice
	}

	if prices[dst] == math.MaxInt {
		return -1
	}
	return prices[dst]
}

func main() {
	// n, src, dst, k, flights := 4, 0, 3, 1, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	n, src, dst, k, flights := 3, 0, 2, 1, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
