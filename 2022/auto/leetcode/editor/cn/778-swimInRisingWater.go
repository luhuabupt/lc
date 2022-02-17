package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func swimInWater(grid [][]int) int {
	n := len(grid)
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	vis := [50][50]bool{}
	var dfs func(i, j, ban int) bool
	dfs = func(i, j, ban int) bool {
		if grid[i][j] > ban {
			return false
		}
		if i == n-1 && j == n-1 {
			return true
		}
		vis[i][j] = true
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < n {
				if !vis[x][y] {
					if dfs(x, y, ban) {
						return true
					}
				}
			}
		}

		return false
	}

	return sort.Search(n*n, func(v int) bool {
		vis = [50][50]bool{}
		return dfs(0, 0, v)
	})
}

//leetcode submit region end(Prohibit modification and deletion)
