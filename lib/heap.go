package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type h []int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type hl []int

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i] > h[j] } // 大顶堆
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type ht [][2]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 小顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type minHeap struct{ sort.IntSlice }

func (minHeap) Push(interface{})     {}
func (minHeap) Pop() (_ interface{}) { return }

type maxHeap struct{ sort.IntSlice }

func (h maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (maxHeap) Push(interface{})     {}
func (maxHeap) Pop() (_ interface{}) { return }



// @see http://cngolib.com/container-heap.html
func main() {
	arr()
	sortSlice()
}

func arr() {
	h := &h{}
	heap.Push(h, 3)
	heap.Push(h, 4)
	heap.Push(h, 2)
	heap.Push(h, 5)
	heap.Push(h, 1)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

func sortSlice() {
	mh := minHeap{[]int{3, 5, 2, 7, 4, 6, 1}}
	heap.Init(&mh)

	// pop push 操作
	// i = 0, 等价于: 先remove(h, i) 再push
	mh.IntSlice[0] = 5
	heap.Fix(&mh, 0)
}
