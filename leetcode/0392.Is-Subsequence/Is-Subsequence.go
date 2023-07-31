package main

import "fmt"

func isSubsequence(s string, t string) bool {
	i, sLen, tLen := 0, len(s), len(t)

	for j := 0; i < sLen && j < tLen; j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}

func main() {
	t := "ahbgdc"
	s := "axc"
	fmt.Println(isSubsequence(s, t))
}
