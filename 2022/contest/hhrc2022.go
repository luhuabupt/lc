package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(longestESR([]int{10, 2, 1, 4, 3, 9, 6, 9, 9}))
	//fmt.Println(longestESR([]int{1, 2, 3}))
	fmt.Println(longestESR([]int{9, 9, 6, 0, 6, 6, 9}))
	fmt.Println(longestESR([]int{9, 9, 9, 6, 6}))
	//fmt.Println(longestESR([]int{1, 1, -1, -1, -1, 1}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minSupplyStationNumber(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var dfs func(p *TreeNode) (int, int, int) // 0 1 2
	dfs = func(p *TreeNode) (int, int, int) {
		if p == nil {
			return 100001, 100001, 100001
		}
		if p.Left == nil && p.Right == nil {
			return 100001, 100001, 1
		}

		return min(x, y)
	}

	return min(dfs(root, 1), dfs(root, 0))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lightDistribution(root *TreeNode) []*TreeNode {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	m := map[int]map[int]map[int][]*TreeNode{} // h, cnt, val
	d := map[*TreeNode][]int{}
	var dfs func(p *TreeNode) (int, int)
	dfs = func(p *TreeNode) (int, int) {
		if p == nil {
			return 0, 0
		}

		hl, cl := dfs(p.Left)
		hr, cr := dfs(p.Right)
		h := max(hl, hr)
		c := cl + cr

		if m[h] == nil {
			m[h] = map[int]map[int][]*TreeNode{}
		}
		if m[h][c] == nil {
			m[h][c] = map[int][]*TreeNode{}
		}
		if m[h][c][p.Val] == nil {
			m[h][c][p.Val] = []*TreeNode{}
		}

		d[p] = []int{h, c}
		m[h][c][p.Val] = append(m[h][c][p.Val], p)
		return h + 1, c + 1
	}
	dfs(root)

	ans := []*TreeNode{}

	var chk func(a, b *TreeNode) bool
	chk = func(a, b *TreeNode) bool {
		if a == nil || b == nil {
			return a == nil && b == nil
		}
		if a.Val != b.Val {
			return false
		}
		if d[a][0] != d[b][0] || d[a][1] != d[b][1] {
			return false
		}

		return chk(a.Left, b.Left) && chk(a.Right, b.Right)
	}

	vis := map[*TreeNode]bool{}
	for _, hv := range m {
		for _, cv := range hv {
			for _, pa := range cv {
				for i, p := range pa {
					if vis[p] {
						continue
					}
					for j := i + 1; j < len(pa); j++ {
						if chk(p, pa[j]) {
							if !vis[p] {
								ans = append(ans, p)
							}
							vis[p] = true
							vis[pa[j]] = true
						}
					}
				}
			}
		}
	}

	return ans
}

func longestESR(a []int) int {
	n := len(a)
	ans := -1
	t := 0
	dp := make([]int, n)

	for i, v := range a {
		if v > 8 {
			t++
		} else {
			t--
		}
		if t > 0 {
			ans = i
		}
		dp[i] = t
	}

	st := []int{}
	stv := []int{}
	for i := n - 1; i >= 0; i-- {
		if len(st) == 0 || dp[i] > dp[st[len(st)-1]] {
			st = append(st, i)
			stv = append(stv, dp[i])
		}
	}

	//fmt.Println(dp)
	//fmt.Println(st, stv)

	for i := 1; i < n; i++ {
		j := sort.SearchInts(stv, dp[i-1]+1)
		//fmt.Println(i, j, dp[i-1]+1)
		if j < len(st) && st[j]-i > ans {
			ans = st[j] - i
		}
	}

	return ans + 1
}
