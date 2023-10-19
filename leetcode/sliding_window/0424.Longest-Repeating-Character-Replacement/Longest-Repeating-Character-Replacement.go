package main

import "fmt"

func characterReplacement(s string, k int) int {
	counts := [26]int{}
	lft, maxFreq, maxRepeats := 0, 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for rgt := range s {
		ch := s[rgt]
		counts[ch-'A']++
		maxFreq = max(maxFreq, counts[ch-'A'])

		if (rgt-lft+1)-maxFreq > k {
			counts[s[lft]-'A']--
			lft++
		}
		maxRepeats = max(maxRepeats, rgt-lft+1)
	}

	return maxRepeats
}

func main() {
	str := "EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH"
	k := 7
	fmt.Println(characterReplacement(str, k))
}
