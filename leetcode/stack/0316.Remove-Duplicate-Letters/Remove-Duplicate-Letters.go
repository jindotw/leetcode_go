package main

import (
	"fmt"
	"strings"
)

func removeDuplicateLetters(s string) string {
	var counts [26]bool
	var lastSeen [26]int
	st := make([]byte, 0)
	// Initialize the lastSeen array to keep track of the last position where each character was seen.
	for i := range s {
		lastSeen[s[i]-'a'] = i
	}

	for i := range s {
		ch := s[i]
		if counts[ch-'a'] == true {
			continue
		}
		counts[ch-'a'] = true
		last := len(st) - 1
		// Check if the current character 'ch' is lexicographically smaller than the top character in the 'st' slice,
		// and whether there are more occurrences of the top character later in the string.
		for last >= 0 && ch < st[last] && i < lastSeen[st[last]-'a'] {
			// Remove the top character from 'st' as it has a greater occurrence later in the string.
			counts[st[last]-'a'] = false
			st = st[:last]
			last--
		}
		st = append(st, ch)
	}
	sb := strings.Builder{}
	sb.Write(st)
	return sb.String()
}

func main() {
	s := "bbbbbaa"
	fmt.Println(removeDuplicateLetters(s))
}
