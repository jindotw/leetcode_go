package main

import (
	"leetcode_go/schema"
)

func reverse(list *schema.ListNode) *schema.ListNode {
	var prev *schema.ListNode

	for list != nil {
		next := list.Next
		list.Next = prev
		prev = list
		list = next
	}

	return prev
}

func addTwoNumbersIter(l1 *schema.ListNode, l2 *schema.ListNode) *schema.ListNode {
	l3 := &schema.ListNode{Val: 0}
	dummy := l3
	carry := 0

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
		carry = sum / 10
		l3 = l3.Next
	}
	if carry > 0 {
		l3.Next = &schema.ListNode{Val: carry}
	}

	return reverse(dummy.Next)
}

func sizeOf(list *schema.ListNode) int {
	count := 0
	for list != nil {
		list = list.Next
		count++
	}
	return count
}

func recur(l1 *schema.ListNode, l2 *schema.ListNode, diff int) int {
	if l1 == nil || l2 == nil {
		return 0
	}

	if diff > 0 {
		carry := recur(l1.Next, l2, diff-1)
		sum := l1.Val + carry
		l1.Val = sum % 10
		return sum / 10
	} else {
		carry := recur(l1.Next, l2.Next, diff)
		sum := l1.Val + l2.Val + carry
		l1.Val = sum % 10
		return sum / 10
	}
}

func addTwoNumbers(l1 *schema.ListNode, l2 *schema.ListNode) *schema.ListNode {
	size1 := sizeOf(l1)
	size2 := sizeOf(l2)
	diff := size1 - size2
	if size2 > size1 {
		l1, l2 = l2, l1
		diff = size2 - size1
	}

	carry := recur(l1, l2, diff)
	if carry > 0 {
		return &schema.ListNode{
			Val:  carry,
			Next: l1,
		}
	}

	return l1
}

func main() {
	l1 := &schema.ListNode{Val: 2}
	l2 := &schema.ListNode{Val: 9}
	l2.Next = &schema.ListNode{Val: 9}
	schema.PrintList(addTwoNumbers(l1, l2))

	l1 = &schema.ListNode{Val: 2}
	l2 = &schema.ListNode{Val: 9}
	l2.Next = &schema.ListNode{Val: 9}
	schema.PrintList(addTwoNumbersIter(l1, l2))
}
