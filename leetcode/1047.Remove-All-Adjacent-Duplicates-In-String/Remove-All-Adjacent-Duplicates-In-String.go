package main

import "fmt"

type Stack []rune

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(val rune) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
func (s *Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func removeDuplicates(s string) string {
	r := []rune(s)
	st := Stack{}
	for i := 0; i < len(r); i++ {
		elem, peeked := st.Peek()
		if !peeked || elem != r[i] {
			st.Push(r[i])
		} else if elem == r[i] {
			st.Pop()
		}
	}

	result := []rune("")
	for !st.IsEmpty() {
		elem, _ := st.Pop()
		result = append(result, elem)
	}

	for i := 0; i < len(result)/2; i++ {
		rgt := len(result) - 1 - i
		result[i], result[rgt] = result[rgt], result[i]
	}

	return string(result)
}

func removeDuplicates2(s string) string {
	var result []rune
	for _, ch := range s {
		size := len(result)
		if size == 0 || result[size-1] != ch {
			result = append(result, ch)
		} else {
			result = result[:size-1]
		}
	}

	return string(result)
}

func main() {
	s := removeDuplicates2("abbaca")
	fmt.Println("s=", s)
}
