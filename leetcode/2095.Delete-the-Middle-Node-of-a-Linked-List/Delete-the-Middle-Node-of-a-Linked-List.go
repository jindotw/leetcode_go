package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	var prev *ListNode
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	if fast == nil {
		slow = prev.Next
	}
	prev.Next = slow.Next

	return head
}

// 1 2 3 4 5 6

func main() {
	// n7 := &ListNode{Val: 6, Next: nil}
	// n6 := &ListNode{Val: 2, Next: nil}
	// n5 := &ListNode{Val: 1, Next: nil}
	// n4 := &ListNode{Val: 7, Next: nil}
	// n3 := &ListNode{Val: 4, Next: n4}
	n2 := &ListNode{Val: 3, Next: nil}
	n1 := &ListNode{Val: 1, Next: n2}
	printList(deleteMiddle(n1))
}

func printList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}
