package main

import "fmt"

func partition(s string) [][]string {
	matrix := make([][]string, 0)
	path := make([]string, 0)

	isPalindrome := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			matrix = append(matrix, tmp)
		}

		for i := start; i < len(s); i++ {
			substring := s[start : i+1]
			if isPalindrome(substring) {
				path = append(path, substring)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return matrix
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
