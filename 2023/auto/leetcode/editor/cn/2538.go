package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxOutput(n int, g [][]int, p []int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	nx := make([][]int, n)
	for _, v := range g {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}

	ans := 0
	vis := make([]bool, n)
	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		vis[i] = true

		wl, nl := 0, 0
		for _, v := range nx[i] {
			if vis[v] {
				continue
			}
			x, y := dfs(v)
			ans = max(ans, y+p[i]+wl)
			ans = max(ans, x+p[i]+nl)
			wl = max(wl, x)
			nl = max(nl, y)
		}

		if nl == 0 {
			nl -= p[i]
		}

		return wl + p[i], nl + p[i]
	}

	dfs(0)

	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
