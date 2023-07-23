package main

func reverseString(s []byte) {
	lft, rgt := 0, len(s)-1
	for lft < rgt {
		s[lft], s[rgt] = s[rgt], s[lft]
		lft++
		rgt--
	}
}

func main() {
	s := "hello"
	reverseString([]byte(s))
}
