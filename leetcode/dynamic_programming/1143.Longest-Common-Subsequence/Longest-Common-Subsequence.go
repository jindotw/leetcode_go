package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
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

func longestCommonSubsequence2(text1, text2 string) int {
	y := len(text2)
	memo := make([]int, len(text2)+1)
	for _, t1 := range text1 {
		prev := 0
		for j, t2 := range text2 {
			tmp := memo[j+1]
			if t1 == t2 {
				memo[j+1] = 1 + prev
			} else {
				memo[j+1] = max(memo[j], memo[j+1])
			}
			fmt.Printf("j=%d, t1=%v, t2=%v, prev=%d, tmp=%d, memo=%v\n", j, string(t1), string(t2), prev, tmp, memo)
			prev = tmp

		}
		fmt.Println()
	}

	return memo[y]
}

func main() {
	str1 := "ezupkr"
	str2 := "ubmrapg"
	// str1, str2 := "abcde", "ace"
	//str1, str2 := "ace", "abcde"
	// fmt.Println(longestCommonSubsequence1(str1, str2))
	fmt.Println(longestCommonSubsequence2(str1, str2))
}
