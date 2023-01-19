package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(" ")
}

func isCircularSentence(s string) bool {
	a := strings.Split(s, " ")
	for i := 1; i < len(s); i++ {
		if a[i-1][len(a[i])-1] != a[i][0] {
			return false
		}
	}
	return a[0][0] == a[len(a)-1][len(a[len(a)-1])-1]
}

func magnificentSets(n int, edges [][]int) int {
	for i, v := range edges {
		edges[i] = []int{v[0] - 1, v[1] - 1}
	}
	nx := make([][]int, n)
	for _, v := range edges {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}
	vis := make([]bool, n)
	var dfs func(i int) []int
	dfs = func(i int) []int {
		r := []int{i}
		vis[i] = true
		for _, x := range nx[i] {
			if !vis[x] {
				r = append(r, dfs(x)...)
			}
		}
		return r
	}

	ans := 0
	for i, v := range vis {
		if !v {
			ar := dfs(i)
			x := cal(ar, nx)
			if x == -1 {
				return -1
			}
			ans += x
		}
	}
	return ans
}

func cal(ar []int, nx [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(ar)
	ga := make([]int, n)
	var dfs func(i, gi int) int
	dfs = func(i, gi int) int {
		r := 0
		td := []int{}
		for _, x := range nx[ar[i]] {
			if ga[x] != -1 {
				diff := x - ga[x]
				if diff != 1 && diff != -1 {
					return -1
				}
			} else {
				ga[x] = gi + 1
				td = append(td, x)
			}
		}
		for _, x := range td {
			rv := dfs(x, gi+1)
			if rv == -1 {
				return -1
			}
			r = max(r, rv)
		}
		return r + 1
	}

	ma := -1
	for _, v := range ar {
		ga = make([]int, n)
		for i, _ := range ga {
			ga[i] = -1
		}
		ma = max(ma, dfs(v, 0))
	}
	return ma
}

func minScore(n int, r [][]int) int {
	fa := make([]int, n+1)
	mi := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i
		mi[i] = 1 << 60
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] == i {
			return i
		}
		return find(i)
	}

	for _, v := range r {
		a, b := find(v[0]), find(v[1])
		if a == b {
			mi[a] = min(mi[a], v[2])
		} else {
			fa[b] = a
			mi[a] = min(mi[a], min(mi[b], v[2]))
		}
	}
	return mi[find(1)]

	//nx := make([][][]int, n)
	//for _, v := range r {
	//	nx[v[0]] = append(nx[v[0]], []int{v[1], v[2]})
	//	nx[v[1]] = append(nx[v[1]], []int{v[0], v[2]})
	//}
	//
	//ans := 1 << 60
	//vis := make([]bool, n)
	//var dfs func(i int)
	//dfs = func(i int) {
	//	vis[i] = true
	//	for _, v := range nx[i] {
	//		if !vis[v[0]] {
	//			if v[1] < ans {
	//				ans = v[1]
	//			}
	//			dfs(v[0])
	//		}
	//	}
	//}

	//dfs(1)
	return ans
}
