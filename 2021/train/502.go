package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(1, 0, []int{1, 2, 3}, []int{1, 1, 2}))
}

type ho struct {
	sort.IntSlice
}

func (h ho) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h ho) Swap(i, j int)       { h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i] }
func (h *ho) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *ho) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(capital)

	list := make([][]int, n)
	for i, v := range capital {
		list[i] = []int{v, profits[i]} // 成本, 收益
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0] || list[i][0] == list[j][0] && list[i][1] > list[j][1]
	})

	h := &ho{}
	for i := 0; i < n; i++ {
		if w < list[i][0] {
			break
		}
		heap.Push(h, list[i][1])
		for (i == n-1 || list[i+1][0] > w) && k > 0 && h.Len() > 0 {
			w += heap.Pop(h).(int)
			k--
		}
		if k == 0 {
			break
		}
	}
	return w
}
