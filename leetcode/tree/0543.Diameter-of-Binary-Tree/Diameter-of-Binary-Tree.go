package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lftHeight := dfs(node.Left)
		rgtHeight := dfs(node.Right)
		maxDiameter = max(maxDiameter, 2+lftHeight+rgtHeight)
		return 1 + max(lftHeight, rgtHeight)
	}
	dfs(root)
	return maxDiameter
}

func main() {
	nums := []interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2}
	root := buildTree(nums)
	fmt.Println(diameterOfBinaryTree(root))
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
