package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minimumWeight(6, [][]int{{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1}}, 0, 1, 5))
	fmt.Println(minimumWeight(6, [][]int{{0, 1, 1}, {2, 1, 1}}, 0, 1, 2))
}

type h [][3]int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i][1] < h[j][1] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.([3]int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	next := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		next[i] = make(map[int]int)
	}

	for _, a := range edges {
		if next[a[0]][a[1]] == 0 || a[2] < next[a[0]][a[1]] {
			next[a[0]][a[1]] = a[2]
		}
	}

	ans := 0

	vis := [2][]bool{}
	vis[0] = make([]bool, n)
	vis[1] = make([]bool, n)

	hp := &h{}
	heap.Push(hp, [3]int{src1, 0, 0})
	heap.Push(hp, [3]int{src2, 0, 1})

	return int64(ans)
}
