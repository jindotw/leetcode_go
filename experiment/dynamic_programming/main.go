package main

import "fmt"

func callFibonacciSequence() {
	fmt.Println(fibonacciSequence(19))
}

func callKnapsack01() {
	values := []int{15, 20, 30}
	weight := []int{1, 3, 4}
	bagSize := 4
	fmt.Println(knapsack01(values, weight, bagSize))
}

func main() {
	callKnapsack01()
}
