package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxPerformance(n int, s []int, e []int, k int) int {
	a := make([][]int, n)
	for i, v := range s {
		a[i] = []int{e[i], v}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})

	ans := 0
	h := &h{}
	t := 0
	for i := 0; i < n; i++ {
		heap.Push(h, a[i][1])
		t += a[i][1]
		if h.Len() > k {
			t -= heap.Pop(h).(int)
		}

		if t*a[i][0] > ans {
			ans = t * a[i][0]
		}
	}

	return ans % int(1e9+7)
}

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

//leetcode submit region end(Prohibit modification and deletion)
