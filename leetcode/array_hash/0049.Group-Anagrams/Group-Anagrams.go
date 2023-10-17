package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	cache := make(map[[26]byte][]string)
	for _, str := range strs {
		count := [26]byte{}
		for _, ch := range str {
			count[ch-'a']++
		}
		cache[count] = append(cache[count], str)
	}

	res := make([][]string, 0, len(cache))
	for _, v := range cache {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := groupAnagrams(strs)
	for _, pairs := range grouped {
		fmt.Println(pairs)
	}
}
