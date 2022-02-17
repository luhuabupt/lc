package main

// number-of-enclaves 飞地的数量  2022-02-12 20:16:41
func main() {
	//fmt.Println(numberOfEnclaves())
}

//leetcode submit region begin(Prohibit modification and deletion)
func numEnclaves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == 0 {
			return
		} else {
			grid[i][j] = 0
		}
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if grid[x][y] == 1 {
					dfs(x, y)
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 1; i < m-1; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for _, arr := range grid {
		for _, v := range arr {
			if v == 1 {
				ans++
			}
		}
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)
