package main

import (
	"errors"
	"fmt"
)

// MonotonicQueue this is an implementation for a decreasing monotonic queue/**
type MonotonicQueue struct {
	elem []int
}

// Enqueue remove all the trailing elements if the value of element to be added is greater than that of those/**
func (queue *MonotonicQueue) Enqueue(val int) {
	for size := len(queue.elem); size > 0 && queue.elem[size-1] < val; {
		size--
		queue.elem = queue.elem[:size]
	}
	queue.elem = append(queue.elem, val)
}

// Dequeue remove the head element if it matches the dequeue value/**
func (queue *MonotonicQueue) Dequeue(val int) {
	if len(queue.elem) > 0 && queue.elem[0] == val {
		queue.elem = queue.elem[1:]
	}
}

func (queue *MonotonicQueue) Front() (int, error) {
	if len(queue.elem) > 0 {
		return queue.elem[0], nil
	}
	return 0, errors.New("empty queue")
}

func (queue *MonotonicQueue) Back() (int, error) {
	size := len(queue.elem)
	if size > 0 {
		return queue.elem[size-1], nil
	}
	return 0, errors.New("empty queue")
}

func (queue *MonotonicQueue) Dump() {
	for i := 0; i < len(queue.elem); i++ {
		fmt.Printf("[%d] %d\n", i, queue.elem[i])
	}
	fmt.Println("### End of Dump ###")
}
