package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kSum([]int{2, 4, -2}, 5))
	fmt.Println(kSum([]int{1, -2, 3, 4, -10, 12}, 16))
}

type h [][]int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func kSum(num []int, k int) int64 {
	sum := 0
	for i, v := range num {
		if v >= 0 {
			sum += v
		} else {
			num[i] = -v
		}
	}
	sort.Ints(num)

	h := &h{}
	heap.Push(h, []int{0, -1})
	x := []int{}
	for i := 0; i < k; i++ { // k 次 pop
		x = heap.Pop(h).([]int)
		if x[1]+1 < len(num) {
			heap.Push(h, []int{x[0] + num[x[1]+1], x[1] + 1})
			if x[1] >= 0 {
				heap.Push(h, []int{x[0] + num[x[1]+1] - num[x[1]], x[1] + 1})
			}
		}
	}
	return int64(sum - x[0])
}

func kSum(num []int, k int) int64 {
	sum := 0
	for i, v := range num {
		if v >= 0 {
			sum += v
		} else {
			num[i] = -v
		}
	}
	sort.Ints(num)

	h := &h{}
	heap.Push(h, 0)
	for i := 0; i < k && i < len(num); i++ {
		v := num[i]
		back := []int{}
		for j := 0; i+j < k && h.Len() > 0; j++ {
			x := heap.Pop(h).(int)
			back = append(back, x)
		}
		for _, vv := range back {
			heap.Push(h, vv)
			heap.Push(h, v+vv)
		}
	}
	x := 0
	for i := 0; i < k; i++ {
		x = heap.Pop(h).(int)
	}
	return int64(sum - x)
}

func minNumberOfHours(e int, ep int, en []int, exp []int) int {
	ans := 0
	for i, v := range en {
		if e <= v {
			ans += v - e + 1
			e = v + 1
		}
		if ep <= exp[i] {
			ans += exp[i] - ep + 1
			ep = exp[i] + 1
		}
		e -= v
		ep += exp[i]
	}
	return ans
}

func largestPalindromic(a string) string {
	c := make([]int, 10)
	for _, v := range a {
		c[v-'0']++
	}
	ans := []byte{}
	for i := 9; i >= 0; i-- {
		for j := 0; j < c[i]/2; j++ {
			ans = append(ans, byte('0'+i))
		}
		c[i] = c[i] % 2
	}
	if ans[0] == '0' {
		ans = []byte{}
	}

	d := len(ans) - 1
	for i := 9; i >= 0; i-- {
		if c[i] > 0 {
			ans = append(ans, byte('0'+i))
			break
		}
	}
	for i := d; i >= 0; i-- {
		ans = append(ans, ans[i])
	}
	return string(ans)
}
