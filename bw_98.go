package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(handleQuery([]int{1, 0, 1}, []int{0, 0, 0}, [][]int{{1, 1, 2}, {2, 1, 0}, {3, 0, 0}}))
}

func beautifulSubsets(a []int, k int) int {
	n := len(a)
	sort.Ints(a)
	con := make([][]int, n)
	for i, v := range a {
		for j := 0; j < i; j++ {
			if v == a[j]-k {
				con[i] = append(con[i], j)
			}
		}
	}

	ans := 0

loop:
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				for _, x := range con[j] {
					if i&(1<<x) > 0 {
						continue loop
					}
				}
			}
			ans++
		}
	}
	return ans
	//c := map[int]int{}
	//for _, v := range a {
	//	c[v]++
	//}
	//
	//dp := make([]int, len(a))
	//ans := 0
	//for i, v := range a {
	//	dp[i] = 1
	//	for j := 0; j < i; j++ {
	//		if a[j] == v-k {
	//			continue
	//		}
	//		dp[i] += dp[j]
	//	}
	//
	//	ans += dp[i]
	//}
	//return ans
}

func squareFreeSubsets(a []int) int {
	chk := func(v int) bool {
		for i := 2; i < 6; i++ {
			if v%(i*i) == 0 {
				return true
			}
		}
		return false
	}
	d := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	M := int(1e9) + 7
	cal := func(v int) int {
		r := 0
		for i, x := range d {
			if v%x == 0 {
				r |= (1 << i)
			}
		}
		return r
	}

	c := [1024]int{}
	ans := 0
	for _, v := range a {
		if chk(v) {
			continue
		}
		ans++

		b := cal(v)
		for j := 0; j < 1024; j++ {
			if b&j > 0 {
				continue
			} else {
				ans += c[j]
				ans %= M
			}
		}
		c[b]++
	}
	return ans
}

func minOperations(n int) int {
	st := []int{}
	vis := make([]bool, 2e5+1)
	for i := 1; i <= 2e5 && i <= n*2; i *= 2 {
		st = append(st, i)
		vis[i] = true
	}

	d := 1
	for len(st) > 0 {
		ns := []int{}
		for _, v := range st {
			if v == n {
				return d
			}
			for i := 1; i <= 2e5 && i <= n*2; i *= 2 {
				x := v + i
				if x <= 2e5 && x <= n*2 && !vis[x] {
					vis[x] = true
					ns = append(ns, x)
				}

				y := v - i
				if y > 0 && !vis[y] {
					vis[y] = true
					ns = append(ns, y)
				}
			}
		}
		d++
		st = ns
	}

	return 0
}

func mergeArrays(a [][]int, b [][]int) [][]int {
	m := map[int]int{}
	for _, v := range a {
		m[v[0]] += v[1]
	}
	for _, v := range b {
		m[v[0]] += v[1]
	}

	ka := []int{}
	for k, _ := range m {
		ka = append(ka, k)
	}
	sort.Ints(ka)

	ans := [][]int{}
	for _, v := range ka {
		ans = append(ans, []int{v, m[v]})
	}
	return ans
}

func handleQuery(a []int, b []int, q [][]int) []int64 {
	s := int64(0)
	for _, v := range b {
		s += int64(v)
	}

	ia := []int64{}
	for _, v := range a {
		ia = append(ia, int64(v))
	}
	t := newLazySegmentTree(ia)

	ans := []int64{}
	for _, v := range q {
		if v[0] == 1 {
			t.update(1, v[1], v[2], 1)
			fmt.Println(t)
		}
		if v[0] == 2 {
			fmt.Println(v, t.queryAll())
			s += t.queryAll() * int64(v[1])
		}
		if v[0] == 3 {
			ans = append(ans, int64(s))
		}
	}
	return ans
}

type lazyST []struct {
	l, r int
	todo int64
	sum  int64
}

func (lazyST) op(a, b int64) int64 {
	return a + b // % mod
}

func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = t.op(lo.sum, ro.sum)
}

func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazyST) do(o int, add int64) {
	to := &t[o]
	//to.todo += add                     // % mod
	//to.sum += int64(to.r-to.l+1) * add // % mod
	to.sum ^= (add % 2)
}

func (t lazyST) spread(o int) {
	if add := t[o].todo; add != 0 {
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
	}
}

// 如果维护的数据（或者判断条件）具有单调性，我们就可以在线段树上二分
// 未找到时返回 n+1
// o=1  [l,r] 1<=l<=r<=n
// https://codeforces.com/problemset/problem/1179/C
func (t lazyST) binarySearch(o, l, r int, val int64) int {
	if t[o].l == t[o].r {
		if val <= t[o].sum {
			return t[o].l
		}
		return t[o].l + 1
	}
	t.spread(o)
	// 注意判断对象是当前节点还是子节点
	if val <= t[o<<1].sum {
		return t.binarySearch(o<<1, l, r, val)
	}
	return t.binarySearch(o<<1|1, l, r, val)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
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

func (t lazyST) queryAll() int64 { return t[1].sum }

// a 从 0 开始
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

func minimizeSum(a []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(a)

	n := len(a)
	if a[1]-a[0] > a[n-1]-a[n-2] {
		a = a[1:]
	} else {
		a = a[:n-1]
	}

	n = len(a)
	if a[1]-a[0] > a[n-1]-a[n-2] {
		a = a[1:]
	} else {
		a = a[:n-1]
	}

	b, c, mi := 0, 0, 1<<60
	for i := 1; i < len(a); i++ {
		x := a[i] - a[i-1]
		if x < mi {
			mi = x
		}
		if x >= b {
			b, c = x, b
		}
		if x > c {
			c = x
		}
	}

	ans := a[len(a)-1] - a[0]
	ans -= min(mi, max(b/3, c/2))

	return ans
}

func minMaxDifference(n int) int {
	if n < 10 {
		return 9
	}
	a := []int{}
	for x := n; x > 0; x /= 10 {
		a = append(a, x%10)
	}

	cal := func(a []int) int {
		r := 0
		for _, v := range a {
			r *= 10
			r += v
		}
		return r
	}

	b := append([]int{}, a...)
	c := append([]int{}, a...)
	for i := 0; i < len(a); i++ {
		if a[i] != 9 {
			for j := i; j < len(a); j++ {
				if a[j] == a[i] {
					b[j] = 9
				}
			}
			break
		}
	}
	for i := 1; i < len(a); i++ {
		if a[i] == a[0] {
			c[i] = 0
		}
	}

	return cal(b) - cal(c)
}
