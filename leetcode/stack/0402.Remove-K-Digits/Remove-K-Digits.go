package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	st := make([]rune, 0)
	for _, number := range num {
		last := len(st) - 1
		for k > 0 && last >= 0 && number < st[last] {
			st = st[:last]
			k--
			last--
		}
		st = append(st, number)
	}
	if k > 0 {
		st = st[:len(st)-k]
	}

	for len(st) > 0 && st[0] == '0' {
		st = st[1:]
	}

	if len(st) == 0 {
		return "0"
	}

	return string(st)
}

func main() {
	num := "9534"
	k := 1
	fmt.Println(removeKdigits(num, k))
}
