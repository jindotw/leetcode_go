package main

import "fmt"

func solution(n int) int {
	if n == 1 {
		return 1
	}
	half := 0
	for i := 1; i < n; i++ {
		half += 1 + (i-1)*2
	}
	middle := 1 + (n-1)*2

	return half*2 + middle
}

func main() {
	area := solution(4)
	fmt.Println(area)
}
