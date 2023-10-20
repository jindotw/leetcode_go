package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	st := make([]int, 0)

	for _, op := range tokens {
		if operands, err := strconv.Atoi(op); err == nil {
			st = append(st, operands)
		} else {
			if len(st) < 2 {
				panic("Invalid RPN expression")
			}
			a, b := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			switch op {
			case "+":
				st = append(st, a+b)
			case "-":
				st = append(st, a-b)
			case "*":
				st = append(st, a*b)
			case "/":
				st = append(st, a/b)
			}
		}
	}
	return st[0]
}

func main() {
	// tokens := []string{"2", "1", "+", "3", "*"}
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// tokens := []string{"4", "3", "-"}
	fmt.Println(evalRPN(tokens))
}
