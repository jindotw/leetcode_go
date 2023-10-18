package main

import (
	"fmt"
	"leetcode_go/schema"
)

func levelOrder(root *schema.TreeNode) [][]int {
	var queue []*schema.TreeNode
	var result [][]int
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		var row []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			row = append(row, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, row)
	}

	return result
}

func main() {
	node1 := schema.TreeNode{Val: 1}
	node3 := schema.TreeNode{Val: 3}
	node2 := schema.TreeNode{Val: 2, Left: &node1, Right: &node3}
	result := levelOrder(&node2)
	fmt.Println(result)
}
