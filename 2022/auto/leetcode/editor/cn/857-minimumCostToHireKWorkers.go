package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mincostToHireWorkers([]int{10, 20, 5}, []int{70, 50, 30}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
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

func mincostToHireWorkers(q []int, w []int, k int) float64 {
	n := len(q)

	type pair struct {
		do, pay, r float64
	}
	p := make([]pair, n)
	for i, v := range q {
		p[i] = pair{float64(v), float64(w[i]), float64(w[i]) / float64(q[i])}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].r < p[j].r
	})

	//fmt.Println(p)

	h := &hl{}
	s := 0.0
	for i := 0; i < k; i++ {
		heap.Push(h, int(p[i].do))
		s += p[i].do
	}
	ans := s * p[k-1].r

	//fmt.Println(s)

	for i := k; i < n; i++ {
		heap.Push(h, int(p[i].do))
		big := float64(heap.Pop(h).(int))
		s += p[i].do - big

		//fmt.Println(i, s, p[i].r)
		if s*p[i].r < ans {
			ans = s * p[i].r
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
