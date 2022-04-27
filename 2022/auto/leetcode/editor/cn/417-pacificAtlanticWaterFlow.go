package main

import "fmt"

func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 1}, {1, 1}, {1, 1}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func pacificAtlantic(g [][]int) [][]int {
	m, n := len(g), len(g[0])
	ans := [][]int{}

	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}

	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(i, j, v int)
	dfs = func(i, j, v int) {
		vis[i][j] |= v
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && g[x][y] >= g[i][j] && vis[x][y]&v == 0 {
				dfs(x, y, v)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, 1)
		dfs(i, n-1, 2)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, 1)
		dfs(m-1, j, 2)
	}

	//fmt.Prinztln(vis)

	for i, ar := range vis {
		for j, v := range ar {
			if v == 3 {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
