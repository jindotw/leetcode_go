package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		lft, lftBalanced := dfs(node.Left)
		if !lftBalanced {
			return 0, false
		}
		rgt, rgtBalanced := dfs(node.Right)
		if !rgtBalanced {
			return 0, false
		}
		if abs(lft-rgt) > 1 {
			return 0, false
		}

		return 1 + max(lft, rgt), true
	}
	_, balanced := dfs(root)
	return balanced
}

func main() {
	// nums := []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := buildTree(nums)
	fmt.Println(isBalanced(root))
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
