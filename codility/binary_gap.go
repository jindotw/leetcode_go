package main

import (
	"fmt"
)

func binaryGap(N int) int {
	curr, max := 0, 0
	started := false
	for N > 0 {
		if (N & 1) == 0 {
			if started {
				curr++
			}
		} else {
			if curr > max {
				max = curr
			}
			curr = 0
			started = true
		}
		N >>= 1
	}
	return max
}

func callBinaryGap() {
	fmt.Println(binaryGap(1041))
}
