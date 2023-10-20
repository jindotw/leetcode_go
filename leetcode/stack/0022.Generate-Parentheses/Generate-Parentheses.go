package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	st := make([]byte, 0)
	var backtracking func(int, int)
	backtracking = func(openN, closeN int) {
		if openN == closeN && openN == n {
			sb := strings.Builder{}
			sb.Write(st)
			result = append(result, sb.String())
			return
		}

		if openN < n {
			st = append(st, '(')
			backtracking(openN+1, closeN)
			st = st[:len(st)-1]
		}
		if closeN < openN {
			st = append(st, ')')
			backtracking(openN, closeN+1)
			st = st[:len(st)-1]
		}
	}
	backtracking(0, 0)

	return result
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
