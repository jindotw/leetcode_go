package main

import "fmt"

func removeStars(s string) string {
	var st []byte
	for _, b := range []byte(s) {
		if b == '*' {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, b)
		}
	}
	return string(st)
}

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}
