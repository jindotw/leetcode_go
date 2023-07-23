package main

import (
	"fmt"
	"leetcode_go/schema"
)

func inOrderRecur(node *schema.TreeNode) {
	if node == nil {
		return
	}
	inOrderRecur(node.Left)
	fmt.Printf("%d ", node.Val)
	inOrderRecur(node.Right)
}

func inOrderIter(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	node := root

	for len(st) > 0 || node != nil {
		if node != nil {
			st = append(st, node)
			node = node.Left
		} else {
			last := len(st) - 1
			node = st[last]
			st = st[:last]
			fmt.Printf("%d ", node.Val)
			node = node.Right
		}
	}
}

func inOrderCommon(root *schema.TreeNode) {
	st := make([]*schema.TreeNode, 0, 10)
	if root != nil {
		st = append(st, root)
	}

	for len(st) > 0 {
		last := len(st) - 1
		node := st[last]
		st = st[:last]

		if node != nil {
			if node.Right != nil {
				st = append(st, node.Right)
			}
			st = append(st, node, nil)
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
