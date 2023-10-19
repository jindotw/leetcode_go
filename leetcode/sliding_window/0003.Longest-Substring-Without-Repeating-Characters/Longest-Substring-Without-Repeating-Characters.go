package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	lft, longest := 0, 0
	set := make(map[byte]struct{})

	for rgt := range s {
		ch := s[rgt]
		if _, present := set[ch]; present {
			for s[lft] != ch {
				delete(set, s[lft])
				lft++
			}
			lft++
		}
		set[ch] = struct{}{}
		longest = max(longest, rgt-lft+1)
	}

	return longest
}

func main() {
	s := "abcbcabb"
	// s := "bbbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
