package classic

import "container/heap"

type pair struct {
	w string
	c int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func topKFrequent(words []string, k int) []string {
	m, ans, h := map[string]int{}, make([]string, k), &hp{}
	for _, v := range words {
		m[v]++
	}
	for w, c := range m { // 入栈
		heap.Push(h, pair{w, c})
		if h.Len() > k { // 超过长度k出栈
			heap.Pop(h)
		}
	}
	for i := k - 1; i >= 0; i-- { // 前k出栈
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}

func _topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
