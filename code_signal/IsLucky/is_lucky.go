package main

import (
	"fmt"
)

func digit(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		n *= -n
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}

	return count
}

func solution(n int) bool {
	digits := digit(n)
	if digits%2 != 0 {
		return false
	}
	digits /= 2
	lft, rgt := 0, 0
	for i := 0; i < digits; i++ {
		rgt += n % 10
		n /= 10
	}
	for i := 0; i < digits; i++ {
		lft += n % 10
		n /= 10
	}

	fmt.Println(lft, rgt)

	return lft == rgt
}

func main() {
	fmt.Println("n digits", digit(-100))
	isLucky := solution(452353)
	fmt.Println(isLucky)
}
