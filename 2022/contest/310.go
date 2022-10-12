package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lengthOfLIS([]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3))
	fmt.Println(lengthOfLIS([]int{1, 3, 3, 4}, 1))
	fmt.Println(lengthOfLIS([]int{7, 4, 5, 1, 8, 12, 4, 7}, 5))
	fmt.Println(lengthOfLIS([]int{1, 5}, 1))
	fmt.Println(lengthOfLIS([]int{8, 13, 18, 23, 19, 6, 7, 8}, 3))
}

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

func lengthOfLIS(nums []int, k int) int {
	a := make([]int, 1e5+1)
	t := make(seg, 6*len(a))
	t.build(a, 1, 1, len(a))

	for _, v := range nums {
		mi := t.query(1, v-k, v-1)
		t.update(1, v, mi+1)
	}

	return t.query(0, 1, 1e5)
}

type h []int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minGroups(arr [][]int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || arr[i][0] == arr[i][0] && arr[i][1] < arr[j][1]
	})

	h := &h{}
	for _, v := range arr {
		if h.Len() == 0 {
			heap.Push(h, v[1])
			continue
		}
		x := heap.Pop(h).(int)
		if x < v[0] {
			heap.Push(h, v[1])
		} else {
			heap.Push(h, x)
			heap.Push(h, v[1])
		}
	}
	return h.Len()
}
