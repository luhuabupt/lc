package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	e38(os.Stdin, os.Stdout)
}

func e38(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int

	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
	}

	M := int64(998244353)
	ans := int64(0)

	o := &stNode{nil, nil, 1, 1e9 + 1, 0}
	for _, v := range a {
		ans += o.query(1, v)
		o.update(v, 1)
		ans %= M
	}

	Fprintln(out, ans)
}

type stNode struct {
	lo, ro *stNode
	l, r   int
	sum    int64
}

func (o *stNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (stNode) op(a, b int64) int64 {
	return a + b //
}

func (o *stNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *stNode) build(a []int64, l, r int) {
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

func (o *stNode) update(i int, add int64) {
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

func (o *stNode) query(l, r int) int64 {
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}
