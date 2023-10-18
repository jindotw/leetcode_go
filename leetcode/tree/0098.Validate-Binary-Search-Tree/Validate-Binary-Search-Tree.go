package main

import (
	"fmt"
	"math"
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

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, lft, rgt int) bool {
		if node == nil {
			return true
		}
		if !(node.Val > lft && node.Val < rgt) {
			return false
		}

		return dfs(node.Left, lft, node.Val) && dfs(node.Right, node.Val, rgt)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func main() {
	// root := buildTree([]interface{}{5, 1, 4, nil, nil, 3, 6})
	// root := buildTree([]interface{}{5, 4, 6, nil, nil, 3, 7})
	root := buildTree([]interface{}{2, 1, 3})
	fmt.Println(isValidBST(root))
}
