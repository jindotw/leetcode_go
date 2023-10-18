package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	depth := 0
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth2(root.Left), maxDepth2(root.Right))
}

func main() {
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}
	fmt.Println(maxDepth(node2))
	fmt.Println(maxDepth2(node2))
}
