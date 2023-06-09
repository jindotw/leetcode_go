package main

import "fmt"

type MyQueue struct {
	elem []int
}

func Constructor() MyQueue {
	return MyQueue{
		elem: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.elem = append(this.elem, x)
}

func (this *MyQueue) Pop() int {
	elem := this.elem[0]
	this.elem = this.elem[1:]
	return elem
}

func (this *MyQueue) Peek() int {
	return this.elem[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.elem) < 1
}

func main() {
	obj := Constructor()
	obj.Push(1)
	fmt.Println(obj.Peek())

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
