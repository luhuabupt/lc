package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDifference([]int{3, 1, 2}))
}

type h []int64

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(int64)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type hl []int64

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i] > h[j] } // 大
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.(int64)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minimumDifference(tmp []int) int64 {
	nums := []int64{}
	for _, v := range tmp {
		nums = append(nums, int64(v))
	}
	n := len(nums) / 3
	l, r := make([]int64, n+1), make([]int64, n+1)

	lh := &hl{}
	sl := int64(0)
	for i := 0; i < n; i++ {
		heap.Push(lh, nums[i])
		sl += nums[i]
	}
	l[0] = sl

	for i := 1; i <= n; i++ {
		l[i] = sl
		sl += nums[n-1+i]
		heap.Push(lh, nums[n-1+i])
		x := heap.Pop(lh).(int64)
		sl -= x
		if sl < l[i] {
			l[i] = sl
		}
	}

	sh := &h{}
	sr := int64(0)
	for i := 2 * n; i < 3*n; i++ {
		heap.Push(sh, nums[i])
		sr += nums[i]
	}
	r[n] = sr

	for i := 1; i <= n; i++ {
		r[n-i] = sr
		sr += nums[n+n-i]
		heap.Push(sh, nums[n+n-i])
		x := heap.Pop(sh).(int64)
		sr -= x

		if sr > r[n-i] {
			r[n-i] = sr
		}
	}

	//fmt.Println(sl, sr, l, r)

	ans := int64(math.MaxInt64)
	for i := 0; i <= n; i++ {
		if l[i]-r[i] < ans {
			ans = l[i] - r[i]
		}
	}
	return ans
}
