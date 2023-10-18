package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	freq := make(map[byte]int)
	for _, r := range t {
		freq[byte(r)]++
	}

	left, have, need := 0, 0, len(freq)
	minLen, start := math.MaxInt, 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if _, present := freq[ch]; present {
			freq[ch]--
			if freq[ch] == 0 {
				have++
			}
		}

		for have == need {
			if strLen := right - left + 1; strLen < minLen {
				minLen = strLen
				start = left
			}

			ch := s[left]
			if _, present := freq[ch]; present {
				freq[ch]++
				if freq[ch] > 0 {
					have--
				}
			}
			left++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}

	return s[start : start+minLen]
}

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	// s, t := "a", "b"
	fmt.Println(minWindow(s, t))
}
