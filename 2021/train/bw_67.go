package main

func goodDaysToRobBank(security []int, time int) []int {
	ans := []int{}
	n := len(security)

	down, up := make([]int, n), make([]int, n)

	down[0] = 0
	up[n-1] = 0

	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			down[i] = down[i-1] + 1
		} else {
			down[i+1] = 0
		}
	}
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			up[i] = up[i+1] + 1
		} else {
			up[i+1] = 0
		}
	}

	for i := 0; i < n; i++ {
		if down[i] >= time && up[i] >= time {
			ans = append(ans, i)
		}
	}

	return ans
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	sub := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if boom(bombs[i][0], bombs[i][1], bombs[i][2], bombs[j][0], bombs[j][0]) {
				sub[i] = append(sub[i], j)
			}
		}
	}

	var vis []bool

	var dfs func(i int) int
	dfs = func(i int) int {
		res := 1
		vis[i] = true
		for _, v := range sub[i] {
			if !vis[v] {
				res += dfs(v)
				vis[v] = true
			}
		}
		return res
	}

	ans := 0
	for i := 0; i < n; i++ {
		vis = make([]bool, n)
		x := dfs(i)
		if x > ans {
			ans = x
		}
	}
	return ans
}

func boom(x0, y0, x, y, r int) bool {
	if (x0-x)*(x0-x)+(y0-y)*(y0-y) <= r*r {
		return true
	}
	return false
}
