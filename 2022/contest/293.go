package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func removeAnagrams(w []string) []string {
	chk := func(a, b string) bool {
		c, d := [26]int{}, [26]int{}
		for _, v := range a {
			c[v-'a']++
		}
		for _, v := range b {
			d[v-'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] != d[i] {
				return false
			}
		}
		return true
	}

	for {
		nw := []string{w[0]}
		for i := 1; i < len(w); i++ {
			if chk(w[i], w[i-1]) {

			} else {
				nw = append(nw, w[i])
			}
		}
		if len(nw) == len(w) {
			break
		}
		w = nw
	}
	return w
}

func maxConsecutive(b int, t int, s []int) int {
	sort.Ints(s)
	ans := s[0] - b
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1]-1 > ans {
			ans = s[i] - s[i-1] - 1
		}
	}
	if t-s[len(s)-1] > ans {
		ans = t - s[len(s)-1]
	}
	return ans
}

func largestCombination(c []int) int {
	d := [31]int{}
	for i := 0; i < 31; i++ {
		for _, v := range c {
			if v&(1<<i) > 0 {
				d[i]++
			}
		}
	}
	ans := 0
	for _, v := range d {
		if v > ans {
			ans = v
		}
	}
	return ans
}

////////////////////////////

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int64
	todo   int64
}

func (o *lazyNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (lazyNode) op(a, b int64) int64 {
	return a + b //
}

func (o *lazyNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *lazyNode) build(a []int64, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &lazyNode{}
	o.lo.build(a, l, m)
	o.ro = &lazyNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *lazyNode) do(add int64) {
	o.todo = 1
	o.sum = int64(o.r - o.l + 1)
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == nil {
		o.lo = &lazyNode{l: o.l, r: m}
	}
	if o.ro == nil {
		o.ro = &lazyNode{l: m + 1, r: o.r}
	}
	if add := o.todo; add != 0 {
		o.lo.do(add)
		o.ro.do(add)
		o.todo = 0 // -1
	}
}

func (o *lazyNode) update(l, r int, add int64) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int64 {
	// 对于不在线段树中的点，应按照题意来返回
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

///////

type CountIntervals struct {
	*lazyNode
}

func Constructor() CountIntervals {
	return CountIntervals{&lazyNode{l: 1, r: 1e9}}
}

func (t *CountIntervals) Add(l int, r int) {
	t.update(l, r, 0)
}

func (t *CountIntervals) Count() int {
	return int(t.sum)
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
