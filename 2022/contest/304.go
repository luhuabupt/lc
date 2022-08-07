package main

import "fmt"

func main() {
	fmt.Println(longestCycle([]int{3, 4, 0, 2, -1, 2}))
	fmt.Println(longestCycle([]int{3, 3, 4, 2, 3}))
}

func longestCycle(g []int) int {
	n := len(g)
	ans := -1
	del := make([]bool, n)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	in := make([]int, n)
	for _, v := range g {
		if v >= 0 {
			in[v]++
		}
	}

	st := []int{}
	for i, v := range in {
		if v == 0 {
			st = append(st, i)
			del[i] = true
		}
	}

	for idx := 0; idx < len(st) && len(st) > 0; idx++ {
		h := st[idx]
		nx := g[h]
		if nx != -1 {
			in[nx]--
			if in[nx] == 0 {
				st = append(st, nx)
				del[nx] = true
			}
		}
	}

	fmt.Println(in)
	fmt.Println(del)

	vis := make([]bool, n)
	var dfs func(i, d int) int
	dfs = func(i, d int) int {
		if vis[i] {
			return d
		}
		vis[i] = true
		if g[i] == -1 {
			return -1
		}
		return dfs(g[i], d+1)
	}

	for i, v := range del {
		if !v {
			ans = max(ans, dfs(i, 0))
		}
	}
	return ans
}

func closestMeetingNode(g []int, a int, b int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(g)
	dp := make([][]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{-1, -1}
	}

	var dfs func(i, d, tt int)
	dfs = func(i, d, tt int) {
		if vis[i] {
			return
		}
		vis[i] = true

		dp[i][tt] = d

		if g[i] == -1 {
			return
		}
		dfs(g[i], d+1, tt)
	}

	dfs(a, 0, 0)

	vis = make([]bool, n)
	dfs(b, 0, 1)

	mi := 1 << 60
	ans := -1
	for i := 0; i < n; i++ {
		if dp[i][0] == -1 || dp[i][1] == -1 {
			continue
		}
		if max(dp[i][0], dp[i][1]) < mi {
			mi = max(dp[i][0], dp[i][1])
			ans = i
		}
	}
	return ans
}
