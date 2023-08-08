package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	last := h.Len() - 1
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	maxHeap := &MaxHeap{}
	*maxHeap = stones
	heap.Init(maxHeap)
	for maxHeap.Len() > 1 {
		stone1 := heap.Pop(maxHeap).(int)
		stone2 := heap.Pop(maxHeap).(int)
		if stone1 != stone2 {
			heap.Push(maxHeap, stone1-stone2)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return (*maxHeap)[0]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
