package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(maximumRobots([]int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20}, []int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20}, 937))
	fmt.Println(maximumRobots([]int{11, 12, 19}, []int{10, 8, 7}, 19))
}

func maximumRobots(c []int, r []int, bu64 int64) int {
	bu := int(bu64)
	n := len(c)
	h := &hl{}
	sum := 0
	s := map[int]int{}
	s[-1] = 0
	for i, v := range r {
		sum += v
		s[i] = sum
	}
	ans := 0
	ma := []int{}
	for i, j := 0, 0; j < n; j++ {
		heap.Push(h, []int{c[j], j})
		//ma := heap.Pop(h).([]int)
		//heap.Push(h, ma)
		for ; i <= j; i++ {
			for h.Len() > 0 {
				ma = heap.Pop(h).([]int)
				if ma[1] >= i {
					break
				}
			}
			heap.Push(h, ma)
			if (j-i+1)*(s[j]-s[i-1])+ma[0] <= bu {
				if j-i+1 > ans {
					ans = j - i + 1
				}
				break
			}
		}
	}

	return ans
}

type hl []int

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i] > h[j] } // 大顶堆
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func maximumRobots_(c []int, r []int, bu64 int64) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	bu := int(bu64)
	n := len(c)
	h := &hl{}
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = []int{c[i], r[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0] || a[i][0] == a[j][0] && a[i][1] < a[j][1]
	})
	s := 0
	for _, v := range a {
		if (s+v[1])*(h.Len()+1)+v[0] <= bu {
			heap.Push(h, v[1])
			s += v[1]
		} else if h.Len() > 0 {
			x := heap.Pop(h).(int)
			heap.Push(h, min(v[1], x))
			s += min(v[1], x) - x
		}
	}
	return h.Len()
}

func maximumRows(g [][]int, c int) int {
	_, n := len(g), len(g[0])
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cal := func(do int) int {
		r := 0
		for _, ar := range g {
			f := true
			for j, v := range ar {
				if v == 1 {
					if do>>j != 1 {
						f = false
						break
					}
				}
			}
			if f {
				r++
			}
		}
		fmt.Println(do, r)
		return r
	}
	for i := 0; i < 1<<n; i++ {
		t := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				t++
			}
		}
		if t == c {
			ans = max(ans, cal(i))
		}
	}
	return ans
}

func isStrictlyPalindromic(n int) bool {
	for l := 2; l < n; l++ {
		a := []int{}
		for x := n; x > 0; x /= l {
			a = append(a, x%l)
		}
		for i := 0; i < len(a)/2; i++ {
			if a[i] != a[len(a)-1-i] {
				return false
			}
		}
	}
	return true
}
