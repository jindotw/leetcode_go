package main

import (
	"fmt"
	"leetcode_go/leetcode/schema"
)

func levelOrderBottom(root *schema.TreeNode) [][]int {
	var result [][]int
	var queue []*schema.TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for true {
		size := len(queue)
		if size == 0 {
			break
		}
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

	arrLen := len(result)
	for i := 0; i < arrLen/2; i++ {
		rgt := arrLen - i - 1
		result[i], result[rgt] = result[rgt], result[i]
	}

	return result
}

func main() {
	node1 := schema.TreeNode{Val: 1}
	node3 := schema.TreeNode{Val: 3}
	node2 := schema.TreeNode{Val: 2, Left: &node1, Right: &node3}
	result := levelOrderBottom(&node2)
	fmt.Println(result)
}
