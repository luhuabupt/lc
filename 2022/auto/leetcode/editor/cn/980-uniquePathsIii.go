package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsIII(g [][]int) int {
	m, n := len(g), len(g[0])

	t := 0
	a, b := 0, 0
	for i, ar := range g {
		for j, v := range ar {
			if v == 0 {
				t++
			}
			if v == 1 {
				a, b = i, j
			}
		}
	}

	ans := 0
	var dfs func(i, j, c int)
	dfs = func(i, j, c int) {
		for _, dv := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			x, y := i+dv[0], j+dv[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if g[x][y] == 2 && c == t {
				ans++
				return
			}
			if g[x][y] != 0 {
				continue
			}

			g[x][y] = -1
			dfs(x, y, c+1)
			g[x][y] = 0
		}
	}

	dfs(a, b, 0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
