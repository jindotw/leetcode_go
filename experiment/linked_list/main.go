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

func main() {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	printList(n1)
	printList(reverseList(n1))
}
