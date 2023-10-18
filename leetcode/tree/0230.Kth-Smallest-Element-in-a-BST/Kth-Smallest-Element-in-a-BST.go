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

func kthSmallest(root *TreeNode, k int) int {
	st := make([]*TreeNode, 0)
	curr := root
	processed := 0

	for len(st) > 0 || curr != nil {
		if curr != nil {
			st = append(st, curr)
			curr = curr.Left
		} else {
			last := len(st) - 1
			node := st[last]
			st = st[:last]
			processed++
			if processed == k {
				return node.Val
			}
			curr = node.Right
		}
	}

	return -1
}

func kthSmallest2(root *TreeNode, k int) int {
	st := make([]*TreeNode, 0)
	if root != nil {
		st = append(st, root)
	}
	processed := 0
	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[:last]

		if node != nil {
			if node.Right != nil {
				st = append(st, node.Right)
			}
			st = append(st, node, nil)
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			last--
			node := st[last]
			st = st[:last]

			processed++
			if processed == k {
				return node.Val
			}
		}
	}

	return -1
}

func main() {
	root := buildTree([]interface{}{5, 3, 6, 2, 4, nil, nil, 1})
	k := 5
	fmt.Println(kthSmallest(root, k))
	fmt.Println(kthSmallest2(root, k))
}
