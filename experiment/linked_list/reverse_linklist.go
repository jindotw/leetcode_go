package main

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for curr := head; curr != nil; {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
