package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.Book(10, 40))
	fmt.Println(c.Book(50, 60))
	fmt.Println(c.Book(20, 30))
	fmt.Println(c.Book(20, 70))
	fmt.Println(c.Book(60, 70))
	fmt.Println(c.Book(30, 1e9))
	fmt.Println(c.Book(80, 1e9))
	fmt.Println(c.Book(90, 1e9))
}

type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int64
	todo   int64
}

// ((
func (o *lazyNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (lazyNode) op(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (o *lazyNode) do(add int64) {
	o.todo += add // % mod
	o.sum += add
}

// ))

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

//leetcode submit region begin(Prohibit modification and deletion)
type MyCalendarTwo struct {
	lazyNode
}

func Constructor() MyCalendarTwo {
	t := lazyNode{
		lo:   nil,
		ro:   nil,
		l:    0,
		r:    1e9,
		sum:  0,
		todo: 0,
	}
	return MyCalendarTwo{t}
}

func (t *MyCalendarTwo) Book(s int, e int) bool {
	t.update(s, e-1, 1)
	if t.query(0, 1e9) > 2 {
		t.update(s, e-1, -1)
		return false
	}
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)
