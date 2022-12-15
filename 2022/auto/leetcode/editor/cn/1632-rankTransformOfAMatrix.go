package main

// rank-transform-of-a-matrix 矩阵转换后的秩  2022-10-06 11:34:16
func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func matrixRankTransform(g [][]int) [][]int {
	type p struct {
		i, j int
		in   int
		nx   []int
		eq   []int
	}

	m, n := len(g), len(g[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	h := make([]p, m*n)
	for i, ar := range g {
		for j, v := range ar {
			tp := p{
				i, j, 0, []int{}, []int{},
			}
			for x := 0; x < m; x++ {
				if x == i {
					continue
				}
				if g[x][j] < v {
					tp.in++
				}
				if g[x][j] == v {
					tp.eq = append(tp.eq, x*n+j)
				}
				if g[x][j] > v {
					tp.nx = append(tp.nx, x*n+j)
				}
			}
			for y := 0; y < n; y++ {
				if y == j {
					continue
				}
				if g[i][y] < v {
					tp.in++
				}
				if g[i][y] == v {
					tp.eq = append(tp.eq, i*n+y)
				}
				if g[i][y] > v {
					tp.nx = append(tp.nx, i*n+y)
				}
			}
			h[i*n+j] = tp
		}
	}

	st := []p{}
	for _, v := range h {
		if v.in == 0 {
			st = append(st, v)
		}
	}

	idx := 1
	for len(st) > 0 {
		ns := []p{}
		for _, t := range st {
			ans[t.i][t.j] = idx
			for _, k := range t.nx {
				h[k].in--
				if h[k].in == 0 {
					if len(h[k].eq) > 0 {
						flag := true
						for _, eqv := range h[k].eq {
							if h[eqv].in > 0 {
								flag = false
								break
							}
						}
						if flag {
							for _, eqv := range h[k].eq {
								ns = append(ns, h[eqv])
							}
						}
					} else {
						ns = append(ns, h[k])
					}
				}
			}
		}
		st = ns
		idx++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
