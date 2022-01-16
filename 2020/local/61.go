package main

type MyCircularQueue struct {
	arr     []int
	l, r, n int
}

func Constructor(k int) MyCircularQueue {
	a := []int{}
	for i := 0; i < k; i++ {
		a = append(a, -1)
	}
	return MyCircularQueue{
		a,
		0,
		0,
		k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.r = (this.r + 1) % this.n
	this.arr[this.r] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.arr[this.l] = -1

	next := (this.l + 1) % this.n
	if this.arr[next] != -1 {
		this.l = next
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	return this.arr[this.l]
}

func (this *MyCircularQueue) Rear() int {
	return this.arr[this.r]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.l == this.r && this.l == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.r+1)%this.n == this.l
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
