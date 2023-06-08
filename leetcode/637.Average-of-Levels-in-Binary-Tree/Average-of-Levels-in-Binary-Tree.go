package main

import (
	"fmt"
	"leetcode_go/leetcode/schema"
)

func averageOfLevels(root *schema.TreeNode) []float64 {
	var result []float64
	var queue []*schema.TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for true {
		size := len(queue)
		if size == 0 {
			break
		}
		sum := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}

	return result
}

func main() {
	node1 := schema.TreeNode{Val: 1}
	node3 := schema.TreeNode{Val: 3}
	node2 := schema.TreeNode{Val: 2, Left: &node1, Right: &node3}
	result := averageOfLevels(&node2)
	fmt.Println(result)
}
