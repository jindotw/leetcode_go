package main

import (
	"fmt"
)

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	} else if n == 4 {
		return 4
	}

	prd := 1
	for n > 4 {
		prd *= 3
		n -= 3
	}
	prd *= n
	return prd
}

func main() {
	prd := integerBreak(10)
	fmt.Println(prd)
}
