package main

import (
	"fmt"
	"sort"
)

func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	lastChild := 0
	childCount := len(g)
	for _, cookie := range s {
		if lastChild >= childCount {
			break
		}
		if cookie >= g[lastChild] {
			lastChild++
		}
	}

	return lastChild
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	cookie := len(s) - 1
	for i := len(g) - 1; cookie >= 0 && i >= 0; i-- {
		if s[cookie] >= g[i] {
			result++
			cookie--
		}
	}

	return result
}

func main() {
	children := []int{1, 2, 3}
	cookies := []int{1, 1}
	num := findContentChildren2(children, cookies)
	fmt.Println("Number of children satisfied:", num)
}
