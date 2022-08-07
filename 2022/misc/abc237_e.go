package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

type pair struct{ d, i int }
type ht []pair

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

// https://atcoder.jp/contests/abc237/tasks/abc237_e
func main() {
	abc237_e(os.Stdin, os.Stdout)
}

func abc237_e(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	var m int
	var x int
	var y int
	var ans int

	Fscan(in, &n)
	Fscan(in, &m)

	w := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		w[i] = x
	}

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &x, &y)
		x--
		y--
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = 1e18
	}

	h := &ht{}
	heap.Push(h, pair{0, 0})

	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		i := p.i
		if p.d > dis[i] {
			continue
		}
		for _, x := range g[i] {
			cost := dis[i] + max(w[x]-w[i], 0)
			if cost < dis[x] {
				dis[x] = cost
				heap.Push(h, pair{dis[x], x})
			}
		}
	}

	for i, v := range dis {
		ans = max(ans, w[0]-w[i]-v)
	}

	Fprintln(out, ans)
}
