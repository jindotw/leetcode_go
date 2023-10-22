package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	k       int
	minHeap *MinHeap
}

type MinHeap []int

func Constructor(k int, nums []int) KthLargest {
	maxHeap := &MinHeap{}
	*maxHeap = nums
	heap.Init(maxHeap)
	return KthLargest{k, maxHeap}
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	last := len(*h) - 1
	if last < 0 {
		return nil
	}
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}

func (h *MinHeap) top() int {
	return (*h)[0]
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	for this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return this.minHeap.top()
}

func main() {
	k := 3
	data := []int{4, 5, 8, 2}
	kth := Constructor(k, data)
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))

}
