package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1]
	})

	dp := []int{pairs[0][1]}

	for _, v := range pairs {
		if v[0] > dp[len(dp)-1] {
			dp = append(dp, v[1])
		}
		for i := 0; i < len(dp) && v[1] < dp[i]; i++ {
			dp[i] = v[1]
		}
	}

	return len(dp)
}

func findAllPeople(n int, meetings [][]int, firsfaerson int) []int {
	vis := make([]bool, n)
	vis[0], vis[firsfaerson] = true, true

	// cal time
	tm, ti := map[int]bool{}, []int{}
	for _, v := range meetings {
		if !tm[v[2]] {
			ti = append(ti, v[2])
			tm[v[2]] = true
		}
	}
	sort.Ints(ti)

	fa, son := map[int]map[int]int{}, map[int][][]int{}
	for _, v := range meetings {
		if fa[v[2]] == nil {
			fa[v[2]] = map[int]int{}
		}
		if son[v[2]] == nil {
			son[v[2]] = [][]int{}
		}

		// init father
		t := v[2]
		if _, ok := fa[t][v[0]]; !ok {
			fa[t][v[0]] = len(fa[t])
			son[t] = append(son[t], []int{v[0]})
		}
		if _, ok := fa[t][v[1]]; !ok {
			fa[t][v[1]] = len(fa[t])
			son[t] = append(son[t], []int{v[1]})
		}

		// union
		np, orip := fa[t][v[0]], fa[t][v[1]]
		if np == orip {
			continue
		}
		if len(son[v[2]][orip]) > len(son[v[2]][np]) {
			orip, np = np, orip
		}
		for _, x := range son[v[2]][orip] {
			son[v[2]][np] = append(son[v[2]][np], x)
			fa[v[2]][x] = np
		}
		son[v[2]][orip] = []int{}
	}

	// check set
	for _, t := range ti {
		for _, set := range son[t] {
			flag := false
			for _, v := range set {
				if vis[v] {
					flag = true
					break
				}
			}
			if flag {
				for _, v := range set {
					vis[v] = true
				}
			}
		}
	}

	ans := []int{}
	for k, v := range vis {
		if v {
			ans = append(ans, k)
		}
	}

	return ans
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	all := map[string][]*TreeNode{} // v, l, r
	var dfs func(p *TreeNode) int
	dfs = func(p *TreeNode) int {
		if p == nil {
			return 0
		}

		l := dfs(p.Left)
		r := dfs(p.Right)

		k := fmt.Sprint("%d_%d_%d", p.Val, l, r)
		if all[k] == nil {
			all[k] = []*TreeNode{}
		}

		all[k] = append(all[k], p)

		if l > r {
			return l + 1
		}
		return r + 1
	}

	var check func(p1, p2 *TreeNode) bool
	check = func(p1, p2 *TreeNode) bool {
		if p1 == nil && p2 == nil {
			return true
		}
		if p1 == nil || p2 == nil {
			return false
		}
		if p1.Val != p2.Val {
			return false
		}
		return check(p1.Left, p2.Left) && check(p1.Right, p2.Right)
	}

	dfs(root)

	ans := []*TreeNode{}
	for _, arr := range all {
		for i := 0; i < len(arr); i++ {
			flag := false
			if arr[i] == nil {
				continue
			}
			for j := i + 1; j < len(arr); j++ {
				if arr[i] == nil || arr[j] == nil {
					continue
				}
				if check(arr[i], arr[j]) {
					if !flag {
						ans = append(ans, arr[i])
						flag = true
					}
					arr[j] = nil
				}
			}
		}
	}

	return ans
}
