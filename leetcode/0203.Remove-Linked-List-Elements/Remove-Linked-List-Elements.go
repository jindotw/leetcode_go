package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}

	return dummy.Next
}

func main() {
	// [1,2,6,3,4,5,6]
	// [1, 2, 2, 1]

	l3 := &ListNode{Val: 1, Next: nil}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 2, Next: l2}
	l0 := &ListNode{Val: 2, Next: l1}
	head := removeElements(l0, 2)
	printList(head)
}

func printList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}
