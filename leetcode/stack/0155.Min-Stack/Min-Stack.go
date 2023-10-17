package main

import "fmt"

type element struct {
	val int
	min int
}

type MinStack struct {
	st []*element
}

func Constructor() MinStack {
	return MinStack{
		st: make([]*element, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.st) > 0 {
		currMin := this.GetMin()
		if currMin < min {
			min = currMin
		}
	}
	this.st = append(this.st, &element{
		val: val,
		min: min,
	})
}

func (this *MinStack) Pop() {
	last := len(this.st) - 1
	if last >= 0 {
		this.st = this.st[:last]
	}
}

func (this *MinStack) Top() int {
	if len(this.st) > 0 {
		return this.st[len(this.st)-1].val
	}
	return 0

}

func (this *MinStack) GetMin() int {
	if len(this.st) > 0 {
		return this.st[len(this.st)-1].min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(-0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
