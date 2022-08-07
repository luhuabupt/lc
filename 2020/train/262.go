package main

import (
	"container/heap"
	"sort"
)

type hs struct{ sort.IntSlice }

func (h hs) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h hs) Swap(i, j int)       { h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i] }
func (h *hs) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hs) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type hs2 struct{ sort.IntSlice }

func (h hs2) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h hs2) Swap(i, j int)       { h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i] }
func (h *hs2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hs2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type IntHeap2 []int

func (h IntHeap2) Len() int { return len(h) }

// 为了实现大根堆，Less在大于时返回小于
func (h IntHeap2) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *IntHeap2) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type StockPrice struct {
	timePrice map[int]int
	bigH      *hs
	smallH    *hs2
	priceMap  map[int]int
	mxTs      int
}

func Constructor() StockPrice {
	return StockPrice{
		timePrice: map[int]int{},
		bigH:      &hs{},
		smallH:    &hs2{},
		priceMap:  map[int]int{},
		mxTs:      0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.timePrice[timestamp] > 0 {
		this.priceMap[this.timePrice[timestamp]]--
	}

	this.timePrice[timestamp] = price
	this.priceMap[price]++

	if timestamp > this.mxTs {
		this.mxTs = timestamp
	}
	heap.Push(this.bigH, price)
	heap.Push(this.smallH, price)
}

func (this *StockPrice) Current() int {
	return this.timePrice[this.mxTs]
}

func (this *StockPrice) Maximum() int {
	h := this.bigH
	for {
		x := h.Pop().(int)
		if this.priceMap[x] > 0 {
			heap.Push(this.bigH, x)
			this.priceMap[x]++
			return x
		}
	}
}

func (this *StockPrice) Minimum() int {
	h := this.smallH
	for {
		x := h.Pop().(int)
		if this.priceMap[x] > 0 {
			heap.Push(this.smallH, x)
			this.priceMap[x]++
			return x
		}
	}
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

func minOperations(grid [][]int, x int) int {
	sum := 0
	cnt := map[int]int{}

	for _, arr := range grid {
		for _, v := range arr {
			sum += v
			cnt[v]++
		}
	}

	arr := []int{}
	for i := 1; i <= 10000; i++ {
		if cnt[i] > 0 {
			arr = append(arr, i)
		}
	}
	for i := 1; i < len(arr); i++ {
		if (arr[i]-arr[i-1])%x != 0 {
			return -1
		}
	}

	ar := []int{}
	for i := 1; i <= 10000; i++ {
		for cnt[i] > 0 {
			ar = append(ar, i)
			cnt[i]--
		}
	}

	ans := 0
	avg := ar[len(ar)/2]

	for _, arr := range grid {
		for _, v := range arr {
			ans += abs(v, avg) / x
		}
	}

	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 19：19
func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	p1, p2, n1, n2 := []int{}, []int{}, []int{}, []int{}
	for _, v := range nums1 {
		if v > 0 {
			p1 = append(p1, v)
		}
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		if nums1[i] < 0 {
			n1 = append(n1, -nums1[i])
		}
	}
	for _, v := range nums2 {
		if v > 0 {
			p2 = append(p2, v)
		}
	}
	for i := len(nums2) - 1; i >= 0; i-- {
		if nums2[i] < 0 {
			n2 = append(n2, -nums2[i])
		}
	}

	neg := int64(len(p1)*len(n2) + len(p2)*len(n1))
	pos := int64(len(p1)*len(p2) + len(n1)*len(n2))
	zero := int64((len(nums1) * len(nums2))) - neg - pos

	if k <= neg {
		k = neg - k + 1
		return -int64(sort.Search(1e10, func(t int) bool {
			cnt := int64(0)
			for _, v := range p1 {
				cnt += int64(sort.SearchInts(n2, t/v+1))
			}
			for _, v := range p2 {
				cnt += int64(sort.SearchInts(n1, t/v+1))
			}
			return cnt >= k
		}))
	} else if k > neg+zero {
		k -= neg + zero
		return int64(sort.Search(1e10, func(t int) bool {
			cnt := int64(0)
			for _, v := range p1 {
				cnt += int64(sort.SearchInts(p2, t/v+1))
			}
			for _, v := range n1 {
				cnt += int64(sort.SearchInts(n2, t/v+1))
			}
			return cnt >= k
		}))
	} else {
		return 0
	}
}
