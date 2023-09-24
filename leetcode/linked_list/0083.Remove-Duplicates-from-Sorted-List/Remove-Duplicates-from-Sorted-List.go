package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	ptr := head

	for ptr != nil {
		for ptr.Next != nil && ptr.Next.Val == ptr.Val {
			ptr.Next = ptr.Next.Next
		}
		ptr = ptr.Next
	}

	return head
}

func main() {
	node32 := &ListNode{Val: 3, Next: nil}
	node31 := &ListNode{Val: 3, Next: node32}
	node21 := &ListNode{Val: 2, Next: node31}
	node13 := &ListNode{Val: 1, Next: node21}
	node12 := &ListNode{Val: 1, Next: node13}
	node11 := &ListNode{Val: 1, Next: node12}

	deleteDuplicates(node11)
	tmp := node11
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
