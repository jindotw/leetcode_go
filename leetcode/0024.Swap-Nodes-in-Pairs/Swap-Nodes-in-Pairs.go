package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}

func swapListInPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next {
		originalNext := curr.Next  // node = 2
		curr.Next = curr.Next.Next // node = 1, its next = 3
		originalNext.Next = curr   // node = 2, its next = 1
		prev.Next = originalNext
		prev = curr
	}

	return dummy.Next
}

// 1 2 3 4
// 2 1 3 4
// 2 1 4 3

func main() {
	// your next points to your prev node
	// your prev node's next point to your next node
	l5 := &ListNode{Val: 5, Next: nil}
	l4 := &ListNode{Val: 4, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	printList(l1)
	printList(swapListInPairs(l1))
}
