package schema

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}
