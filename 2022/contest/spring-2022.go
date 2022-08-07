package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("")

}

func giveGem(g []int, o [][]int) int {
	for _, v := range o {
		d := g[v[0]] / 2
		g[v[0]] -= d
		g[v[1]] += d
	}

	mi, ma := 1<<60, -1
	for _, v := range g {
		if v > ma {
			ma = v
		}
		if v < mi {
			mi = v
		}
	}

	return ma - mi
}

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	n := len(cookbooks)

	ans := -1

loop:
	for i := 0; i < 1<<n; i++ {
		use := append([]int{}, materials...)
		a, b := 0, 0
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				for x := 0; x < 5; x++ {
					use[x] -= cookbooks[j][x]
					if use[x] < 0 {
						continue loop
					}
				}
				a += attribute[j][0]
				b += attribute[j][1]
			}
		}
		if b >= limit && a > ans {
			ans = a
		}
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ht [][]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func getNumber(root *TreeNode, ops [][]int) (ans int) {
	a := []int{}

	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		dfs(p.Left)
		a = append(a, p.Val)
		dfs(p.Right)
	}
	dfs(root)

	do := map[int][][]int{} // opIdx, start/end, color
	for i, v := range ops {
		if do[v[1]] == nil {
			do[v[1]] = [][]int{}
		}
		//if do[v[2]] == nil {
		//	do[v[2]] = [][]int{}
		//}

		do[v[1]] = append(do[v[1]], []int{i, 0, v[0]})
		//do[v[2]] = append(do[v[2]], []int{1, v[0], i})
	}

	h := &ht{}
loop:
	for _, v := range a {
		for _, x := range do[v] {
			heap.Push(h, x)
		}
		if h.Len() == 0 {
			continue
		}

		cur := heap.Pop(h).([]int)
		for ops[cur[0]][2] < v {
			if h.Len() == 0 {
				continue loop
			}
			cur = heap.Pop(h).([]int)
		}
		if cur[2] == 1 {
			ans++
		}
		heap.Push(h, cur)
	}

	return ans
}

func defendSpaceCity(time []int, pos []int) (ans int) {
	n := len(time)
	dp := make([][6][]int, n)
	do := make([][6]bool, n)
	mi := make([][6]int, n)

	//for i,v := range time {
	//	do[i] = [2]int{v, pos[i]}
	//}
	//sort.Slice(do, func(i, j int) bool {
	//	return do[i][0] < do[j][0] || do[i][0] == do[j][0] && do[i][1] < do[j][1]
	//})
	for i, v := range time {
		do[pos[i]][v] = true
	}

	for t := 1; t < 6; t++ {
		for j := 0; j < 4; j++ {
			dp[0][t][j] = 1
		}
	}

	for i := 1; i < n; i++ {
		for t := 1; t < 6; t++ {
			dp[i][t][0] = mi[i][t-1] + 2

		}
	}

	return ans
}
