package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//  2  0
//  0  2
// -4  -4
//  2

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow = head; slow != fast; {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	n2 := &ListNode{Val: 2, Next: nil}
	n4 := &ListNode{Val: -4, Next: n2}
	n3 := &ListNode{Val: 0, Next: n4}
	n1 := &ListNode{Val: 3, Next: n2}
	n2.Next = n3

	node := detectCycle(n1)
	if node != nil {
		fmt.Println(node.Val)
	}
}
