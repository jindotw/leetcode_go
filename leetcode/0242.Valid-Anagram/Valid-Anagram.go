package main

import "fmt"

func isAnagramSlice(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr1, arr2 := [26]int{}, [26]int{}

	for _, ch := range s {
		arr1[ch-'a']++
	}
	for _, ch := range t {
		arr2[ch-'a']++
	}
	return arr1 == arr2
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[rune]int)
	for _, ch := range s {
		counter[ch]++
	}
	for _, ch := range t {
		if count, ok := counter[ch]; !ok {
			return false
		} else {
			if count = count - 1; count == 0 {
				delete(counter, ch)
			} else {
				counter[ch] = count
			}
		}
	}
	return len(counter) == 0
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Printf("%t\n", isAnagramSlice(s, t))
}
