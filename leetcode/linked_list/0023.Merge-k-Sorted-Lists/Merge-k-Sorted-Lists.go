package main

import (
	"container/heap"
	"fmt"
)

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

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return ((*h)[i]) < ((*h)[j])
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	last := h.Len() - 1
	if last < 0 {
		return nil
	}
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	for _, node := range lists {
		for node != nil {
			*minHeap = append(*minHeap, node.Val)
			node = node.Next
		}
	}
	heap.Init(minHeap)
	head := &ListNode{}
	curr := head
	for minHeap.Len() > 0 {
		curr.Next = &ListNode{Val: heap.Pop(minHeap).(int)}
		curr = curr.Next
	}

	return head.Next
}

func main() {
	// [[1,4,5],[1,3,4],[2,6]]
	node13 := &ListNode{Val: 5}
	node12 := &ListNode{Val: 4, Next: node13}
	node11 := &ListNode{Val: 1, Next: node12}

	node23 := &ListNode{Val: 4}
	node22 := &ListNode{Val: 3, Next: node23}
	node21 := &ListNode{Val: 1, Next: node22}

	node32 := &ListNode{Val: 6}
	node31 := &ListNode{Val: 2, Next: node32}

	lists := []*ListNode{node11, node21, node31}
	// var lists []*ListNode

	PrintList(mergeKLists(lists))
}
