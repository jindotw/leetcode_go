package main

import "fmt"

// TreeNode /*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func main() {
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}
	root := invertTree(node2)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}
