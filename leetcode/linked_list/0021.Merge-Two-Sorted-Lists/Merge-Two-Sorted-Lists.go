package main

import "leetcode_go/schema"

func mergeTwoLists(list1 *schema.ListNode, list2 *schema.ListNode) *schema.ListNode {
	if list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}

	if list1 == nil {
		return list2
	}
	return list1
}

func main() {
	l1 := &schema.ListNode{
		Val: 1,
		Next: &schema.ListNode{
			Val: 2,
			Next: &schema.ListNode{
				Val: 4,
			},
		},
	}
	l2 := &schema.ListNode{
		Val: 2,
		Next: &schema.ListNode{
			Val: 3,
			Next: &schema.ListNode{
				Val: 5,
			},
		},
	}
	schema.PrintList(mergeTwoLists(l1, l2))
}
