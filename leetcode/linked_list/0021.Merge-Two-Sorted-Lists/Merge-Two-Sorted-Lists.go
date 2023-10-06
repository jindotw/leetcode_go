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

func mergeTwoListsIter(list1 *schema.ListNode, list2 *schema.ListNode) *schema.ListNode {
	list3 := &schema.ListNode{}
	dummy := list3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list3.Next = list1
			list1 = list1.Next
		} else {
			list3.Next = list2
			list2 = list2.Next
		}
		list3 = list3.Next
	}

	if list1 != nil {
		list3.Next = list1
	}
	if list2 != nil {
		list3.Next = list2
	}

	return dummy.Next
}

func main() {
	l1 := &schema.ListNode{
		Val: 1,
		Next: &schema.ListNode{
			Val: 3,
			Next: &schema.ListNode{
				Val: 5,
			},
		},
	}
	l2 := &schema.ListNode{
		Val: 2,
		Next: &schema.ListNode{
			Val: 4,
			Next: &schema.ListNode{
				Val: 6,
			},
		},
	}
	schema.PrintList(mergeTwoListsIter(l1, l2))
}
