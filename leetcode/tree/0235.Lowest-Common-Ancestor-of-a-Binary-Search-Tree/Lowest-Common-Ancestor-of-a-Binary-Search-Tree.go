package main

import "fmt"

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root

	for curr != nil {
		if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else {
			break
		}
	}
	return curr
}

func main() {
	root := buildTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	node := lowestCommonAncestor(root, root.Left, root.Right)
	if node != nil {
		fmt.Println(node.Val)
	}
}
