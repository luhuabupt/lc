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

//func finddegplicateSubtrees(root *TreeNode) []*TreeNode {
//	all := map[string][]*TreeNode{} // v, l, r
//	var dfs func(p *TreeNode) int
//	dfs = func(p *TreeNode) int {
//		if p == nil {
//			return 0
//		}
//
//		l := dfs(p.Left)
//		r := dfs(p.Right)
//
//		k := fmt.Sprint("%d_%d_%d", p.Val, l, r)
//		if all[k] == nil {
//			all[k] = []*TreeNode{}
//		}
//
//		all[k] = append(all[k], p)
//
//		if l > r {
//			return l + 1
//		}
//		return r + 1
//	}
//
//	var check func(p1, p2 *TreeNode) bool
//	check = func(p1, p2 *TreeNode) bool {
//		if p1 == nil && p2 == nil {
//			return true
//		}
//		if p1 == nil || p2 == nil {
//			return false
//		}
//		if p1.Val != p2.Val {
//			return false
//		}
//		return check(p1.Left, p2.Left) && check(p1.Right, p2.Right)
//	}
//
//	dfs(root)
//
//	ans := []*TreeNode{}
//	for _, arr := range all {
//		for i := 0; i < len(arr); i++ {
//			flag := false
//			if arr[i] == nil {
//				continue
//			}
//			for j := i + 1; j < len(arr); j++ {
//				if arr[i] == nil || arr[j] == nil {
//					continue
//				}
//				if check(arr[i], arr[j]) {
//					if !flag {
//						ans = append(ans, arr[i])
//						flag = true
//					}
//					arr[j] = nil
//				}
//			}
//		}
//	}
//
//	return ans
//}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	oc := grid[row][col]

	graph := make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		graph[i][j] = 1
		for _, x := range dirs {
			ni, nj := i+x[0], j+x[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n {
				if graph[ni][nj] == 0 {
					if grid[ni][nj] == oc {
						dfs(ni, nj)
					} else {
						graph[ni][nj] = -1
					}
				}
				if grid[ni][nj] != oc {
					graph[i][j] = 2
				}
			} else {
				graph[i][j] = 2
			}
		}
	}

	dfs(row, col)

	for i, arr := range graph {
		for j, v := range arr {
			if v == 2 {
				grid[i][j] = color
			}
		}
	}

	return grid
}

func validArrangement(pairs [][]int) [][]int {
	deg := map[int][]int{}
	next := map[int][]int{}
	for _, v := range pairs {
		if deg[v[0]] == nil {
			deg[v[0]] = []int{0, 0}
		}
		if deg[v[1]] == nil {
			deg[v[1]] = []int{0, 0}
		}
		deg[v[0]][0]++
		deg[v[1]][1]++

		if next[v[0]] == nil {
			next[v[0]] = []int{}
		}
		next[v[0]] = append(next[v[0]], v[1])
	}

	s := pairs[0][0]
	for k, v := range deg {
		if v[0] == v[1]+1 {
			s = k
			break
		}
	}

	points := []int{}
	var dfs func(s int)
	dfs = func(s int) {
		if _, ok := next[s]; ok {
			for len(next[s]) > 0 {
				top := next[s][len(next[s])-1]
				next[s] = next[s][:len(next[s])-1]
				dfs(top)
			}
		}

		points = append(points, s)
	}
	dfs(s)

	ans := [][]int{}
	for i := len(points) - 2; i >= 0; i-- {
		ans = append(ans, []int{points[i+1], points[i]})
	}

	return ans
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	i, j, ans := 0, 0, 0
	for ; i < len(houses); i++ {
		for ; j < len(heaters)-1 && heaters[j] < houses[i]; j++ {
		}
		l := d(heaters[j], houses[i])
		if j > 0 && d(houses[i], heaters[j-1]) < l {
			l = d(houses[i], heaters[j-1])
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}

func d(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func minWindow(s string, t string) string {
	cs, ct := make([]int, 52), make([]int, 52)

	gk := func(a uint8) int {
		if a >= 'a' && a <= 'z' {
			return int(a - 'a')
		}
		return int(a - 'A' + 26)
	}
	chk := func() bool {
		flag := true
		for i := 0; i < 52; i++ {
			if cs[i] < ct[i] {
				flag = false
				break
			}
		}
		return flag
	}

	for i := 0; i < len(t); i++ {
		ct[gk(t[i])]++
	}

	l0, r0 := 0, len(s)
	for l, r := 0, 0; r < len(s); r++ {
		cs[gk(s[r])]++
		for ; l < len(s) && chk(); l++ {
			if r-l < r0-l0 {
				r0, l0 = r, l
			}
			cs[gk(s[l])]--
		}
	}

	if r0 == len(s) {
		return ""
	}
	return s[l0 : r0+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

// 632
func smallestRange(nums [][]int) []int {
	type p struct {
		k, val int
	}

	arr := []p{}
	for k, v := range nums {
		for _, x := range v {
			arr = append(arr, p{k, x})
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val || arr[i].val == arr[j].val && arr[i].k == arr[j].k
	})

	ans := []int{}
	return ans
}

func repeatedStringMatch(a string, b string) int {
	arr, an, bn := []byte(a), len(a), len(b) // arr: a重复后的字符数组

	// 判断从start开始是否相等
	eq := func(start int) bool {
		for i := 0; i < len(b); i++ {
			if arr[i+start] != b[i] {
				return false
			}
		}
		return true
	}

	ans := 1
	for i := 0; i < an; i++ { // 最多an次, 再后面就和之前重复了
		for len(arr)-i < bn { // arr不够长, 则重复a
			arr = append(arr, []byte(a)...)
			ans++
		}
		if eq(i) {
			return ans
		}
	}

	return -1
}

func gcdSort(nums []int) bool {
	n := len(nums)
	pa := make([]int, n)

	var find func(i int) int
	find = func(i int) int {
		if i == pa[i] {
			return i
		}
		return find(pa[i])
	}

	union := func(i, j int) {
		if pa[i] == pa[j] {
			return
		}

	}
}
