package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB
	countLength := func(head *ListNode) (count int) {
		for ; head != nil; head, count = head.Next, count+1 {
		}
		return count
	}
	len1 := countLength(l1)
	len2 := countLength(l2)

	if len1 > len2 {
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}
	for diff := len2 - len1; diff > 0; diff-- {
		l2 = l2.Next
	}
	for ptr1, ptr2 := l1, l2; ptr1 != nil && ptr2 != nil; ptr1, ptr2 = ptr1.Next, ptr2.Next {
		if ptr1 == ptr2 {
			return ptr1
		}
	}

	return nil
}

func printList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d ", curr.Val)
	}
	fmt.Println()
}

func main() {
	//   4 1 8 4 5
	// 5 6 1 8 4 5

	a5 := &ListNode{Val: 5, Next: nil}
	a4 := &ListNode{Val: 4, Next: a5}
	a3 := &ListNode{Val: 8, Next: a4}
	a2 := &ListNode{Val: 1, Next: a3}
	a1 := &ListNode{Val: 4, Next: a2}

	b3 := &ListNode{Val: 1, Next: a3}
	b2 := &ListNode{Val: 6, Next: b3}
	b1 := &ListNode{Val: 5, Next: b2}

	intersected := getIntersectionNode(a1, b1)
	printList(intersected)
}
