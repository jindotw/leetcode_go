package main

import (
	"fmt"
	"leetcode_go/leetcode/schema"
)

func longestZigZag(root *schema.TreeNode) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(*schema.TreeNode, int, int, *int)
	dfs = func(node *schema.TreeNode, left int, right int, path *int) {
		if node == nil {
			return
		}

		*path = max(*path, max(left, right))
		dfs(node.Left, 0, left+1, path)
		dfs(node.Right, right+1, 0, path)
	}

	path := 0
	dfs(root, 0, 0, &path)
	return path
}

func main() {
	root := schema.MakeTree()
	zigZag := longestZigZag(root)
	fmt.Println("Zigzag:", zigZag)
}
