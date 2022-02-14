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

// 小顶堆
type sh struct{ sort.IntSlice }

func (sh) Push(interface{})     {}
func (sh) Pop() (_ interface{}) { return }

// 大顶堆
type bh struct{ sort.IntSlice }

func (h bh) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (bh) Push(interface{})     {}
func (bh) Pop() (_ interface{}) { return }

// @see http://cngolib.com/container-heap.html
func main() {
	h := &bh{}
	heap.Push(h, 3)
	heap.Push(h, 8)

	heap.Push(h, 2)
	heap.Push(h, 7)
	heap.Push(h, 1)
	heap.Push(h, 4)
	x := heap.Pop(h).(int)
	fmt.Println(x)

	x = heap.Pop(h).(int)
	fmt.Println(x)

	x = heap.Pop(h).(int)
	fmt.Println(x)
	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
