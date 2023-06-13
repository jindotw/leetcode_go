package main

import "fmt"

func reverseVowels(s string) string {
	r := []rune(s)
	lft := 0
	rgt := len(s) - 1

	for lft < rgt {
		if !isVowel(r[lft]) {
			lft++
			continue
		}

		if !isVowel(r[rgt]) {
			rgt--
			continue
		}

		r[lft], r[rgt] = r[rgt], r[lft]
		lft++
		rgt--
	}
	return string(r)
}

func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func main() {
	s := "hello"
	reversed := reverseVowels(s)
	fmt.Println(s, "=>", reversed)
}
