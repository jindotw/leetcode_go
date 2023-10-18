package main

import (
	"fmt"
	"leetcode_go/schema"
)

func rightSideView(root *schema.TreeNode) []int {
	var result []int
	var queue []*schema.TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for true {
		size := len(queue)
		if size == 0 {
			break
		}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i+1 == size {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

func main() {
	node1 := schema.TreeNode{Val: 1}
	node3 := schema.TreeNode{Val: 3}
	node2 := schema.TreeNode{Val: 2, Left: &node1, Right: &node3}
	result := rightSideView(&node2)
	fmt.Println(result)
}
