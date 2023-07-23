package main

import (
	"fmt"
	"leetcode_go/schema"
)

func bfs(root *schema.TreeNode) {
	res := make([][]int, 0)
	queue := make([]*schema.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		size = len(level)
		for lft := 0; lft < size/2; lft++ {
			rgt := size - 1 - lft
			level[lft], level[rgt] = level[rgt], level[lft]
		}

		res = append(res, level)
	}

	fmt.Println(res)
}
