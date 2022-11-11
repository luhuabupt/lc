package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(a []int) int {
	t := stNode{nil, nil, -2e10, 2e10, 0}
	ans := 0
	for _, v := range a {
		ans += t.query(v+1, 2e10)
		t.update(v, 1)
	}
	return ans
}

type stNode struct {
	lo, ro *stNode
	l, r   int
	sum    int
}

func (o *stNode) get() int {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (stNode) op(a, b int) int {
	return a + b //
}

func (o *stNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *stNode) build(a []int, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &stNode{}
	o.lo.build(a, l, m)
	o.ro = &stNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *stNode) update(i int, add int) {
	if o.l == o.r {
		o.sum += add //
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &stNode{l: o.l, r: m}
		}
		o.lo.update(i, add)
	} else {
		if o.ro == nil {
			o.ro = &stNode{l: m + 1, r: o.r}
		}
		o.ro.update(i, add)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) int {
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

//leetcode submit region end(Prohibit modification and deletion)
