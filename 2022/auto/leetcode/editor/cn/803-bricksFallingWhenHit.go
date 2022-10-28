package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type uf803 struct {
	fa  []int
	son []int
}

func (t *uf803) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *uf803) union(i, j int) {
	fi, fj := t.find(i), t.find(j)
	if fi == fj {
		return
	}

	if fi > fj {
		fi, fj = fj, fi
	}

	t.son[fi] += t.son[fj]
	t.fa[fj] = fi
}

func hitBricks(g [][]int, h [][]int) []int {
	m := len(g)
	n := len(g[0])
	ans := make([]int, len(h))

	og := make([][]int, m)
	for i, v := range g {
		og[i] = append([]int{}, v...)
	}

	for _, v := range h {
		g[v[0]][v[1]] = 0
	}

	fa := make([]int, m*n)
	son := make([]int, m*n)
	for i := 0; i < n; i++ {
		fa[i] = 0
		son[0] += g[0][i]
	}
	for i, ar := range og {
		if i == 0 {
			continue
		}
		for j, v := range ar {
			fa[i*n+j] = i*n + j
			son[i*n+j] = v
		}
	}
	t := uf803{fa, son}

	var d0 func(i, j int)
	d0 = func(i, j int) {
		for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if g[x][y] == 1 {
					t.union(i*n+j, x*n+y)
				}
			}
		}
	}
	for i, ar := range g {
		for j, v := range ar {
			if v == 1 {
				d0(i, j)
			}
		}
	}

	//fmt.Println(og)
	//fmt.Println(g)
	//fmt.Println(t.fa)
	//fmt.Println(t.son)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		ori := t.son[0]
		if t.find(i*n+j) == 0 {
			t.son[0]++
		}
		for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if g[x][y] == 1 {
					t.union(i*n+j, x*n+y)
				}
			}
		}
		//fmt.Println(i, j, t.son[t.find(i*n+j)], ori, t.son[0])
		if t.son[0] == ori {
			return 0
		}
		return t.son[0] - ori - 1
	}

	for i := len(h) - 1; i >= 0; i-- {
		v := h[i]
		if og[v[0]][v[1]] == 1 {
			ans[i] = dfs(v[0], v[1])
			g[v[0]][v[1]] = 1
		}
	}

	//fmt.Println(t.fa)
	//fmt.Println(t.son)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
