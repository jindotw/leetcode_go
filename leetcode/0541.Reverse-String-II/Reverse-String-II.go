package main

import "fmt"

func reverseStr(s string, k int) string {
	bytes := []byte(s)

	reverse := func(lft, rgt int) {
		for lft < rgt {
			bytes[lft], bytes[rgt] = bytes[rgt], bytes[lft]
			lft++
			rgt--
		}
	}

	span := 2 * k
	for i := 0; i < len(bytes); i = i + span {
		remaining := len(bytes) - i
		if remaining >= span || remaining >= k {
			reverse(i, i+k-1)
		} else if remaining < k {
			reverse(i, len(bytes)-1)
		}
	}
	return string(bytes)
}

func main() {
	str := "abcd"
	k := 4
	fmt.Printf("%s\n", reverseStr(str, k))
}
