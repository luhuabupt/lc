package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func answerQueries(a []int, q []int) []int {
	sort.Ints(a)
	b := []int{}
	s := 0
	for _, v := range a {
		s += v
		b = append(b, v)
	}
	ans := make([]int, len(q))
	for i, v := range q {
		for j, x := range b {
			if x > v {
				ans[i] = j
				break
			}
		}
		if v >= b[len(b)-1] {
			ans[i] = len(b)
		}
	}
	return ans
}

func removeStars(s string) string {
	n := len(s)
	ans := []byte{}
	t := 0
	for i := n - 1; i >= 0; i-- {
		v := s[i]
		if v == '*' {
			t++
		} else {
			if t > 0 {
				t--
			} else {
				ans = append(ans, v)
			}
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}

func garbageCollection(g []string, t []int) int {
	ans := 0
	a, b, c := 0, 0, 0
	for i, s := range g {
		for _, v := range s {
			if v == 'M' {
				a = i
			} else if v == 'P' {
				b = i
			} else {
				c = i
			}
		}
		ans += len(s)
	}

	cal := func(x int) int {
		r := 0
		for i, v := range t {
			if x == i {
				break
			}
			r += v
		}
		return r
	}

	ans += cal(a)
	ans += cal(b)
	ans += cal(c)
	return ans
}

func buildMatrix(k int, r [][]int, c [][]int) [][]int {
	a := cal(k, r)
	b := cal(k, c)

	if len(a) == 0 || len(b) == 0 {
		return [][]int{}
	}

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = make([]int, k)
	}

	for v := 1; v <= k; v++ {
		ans[a[v]][b[v]] = v
	}

	return ans
}

func cal(k int, r [][]int) []int {
	a := []int{}
	for i := 1; i <= k; i++ {
		a = append(a, i)
	}
	rd := make([]map[int]bool, k+1)
	for _, v := range r {
		if rd[v[0]] == nil {
			rd[v[0]] = map[int]bool{}
		}
		rd[v[0]][v[1]] = true
	}

	dp := make([][]bool, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]bool, k+1)
	}

	var dfs func(i, f int) bool
	dfs = func(i, f int) bool {
		if dp[i][f] {
			return true
		}
		for x, _ := range rd[i] {
			if x == f {
				dp[i][f] = true
				return true
			}
			as := dfs(x, f)
			if as {
				dp[i][f] = true
				return true
			}
		}
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		return dfs(x, y)
	})
	d := make([]int, k+1)
	for i, v := range a {
		d[v] = i
	}

	for _, v := range r {
		if d[v[0]] > d[v[1]] {
			return []int{}
		}
	}

	return d
}
