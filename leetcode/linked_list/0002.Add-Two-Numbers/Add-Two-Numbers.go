package main

import "leetcode_go/schema"

func addTwoNumbers(l1 *schema.ListNode, l2 *schema.ListNode) *schema.ListNode {
	carry := 0
	l3 := &schema.ListNode{Val: 0}
	dummy := l3

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		l3.Next = &schema.ListNode{Val: sum % 10}
		l3 = l3.Next
		carry = sum / 10
	}
	if carry > 0 {
		l3.Next = &schema.ListNode{Val: carry}
	}

	return dummy.Next
}

func main() {
	l1 := &schema.ListNode{Val: 2}
	l1.Next = &schema.ListNode{Val: 4}
	l1.Next.Next = &schema.ListNode{Val: 5}

	l2 := &schema.ListNode{Val: 5}
	l2.Next = &schema.ListNode{Val: 6}
	l2.Next.Next = &schema.ListNode{Val: 4}

	schema.PrintList(addTwoNumbers(l1, l2))
}
