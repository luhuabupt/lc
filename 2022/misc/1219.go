package main

func getMaximumGold(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	g, vis := make([][]int, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		g[i] = make([]int, n)
		vis[i] = make([]bool, n)
	}

	type pair struct {
		x, y int
	}
	dir := []pair{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) (mx int) {
		vis[i][j] = true

		for _, d := range dir {
			x, y := i+d.x, j+d.y
			if x >= 0 && x < m && y >= 0 && y < n {
				if grid[x][y] > 0 && !vis[x][y] {
					re := dfs(x, y)
					if re > mx {
						mx = re
					}
				}
			}
		}

		vis[i][j] = false
		return mx + grid[i][j]
	}

	for i, arr := range grid {
		for j, v := range arr {
			if v == 0 {
				continue
			}

			for x := 0; x < m; x++ {
				for y := 0; y < n; y++ {
					vis[x][y] = false
				}
			}
			x := dfs(i, j)
			if x > ans {
				ans = x
			}
		}
	}

	return ans
}
