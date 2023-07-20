package main

import "fmt"

func isHappy(n int) bool {
	getSum := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}
	seen := make(map[int]*struct{})
	for n != 1 && seen[n] == nil {
		n, seen[n] = getSum(n), &struct{}{}
	}
	return n == 1
}

func main() {
	n := 19
	happy := isHappy(n)
	fmt.Printf("%d is happy? %t\n", n, happy)
}
