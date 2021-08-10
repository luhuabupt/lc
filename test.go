package main

import (
	"container/heap"
	"fmt"
	"sort"
)

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

func main() {

}

func testVar() {
	a := []int{0, 1, 2, 3}
	for len(a) > 0 {
		var t int
		t, a = a[0], a[1:]
		t, a := a[0], a[1:]
		fmt.Println(t, a)
	}
}

func topKFrequent(words []string, k int) []string {
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

func isPrefixString(s string, words []string) bool {
	idx := 0
	for _, str := range words {
		if idx+len(str) > len(s) {
			return false
		}
		for i := 0; i < len(str); i++ {
			if str[i] != s[idx] {
				return false
			}
			idx++
		}
		if idx == len(s) {
			return true
		}
	}
	return idx == len(s)
}

func minStoneSum(piles []int, k int) int {
	sort.Ints(piles)
	for i := len(piles) - 1; i >= 0; i++ {

	}
}
