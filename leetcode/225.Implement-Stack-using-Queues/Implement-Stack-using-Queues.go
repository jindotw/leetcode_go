package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	pos := len(this.queue) - 1
	elem := this.queue[pos]
	this.queue = this.queue[:pos]
	return elem
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
