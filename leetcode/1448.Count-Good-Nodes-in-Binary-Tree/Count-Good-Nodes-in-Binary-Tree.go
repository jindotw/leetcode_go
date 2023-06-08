package main

import (
	"fmt"
	"leetcode_go/leetcode/schema"
	"math"
)

func goodNodesRecur(node *schema.TreeNode, maxSeen int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val >= maxSeen {
		maxSeen = node.Val
		count++
	}
	count += goodNodesRecur(node.Left, maxSeen)
	count += goodNodesRecur(node.Right, maxSeen)

	return count
}

func goodNodes(root *schema.TreeNode) int {
	return goodNodesRecur(root, math.MinInt)
}

func main() {
	num := goodNodes(schema.MakeTree())
	fmt.Println(num)
}
