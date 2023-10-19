package main

import "fmt"

func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var s1Count, s2Count [26]int
	matches := 0
	for i := range s1 {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	for lft, rgt := 0, len(s1); rgt < len(s2); lft, rgt = lft+1, rgt+1 {
		if matches == 26 {
			return true
		}
		index := s2[rgt] - 'a'
		s2Count[index]++
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			matches--
		}
		index = s2[lft] - 'a'
		s2Count[index]--
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			matches--
		}
	}

	return matches == 26
}

func checkInclusion2(s1 string, s2 string) bool {
	lft, count, empty := 0, [26]int{}, [26]int{}
	for i := range s1 {
		count[s1[i]-'a']++
	}

	for rgt := range s2 {
		count[s2[rgt]-'a']--
		if count == empty {
			return true
		}

		if rgt+1 >= len(s1) {
			count[s2[lft]-'a']++
			lft++
		}
	}
	return false
}

func main() {
	s1, s2 := "ab", "eidoooba"
	fmt.Println(checkInclusion(s1, s2))
	fmt.Println(checkInclusion2(s1, s2))
}
