package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	max, currCount := 0, 0
	lft, rgt := 0, 0
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	for rgt < len(s) {
		if isVowel(s[rgt]) {
			currCount++
		}

		if rgt-lft+1 == k {
			if currCount > max {
				max = currCount
			}
			if isVowel(s[lft]) {
				currCount--
			}
			lft++
		}
		rgt++
	}

	return max
}

func main() {
	s := "leetcode"
	k := 3
	fmt.Println(maxVowels(s, k))
}
