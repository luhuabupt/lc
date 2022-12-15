package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func maxPalindromes(s string, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	dp := make([][]bool, n)
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			if l == 0 {
				dp[i] = make([]bool, n)
				dp[i][l] = true
			}
			if l == 1 {
				dp[i][l] = s[i] == s[i+1]
			}
			dp[i][l] = dp[i+1][i+l-1] && s[i] == s[i+l]
		}
	}

	ma := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := 0
		for r, x := range dp[i] {
			if x {
				t = max(t, ma[r+1])
			}
		}
		ma[i] = t
	}
	return ma[0]
}

func minimumOperations(root *TreeNode) int {
	st := []*TreeNode{root}
	ans := 0

	cal := func(a []int) int {
		//b := [][]int{}
		p := map[int]int{}
		c := append([]int{}, a...)
		for i, v := range a {
			p[v] = i
		}
		sort.Ints(c)
		//for i, v := range a {
		//	b = append(b, []int{v, i})
		//}
		//sort.Slice(b, func(i, j int) bool {
		//	return b[i][0] < b[j][0]
		//})
		r := 0
		for i, v := range a {
			if v != c[i] {
				a[i], a[p[c[i]]], p[v], p[c[i]] = a[p[c[i]]], a[i], p[c[i]], p[v]
				r++
			}
		}
		return r
	}

	for len(st) > 0 {
		ns := []*TreeNode{}
		a := []int{}
		for _, p := range st {
			a = append(a, p.Val)
			if p.Left != nil {
				ns = append(ns, p.Left)
			}
			if p.Right != nil {
				ns = append(ns, p.Right)
			}
		}
		ans += cal(a)
		st = ns
	}

	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeQueries(r *TreeNode, q []int) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := int(1e5 + 2)
	level := make([][2][2]int, n) // first, second : sonDeep, val
	deepArr := make([][]int, n)
	maxDeep := 0

	var cal3 func(a, b, v [2]int) [2][2]int
	cal3 = func(a, b, v [2]int) [2][2]int {
		if v[0] >= a[0] {
			a, b = v, a
		} else if v[0] > b[0] {
			b = v
		}
		return [2][2]int{a, b}
	}

	var dfs func(p *TreeNode, d int) int
	dfs = func(p *TreeNode, d int) int {
		if p == nil {
			return 0
		}
		if d > maxDeep {
			maxDeep = d
		}
		deepArr[d] = append(deepArr[d], p.Val)
		x := max(dfs(p.Left, d+1), dfs(p.Right, d+1))
		level[d] = cal3(level[d][0], level[d][1], [2]int{x, p.Val})
		return x
	}
	dfs(r, 0)

	dp := make([]int, n)
	for i := 0; i < maxDeep; i++ {
		a, b := level[i][0], level[i][1]
		for _, x := range deepArr[i] {
			dv := i + a[0]
			if x == a[1] {
				dv = i + b[0]
			}
			dp[x] = dv
		}
	}

	ans := []int{}
	for _, v := range q {
		ans = append(ans, dp[v])
	}
	return ans
}

func makeIntegerBeautiful(ni int64, target int) int64 {
	n := int(ni)
	a := []int{}
	s := 0
	for x := n; x > 0; x /= 10 {
		a = append(a, n%10)
		s += n % 10
	}

	cal := func() int {
		t := 0
		for i := 0; i < len(a); i++ {
			a[i] += t
			t = a[i] / 10
			a[i] = a[i] % 10
		}
		if t > 0 {
			a = append(a, 1)
		}
		ans := 0
		for _, v := range a {
			ans += v
		}
		return ans
	}
	ans := 0
	for i := 0; ; i++ {
		ans += 10 - a[i]
		a[0] = 0
		a[1]++
		if cal() <= target {
			break
		}
	}
	return int64(ans)
}

func mostPopularCreator(ct []string, ids []string, views []int) [][]string {
	c := map[string]int{}
	ci := map[string][]int{} // views ids
	ma := 0
	for i, v := range ct {
		c[v] += views[i]
		if c[v] > ma {
			ma = c[v]
		}
		if ci[v][0] < views[i] || ci[v][0] == views[i] && ids[i] < ids[ci[v][1]] {
			ci[v] = []int{views[i], i}
		}
	}

	ans := [][]string{}
	for i, v := range c {
		if v == ma {
			ans = append(ans, []string{i, ids[ci[i][1]]})
		}
	}

	return ans
}
