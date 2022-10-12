package main

import (
	"container/heap"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {

}

type hl [][]int

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i][0] > h[j][0] } // 大顶堆
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func maximumSegmentSum(a []int, rm []int) []int64 {
	n := len(a)
	ans := make([]int64, n)
	sum := make([]int, n)
	del := map[int]bool{0: true}
	h := &hl{}
	rbt := redblacktree.NewWithIntComparator()

	sst := 0
	for i, v := range a {
		sst += v
		sum[i] = sst
	}
	cal := func(l, r int) []int {
		if r < l {
			return []int{0, 0, 0}
		}
		if l == 0 {
			return []int{sum[r], l, r}
		}
		return []int{sum[r] - sum[l-1], l, r}
	}

	heap.Push(h, []int{sst, 0, n - 1})
	for i, v := range rm {
		r, hr := rbt.Floor(v)
		l, hl := rbt.Ceiling(v)
		lv, rv := 0, n-1
		if hr {
			rv = r.Value.(int)
		}
		if hl {
			lv = l.Value.(int)
		}
		rbt.Put(v, v)

		del[n*lv+rv] = true
		heap.Push(h, cal(lv+1, v-1))
		heap.Push(h, cal(rv-1, v+1))
		for h.Len() > 0 {
			x := heap.Pop(h).([]int)
			if !del[x[1]*n+x[2]] {
				ans[i] = int64(x[0])
				heap.Push(h, x)
			}
		}
	}
	return ans
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	do := make([]int, n+1)
	cal := func(v byte, k int) byte {
		k %= 26
		if k < 0 {
			k += 26
		}
		v = byte('a' + (int(v-'a')+k)%26)
		return v
	}
	for _, v := range shifts {
		if v[2] == 0 {
			v[2] = -1
		}
		do[v[0]] += v[2]
		do[v[1]+1] -= v[2]
	}
	t := 0
	for i, v := range do {
		t += v
		do[i] = v
	}
	ans := []byte(s)
	for i, v := range ans {
		ans[i] = cal(v, do[i])
	}
	return string(ans)
}

func secondsToRemoveOccurrences(s string) int {
	n := len(s)
	var dfs func(sa []byte) int
	dfs = func(sa []byte) int {
		ns := append([]byte{}, sa...)
		flag := false
		for i := 1; i < n; i++ {
			if sa[i] == '1' && sa[i-1] == '0' {
				ns[i-1] = '1'
				ns[i] = '0'
				i++
				flag = true
			}
		}
		if !flag {
			return 0
		}
		return dfs(ns) + 1
	}

	return dfs([]byte(s))
}

func minimumRecolors(a string, k int) int {
	n := len(a)
	b := make([]int, n)
	sb := 0
	ans := n
	for i, v := range a {
		if v == 'W' {
			sb++
		}
		b[i] = sb
		if i == k-1 {
			ans = b[i]
		}
		if i >= k && b[i]-b[i-k] < ans {
			ans = b[i] - b[i-k]
		}
	}
	if ans < 0 {
		ans = 0
	}
	return ans
}
