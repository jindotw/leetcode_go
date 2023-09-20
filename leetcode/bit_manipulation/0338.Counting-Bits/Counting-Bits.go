package main

import "fmt"

func countBits(n int) []int {
	bits := make([]int, n+1)

	countBits := func(i int) int {
		count := 0
		for i > 0 {
			count++
			i = i & (i - 1)
		}
		return count
	}
	for i := 0; i <= n; i++ {
		bits[i] = countBits(i)
	}

	return bits
}

func countBits2(n int) []int {
	bits := make([]int, n+1)
	bits[0] = 0
	if n >= 1 {
		bits[1] = 1
	}
	power := 2
	for i := 2; i <= n; i++ {
		if i/2 == power {
			power = i
			bits[i] = 1
			continue
		}
		bits[i] = bits[i-power] + 1
	}

	return bits
}

func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}

func main() {
	fmt.Println(countBits(20))
	fmt.Println(countBits2(20))
	fmt.Println(countBits3(20))
}
