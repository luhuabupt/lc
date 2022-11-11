package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func orderOfLargestPlusSign(n int, set [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	g := make([][]int, n)
	for i, _ := range g {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = 1
		}
	}
	for _, v := range set {
		g[v[0]][v[1]] = 0
	}

	cal := func(i, j int) int {
		for do := 0; do < n; do++ {
			for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				x, y := i+do*d[0], j+do*d[1]
				if x >= 0 && x < n && y >= 0 && y < n {
					if g[x][y] == 0 {
						return do
					}
				} else {
					return do
				}
			}
		}
		return 1
	}

	ans := 0
	for i, a := range g {
		for j, v := range a {
			if v == 1 {
				ans = max(ans, cal(i, j))
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
