package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	isAlphaNumerical := func(b byte) bool {
		return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') || ('0' <= b && b <= '9')
	}
	toLowerCase := func(b byte) byte {
		if 'A' <= b && b <= 'Z' {
			return b + 32
		}
		return b
	}
	for lft, rgt := 0, len(s)-1; lft < rgt; lft, rgt = lft+1, rgt-1 {
		for lft < rgt && !isAlphaNumerical(s[lft]) {
			lft++
		}
		for lft < rgt && !isAlphaNumerical(s[rgt]) {
			rgt--
		}
		if toLowerCase(s[lft]) != toLowerCase(s[rgt]) {
			return false
		}

	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
