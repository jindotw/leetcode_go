package main

import (
	"container/heap"
	"fmt"
)

type vertex struct {
	node   int
	weight int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type minHeap []*vertex

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return ((*h)[i]).weight < ((*h)[j]).weight
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*vertex))
}

func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	if last < 0 {
		panic("Popping an empty heap")
	}
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]*vertex)
	for _, time := range times {
		src := time[0]
		dst := time[1]
		weight := time[2]
		adj[src] = append(adj[src], &vertex{
			node:   dst,
			weight: weight,
		})
	}
	visited := make(map[int]struct{})
	maxWeight := 0
	minHeap := &minHeap{}
	heap.Push(minHeap, &vertex{
		node:   k,
		weight: 0,
	})

	for minHeap.Len() > 0 {
		v := heap.Pop(minHeap).(*vertex)
		if _, present := visited[v.node]; present {
			continue
		}
		visited[v.node] = struct{}{}
		maxWeight = max(maxWeight, v.weight)
		for _, dst := range adj[v.node] {
			heap.Push(minHeap, &vertex{
				node:   dst.node,
				weight: dst.weight + v.weight,
			})
		}
	}

	if len(visited) == n {
		return maxWeight
	}
	return -1
}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	fmt.Println(networkDelayTime(times, n, k))
}
