package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(" ")
}

type ht [][]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] > h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1] }
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func findCrossingTime(n int, k int, t [][]int) int {
	l := &ht{}
	r := &ht{}
	for i, v := range t {
		heap.Push(l, []int{v[0] + v[2], i})
	}

	t := 0
	for do := 0; do < n; do++ {
		if l.Len() > 0 {
			x := l.Pop().([]int)
			
		} else {

		}
	}
}

func maxKelements(a []int, k int) int64 {
	h := &hl{}
	for _, v := range a {
		heap.Push(h, v)
	}
	ans := 0
	for i := 0; i < k; i++ {
		x := heap.Pop(h).(int)
		ans += x
		nx := x / 3
		if x%3 != 0 {
			nx++
		}
		heap.Push(h, nx)
	}
	return int64(ans)
}
