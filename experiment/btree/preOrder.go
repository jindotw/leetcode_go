package main

import (
	"fmt"
	"leetcode_go/schema"
)

func preOrderRecur(node *schema.TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	preOrderRecur(node.Left)
	preOrderRecur(node.Right)
}

func preOrderIter(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[0 : len(st)-1]
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
		fmt.Printf("%d ", node.Val)
	}
}

func preOrderCommon(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[0:last]

		if node != nil {
			if node.Right != nil {
				st = append(st, node.Right)
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
			st = append(st, node, nil)
		} else {
			last--
			node := st[last]
			st = st[0:last]
			fmt.Printf("%d ", node.Val)
		}
	}
}
