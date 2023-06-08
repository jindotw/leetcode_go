package main

import (
	"container/list"
	"fmt"
	"leetcode_go/leetcode/schema"
)

func rightSideView(root *schema.TreeNode) []int {
	var result []int
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	for true {
		size := queue.Len()
		if size == 0 {
			break
		}
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*schema.TreeNode)
			if i+1 == size {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
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
