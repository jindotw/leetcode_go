package main

import "fmt"

func solution(statues []int) int {
	min := 21
	max := 0
	for _, val := range statues {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return (max - min + 1) - len(statues)
}

func main() {
	statues := []int{6, 2, 3, 8}
	missing := solution(statues)
	fmt.Println("need", missing, "more statues")
}
