package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	x := len(text1)
	y := len(text2)
	memo := make([][]int, x+1)
	for i, _ := range memo {
		memo[i] = make([]int, y+1)
	}
	dump := func() {
		for _, r := range memo {
			fmt.Println(r)
		}
		fmt.Println()
	}
	for i, t1 := range text1 {
		for j, t2 := range text2 {
			if t1 == t2 {
				memo[i+1][j+1] = memo[i][j] + 1
			} else {
				memo[i+1][j+1] = max(memo[i][j+1], memo[i+1][j])
			}
		}
		fmt.Println("---", string(t1), "---")
		dump()
	}

	return memo[x][y]
}

func main() {
	str1 := "ezupkr"
	str2 := "ubmrapg"
	fmt.Println(longestCommonSubsequence(str1, str2))
}
