package main

import (
	"fmt"
	"leetcode_go/schema"
)

func postOrderRecur(node *schema.TreeNode) {
	if node == nil {
		return
	}

	postOrderRecur(node.Left)
	postOrderRecur(node.Right)
	fmt.Printf("%d ", node.Val)
}

func PostOrderCommon(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[:last]

		if node != nil {
			st = append(st, node, nil)
			if node.Right != nil {
				st = append(st, node.Right)
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			last--
			node := st[last]
			st = st[:last]
			fmt.Printf("%d ", node.Val)
		}
	}
}

func PostOrderIter(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[:last]

		if node.Left == nil && node.Right == nil {
			fmt.Printf("%d ", node.Val)
			continue
		}
		st = append(st, node)
		if node.Right != nil {
			st = append(st, node.Right)
			node.Right = nil
		}
		if node.Left != nil {
			st = append(st, node.Left)
			node.Left = nil
		}
	}
}
