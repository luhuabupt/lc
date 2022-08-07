package main

import (
	"container/heap"
	"fmt"
)

type p struct {
	f  string
	sc int
}
type ht []p

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i].sc > h[j].sc || h[i].sc == h[j].sc && h[i].f < h[j].f } // da顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.(p)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type FoodRatings struct {
	d  map[string]*ht
	h  map[string]int
	fc map[string]string
}

func Constructor(f []string, c []string, r []int) FoodRatings {
	t := map[string]*ht{}
	h := map[string]int{}
	fc := map[string]string{}
	n := len(f)
	for i := 0; i < n; i++ {
		if _, ok := t[c[i]]; !ok {
			t[c[i]] = &ht{}
		}
		heap.Push(t[c[i]], p{f[i], r[i]})
		h[f[i]] = r[i]
		fc[f[i]] = c[i]
	}
	return FoodRatings{t, h, fc}
}

func (t *FoodRatings) ChangeRating(food string, newRating int) {
	t.h[food] = newRating
	heap.Push(t.d[t.fc[food]], p{food, newRating})
}

func (t *FoodRatings) HighestRated(v string) string {
	for t.d[v].Len() > 0 {
		fmt.Println(t.d[v].Len())
		x := heap.Pop(t.d[v]).(p)
		if t.h[x.f] != x.sc {
			continue
		} else {
			heap.Push(t.d[v], x)
			return x.f
		}
	}
	return ""
}
