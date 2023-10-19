package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// Create the left child
		if i < len(nums) && nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Create the right child
		if i < len(nums) && nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lft := max(dfs(node.Left), 0)
		rgt := max(dfs(node.Right), 0)
		// compute the max path sum WITH split
		maxSum = max(maxSum, node.Val+lft+rgt)

		// return the max value WITHOUT split
		return node.Val + max(lft, rgt)
	}
	dfs(root)
	return maxSum
}

func main() {
	// root := []interface{}{8, 9, 20, nil, nil, 6, 7}
	root := []interface{}{2, 1, 3}
	fmt.Println(maxPathSum(buildTree(root)))
}
