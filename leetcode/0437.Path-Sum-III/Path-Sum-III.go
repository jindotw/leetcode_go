package main

import (
	"fmt"
	"leetcode_go/schema"
)

func preOrder(node *schema.TreeNode, targetSum int, cache []int) int {
	if node == nil {
		return 0
	}
	cache = append(cache, node.Val)
	sum := 0
	count := 0
	for i := len(cache) - 1; i >= 0; i-- {
		sum += cache[i]
		if targetSum == sum {
			count++
		}
	}
	count += preOrder(node.Left, targetSum, cache)
	count += preOrder(node.Right, targetSum, cache)

	return count
}

func pathSum(root *schema.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return preOrder(root, targetSum, []int{})
}

func main() {
	root := schema.MakeTreePathSum()
	count := pathSum(root, 8)
	fmt.Println("There are", count, "paths")
}
