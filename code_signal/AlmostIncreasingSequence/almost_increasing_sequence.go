package main

import "fmt"

func solution(sequence []int) bool {
	tail, count := 0, 0
	for i := 1; i < len(sequence); i++ {
		if sequence[i] > sequence[tail] {
			tail = i
		} else {
			if tail == 0 || sequence[i] > sequence[tail-1] {
				tail = i
			}
			if count++; count > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	arr := []int{1, 3, 4, 3, 5}
	// arr := []int{4, 3, 5, 67, 98}
	asc := solution(arr)
	fmt.Println(asc)
}
