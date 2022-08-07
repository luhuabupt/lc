package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumWeight(6, [][]int{{0, 1, 22}, {0, 5, 11}, {0, 2, 7}, {4, 3, 4}, {5, 3, 48}, {5, 3, 17}, {4, 3, 45}}, 4, 0, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
type h [][2]int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i][1] < h[j][1] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	next := make([][][2]int, n)
	re := make([][][2]int, n)
	for _, v := range edges {
		next[v[0]] = append(next[v[0]], [2]int{v[1], v[2]})
		re[v[1]] = append(re[v[1]], [2]int{v[0], v[2]})
	}
	for k := 0; k < n; k++ {
		if len(next[k]) > 0 {
			sort.Slice(next[k], func(i, j int) bool {
				return next[k][i][1] < next[k][j][1]
			})
		}
		if len(re[k]) > 0 {
			sort.Slice(re[k], func(i, j int) bool {
				return re[k][i][1] < re[k][j][1]
			})
		}
	}

	dp := [3][]int{}
	for i, _ := range dp {
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], -1)
		}
	}

	dijkstra := func(s int, idx int) {
		hp := &h{}
		heap.Push(hp, [2]int{s, 0})
		dp[idx][s] = 0
		for hp.Len() > 0 {
			cur := heap.Pop(hp).([2]int)
			nex := next[cur[0]]
			if idx == 2 {
				nex = re[cur[0]]
			}
			for _, ne := range nex {
				if dp[idx][ne[0]] == -1 || cur[1]+ne[1] < dp[idx][ne[0]] {
					dp[idx][ne[0]] = cur[1] + ne[1]
					heap.Push(hp, [2]int{ne[0], cur[1] + ne[1]})
				}
			}
		}
	}

	dijkstra(src1, 0)
	dijkstra(src2, 1)
	dijkstra(dest, 2)

	//fmt.Println(dp[0])
	//fmt.Println(dp[1])
	//fmt.Println(dp[2])

	ans := -1
	for i := 0; i < n; i++ {
		if dp[0][i] != -1 && dp[1][i] != -1 && dp[2][i] != -1 {
			if ans == -1 || dp[0][i]+dp[1][i]+dp[2][i] < ans {
				ans = dp[0][i] + dp[1][i] + dp[2][i]
			}
		}
	}

	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
