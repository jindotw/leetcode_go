package main

import (
	"leetcode_go/schema"
)

func reorderList(head *schema.ListNode) {
	// Step 1: Create a slice to store all nodes from the linked list
	nodes := make([]*schema.ListNode, 0)
	for curr := head; curr != nil; curr = curr.Next {
		nodes = append(nodes, curr)
	}

	// Calculate the number of iterations needed to reorder the list
	loops := len(nodes) / 2
	if len(nodes)%2 == 0 {
		loops--
	}

	// Step 2: Reorder the linked list
	for curr, i := head, 0; i < loops; i++ {
		// Get the last node from the slice
		lastIndex := len(nodes) - 1
		tail := nodes[lastIndex]

		// Remove the last node from the slice
		nodes = nodes[:lastIndex]

		// Save a reference to the next node in the original order
		next := curr.Next

		// Update the current node's Next pointer to point to the tail node
		curr.Next = tail

		// Update the tail node's Next pointer to point to the next node in the original order
		tail.Next = next

		// Move to the next pair of nodes for reordering
		curr = next
	}

	// Step 3: Set the Next pointer of the last node to nil to terminate the list
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Next = nil
	}
}

func reorderList2(head *schema.ListNode) {
	// Step 1: Find the middle of the linked list using slow and fast pointers
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the linked list (from slow.Next onwards)
	var second *schema.ListNode
	for curr := slow.Next; curr != nil; {
		next := curr.Next
		curr.Next = second
		second = curr
		curr = next
	}

	// Step 3: Cut the linked list into two halves by setting slow.Next to nil
	slow.Next = nil

	// Step 4: Merge the first and reversed second halves of the linked list
	for first := head; second != nil; {
		firstNext, secondNext := first.Next, second.Next
		first.Next = second
		second.Next = firstNext
		first, second = firstNext, secondNext
	}
}

func main() {
	head := &schema.ListNode{
		Val: 1,
		Next: &schema.ListNode{
			Val: 2,
			Next: &schema.ListNode{
				Val: 3,
				Next: &schema.ListNode{
					Val: 4,
					Next: &schema.ListNode{
						Val: 5,
						Next: &schema.ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	reorderList(head)
	schema.PrintList(head)
}
