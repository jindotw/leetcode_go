package main

import "fmt"

/*
brute force version
*/
func canCompleteCircuit(gas []int, cost []int) int {
	loops := len(gas)
	for i := 0; i < loops; i++ {
		tank := gas[i] - cost[i]
		j := (i + 1) % loops
		for tank >= 0 && j != i {
			tank += gas[j] - cost[j]
			j = (j + 1) % loops
		}
		if tank >= 0 && j == i {
			return i
		}
	}

	return -1
}

func canCompleteCircuit2(gas []int, cost []int) int {
	curr, total, start := 0, 0, 0
	for idx, val := range gas {
		curr += val - cost[idx]
		total += val - cost[idx]

		if curr < 0 {
			curr = 0
			start = (idx + 1) % len(gas)
		}
	}
	if total < 0 {
		return -1
	}

	return start
}

func main() {
	gas := []int{3, 1, 1}
	cost := []int{1, 2, 2}
	station := canCompleteCircuit2(gas, cost)
	fmt.Println("station:", station)
}
