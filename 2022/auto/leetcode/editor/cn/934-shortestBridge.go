package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestBridge(g [][]int) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		g[i][j] = 2
		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < len(g) && y >= 0 && y < len(g[0]) {
				if g[x][y] == 1 {
					dfs(x, y)
				}
			}
		}
	}

	flag := false
	for i, a := range g {
		if flag {
			break
		}
		for j, v := range a {
			if v == 1 {
				dfs(i, j)
				flag = true
				break
			}
		}
	}

	st := [][2]int{}
	for i, a := range g {
		for j, v := range a {
			if v == 1 {
				g[i][j] = 3
				st = append(st, [2]int{i, j})
			}
		}
	}

	deep := 0
	for len(st) > 0 {
		ns := [][2]int{}
		for _, t := range st {
			i, j := t[0], t[1]
			for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < len(g) && y >= 0 && y < len(g[0]) {
					if g[x][y] == 2 {
						return deep
					}
					if g[x][y] == 0 {
						g[x][y] = 3
						ns = append(ns, [2]int{x, y})
					}
				}
			}
		}
		st = ns
		deep++
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
