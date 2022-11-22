package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(totalCost([]int{18, 64, 12, 21, 21, 78, 36, 58, 88, 58, 99, 26, 92, 91, 53, 10, 24, 25, 20, 92, 73, 63, 51, 65, 87, 6, 17, 32, 14, 42, 46, 65, 43, 9, 75}, 13, 23))
	fmt.Println(minimumTotalDistance([]int{0, 4, 6}, [][]int{}))
}

func minimumTotalDistance(a []int, c [][]int) int64 {
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	sort.Ints(a)
	b := []int{}
	for _, v := range c {
		for i := 0; i < v[1]; i++ {
			b = append(b, v[0])
		}
	}
	sort.Ints(b)
	n, m := len(a), len(b)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == -1 {
			return 0
		}
		if i < 0 || j < 0 {
			return 2e15
		}
		if i > j {
			return 2e15
		}
		if i == 0 && j == 0 {
			return abs(a[0] - b[0])
		}
		if dp[i][j] > 0 {
			return dp[i][j]
		}

		dp[i][j] = min(abs(a[i]-b[j])+dfs(i-1, j-1), dfs(i, j-1))
		return dp[i][j]
	}

	return int64(dfs(n-1, m-1))
}

type ht [][]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][2] < h[j][2] } // 小顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func totalCost(a []int, k int, cd int) int64 {
	h := &ht{}
	n := len(a)
	l := cd - 1
	r := n - cd

	for i := 0; i < cd; i++ {
		heap.Push(h, []int{a[i], 0, i})
	}
	for i := 0; i < cd && n-1-i > l; i++ {
		heap.Push(h, []int{a[n-1-i], 1, n - 1 - i})
	}
	fmt.Println(n, l, r)

	ans := 0
	for do := 0; do < k; do++ {
		x := heap.Pop(h).([]int)
		ans += x[0]
		if x[1] == 0 {
			l++
			if l < r {
				heap.Push(h, []int{a[l], 0, l})
			}
		} else {
			r--
			if l < r {
				heap.Push(h, []int{a[r], 1, r})
			}
		}
	}

	return int64(ans)
}
