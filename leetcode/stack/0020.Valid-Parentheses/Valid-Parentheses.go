package main

import "fmt"

func isValid(s string) bool {
	var stack []byte
	counterpart := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := range s {
		open, present := counterpart[s[i]]

		if !present {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // Pop the matching open parenthesis from the stack
		}
	}

	return len(stack) == 0 // Check if the stack is empty after processing all characters
}

func main() {
	s := "]"
	fmt.Println(isValid(s))
}
