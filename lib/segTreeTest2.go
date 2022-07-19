package main

// 动态开点线段树·其二·延迟标记（区间修改）
// https://codeforces.com/problemset/problem/915/E（注：此题有多种解法）
// https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/F https://www.luogu.com.cn/problem/P5848
//（内存受限）https://codeforces.com/problemset/problem/1557/D
// rt := &lazyNode{l: 1, r: 1e9}
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
	o.todo += add                   // % mod
	o.sum += int64(o.r-o.l+1) * add // % mod
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
