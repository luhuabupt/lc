package main

import (
	"fmt"
	"sort"
)

type BookMyShow struct {
	s   seg
	cap int
	n   int
}

func Constructor(n int, m int) BookMyShow {
	a := make([]int, n)
	s := newSegmentTree(a)
	return BookMyShow{s, m, n}
}

func (t *BookMyShow) Gather(k int, maxRow int) []int {
	idx := t.s.queryFirstLessPos(1, t.cap-k+1)
	fmt.Println("gat idx", idx)
	if idx == -1 || idx > maxRow {
		return []int{}
	}
	ori := t.s.query(1, idx, idx)
	t.s.update(1, idx, ori+k)
	fmt.Println("gat ori", ori)
	return []int{idx, ori + 1}
}

func (t *BookMyShow) Scatter(k int, maxRow int) bool {
	idx := sort.Search(t.n, func(i int) bool {
		sum := t.s.query(1, 0, i)
		return (i+1)*t.cap-sum >= k
	})
	fmt.Println("Scatter idx", idx)
	if idx > maxRow {
		return false
	}
	lef := (idx+1)*t.cap - t.s.query(1, 0, idx)
	t.s.update(1, idx, t.cap-lef)
	for i := 0; i < idx-1; i++ {
		t.s.update(1, i, t.cap)
	}
	fmt.Println("Scatter lef", lef)
	return true
}

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

type seg []struct {
	l, r int
	val  int
}

// 单点更新：build 和 update 通用
func (t seg) set(o, val int) {
	t[o].val = val
}

// 合并两个节点上的数据：maintain 和 query 通用
// 要求操作满足区间可加性
// 例如 + * | & ^ min max gcd mulMatrix 摩尔投票 最大子段和 ...
func (seg) op(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.op(lo.val, ro.val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	if m := (t[o].l + t[o].r) >> 1; i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t seg) queryAll() int { return t[1].val }

// EXTRA: 查询整个区间小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 需要先判断 t[1].min < v
func (t seg) queryFirstLessPos(o, v int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	if t[o<<1].val < v {
		return t.queryFirstLessPos(o<<1, v)
	}
	return t.queryFirstLessPos(o<<1|1, v)
}
