package local

import (
	"container/heap"
	"fmt"
	"sort"
)

func getDescentPeriods(prices []int) int64 {
	ans := int64(0)
	n := len(prices)
	tmp := 1
	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1] {
			tmp++
		} else {
			ans += int64((tmp - 1) * tmp / 2)
			tmp = 1
		}
	}
	return ans
}

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	ans := 0

	for l := 0; l < k; l++ {
		x := []int{}
		cnt := 0
		for i := 0; i < n; i += k {
			x = append(x, arr[i])
			cnt++
		}
		ans += cnt - findMaxUp(x)
	}

	return ans
}

// 最长上升子序列
func findMaxUp(arr []int) int {
	dp := []int{}
	for _, v := range arr {
		if p := sort.SearchInts(dp, v); p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return len(dp)
}

type ap struct {
	n, d int
}
type h []ap

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i].d > h[j].d }
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(ap)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func eatenApples(apples []int, days []int) int {
	hp := &h{}
	ans := 0
	for i := 0; i < 40000; i++ {
		if i < len(apples) && apples[i] > 0 {
			heap.Push(hp, ap{apples[i], i + days[i]})
		}

		x := ap{}
		for hp.Len() > 0 {
			x = heap.Pop(hp).(ap)
			if x.d <= i && x.n > 0 {
				break
			}
		}

		if x.n > 0 {
			x.n--
			ans++
			if x.n > 0 {
				heap.Push(hp, x)
			}
		}

	}
	return ans
}

func cutOffTree(forest [][]int) int {
	type p struct {
		i, j, v int
	}
	fmt.Println(len(forest), len(forest[0]))

	all := []p{{0, 0, -1}}
	for i, arr := range forest {
		for j, v := range arr {
			if v <= 1 {
				continue
			}
			all = append(all, p{i, j, v})
		}
	}
	if len(all) == 0 {
		return 0
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].v < all[j].v
	})
	fmt.Println(all)

	dir := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	find := func(pi, pj, i, j int) int {
		if pi == i && pj == j {
			return 0
		}
		sub := [][]int{{pi, pj, 0}}
		vis := [50][50]bool{}
		vis[pi][pj] = true
		fmt.Println(pi, pj, i, j)
		for len(sub) > 0 {
			cur := sub[0]
			sub = sub[1:]
			fmt.Println(sub, cur, vis)
			for _, d := range dir {
				ni, nj, nd := cur[0]+d[0], cur[1]+d[1], cur[2]+1
				if ni == i && nj == j {
					return nd
				}
				if !vis[ni][nj] && ni >= 0 && ni < len(forest) && nj < len(forest[0]) && nj >= 0 && forest[ni][nj] > 0 {
					vis[ni][nj] = true
					sub = append(sub, []int{ni, nj, nd})
				}
			}
		}
		return -1
	}

	ans := 0
	for i := 1; i < len(all); i++ {
		x := find(all[i-1].i, all[i-1].j, all[i].i, all[i].j)
		if x == -1 {
			return -1
		}
		ans += x
	}

	return ans
}
