package main

import "fmt"

func solution(inputString string) bool {
	for i, j := 0, len(inputString)-1; i < j; i, j = i+1, j-1 {
		if inputString[i] != inputString[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "abc"
	isPalindrome := solution(str)
	fmt.Println(isPalindrome)
}
