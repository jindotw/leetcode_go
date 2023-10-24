package main

import "fmt"

type Stack struct {
	arr []int
}

func (st *Stack) Push(val int) {
	st.arr = append(st.arr, val)
}

func (st *Stack) Pop() int {
	last := len(st.arr) - 1
	if last < 0 {
		panic("Stack is empty when popping")
	}
	val := st.arr[last]
	st.arr = st.arr[:last]
	return val
}

func (st *Stack) Peek() int {
	last := len(st.arr) - 1
	if last < 0 {
		panic("Stack is empty when peeking")
	}
	return st.arr[last]
}

func (st *Stack) IsEmpty() bool {
	return len(st.arr) == 0
}

func (st *Stack) IsNotEmpty() bool {
	return len(st.arr) > 0
}

func nextGreaterElements(nums []int) []int {
	//size := len(nums)
	//doubled := make([]int, size*2)
	//for idx, num := range nums {
	//	doubled[idx] = num
	//	doubled[idx+size] = num
	//}
	//
	//st := make([]int, 0)
	//result := make([]int, size*2)
	//for idx, num := range doubled {
	//	result[idx] = -1
	//	last := len(st) - 1
	//	for last >= 0 && num > doubled[st[last]] {
	//		result[st[last]] = num
	//		st = st[:last]
	//		last--
	//	}
	//	st = append(st, idx)
	//}
	//
	//return result[:size]
	n := len(nums)
	st := &Stack{}
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}
	for i := 0; i < n*2; i++ {
		num := nums[i%n]
		// while stack is not empty and the iterated number is greater than the stack's top element
		// for last := len(st) - 1; last >= 0 && num > nums[st[last]]; last-- {
		for st.IsNotEmpty() && num > nums[st.Peek()] {
			result[st.Pop()] = num
		}
		if i < n {
			st.Push(i)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(nextGreaterElements(nums))
}
