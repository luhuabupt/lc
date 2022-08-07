package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(maximumMinutes([][]int{{0, 2, 0, 0, 0, 0, 0}, {0, 0, 0, 2, 2, 1, 0}, {0, 2, 0, 0, 1, 2, 0}, {0, 0, 2, 2, 2, 0, 2}, {0, 0, 0, 0, 0, 0, 0}}))
	//fmt.Println(maximumMinutes([][]int{{0, 0, 0}, {2, 2, 0}, {1, 2, 0}}))
	//fmt.Println(maximumMinutes([][]int{{0, 2, 0, 0, 1}, {0, 2, 0, 2, 2}, {0, 2, 0, 0, 0}, {0, 0, 2, 2, 0}, {0, 0, 0, 0, 0}}))
}

func maximumMinutes(g [][]int) int {
	dir := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	not := 1000000000

	m, n := len(g), len(g[0])
	f := make([][]int, m)
	sf := [][]int{}

	for i, a := range g {
		f[i] = make([]int, n)
		for j, v := range a {
			if v == 1 {
				sf = append(sf, []int{i, j})
				f[i][j] = 0
			} else {
				f[i][j] = -1
			}
		}
	}

	t := 0
	for len(sf) > 0 {
		t++
		ns := [][]int{}
		for _, sfv := range sf {
			for _, d := range dir {
				x, y := sfv[0]+d[0], sfv[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if f[x][y] == -1 && g[x][y] == 0 {
						f[x][y] = t
						ns = append(ns, []int{x, y})
					}
				}
			}
		}
		sf = ns
	}

	ans := sort.Search(m*n+2, func(t int) bool {
		dp := make([][]int, m)
		for i := 0; i < m; i++ {
			dp[i] = make([]int, n)
		}
		die := true
		dp[0][0] = t
		var dfs func(i, j int)
		dfs = func(i, j int) {
			//fmt.Println(t, i, j)
			if i == m-1 && j == n-1 {
				die = false
				return
			}
			for _, d := range dir {
				x, y := d[0]+i, d[1]+j
				if x >= 0 && x < m && y >= 0 && y < n {
					//fmt.Println(x, y, dp[i][j]+1, f[x][y])
					if g[x][y] == 0 &&
						(dp[i][j]+1 < f[x][y] || f[x][y] == -1 || (x == m-1 && y == n-1 && dp[i][j]+1 == f[x][y])) &&
						dp[x][y] == 0 {       // grass && arrive < fire
						if x == 0 && y == 0 { // 不能回到0,0
							continue
						}
						dp[x][y] = dp[i][j] + 1
						dfs(x, y)
					}
				}
			}
		}
		dfs(0, 0)
		return die
	})

	if ans > m*n {
		return not
	}
	return ans - 1
}

func countUnguarded(m int, n int, g [][]int, w [][]int) int {
	h := make([][]int, m)
	for i := 0; i < m; i++ {
		h[i] = make([]int, n)
	}

	// 0, 1 - g, 2 - w, 3, vis
	for _, v := range g {
		h[v[0]][v[1]] = 1
	}
	for _, v := range w {
		h[v[0]][v[1]] = 2
	}

	pre := 0
	for i := 0; i < m; i++ {
		pre = 0
		for j := 0; j < n; j++ {
			if h[i][j] == 1 || h[i][j] == 2 {
				pre = h[i][j]
			} else if h[i][j] == 0 {
				if pre == 1 {
					h[i][j] = 3
				}
			}
		}
		pre = 0
		for j := n - 1; j >= 0; j-- {
			if h[i][j] == 1 || h[i][j] == 2 {
				pre = h[i][j]
			} else if h[i][j] == 0 {
				if pre == 1 {
					h[i][j] = 3
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		pre = 0
		for i := 0; i < m; i++ {
			if h[i][j] == 1 || h[i][j] == 2 {
				pre = h[i][j]
			} else if h[i][j] == 0 {
				if pre == 1 {
					h[i][j] = 3
				}
			}
		}
		pre = 0
		for i := m - 1; i >= 0; i-- {
			if h[i][j] == 1 || h[i][j] == 2 {
				pre = h[i][j]
			} else if h[i][j] == 0 {
				if pre == 1 {
					h[i][j] = 3
				}
			}
		}
	}

	fmt.Println(h)

	ans := 0
	for _, a := range h {
		for _, v := range a {
			if v == 0 {
				ans++
			}
		}
	}

	return ans
}

func countUnguarded_(m int, n int, g [][]int, w [][]int) int {
	h := make([][]int, m)
	for i := 0; i < m; i++ {
		h[i] = make([]int, n)
	}

	// 0, 1 - g, 2 - vis, 3, w
	for _, v := range w {
		h[v[0]][v[1]] = 3
	}

	dir := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for _, v := range g {
		h[v[0]][v[1]] = 1
		for _, d := range dir {
			for x, y := v[0]+d[0], v[1]+d[1]; x >= 0 && x < m && y >= 0 && y < n; {
				if h[x][y] == 1 || h[x][y] == 3 {
					break
				}
				h[x][y] = 2
				x += d[0]
				y += d[1]
			}
		}
	}

	ans := 0
	for _, ar := range h {
		for _, v := range ar {
			if v == 0 {
				ans++
			}
		}
	}
	return ans
}
