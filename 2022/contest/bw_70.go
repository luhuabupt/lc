package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(123)
}

var dir = []struct{ x, y int }{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	ans := [][]int{}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dp[start[0]][start[1]] = 0

	st := [][2]int{{start[0], start[1]}}
	if grid[start[0]][start[1]] >= pricing[0] && grid[start[0]][start[1]] <= pricing[1] {
		ans = append(ans, start)
	}
	if grid[start[0]][start[1]] == 0 {
		return ans
	}

	i := 0
	for len(st) > i {
		cur := st[i]
		for _, d := range dir {
			if x, y := cur[0]+d.x, cur[1]+d.y; x >= 0 && x < m && y >= 0 && y < n {
				if grid[x][y] > 0 && dp[x][y] == -1 {
					if grid[x][y] >= pricing[0] && grid[x][y] <= pricing[1] {
						ans = append(ans, []int{x, y})
					}

					st = append(st, [2]int{x, y})
					dp[x][y] = dp[cur[0]][cur[1]] + 1
				}
			}
		}

		i++
	}

	sort.Slice(ans, func(i, j int) bool {
		ix, iy, jx, jy := ans[i][0], ans[i][1], ans[j][0], ans[j][1]
		if dp[ix][iy] != dp[jx][jy] {
			return dp[ix][iy] < dp[jx][jy]
		}
		if grid[ix][iy] != grid[jx][jy] {
			return grid[ix][iy] < grid[jx][jy]
		}
		return ix < jx || ix == jx && iy < jy
	})

	if len(ans) < k {
		return ans
	}

	return ans[:k]
}

func numberOfWays(corridor string) int {
	n := len(corridor)
	M := int(1e9 + 7)

	sum := strings.Count(corridor, "S")
	if sum == 0 || sum&1 == 1 {
		return 0
	}

	ans := 1
	s, p := 0, 0
	for i := 0; i < n; i++ {
		v := corridor[i]

		if v == 'S' {
			if s == 2 {
				ans *= p + 1
				ans %= M
				p, s = 0, 0
			}
			s++
		}
		if s == 2 && v == 'P' {
			p++
		}
	}

	return ans
}
