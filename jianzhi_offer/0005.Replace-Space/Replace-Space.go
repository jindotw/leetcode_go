package main

import "fmt"

func replaceSpace(s string) string {
	origin := []byte(s)
	count := 0
	for _, ch := range origin {
		if ch == ' ' {
			count++
		}
	}
	bytes := make([]byte, len(s)+count*2)

	for i, j := len(origin)-1, len(bytes)-1; i >= 0; i-- {
		if origin[i] != ' ' {
			bytes[j] = origin[i]
			j--
		} else {
			bytes[j], bytes[j-1], bytes[j-2] = '0', '2', '%'
			j = j - 3
		}
	}

	return string(bytes)
}

func main() {
	str := "Hello World"
	fmt.Println(replaceSpace(str))
}
