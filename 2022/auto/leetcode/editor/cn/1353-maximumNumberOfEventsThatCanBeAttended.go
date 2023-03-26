package main

import (
	"container/heap"
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxEvents(a [][]int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	h := &h{}
	ans := 0
	i := 0
	t := a[0][0]

	for !(i == len(a) && h.Len() == 0) {
		for ; i < len(a) && a[i][0] == t; i++ {
			heap.Push(h, a[i][1])
		}

		for h.Len() > 0 {
			if heap.Pop(h).(int) >= t {
				ans++
				break
			}
		}
		if h.Len() == 0 && i < len(a) {
			t = a[i][0]
		} else {
			t++
		}
	}

	return ans
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
